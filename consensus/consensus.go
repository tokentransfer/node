package consensus

import (
	"errors"
	"fmt"
	"time"

	"github.com/caivega/glog"
	"github.com/tokentransfer/node/block"
	"github.com/tokentransfer/node/core"
	"github.com/tokentransfer/node/core/pb"
	"github.com/tokentransfer/node/util"

	libaccount "github.com/tokentransfer/interfaces/account"
	libblock "github.com/tokentransfer/interfaces/block"
	libcore "github.com/tokentransfer/interfaces/core"
	libcrypto "github.com/tokentransfer/interfaces/crypto"
	libstore "github.com/tokentransfer/interfaces/store"
)

type ConsensusService struct {
	CryptoService  libcrypto.CryptoService
	MerkleService  libstore.MerkleService
	Config         libcore.Config
	AccountService libaccount.AccountService
	StorageService *StorageService

	ValidatedBlock libblock.Block
}

func (service *ConsensusService) GetBlockNumber() int64 {
	if service.ValidatedBlock != nil {
		return int64(service.ValidatedBlock.GetIndex())
	}
	return -1
}

func (service *ConsensusService) GetBlockHash() string {
	if service.ValidatedBlock != nil {
		return service.ValidatedBlock.GetHash().String()
	}
	return ""
}

func (service *ConsensusService) GetBlock() libblock.Block {
	if service.ValidatedBlock != nil {
		return service.ValidatedBlock
	}
	return nil
}

