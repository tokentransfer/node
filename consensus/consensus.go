package consensus

import (
	"fmt"
	"time"

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

	lastData *block.DataInfo
	data     *block.DataInfo

	lastGas *util.Value
	gas     *util.Value
}

type ConsensusService struct {
	n *Node
}

func (service *ConsensusService) GetAccountInfo(rootAccount libcore.Address, account libcore.Address) (*block.AccountState, error) {
	entry, err := service.n.GetEntry(rootAccount)
	if err != nil {
		return nil, err
	}
	ms := entry.merkle

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

func (service *ConsensusService) VerifyTransaction(rootAccount libcore.Address, t libblock.Transaction) (bool, error) {
	entry, err := service.n.GetEntry(rootAccount)
	if err != nil {
		return false, err
	}
	cs := service.n.cryptoService
	ss := entry.storage
	ms := entry.merkle

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

	fromAccount := tx.Account
	sequence := libstore.GetSequence(ms, fromAccount) + 1
	if tx.Sequence != sequence {
		return false, util.ErrorOfInvalid("sequence", fmt.Sprintf("%d != %d", tx.Sequence, sequence))
	}

	info, err := service.GetAccountInfo(rootAccount, fromAccount)
	if err != nil || info == nil {
		return false, util.ErrorOfNotFound("account", fromAccount.String())
	}
	if rootAccount != nil {
		if tx.Gas < 10 {
			return false, util.ErrorOf("insufficient", "gas", fmt.Sprintf("%d < 10", tx.Gas))
		}
		gasValue, err := util.NewValue(fmt.Sprintf("%d", tx.Gas))
		if err != nil {
			return false, err
		}
		fromValue, err := ss.GetGas(fromAccount)
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
				if !(info.Account != nil && len(info.Account.Data) > 0 && len(info.Account.Code) > 0 && info.Data != nil && (len(info.Data.Content) > 0 || len(info.Data.Hash) > 0)) {
					return false, util.ErrorOfInvalid("format", "user info")
				}

			case core.CORE_META_INFO:
				info := msg.(*pb.MetaInfo)
				if len(info.Symbol) == 0 {
					return false, util.ErrorOfEmpty("symbol", "meta info")
				}
				for i := 0; i < len(info.Items); i++ {
					item := info.Items[i]
					if item == nil {
						return false, util.ErrorOfEmpty("meta item", fmt.Sprintf("%d", i))
					}
					_, err := core.GetDataTypeByName(item.Type)
					if err != nil {
						return false, err
					}
				}
				_, a, _, _ := ss.ReadMeta(libcore.Symbol(info.Symbol))
				if a != nil && a.String() != fromAccount.String() {
					return false, util.ErrorOfUnmatched("account", "meta info", a.String(), fromAccount.String())
				}

			case core.CORE_TOKEN_INFO:
				info := msg.(*pb.TokenInfo)
				if len(info.Symbol) == 0 {
					return false, util.ErrorOfEmpty("symbol", "token info")
				}
				index, a, meta, err := ss.ReadMeta(libcore.Symbol(info.Symbol))
				if err != nil {
					return false, err
				}
				if a != nil && a.String() != fromAccount.String() {
					return false, util.ErrorOfUnmatched("account", "token info", a.String(), fromAccount.String())
				}
				if !(meta.Total < 0 || (info.Index == uint64(index+1) && info.Index < uint64(meta.Total))) {
					return false, util.ErrorOfInvalid("token index", fmt.Sprintf("%d, %d", info.Index, meta.Total))
				}
				err = service.verifyToken(meta, info)
				if err != nil {
					return false, err
				}

			case core.CORE_DATA_INFO:
			case core.CORE_PEER_INFO:
			default:
				return false, util.ErrorOfInvalid("format", "info")
			}
		}
	}

	return true, nil
}

