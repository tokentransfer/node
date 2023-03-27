package consensus

import (
	"errors"
	"fmt"
	"time"

	"github.com/caivega/glog"
	"github.com/tokentransfer/node/block"
	"github.com/tokentransfer/node/core"
	"github.com/tokentransfer/node/core/pb"

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
	// ss := service.StorageService

	v := service.ValidatedBlock

	var b *block.Block
	if service.ValidatedBlock == nil { //genesis
		if len(list) > 0 {
			return nil, errors.New("error genesis block")
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

			Timestamp: time.Now().UnixNano(),
		}

		err = ms.Cancel()
		if err != nil {
			return nil, err
		}
	} else {
		fmt.Printf("=== package %d transactions in block %d\n", len(list), v.GetIndex()+1)

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

			fmt.Printf("=== %d %s\n", i, txWithData.GetTransaction().GetHash().String())
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
			err = errors.New("verify transaction failed")
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
		if !arh.Equals(brh) {
			ok = false
			err = errors.New("process transaction receipt failed")
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
		if !b.GetParentHash().Equals(service.ValidatedBlock.GetHash()) {
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
		return false, core.ErrorOfInvalid("amount", tx.Amount.String())
	}

	account := tx.Account
	sequence := libstore.GetSequence(ms, account) + 1
	if tx.Sequence != sequence {
		return false, fmt.Errorf("error sequence: %d != %d", tx.Sequence, sequence)
	}

	info, err := service.GetAccountInfo(account, nil, nil)
	if err != nil || info == nil {
		return false, core.ErrorOfNonexists("account", account.String())
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
		return false, core.ErrorOf("insuffient", "amount", remain.String())
	}

	if tx.Payload != nil && len(tx.Payload) > 0 {
		meta, msg, err := core.Unmarshal(tx.Payload)
		if err != nil {
			return false, err
		}
		if meta != core.CORE_PAYLOAD_INFO {
			return false, core.ErrorOfInvalid("format", "payload")
		}
		info := msg.(*pb.PayloadInfo)
		for _, payload := range info.Payload {
			meta, _, err = core.Unmarshal(payload)
			if err != nil {
				return false, err
			}
			if meta != core.CORE_PAYLOAD_INFO && meta != core.CORE_CONTRACT_INFO && meta != core.CORE_META_INFO && meta != core.CORE_TOKEN_INFO && meta != core.CORE_DATA_INFO && meta != core.CORE_PEER_INFO && meta != core.CORE_PAGE_INFO {
				return false, core.ErrorOfInvalid("format", "info")
			}
		}
	}

	return true, nil
}

func (service *ConsensusService) addBalance(account libcore.Address, amount *core.Amount, sequence uint64) (*block.AccountState, error) {
	info, _ := service.GetAccountInfo(account, amount.Currency, amount.Issuer)
	if info != nil {
		s, err := block.CloneState(info)
		if err != nil {
			return nil, err
		}
		newAmount, err := info.Amount.Add(*amount)
		if err != nil {
			return nil, err
		}

		info = s.(*block.AccountState)
		info.Amount = *newAmount
		info.Sequence = sequence
	} else {
		info = &block.AccountState{
			State: block.State{
				StateType: block.ACCOUNT_STATE,
				Account:   account,
				Sequence:  sequence,
			},
			Amount: *amount,
		}
	}
	return info, nil
}

func (service *ConsensusService) removeBalance(account libcore.Address, amount *core.Amount, sequence uint64) (*block.AccountState, error) {
	info, _ := service.GetAccountInfo(account, amount.Currency, amount.Issuer)
	if info != nil {
		s, err := block.CloneState(info)
		if err != nil {
			return nil, err
		}
		newAmount, err := info.Amount.Subtract(*amount)
		if err != nil {
			return nil, err
		}
		if newAmount.Less(*amount.ZeroClone()) {
			return nil, errors.New("insuffient amount")
		}
		info = s.(*block.AccountState)
		info.Amount = *newAmount
		info.Sequence = sequence
		return info, nil
	}
	return nil, errors.New("insuffient amount")
}

func (service *ConsensusService) ProcessTransaction(t libblock.Transaction) (libblock.TransactionWithData, error) {
	config := service.Config
	ms := service.MerkleService

	tx, ok := t.(*block.Transaction)
	if !ok {
		return nil, errors.New("error transaction")
	}
	if !tx.Amount.IsNative() {
		return nil, core.ErrorOfInvalid("amount", tx.Amount.String())
	}

	sequenceMap := map[string]uint64{}
	accountMap := map[string]*block.AccountState{}

	account := tx.GetAccount()
	sequence := libstore.GetSequence(ms, account) + 1 // sequence of from account, +1
	if tx.Sequence != sequence {
		return nil, core.ErrorOf("error", "sequence", fmt.Sprintf("%d != %d", tx.Sequence, sequence))
	}
	sequenceMap[account.String()] = sequence

	gasAmount, err := core.NewAmount(int64(tx.Gas))
	if err != nil {
		return nil, err
	}

	states := make([]libblock.State, 0)

	totalAmount, err := tx.Amount.Add(*gasAmount)
	if err != nil {
		return nil, err
	}
	s, err := service.removeBalance(tx.Account, totalAmount, sequence)
	if err != nil {
		return nil, err
	}
	states = append(states, s)
	accountMap[tx.Account.String()] = s

	gasAccount := config.GetGasAccount()
	gasSequence := libstore.GetSequence(ms, gasAccount)
	s, err = service.addBalance(gasAccount, gasAmount, gasSequence) // sequence of gas account, keep unchanged
	if err != nil {
		return nil, err
	}
	sequenceMap[gasAccount.String()] = gasSequence
	states = append(states, s)
	accountMap[gasAccount.String()] = s

	destination := tx.Destination
	destinationSequence := libstore.GetSequence(ms, destination)
	s, err = service.addBalance(destination, &tx.Amount, destinationSequence) // sequence of destination account, keep unchanged
	if err != nil {
		return nil, err
	}
	sequenceMap[destination.String()] = destinationSequence
	states = append(states, s)
	accountMap[destination.String()] = s

	if tx.Payload != nil && len(tx.Payload) > 0 {
		meta, msg, err := core.Unmarshal(tx.Payload)
		if err != nil {
			return nil, err
		}
		if meta != core.CORE_PAYLOAD_INFO {
			return nil, core.ErrorOfInvalid("format", "payload")
		}
		info := msg.(*pb.PayloadInfo)
		payloadStates, err := service.ProcessPayload(tx, info, accountMap)
		if err != nil {
			return nil, err
		}
		states = append(states, payloadStates...)
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

func (service *ConsensusService) ProcessPayload(tx *block.Transaction, info *pb.PayloadInfo, accountMap map[string]*block.AccountState) ([]libblock.State, error) {
	as := service.AccountService
	ss := service.StorageService

	states := make([]libblock.State, 0)
	for _, payload := range info.Payload {
		meta, msg, err := core.Unmarshal(payload)
		if err != nil {
			return nil, err
		}
		switch meta {
		case core.CORE_PAYLOAD_INFO:
			info := msg.(*pb.PayloadInfo)
			payloadStates, err := service.ProcessPayload(tx, info, accountMap)
			if err != nil {
				return nil, err
			}
			states = append(states, payloadStates...)

		case core.CORE_CONTRACT_INFO:
			info := msg.(*pb.ContractInfo)
			if len(info.Code) > 0 {
				rootHash, codeHash, err := ss.CreateContract(tx.Destination, info.Code)
				if err != nil {
					return nil, err
				}
				accountInfo, ok := accountMap[tx.Destination.String()]
				if ok {
					accountInfo.Code = &block.DataInfo{
						GroupHash: rootHash,
						DataHash:  codeHash,
					}
				}
			} else {
				var account libcore.Address
				if len(info.Account) > 0 {
					_, account, _ = as.NewAccountFromBytes(info.Account[:])
				} else {
					account = tx.Account
				}
				rootHash, codeHash, result, err := ss.RunContract(tx.Gas, tx.Destination, account, info.Method, info.Params)
				if err != nil {
					return nil, err
				}
				accountInfo, ok := accountMap[account.String()]
				if ok {
					accountInfo.Data = &block.DataInfo{
						GroupHash: rootHash,
						DataHash:  codeHash,
						Content:   result,
					}
				}
			}

		case core.CORE_PAGE_INFO:
			info := msg.(*pb.PageInfo)
			if len(info.Data) > 0 {
				rootHash, pageHash, err := ss.CreatePage(tx.Destination, info.Data)
				if err != nil {
					return nil, err
				}
				accountInfo, ok := accountMap[tx.Destination.String()]
				if ok {
					accountInfo.Page = &block.DataInfo{
						GroupHash: rootHash,
						DataHash:  pageHash,
					}
				}
			}

		case core.CORE_META_INFO:
		case core.CORE_TOKEN_INFO:
		case core.CORE_DATA_INFO:
		default:
			return nil, core.ErrorOfUnknown("format", "info")
		}
	}

	return states, nil
}