func (service *ConsensusService) GenerateBlock(list []libblock.TransactionWithData) (libblock.Block, error) {
	cs := service.CryptoService
	as := service.AccountService
	ms := service.MerkleService
	ss := service.StorageService

	v := service.ValidatedBlock

	var b *block.Block
	if service.ValidatedBlock == nil { //genesis
		if len(list) > 0 {
			return nil, util.ErrorOfInvalid("transactions", "genesis block")
		}
		_, rootKey, err := as.GenerateFamilySeed("masterpassphrase")
		if err != nil {
			return nil, err
		}
		rootAccount, err := rootKey.GetAddress()
		if err != nil {
			return nil, err
		}
		a, err := core.NewAmount("100000000000000")
		if err != nil {
			return nil, err
		}
		states := []libblock.State{
			&block.AccountState{
				State: block.State{
					StateType:  block.ACCOUNT_STATE,
					Account:    rootAccount,
					Sequence:   uint64(0),
					BlockIndex: uint64(0),
				},
				Amount: *a,
			},
		}

		for i := 0; i < len(states); i++ {
			state := states[i]
			err := ms.PutState(state)
			if err != nil {
				return nil, err
			}
		}

		b = &block.Block{
			BlockIndex: uint64(0),
			ParentHash: libcrypto.ZeroHash(cs),

			Transactions:    []libblock.TransactionWithData{},
			TransactionHash: ms.GetTransactionRoot(),

			States:    states,
			StateHash: ms.GetStateRoot(),

			RootHash: libcore.Hash(ss.storage.Root()),

			Timestamp: time.Now().UnixNano(),
		}

		err = ms.Cancel()
		if err != nil {
			return nil, err
		}
	} else {
		glog.Infof("=== package %d transactions in block %d\n", len(list), v.GetIndex()+1)

		stateMap := map[string][]uint64{}
		for i := 0; i < len(list); i++ {
			txWithData := list[i]

			r := txWithData.GetReceipt()
			r.SetTransactionIndex(uint32(i))
			r.SetBlockIndex(v.GetIndex() + 1)
			states := r.GetStates()
			for j := 0; j < len(states); j++ {
				s := states[j]
				s.SetBlockIndex(v.GetIndex() + 1)

				key := fmt.Sprintf("%d-%s", s.GetStateType(), s.GetStateKey())
				index := s.GetIndex()
				stateMap[key] = []uint64{uint64(i), index}
			}

			err := ms.PutTransaction(txWithData)
			if err != nil {
				return nil, err
			}

			glog.Infof("=== %d %s\n", i, txWithData.GetTransaction().GetHash().String())
		}

		states := make([]libblock.State, 0)
		for i := 0; i < len(list); i++ {
			txWithData := list[i]

			r := txWithData.GetReceipt()
			rs := r.GetStates()
			for j := 0; j < len(rs); j++ {
				s := rs[j]

				key := fmt.Sprintf("%d-%s", s.GetStateType(), s.GetStateKey())
				item, ok := stateMap[key]
				if ok && item[0] == uint64(i) && item[1] == s.GetIndex() {
					states = append(states, s)
				}
			}
		}

		for i := 0; i < len(states); i++ {
			state := states[i]
			err := ms.PutState(state)
			if err != nil {
				return nil, err
			}
		}

		b = &block.Block{
			BlockIndex: v.GetIndex() + 1,
			ParentHash: v.GetHash(),

			Transactions:    list,
			TransactionHash: ms.GetTransactionRoot(),

			States:    states,
			StateHash: ms.GetStateRoot(),

			RootHash: libcore.Hash(ss.storage.Root()),

			Timestamp: time.Now().UnixNano(),
		}

		err := ms.Cancel()
		if err != nil {
			return nil, err
		}
	}

	_, _, err := cs.Raw(b, libcrypto.RawBinary)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (service *ConsensusService) VerifyBlock(b libblock.Block) (ok bool, err error) {
	ms := service.MerkleService
	cs := service.CryptoService
	ss := service.StorageService

	ok = true
	err = nil

	defer func() {
		if !ok || err != nil {
			ss.CancelSandbox()
			ms.Cancel()
		} else {
			ss.CommitSandbox()
		}
	}()

	err = ss.CreateSandbox()
	if err != nil {
		return
	}

	transactions := b.GetTransactions()
	l := len(transactions)
	for i := 0; i < l; i++ {
		txWithData := transactions[i]
		tx := txWithData.GetTransaction()

		ok, err = service.VerifyTransaction(tx)
		if err != nil {
			return
		}
		if !ok {
			err = util.ErrorOfInvalid("transaction", tx.GetHash().String())
			return
		}

		err = ss.CreateSandbox()
		if err != nil {
			return
		}
		newWithData, e := service.ProcessTransaction(tx)
		if e != nil {
			err = ss.CancelSandbox()
			if err != nil {
				glog.Error(err)
			}
			ok = false
			err = e
			return
		}
		err = ss.CommitSandbox()
		if err != nil {
			return
		}

		arh, _, e := cs.Raw(txWithData.GetReceipt(), libcrypto.RawBinary)
		if e != nil {
			ok = false
			err = e
			return
		}
		brh, _, e := cs.Raw(newWithData.GetReceipt(), libcrypto.RawBinary)
		if e != nil {
			ok = false
			err = e
			return
		}
		if arh.String() != brh.String() {
			ok = false
			err = util.ErrorOf("unmatched", "transaction receipt", fmt.Sprintf("%s != %s", arh.String(), brh.String()))
			return
		}

		err = ms.PutTransaction(txWithData)
		if err != nil {
			ok = false
			return
		}
	}

	if service.ValidatedBlock != nil {
		if b.GetIndex() != (service.ValidatedBlock.GetIndex() + 1) {
			ok = false
			err = fmt.Errorf("error block index: %d != %d", b.GetIndex(), (service.ValidatedBlock.GetIndex() + 1))
			return
		}
		if b.GetParentHash().String() != service.ValidatedBlock.GetHash().String() {
			ok = false
			err = fmt.Errorf("error parent hash: %s != %s", b.GetParentHash().String(), service.ValidatedBlock.GetHash().String())
			return
		}
	} else {
		if b.GetIndex() != 0 {
			ok = false
			err = errors.New("error block index")
			return
		}
		if !b.GetParentHash().IsZero() {
			ok = false
			err = errors.New("error parent hash")
			return
		}
	}

	states := b.GetStates()
	l = len(states)
	for i := 0; i < l; i++ {
		state := states[i]
		err = ms.PutState(state)
		if err != nil {
			ok = false
			return
		}
	}

	transactionHash := ms.GetTransactionRoot()
	stateHash := ms.GetStateRoot()
	if !b.GetTransactionHash().Equals(transactionHash) {
		ok = false
		err = fmt.Errorf("error transaction hash: %s != %s", b.GetTransactionHash().String(), transactionHash.String())
		return
	}
	if !b.GetStateHash().Equals(stateHash) {
		ok = false
		err = fmt.Errorf("error state hash: %s != %s", b.GetStateHash().String(), stateHash.String())
		return
	}
	return
}

func (service *ConsensusService) AddBlock(b libblock.Block) error {
	ms := service.MerkleService

	err := ms.PutBlock(b)
	if err != nil {
		return err
	}
	err = ms.Commit()
	if err != nil {
		return err
	}
	service.ValidatedBlock = b

	return nil
}

func (service *ConsensusService) GetAccountInfo(account libcore.Address, currency *libcore.Symbol, issuer libcore.Address) (*block.AccountState, error) {
	ms := service.MerkleService

	accountKey := core.GetAccountKey(account, currency, issuer, "-")
	state, err := ms.GetStateByTypeAndKey(block.ACCOUNT_STATE, accountKey)
	if err != nil {
		return nil, err
	}
	info, ok := state.(*block.AccountState)
	if !ok {
		return nil, errors.New("error account state")
	}
	return info, nil
}

func (service *ConsensusService) VerifyTransaction(t libblock.Transaction) (bool, error) {
	cs := service.CryptoService
	ms := service.MerkleService

	ok, err := cs.Verify(t)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, errors.New("error transaction")
	}
	tx, ok := t.(*block.Transaction)
	if !ok {
		return false, errors.New("error transaction")
	}
	if !tx.Amount.IsNative() {
		return false, util.ErrorOfInvalid("amount", tx.Amount.String())
	}

	account := tx.Account
	sequence := libstore.GetSequence(ms, account) + 1
	if tx.Sequence != sequence {
		return false, fmt.Errorf("error sequence: %d != %d", tx.Sequence, sequence)
	}

	info, err := service.GetAccountInfo(account, nil, nil)
	if err != nil || info == nil {
		return false, util.ErrorOfNonexists("account", account.String())
	}

	if tx.Gas < 10 {
		return false, util.ErrorOf("insufficient", "gas", fmt.Sprintf("%d < 10", tx.Gas))
	}
	if tx.Account.String() == tx.Destination.String() && !tx.Amount.IsZero() {
		return false, util.ErrorOfInvalid("amount", "should be 0 when using same accounts")
	}
	gasAmount, err := core.NewAmount(int64(tx.Gas))
	if err != nil {
		return false, err
	}
	totalAmount, err := tx.Amount.Add(*gasAmount)
	if err != nil {
		return false, err
	}
	remain, err := info.Amount.Subtract(*totalAmount)
	if err != nil {
		return false, err
	}
	if remain.Less(*gasAmount.ZeroClone()) {
		return false, util.ErrorOf("insuffient", "amount", remain.String())
	}

	if tx.Payload != nil && len(tx.Payload.Infos) > 0 {
		for _, payload := range tx.Payload.Infos {
			meta, msg, err := core.Unmarshal(payload.Content)
			if err != nil {
				return false, err
			}
			switch meta {
			case core.CORE_PAYLOAD_INFO:
			case core.CORE_CONTRACT_INFO:
				info := msg.(*pb.ContractInfo)
				if !(len(info.Method) > 0 && info.Params != nil) {
					return false, util.ErrorOfInvalid("format", "contract info")
				}
			case core.CORE_PAGE_INFO:
			case core.CORE_CODE_INFO:
				info := msg.(*pb.CodeInfo)
				if !(len(info.Code) > 0) {
					return false, util.ErrorOfInvalid("format", "code info")
				}
			case core.CORE_META_INFO:
			case core.CORE_TOKEN_INFO:
			case core.CORE_DATA_INFO:
			case core.CORE_PEER_INFO:
			default:
				return false, util.ErrorOfInvalid("format", "info")
			}
		}
	}

	return true, nil
}

