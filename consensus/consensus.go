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

type accountEntry struct {
	lastInfo *block.AccountState
	info     *block.AccountState

	lastGas *util.Value
	gas     *util.Value

	lastLocalGas *util.Value
	localGas     *util.Value
}

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

func (service *ConsensusService) GenerateBlock(rootAccount libcore.Address, list []libblock.TransactionWithData) (libblock.Block, error) {
	cs := service.CryptoService
	ms := service.MerkleService
	ss := service.StorageService

	v := service.ValidatedBlock

	var b *block.Block
	if service.ValidatedBlock == nil { //genesis
		if len(list) > 0 {
			return nil, util.ErrorOfInvalid("transactions", "genesis block")
		}
		var states []libblock.State
		if rootAccount != nil {
			v, err := util.NewValue("100000000000000")
			if err != nil {
				return nil, err
			}
			err = ss.UpdateGas(rootAccount, rootAccount, *v)
			if err != nil {
				return nil, err
			}
			gasAccount := service.Config.GetGasAccount()
			err = ss.UpdateGas(nil, gasAccount, v.ZeroClone())
			if err != nil {
				return nil, err
			}
			states = []libblock.State{
				&block.AccountState{
					State: block.State{
						StateType:  block.ACCOUNT_STATE,
						Account:    rootAccount,
						Sequence:   uint64(0),
						BlockIndex: uint64(0),
					},
				},
			}
		} else {
			states = make([]libblock.State, 0)
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

		err := ms.Cancel()
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

	fromAccount := tx.From
	sequence := libstore.GetSequence(ms, fromAccount) + 1
	if tx.Sequence != sequence {
		return false, util.ErrorOfInvalid("sequence", fmt.Sprintf("%d != %d", tx.Sequence, sequence))
	}

	info, err := service.GetAccountInfo(fromAccount)
	if err != nil || info == nil {
		return false, util.ErrorOfNotFound("account", fromAccount.String())
	}

	if tx.Gas < 10 {
		return false, util.ErrorOf("insufficient", "gas", fmt.Sprintf("%d < 10", tx.Gas))
	}

	account := tx.Account
	gasValue, err := util.NewValue(fmt.Sprintf("%d", tx.Gas))
	if err != nil {
		return false, err
	}
	fromValue, err := service.StorageService.GetGas(account, fromAccount)
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
				info := msg.(*pb.PageInfo)
				if !(len(info.Data) > 0) {
					return false, util.ErrorOfInvalid("format", "page info")
				}
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

func (service *ConsensusService) getAccountGas(account libcore.Address) (*util.Value, *util.Value, error) {
	value, err := service.StorageService.GetGas(nil, account)
	if err != nil {
		return nil, nil, err
	}
	localValue, err := service.StorageService.GetGas(account, account)
	if err != nil {
		return nil, nil, err
	}
	return value, localValue, nil
}

func (service *ConsensusService) addBalance(oldValue *util.Value, value util.Value) (*util.Value, error) {
	isNegative, err := value.IsNegative()
	if err != nil {
		return nil, err
	}
	if isNegative {
		return nil, util.ErrorOfInvalid("gas", "nagative")
	}
	if oldValue == nil {
		zero := value.ZeroClone()
		oldValue = &zero
	}
	newValue, err := oldValue.Add(value)
	if err != nil {
		return nil, err
	}
	return &newValue, nil
}

func (service *ConsensusService) removeBalance(oldValue *util.Value, value util.Value) (*util.Value, error) {
	isNegative, err := value.IsNegative()
	if err != nil {
		return nil, err
	}
	if isNegative {
		return nil, util.ErrorOfInvalid("gas", "nagative")
	}
	if oldValue == nil {
		zero := value.ZeroClone()
		oldValue = &zero
	}
	newValue, err := oldValue.Subtract(value)
	if err != nil {
		return nil, err
	}
	isNegative, err = newValue.IsNegative()
	if err != nil {
		return nil, err
	}
	if isNegative {
		return nil, util.ErrorOf("insuffient gas", "nagative", "balance")
	}
	return &newValue, nil
}

func (service *ConsensusService) getAccountEntry(account libcore.Address, accountMap map[string]*accountEntry) (bool, *accountEntry, error) {
	ms := service.MerkleService

	entry, ok := accountMap[account.String()]
	if !ok {
		sequence := libstore.GetSequence(ms, account)
		lastInfo, info, err := service.getAccountInfo(account, sequence)
		if err != nil {
			return false, nil, err
		}
		lastGas, lastLocalGas, err := service.getAccountGas(account)
		if err != nil {
			return false, nil, err
		}
		entry = &accountEntry{
			lastInfo: lastInfo,
			info:     info,

			lastGas:      lastGas,
			lastLocalGas: lastLocalGas,
		}
		accountMap[account.String()] = entry
	}
	return ok, entry, nil
}

func (service *ConsensusService) ProcessTransaction(t libblock.Transaction) (libblock.TransactionWithData, error) {
	config := service.Config

	tx, ok := t.(*block.Transaction)
	if !ok {
		return nil, util.ErrorOfInvalid("transaction", t.GetHash().String())
	}

	accountMap := make(map[string]*accountEntry)
	accounts := make([]string, 0)

	fromAccount := tx.GetAccount()
	ok, fromEntry, err := service.getAccountEntry(fromAccount, accountMap)
	if err != nil {
		return nil, err
	}
	if !ok {
		accounts = append(accounts, fromAccount.String())
	}
	fromSequence := fromEntry.info.Sequence + 1 // sequence of from from account, +1
	fromEntry.info.Sequence = fromSequence

	if tx.Sequence != fromSequence {
		return nil, util.ErrorOf("error", "sequence", fmt.Sprintf("%d != %d", tx.Sequence, fromSequence))
	}

	toAccount := tx.To
	existDest, _, err := service.getAccountEntry(toAccount, accountMap)
	if err != nil {
		return nil, err
	}
	if !existDest {
		accounts = append(accounts, toAccount.String())
	}

	gasAccount := config.GetGasAccount()
	existGas, gasEntry, err := service.getAccountEntry(gasAccount, accountMap)
	if err != nil {
		return nil, err
	}
	if !existGas {
		accounts = append(accounts, gasAccount.String())
	}

	gas := uint64(0) //gasPrice gasLimit
	if tx.Payload != nil && len(tx.Payload.Infos) > 0 {
		dataCost := uint64(0)
		for _, info := range tx.Payload.Infos {
			dataCost += uint64(len(info.Content)) // 1:1
		}
		wasmCost, payloadStates, err := service.ProcessPayload(tx.Gas-dataCost, tx, tx.Payload, accountMap)
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
	fromEntry.localGas, err = service.removeBalance(fromEntry.lastLocalGas, *gasAmount)
	if err != nil {
		return nil, err
	}
	gasEntry.localGas, err = service.addBalance(gasEntry.lastLocalGas, *gasAmount)
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
		entry := accountMap[a]
		if entry.lastInfo == nil || entry.info == nil {
			return nil, util.ErrorOfInvalid("account info", a)
		} else {
			if entry.lastInfo == nil {
				states = append(states, entry.info)
			} else {
				lastHash, err := service.HashState(entry.lastInfo)
				if err != nil {
					return nil, err
				}
				hash, err := service.HashState(entry.info)
				if err != nil {
					return nil, err
				}
				if lastHash.String() != hash.String() {
					entry.info.Version = entry.info.Version + 1
					states = append(states, entry.info)
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

func (service *ConsensusService) ProcessPayload(remainCost uint64, tx *block.Transaction, info *block.PayloadInfo, accountMap map[string]*accountEntry) (uint64, []string, error) {
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
			usedCost, payloadStates, err := service.ProcessPayload(remainCost-cost, tx, payloadInfo, accountMap)
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
			usedCost, retAccount, _, retHash, retContent, err := ss.RunContract(cs, remainCost, tx.Account, dataAccount, tx.To, info.Method, info.Params, inputs, outputs)
			if err != nil {
				return cost, nil, err
			}
			ok, retEntry, err := service.getAccountEntry(retAccount, accountMap)
			if err != nil {
				return cost, nil, err
			}
			if !ok {
				accounts = append(accounts, retAccount.String())
			}
			retEntry.info.Data = &block.DataInfo{
				Hash:    retHash,
				Content: retContent,
			}
			cost += uint64(usedCost)

		case core.CORE_PAGE_INFO:
			info := msg.(*pb.PageInfo)
			_, pageHash, err := ss.WritePage(info.Name, tx.To, info.Data)
			if err != nil {
				return cost, nil, err
			}
			accountEntry, ok := accountMap[tx.To.String()]
			if !ok {
				return cost, nil, util.ErrorOfNotFound("account entry", tx.To.String())
			}
			accountEntry.info.Page = &block.DataInfo{
				Hash: pageHash,
			}

		case core.CORE_CODE_INFO:
			info := msg.(*pb.CodeInfo)
			_, codeHash, err := ss.CreateContract(tx.To, info.Code, info.Abi)
			if err != nil {
				return cost, nil, err
			}
			accountEntry, ok := accountMap[tx.To.String()]
			if !ok {
				return cost, nil, util.ErrorOfNotFound("account entry", tx.To.String())
			}
			accountEntry.info.Code = &block.DataInfo{
				Hash: codeHash,
			}

		case core.CORE_USER_INFO:
			info := msg.(*pb.UserInfo)
			_, account, err := as.NewAccountFromBytes(info.Account)
			if err != nil {
				return cost, nil, err
			}
			userHash, userAccount, err := ss.WriteUser(tx.Account, tx.To, account, info)
			if err != nil {
				return cost, nil, err
			}
			accountEntry, ok := accountMap[userAccount.String()]
			if !ok {
				return cost, nil, util.ErrorOfNotFound("account entry", userAccount.String())
			}
			accountEntry.info.User = &block.DataInfo{
				Hash: userHash,
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