func (service *ConsensusService) verifyToken(meta *pb.MetaInfo, token *pb.TokenInfo) error {
	ml := len(meta.Items)
	tl := len(token.Items)
	if ml != tl {
		return util.ErrorOfUnmatched("token item", "length", ml, tl)
	}
	for i := 0; i < ml; i++ {
		metaItem := meta.Items[i]
		tokenItem := token.Items[i]
		if tokenItem == nil {
			return util.ErrorOfEmpty("token item", fmt.Sprintf("%d", i))
		}
		if tokenItem.Name != metaItem.Name {
			return util.ErrorOfUnmatched("token item", "name", tokenItem.Name, metaItem.Name)
		}
		t, err := core.GetDataTypeByName(metaItem.Type)
		if err != nil {
			return err
		}
		required := getOptions(metaItem.Options)
		if required {
			_, err = t.FromString(tokenItem.Value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func getOptions(options []string) bool {
	required := true
	for _, o := range options {
		if o == "optional" {
			required = false
		}
	}
	return required
}

func (service *ConsensusService) getAccountInfo(rootAccount libcore.Address, account libcore.Address, sequence uint64) (*block.AccountState, *block.AccountState, error) {
	lastInfo, _ := service.GetAccountInfo(rootAccount, account)
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

func (service *ConsensusService) getAccountEntry(rootAccount libcore.Address, account libcore.Address, accountMap map[string]*accountEntry) (bool, *accountEntry, error) {
	entry, err := service.n.GetEntry(rootAccount)
	if err != nil {
		return false, nil, err
	}
	ss := entry.storage
	ms := entry.merkle

	e, ok := accountMap[account.String()]
	if !ok {
		sequence := libstore.GetSequence(ms, account)
		lastInfo, info, err := service.getAccountInfo(rootAccount, account, sequence)
		if err != nil {
			return false, nil, err
		}
		lastGas, _ := ss.GetGas(account)
		e = &accountEntry{
			lastInfo: lastInfo,
			info:     info,

			lastGas: lastGas,
		}
		accountMap[account.String()] = e
	}
	return ok, e, nil
}

func (service *ConsensusService) ProcessTransaction(rootAccount libcore.Address, t libblock.Transaction) (libblock.TransactionWithData, error) {
	config := service.n.config

	tx, ok := t.(*block.Transaction)
	if !ok {
		return nil, util.ErrorOfInvalid("transaction", t.GetHash().String())
	}

	accountMap := make(map[string]*accountEntry)
	accounts := make([]string, 0)

	fromAccount := tx.GetAccount()
	ok, fromEntry, err := service.getAccountEntry(rootAccount, fromAccount, accountMap)
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

	gasAccount := config.GetGasAccount()
	existGas, gasEntry, err := service.getAccountEntry(rootAccount, gasAccount, accountMap)
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
		wasmCost, payloadStates, err := service.ProcessPayload(rootAccount, tx.Gas-dataCost, tx, tx.Payload, accountMap)
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
	gasAmount, err := util.NewValue(fmt.Sprintf("%d", gas))
	if err != nil {
		return nil, err
	}
	if rootAccount != nil {
		if tx.Gas < gas {
			return nil, util.ErrorOf("insufficient", "gas", fmt.Sprintf("%d < %d", tx.Gas, gas))
		}
		fromEntry.gas, err = service.removeBalance(fromEntry.lastGas, *gasAmount)
		if err != nil {
			return nil, err
		}
		gasEntry.gas, err = service.addBalance(gasEntry.lastGas, *gasAmount)
		if err != nil {
			return nil, err
		}
	} else {
		fromEntry.gas, err = service.addBalance(fromEntry.lastGas, *gasAmount)
		if err != nil {
			return nil, err
		}
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

	dataMap := make(map[string]*block.DataInfo)
	states := make([]libblock.State, 0)
	datas := make([]*block.DataInfo, 0)
	for _, a := range stateList {
		entry := accountMap[a]
		if entry.lastInfo == nil && entry.info == nil {
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
		if entry.data != nil {
			info := entry.data
			if len(info.Hash) == 0 {
				return nil, util.ErrorOfEmpty("data hash", "account entry")
			}
			_, ok := dataMap[info.Hash.String()]
			if !ok {
				dataMap[info.Hash.String()] = info
				datas = append(datas, info)
			}
		}
	}

	r := &block.Receipt{
		TransactionResult: 0,
		States:            states,
		Datas:             datas,
	}

	return &block.TransactionWithData{
		Transaction: t,
		Receipt:     r,
		Date:        int64(time.Now().UnixNano() / 1e9),
	}, nil
}

func (service *ConsensusService) ProcessPayload(rootAccount libcore.Address, remainCost uint64, tx *block.Transaction, info *block.PayloadInfo, accountMap map[string]*accountEntry) (uint64, []string, error) {
	cost := uint64(0)

	entry, err := service.n.GetEntry(rootAccount)
	if err != nil {
		return cost, nil, err
	}
	as := service.n.accountService
	ss := entry.storage

	accounts := make([]string, 0)

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
			usedCost, payloadStates, err := service.ProcessPayload(rootAccount, remainCost-cost, tx, payloadInfo, accountMap)
			if err != nil {
				return 0, nil, err
			}
			accounts = append(accounts, payloadStates...)
			cost += usedCost

		case core.CORE_CONTRACT_INFO:
			info := msg.(*pb.ContractInfo)

			dataAccount := tx.Account
			codeAccount := tx.Account
			if info.Account != nil {
				if len(info.Account.Data) > 0 {
					_, a, err := as.NewAccountFromBytes(info.Account.Data)
					if err != nil {
						return 0, nil, err
					}
					dataAccount = a
				}
				if len(info.Account.Code) > 0 {
					_, a, err := as.NewAccountFromBytes(info.Account.Code)
					if err != nil {
						return 0, nil, err
					}
					codeAccount = a
				}
			}

			inputs, err := getAccounts(as, info.Inputs)
			if err != nil {
				return 0, nil, err
			}
			outputs, err := getAccounts(as, info.Outputs)
			if err != nil {
				return 0, nil, err
			}
			usedCost, retAccount, _, retHash, retContent, err := ss.RunContract(remainCost, tx.Account, dataAccount, codeAccount, info.Method, info.Params, inputs, outputs)
			if err != nil {
				return cost, nil, err
			}
			ok, retEntry, err := service.getAccountEntry(rootAccount, retAccount, accountMap)
			if err != nil {
				return cost, nil, err
			}
			if !ok {
				accounts = append(accounts, retAccount.String())
			}
			retEntry.info.Memory = &block.DataInfo{
				Hash:    retHash,
				Content: retContent,
			}
			cost += uint64(usedCost)

		case core.CORE_PAGE_INFO:
			info := msg.(*pb.PageInfo)
			_, pageHash, err := ss.WritePage(info.Name, tx.Account, info.Data)
			if err != nil {
				return cost, nil, err
			}
			accountEntry, ok := accountMap[tx.Account.String()]
			if !ok {
				return cost, nil, util.ErrorOfNotFound("account entry", tx.Account.String())
			}
			accountEntry.info.Page = &block.DataInfo{
				Hash: pageHash,
			}

		case core.CORE_CODE_INFO:
			info := msg.(*pb.CodeInfo)
			_, codeHash, err := ss.CreateContract(tx.Account, info.Code, info.Abi)
			if err != nil {
				return cost, nil, err
			}
			accountEntry, ok := accountMap[tx.Account.String()]
			if !ok {
				return cost, nil, util.ErrorOfNotFound("account entry", tx.Account.String())
			}
			accountEntry.info.Code = &block.DataInfo{
				Hash: codeHash,
			}

		case core.CORE_USER_INFO:
			info := msg.(*pb.UserInfo)

			dataAccount := tx.Account
			codeAccount := tx.Account
			if info.Account != nil {
				if len(info.Account.Data) > 0 {
					_, a, err := as.NewAccountFromBytes(info.Account.Data)
					if err != nil {
						return 0, nil, err
					}
					dataAccount = a
				}
				if len(info.Account.Code) > 0 {
					_, a, err := as.NewAccountFromBytes(info.Account.Code)
					if err != nil {
						return 0, nil, err
					}
					codeAccount = a
				}
			}

			userHash, userAccount, err := ss.WriteUser(tx.Account, dataAccount, codeAccount, info)
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
			info := msg.(*pb.MetaInfo)
			dataInfo, err := ss.WriteMeta(tx.Account, info)
			if err != nil {
				return cost, nil, err
			}
			accountEntry, ok := accountMap[tx.Account.String()]
			if !ok {
				return cost, nil, util.ErrorOfNotFound("account entry", tx.Account.String())
			}
			accountEntry.info.Token = dataInfo

		case core.CORE_TOKEN_INFO:
			info := msg.(*pb.TokenInfo)

			dataInfo, err := ss.WriteToken(tx.Account, info)
			if err != nil {
				return cost, nil, err
			}
			accountEntry, ok := accountMap[tx.Account.String()]
			if !ok {
				return cost, nil, util.ErrorOfNotFound("account entry", tx.Account.String())
			}
			accountEntry.info.Token = dataInfo

			stateType := libblock.GetStateTypeByName(info.Symbol)
			switch stateType {
			case block.ACCOUNT_STATE:
				_, createAccount, _ := as.NewAccountFromAddress(info.Items[0].Value)
				if createAccount != nil {
					ok, _, err = service.getAccountEntry(rootAccount, createAccount, accountMap)
					if err != nil {
						return cost, nil, err
					}
					if !ok {
						accounts = append(accounts, createAccount.String())
					}
				}
			}

		case core.CORE_DATA_INFO:
		case core.CORE_GROUP_INFO:
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
	cs := service.n.cryptoService

	h, _, err := cs.Raw(b, libcrypto.RawBinary)
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
	cs := service.n.cryptoService

	_, _, err := cs.Raw(txWithData, libcrypto.RawBinary)
	if err != nil {
		return nil, err
	}

	h, _, err := cs.Raw(txWithData.GetTransaction(), libcrypto.RawBinary)
	if err != nil {
		return nil, err
	}

	_, _, err = cs.Raw(txWithData.GetReceipt(), libcrypto.RawBinary)
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
	cs := service.n.cryptoService

	h, _, err := cs.Raw(r, libcrypto.RawBinary)
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
	cs := service.n.cryptoService

	h, _, err := cs.Raw(s, libcrypto.RawBinary)
	if err != nil {
		return nil, err
	}
	return h, nil
}