func (service *ConsensusService) getAccountInfo(account libcore.Address, a core.Amount, sequence uint64) (*block.AccountState, error) {
	lastInfo, _ := service.GetAccountInfo(account, a.Currency, a.Issuer)
	if lastInfo != nil {
		s, err := block.CloneState(lastInfo)
		if err != nil {
			return nil, err
		}
		h, err := service.HashState(s)
		if err != nil {
			return nil, err
		}
		info := s.(*block.AccountState)
		info.Sequence = sequence
		info.Previous = h
		info.Version = lastInfo.Version + 1
		return info, nil
	} else {
		info := &block.AccountState{
			State: block.State{
				StateType: block.ACCOUNT_STATE,
				Account:   account,
				Sequence:  sequence,
			},
			Amount: *a.ZeroClone(),
		}
		return info, nil
	}
}

func (service *ConsensusService) addBalance(info *block.AccountState, amount *core.Amount) error {
	newAmount, err := info.Amount.Add(*amount)
	if err != nil {
		return err
	}
	info.Amount = *newAmount
	return nil
}

func (service *ConsensusService) removeBalance(info *block.AccountState, amount *core.Amount) error {
	newAmount, err := info.Amount.Subtract(*amount)
	if err != nil {
		return err
	}
	if newAmount.Less(*amount.ZeroClone()) {
		return errors.New("insuffient amount")
	}
	info.Amount = *newAmount
	return nil
}

