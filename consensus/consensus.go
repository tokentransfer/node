package consensus

import (
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
		rootPublic, err := rootKey.GetPublic()
		if err != nil {
			return nil, err
		}
		rootPublicKey, err := rootPublic.MarshalBinary()
		if err != nil {
			return nil, err
		}
		rootAccount, err := rootKey.GetAddress()
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

				PublicKey: libcore.PublicKey(rootPublicKey),
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

				keys := s.GetStateKey()
				for k := 0; k < len(keys); k++ {
					key := fmt.Sprintf("%d-%s", s.GetStateType(), keys[k])
					index := s.GetIndex()
					stateMap[key] = []uint64{uint64(i), index}
				}
			}

			err := ms.PutTransaction(txWithData)
			if err != nil {
				return nil, err
			}

			glog.Infof("=== %d %s\n", i, txWithData.GetTransaction().GetHash().String())
		}

		states := make([]libblock.State, 0)
		hashMap := make(map[string]libblock.State)
		for i := 0; i < len(list); i++ {
			txWithData := list[i]

			r := txWithData.GetReceipt()
			rs := r.GetStates()
			for j := 0; j < len(rs); j++ {
				s := rs[j]

				keys := s.GetStateKey()
				for k := 0; k < len(keys); k++ {
					key := fmt.Sprintf("%d-%s", s.GetStateType(), keys[k])
					item, sok := stateMap[key]
					_, hok := hashMap[s.GetHash().String()]
					if sok && !hok && item[0] == uint64(i) && item[1] == s.GetIndex() {
						states = append(states, s)
						hashMap[s.GetHash().String()] = s
					}
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

func (service *ConsensusService) AddBlock(b libblock.Block) error {
	ms := service.MerkleService
	ss := service.StorageService

	err := ss.CommitSandbox()
	if err != nil {
		return err
	}
	err = ms.PutBlock(b)
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
			err = util.ErrorOfUnmatched("raw hash", "transaction receipt", arh.String(), brh.String())
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
			err = util.ErrorOfUnmatched("index", "block", (service.ValidatedBlock.GetIndex() + 1), b.GetIndex())
			return
		}
		if b.GetParentHash().String() != service.ValidatedBlock.GetHash().String() {
			ok = false
			err = util.ErrorOfUnmatched("parent hash", "block", service.ValidatedBlock.GetHash().String(), b.GetParentHash().String())
			return
		}
	} else {
		if b.GetIndex() != 0 {
			ok = false
			err = util.ErrorOfInvalid("index", "block")
			return
		}
		if !b.GetParentHash().IsZero() {
			ok = false
			err = util.ErrorOfInvalid("parent hash", "block")
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
	if b.GetTransactionHash().String() != transactionHash.String() {
		ok = false
		err = util.ErrorOfUnmatched("hash", "transaction", b.GetTransactionHash().String(), transactionHash.String())
		return
	}
	if b.GetStateHash().String() != stateHash.String() {
		ok = false
		err = util.ErrorOfUnmatched("hash", "state", b.GetStateHash().String(), stateHash.String())
		return
	}
	return
}

func (service *ConsensusService) GetAccountInfo(account libcore.Address) (*block.AccountState, error) {
	ms := service.MerkleService

	accountKey := account.String()
	state, err := ms.GetStateByTypeAndKey(block.ACCOUNT_STATE, accountKey)
	if err != nil {
		return nil, err
	}
	info, ok := state.(*block.AccountState)
	if !ok {
		return nil, util.ErrorOfInvalid("data", "account state")
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
		return false, util.ErrorOfInvalid("verify", "transaction")
	}
	tx, ok := t.(*block.Transaction)
	if !ok {
		return false, util.ErrorOfInvalid("data", "transaction")
	}

	account := tx.Account
	sequence := libstore.GetSequence(ms, account) + 1
	if tx.Sequence != sequence {
		return false, util.ErrorOfInvalid("sequence", fmt.Sprintf("%d != %d", tx.Sequence, sequence))
	}

	info, err := service.GetAccountInfo(account)
	if err != nil || info == nil {
		return false, util.ErrorOfNotFound("account", account.String())
	}

	if tx.Gas < 10 {
		return false, util.ErrorOf("insufficient", "gas", fmt.Sprintf("%d < 10", tx.Gas))
	}
	gasValue, err := util.NewValue(fmt.Sprintf("%d", tx.Gas))
	if err != nil {
		return false, err
	}
	fromValue, err := service.StorageService.GetGas(account)
	if err != nil {
		return false, err
	}
	remain, err := fromValue.Subtract(*gasValue)
	if err != nil {
		return false, err
	}
	isNegative, err := remain.IsNegative()
	if err != nil {
		return false, err
	}
	if isNegative {
		return false, util.ErrorOf("insuffient", "gas", remain.String())
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
				if !(len(info.Method) > 0) {
					return false, util.ErrorOfInvalid("format", "contract info")
				}
			case core.CORE_PAGE_INFO:
			case core.CORE_CODE_INFO:
				info := msg.(*pb.CodeInfo)
				if !(len(info.Code) > 0) {
					return false, util.ErrorOfInvalid("format", "code info")
				}
			case core.CORE_USER_INFO:
				info := msg.(*pb.UserInfo)
				if !(len(info.Account) > 0 && info.Data != nil && (len(info.Data.Content) > 0 || len(info.Data.Hash) > 0)) {
					return false, util.ErrorOfInvalid("format", "user info")
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

func (service *ConsensusService) getAccountInfo(account libcore.Address, sequence uint64) (*block.AccountState, *block.AccountState, error) {
	lastInfo, _ := service.GetAccountInfo(account)
	if lastInfo != nil {
		s, err := block.CloneState(lastInfo)
		if err != nil {
			return nil, nil, err
		}
		h, err := service.HashState(s)
		if err != nil {
			return nil, nil, err
		}
		info := s.(*block.AccountState)
		if sequence < info.Sequence {
			sequence = info.Sequence
		}
		info.Sequence = sequence
		info.Previous = h
		info.Version = lastInfo.Version
		return lastInfo, info, nil
	} else {
		info := &block.AccountState{
			State: block.State{
				StateType: block.ACCOUNT_STATE,
				Account:   account,
				Sequence:  sequence,
			},
		}
		return nil, info, nil
	}
}

func (service *ConsensusService) addBalance(info *block.AccountState, value util.Value) (*util.Value, error) {
	oldValue, err := service.StorageService.GetGas(info.Account)
	if err != nil {
		return nil, err
	}
	newValue, err := oldValue.Add(value)
	if err != nil {
		return nil, err
	}
	err = service.StorageService.UpdateGas(info.Account, newValue)
	if err != nil {
		return nil, err
	}
	return &newValue, nil
}

func (service *ConsensusService) removeBalance(info *block.AccountState, value util.Value) (*util.Value, error) {
	oldValue, err := service.StorageService.GetGas(info.Account)
	if err != nil {
		return nil, err
	}
	newValue, err := oldValue.Subtract(value)
	if err != nil {
		return nil, err
	}
	isNegative, err := newValue.IsNegative()
	if err != nil {
		return nil, err
	}
	if isNegative {
		return nil, util.ErrorOf("insuffient gas", "nagative", "balance")
	}
	err = service.StorageService.UpdateGas(info.Account, newValue)
	if err != nil {
		return nil, err
	}
	return &newValue, nil
}

func (service *ConsensusService) ProcessTransaction(t libblock.Transaction) (libblock.TransactionWithData, error) {
	config := service.Config
	ms := service.MerkleService

	tx, ok := t.(*block.Transaction)
	if !ok {
		return nil, util.ErrorOfInvalid("transaction", t.GetHash().String())
	}

	sequenceMap := map[string]uint64{}
	accountMap := map[string]*block.AccountState{}
	lastMap := map[string]*block.AccountState{}
	accounts := make([]string, 0)

	fromAccount := tx.GetAccount()
	fromSequence := libstore.GetSequence(ms, fromAccount) + 1 // sequence of from from account, +1
	if tx.Sequence != fromSequence {
		return nil, util.ErrorOf("error", "sequence", fmt.Sprintf("%d != %d", tx.Sequence, fromSequence))
	}
	sequenceMap[fromAccount.String()] = fromSequence

	lastFromInfo, fromInfo, err := service.getAccountInfo(fromAccount, fromSequence)
	if err != nil {
		return nil, err
	}
	fromInfo.PublicKey = tx.PublicKey
	accountMap[tx.Account.String()] = fromInfo
	lastMap[tx.Account.String()] = lastFromInfo
	accounts = append(accounts, tx.Account.String())

	destAccount := tx.Destination
	destSequence := libstore.GetSequence(ms, destAccount) // sequence of destination account, keep unchanged
	lastDestInfo, destInfo, err := service.getAccountInfo(destAccount, destSequence)
	if err != nil {
		return nil, err
	}
	isSame := (fromAccount.String() == destAccount.String())
	if !isSame {
		sequenceMap[destAccount.String()] = destSequence
		accountMap[destAccount.String()] = destInfo
		lastMap[destAccount.String()] = lastDestInfo
		accounts = append(accounts, destAccount.String())
	}

	gasAccount := config.GetGasAccount()
	gasSequence := libstore.GetSequence(ms, gasAccount)
	lastGasInfo, gasInfo, err := service.getAccountInfo(gasAccount, gasSequence) // sequence of gas account, keep unchanged
	if err != nil {
		return nil, err
	}
	sequenceMap[gasAccount.String()] = gasSequence
	accountMap[gasAccount.String()] = gasInfo
	lastMap[gasAccount.String()] = lastGasInfo
	accounts = append(accounts, gasAccount.String())

	gas := uint64(0) //gasPrice gasLimit
	if tx.Payload != nil && len(tx.Payload.Infos) > 0 {
		dataCost := uint64(0)
		for _, info := range tx.Payload.Infos {
			dataCost += uint64(len(info.Content)) // 1:1
		}
		wasmCost, payloadStates, err := service.ProcessPayload(tx.Gas-dataCost, tx, tx.Payload, accountMap, lastMap)
		if err != nil {
			return nil, err
		}
		gas += dataCost
		gas += wasmCost
		accounts = append(accounts, payloadStates...)
	}
	if gas < 10 {
		gas = 10
	}
	if tx.Gas < gas {
		return nil, util.ErrorOf("insufficient", "gas", fmt.Sprintf("%d < %d", tx.Gas, gas))
	}

	gasAmount, err := util.NewValue(fmt.Sprintf("%d", gas))
	if err != nil {
		return nil, err
	}
	_, err = service.removeBalance(fromInfo, *gasAmount)
	if err != nil {
		return nil, err
	}
	_, err = service.addBalance(gasInfo, *gasAmount)
	if err != nil {
		return nil, err
	}

	stateMap := make(map[string]string)
	stateList := make([]string, 0)
	for _, a := range accounts {
		_, ok := stateMap[a]
		if !ok {
			stateMap[a] = a
			stateList = append(stateList, a)
		}
	}

	states := make([]libblock.State, 0)
	for _, a := range stateList {
		lastInfo := lastMap[a]
		info := accountMap[a]
		if lastInfo == nil || info == nil {
			return nil, util.ErrorOfInvalid("account info", a)
		} else {
			if lastInfo == nil {
				states = append(states, info)
			} else {
				lastHash, err := service.HashState(lastInfo)
				if err != nil {
					return nil, err
				}
				hash, err := service.HashState(info)
				if err != nil {
					return nil, err
				}
				if lastHash.String() != hash.String() {
					info.Version = info.Version + 1
					states = append(states, info)
				}
			}
		}
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

func (service *ConsensusService) ProcessPayload(remainCost uint64, tx *block.Transaction, info *block.PayloadInfo, accountMap map[string]*block.AccountState, lastMap map[string]*block.AccountState) (uint64, []string, error) {
	cs := service.CryptoService
	as := service.AccountService
	ss := service.StorageService

	accounts := make([]string, 0)
	cost := uint64(0)
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
			usedCost, payloadStates, err := service.ProcessPayload(remainCost-cost, tx, payloadInfo, accountMap, lastMap)
			if err != nil {
				return 0, nil, err
			}
			accounts = append(accounts, payloadStates...)
			cost += usedCost

		case core.CORE_CONTRACT_INFO:
			info := msg.(*pb.ContractInfo)

			dataAccount := tx.Account
			if len(info.Account) > 0 {
				_, account, err := as.NewAccountFromBytes(info.Account)
				if err != nil {
					return 0, nil, err
				}
				dataAccount = account
			}

			inputs, err := getAccounts(as, info.Inputs)
			if err != nil {
				return 0, nil, err
			}
			outputs, err := getAccounts(as, info.Outputs)
			if err != nil {
				return 0, nil, err
			}
			usedCost, retAccount, _, retHash, retContent, err := ss.RunContract(cs, remainCost, tx.Account, dataAccount, tx.Destination, info.Method, info.Params, inputs, outputs)
			if err != nil {
				return cost, nil, err
			}
			accountInfo, ok := accountMap[retAccount.String()]
			if !ok {
				lastInfo, info, err := service.getAccountInfo(retAccount, 0)
				if err != nil {
					return cost, nil, err
				}
				accountMap[retAccount.String()] = info
				lastMap[retAccount.String()] = lastInfo
				accounts = append(accounts, retAccount.String())

				accountInfo = info
			}
			accountInfo.Data = &block.DataInfo{
				Hash:    retHash,
				Content: retContent,
			}
			cost += uint64(usedCost)

		case core.CORE_PAGE_INFO:
			info := msg.(*pb.PageInfo)
			if len(info.Data) > 0 {
				_, pageHash, err := ss.WritePage(info.Name, tx.Destination, info.Data)
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

		case core.CORE_USER_INFO:
			info := msg.(*pb.UserInfo)
			_, account, err := as.NewAccountFromBytes(info.Account)
			if err != nil {
				return cost, nil, err
			}
			userHash, userAccount, err := ss.WriteUser(tx.Account, tx.Destination, account, info)
			if err != nil {
				return cost, nil, err
			}
			accountInfo, ok := accountMap[userAccount.String()]
			if ok {
				accountInfo.User = &block.DataInfo{
					Hash: userHash,
				}
			}

		case core.CORE_META_INFO:
		case core.CORE_TOKEN_INFO:
		case core.CORE_DATA_INFO:
		default:
			return cost, nil, util.ErrorOfUnknown("format", "info")
		}
	}

	return cost, accounts, nil
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