func (service *ConsensusService) ProcessTransaction(t libblock.Transaction) (libblock.TransactionWithData, error) {
	config := service.Config
	ms := service.MerkleService

	tx, ok := t.(*block.Transaction)
	if !ok {
		return nil, util.ErrorOfInvalid("transaction", t.GetHash().String())
	}
	if !tx.Amount.IsNative() {
		return nil, util.ErrorOfInvalid("amount", tx.Amount.String())
	}

	sequenceMap := map[string]uint64{}
	accountMap := map[string]*block.AccountState{}

	fromAccount := tx.GetAccount()
	fromSequence := libstore.GetSequence(ms, fromAccount) + 1 // sequence of from from account, +1
	if tx.Sequence != fromSequence {
		return nil, util.ErrorOf("error", "sequence", fmt.Sprintf("%d != %d", tx.Sequence, fromSequence))
	}
	sequenceMap[fromAccount.String()] = fromSequence

	states := make([]libblock.State, 0)

	fromInfo, err := service.getAccountInfo(fromAccount, tx.Amount, fromSequence)
	if err != nil {
		return nil, err
	}
	states = append(states, fromInfo)
	accountMap[tx.Account.String()] = fromInfo

	destAccount := tx.Destination
	destSequence := libstore.GetSequence(ms, destAccount) // sequence of destination account, keep unchanged
	destInfo, err := service.getAccountInfo(destAccount, tx.Amount, destSequence)
	if err != nil {
		return nil, err
	}
	isSame := (fromAccount.String() == destAccount.String())
	if !isSame {
		sequenceMap[destAccount.String()] = destSequence
		states = append(states, destInfo)
		accountMap[destAccount.String()] = destInfo
	}

	gasAccount := config.GetGasAccount()
	gasSequence := libstore.GetSequence(ms, gasAccount)
	gasInfo, err := service.getAccountInfo(gasAccount, tx.Amount, gasSequence) // sequence of gas account, keep unchanged
	if err != nil {
		return nil, err
	}
	sequenceMap[gasAccount.String()] = gasSequence
	states = append(states, gasInfo)
	accountMap[gasAccount.String()] = gasInfo

	gas := int64(0) //gasPrice gasLimit
	if tx.Payload != nil && len(tx.Payload.Infos) > 0 {
		dataCost := int64(0)
		for _, info := range tx.Payload.Infos {
			dataCost += int64(len(info.Content)) // 1:1
		}
		wasmCost, payloadStates, err := service.ProcessPayload(tx.Gas-dataCost, tx, tx.Payload, accountMap)
		if err != nil {
			return nil, err
		}
		gas += dataCost
		gas += wasmCost
		states = append(states, payloadStates...)
	}
	if gas < 10 {
		gas = 10
	}
	if tx.Gas < gas {
		return nil, util.ErrorOf("insufficient", "gas", fmt.Sprintf("%d < %d", tx.Gas, gas))
	}

	gasAmount, err := core.NewAmount(int64(gas))
	if err != nil {
		return nil, err
	}
	totalAmount, err := tx.Amount.Add(*gasAmount)
	if err != nil {
		return nil, err
	}
	err = service.removeBalance(fromInfo, totalAmount)
	if err != nil {
		return nil, err
	}
	err = service.addBalance(destInfo, &tx.Amount)
	if err != nil {
		return nil, err
	}
	err = service.addBalance(gasInfo, gasAmount)
	if err != nil {
		return nil, err
	}

	r := &block.Receipt{
		TransactionResult: 0,
		States:            states,
	}

	return &block.TransactionWithData{
		Transaction: t,
		Receipt:     r,
		Date:        int64(time.Now().UnixNano() / 1e9),
	}, nil
}

func (service *ConsensusService) ProcessPayload(remainCost int64, tx *block.Transaction, info *block.PayloadInfo, accountMap map[string]*block.AccountState) (int64, []libblock.State, error) {
	cs := service.CryptoService
	as := service.AccountService
	ss := service.StorageService

	states := make([]libblock.State, 0)
	cost := int64(0)
	for _, payload := range info.Infos {
		meta, msg, err := core.Unmarshal(payload.Content)
		if err != nil {
			return cost, nil, err
		}
		switch meta {
		case core.CORE_PAYLOAD_INFO:
			info := msg.(*pb.PayloadInfo)
			infoData, err := core.Marshal(info)
			if err != nil {
				return cost, nil, err
			}
			payloadInfo := &block.PayloadInfo{}
			err = payloadInfo.UnmarshalBinary(infoData)
			if err != nil {
				return cost, nil, err
			}
			usedCost, payloadStates, err := service.ProcessPayload(remainCost-cost, tx, payloadInfo, accountMap)
			if err != nil {
				return 0, nil, err
			}
			states = append(states, payloadStates...)
			cost += usedCost

		case core.CORE_CONTRACT_INFO:
			info := msg.(*pb.ContractInfo)
			inputs, err := getAccounts(as, info.Inputs)
			if err != nil {
				return 0, nil, err
			}
			outputs, err := getAccounts(as, info.Outputs)
			if err != nil {
				return 0, nil, err
			}
			usedCost, _, dataHash, dataContent, err := ss.RunContract(cs, remainCost, tx.Account, tx.Destination, info.Method, info.Params, inputs, outputs)
			if err != nil {
				return cost, nil, err
			}
			accountInfo, ok := accountMap[tx.Destination.String()]
			if ok {
				accountInfo.Data = &block.DataInfo{
					Hash:    dataHash,
					Content: dataContent,
				}
			}
			cost += usedCost

		case core.CORE_PAGE_INFO:
			info := msg.(*pb.PageInfo)
			if len(info.Data) > 0 {
				_, pageHash, err := ss.CreatePage(info.Name, tx.Destination, info.Data)
				if err != nil {
					return cost, nil, err
				}
				accountInfo, ok := accountMap[tx.Destination.String()]
				if ok {
					accountInfo.Page = &block.DataInfo{
						Hash: pageHash,
					}
				}
			}

		case core.CORE_CODE_INFO:
			info := msg.(*pb.CodeInfo)
			_, codeHash, err := ss.CreateContract(tx.Destination, info.Code, info.Abi)
			if err != nil {
				return cost, nil, err
			}
			accountInfo, ok := accountMap[tx.Destination.String()]
			if ok {
				accountInfo.Code = &block.DataInfo{
					Hash: codeHash,
				}
			}

		case core.CORE_META_INFO:
		case core.CORE_TOKEN_INFO:
		case core.CORE_DATA_INFO:
		default:
			return cost, nil, util.ErrorOfUnknown("format", "info")
		}
	}

	return cost, states, nil
}

func getAccounts(as libaccount.AccountService, list [][]byte) ([]libcore.Address, error) {
	if list == nil {
		return nil, nil
	}
	accounts := make([]libcore.Address, 0)
	for _, item := range list {
		_, account, err := as.NewAccountFromBytes(item)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func (service *ConsensusService) HashBlock(b libblock.Block) (libcore.Hash, error) {
	h, _, err := service.CryptoService.Raw(b, libcrypto.RawBinary)
	if err != nil {
		return nil, err
	}
	transactions := b.GetTransactions()
	for i := 0; i < len(transactions); i++ {
		txWithData := transactions[i]
		_, err := service.HashTransaction(txWithData)
		if err != nil {
			return nil, err
		}
	}
	states := b.GetStates()
	for i := 0; i < len(states); i++ {
		state := states[i]
		_, err := service.HashState(state)
		if err != nil {
			return nil, err
		}
	}
	return h, nil
}

func (service *ConsensusService) HashTransaction(txWithData libblock.TransactionWithData) (libcore.Hash, error) {
	_, _, err := service.CryptoService.Raw(txWithData, libcrypto.RawBinary)
	if err != nil {
		return nil, err
	}

	h, _, err := service.CryptoService.Raw(txWithData.GetTransaction(), libcrypto.RawBinary)
	if err != nil {
		return nil, err
	}

	_, _, err = service.CryptoService.Raw(txWithData.GetReceipt(), libcrypto.RawBinary)
	if err != nil {
		return nil, err
	}

	_, err = service.HashReceipt(txWithData.GetReceipt())
	if err != nil {
		return nil, err
	}
	return h, nil
}

func (service *ConsensusService) HashReceipt(r libblock.Receipt) (libcore.Hash, error) {
	h, _, err := service.CryptoService.Raw(r, libcrypto.RawBinary)
	if err != nil {
		return nil, err
	}

	states := r.GetStates()
	for i := 0; i < len(states); i++ {
		state := states[i]
		_, err := service.HashState(state)
		if err != nil {
			return nil, err
		}
	}
	return h, nil
}

func (service *ConsensusService) HashState(s libblock.State) (libcore.Hash, error) {
	h, _, err := service.CryptoService.Raw(s, libcrypto.RawBinary)
	if err != nil {
		return nil, err
	}
	return h, nil
}
