package consensus

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"os"
	"path"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/spkg/zipfs"

	"chainmaker.org/chainmaker/protocol/v2"
	"github.com/caivega/glog"
	"github.com/tokentransfer/node/account"
	"github.com/tokentransfer/node/block"
	"github.com/tokentransfer/node/config"
	"github.com/tokentransfer/node/core"
	"github.com/tokentransfer/node/core/pb"
	"github.com/tokentransfer/node/crypto"
	"github.com/tokentransfer/node/merkle"
	"github.com/tokentransfer/node/p2p"
	"github.com/tokentransfer/node/storage"
	"github.com/tokentransfer/node/util"

	libaccount "github.com/tokentransfer/interfaces/account"
	libblock "github.com/tokentransfer/interfaces/block"
	libcore "github.com/tokentransfer/interfaces/core"
	libcrypto "github.com/tokentransfer/interfaces/crypto"
	libstore "github.com/tokentransfer/interfaces/store"
)

const (
	TOPIC_PEER_DISCOVERY = "peer_discovery"
	TOPIC_PEER_DATA      = "peer_data"
	TOPIC_PEER_MESSAGE   = "peer_message"
)

type Node struct {
	messageId uint64

	cryptoService    libcrypto.CryptoService
	accountService   libaccount.AccountService
	consensusService *ConsensusService

	entryMap    map[string]*Entry
	entryLocker *sync.RWMutex

	storageService    *storage.StorageService
	transactionLocker *sync.Mutex

	net protocol.Net

	outs chan *pb.Message
	ins  chan *pb.Message

	config *config.Config
	key    libaccount.Key

	self       *Peer
	peers      map[uint64]*Peer
	peerLocker *sync.RWMutex
}

func NewNode() *Node {
	ready := &sync.WaitGroup{}
	ready.Add(1)

	return &Node{
		messageId: uint64(0),

		outs: make(chan *pb.Message, 8),
		ins:  make(chan *pb.Message, 8),

		peers:      map[uint64]*Peer{},
		peerLocker: &sync.RWMutex{},

		transactionLocker: &sync.Mutex{},

		entryMap:    make(map[string]*Entry),
		entryLocker: &sync.RWMutex{},

		accountService: account.NewAccountService(),
		cryptoService:  &crypto.CryptoService{},
	}
}

func (n *Node) Init(c libcore.Config) error {
	n.config = c.(*config.Config)
	util.Init(n.config.GetMode())

	_, key, err := n.accountService.NewKeyFromSecret(n.config.GetSecret())
	if err != nil {
		return err
	}
	n.key = key

	pubKey, err := n.key.GetPublic()
	if err != nil {
		return err
	}
	n.self = &Peer{
		Key:     pubKey,
		peermap: make(map[string]*peerEntry),
	}

	storageService, err := storage.NewCategoryService(n.config, nil)
	if err != nil {
		return err
	}

	consensusService := &ConsensusService{
		n: n,
	}
	n.consensusService = consensusService
	n.storageService = storageService

	return nil
}

func (n *Node) getContractData(rootAccount libcore.Address, txm map[string]interface{}) (int64, interface{}, error) {
	entry, err := n.GetEntry(rootAccount)
	if err != nil {
		return 0, nil, err
	}
	as := n.accountService
	ss := entry.storage

	from := util.ToString(&txm, "from")
	to := util.ToString(&txm, "to")

	_, fromAccount, err := as.NewAccountFromAddress(from)
	if err != nil {
		return 0, nil, err
	}

	_, toAccount, err := as.NewAccountFromAddress(to)
	if err != nil {
		return 0, nil, err
	}

	format := util.ToString(&txm, "format")
	usedCost, r, err := ss.GetContractData(fromAccount, toAccount, format)
	if err != nil {
		return 0, nil, err
	}
	return usedCost, r, nil
}

func (n *Node) callContract(rootAccount libcore.Address, txm map[string]interface{}) (int64, interface{}, error) {
	entry, err := n.GetEntry(rootAccount)
	if err != nil {
		return 0, nil, err
	}
	as := n.accountService
	ss := entry.storage

	from := util.ToString(&txm, "from")
	to := util.ToString(&txm, "to")

	_, fromAccount, err := as.NewAccountFromAddress(from)
	if err != nil {
		return 0, nil, err
	}

	_, toAccount, err := as.NewAccountFromAddress(to)
	if err != nil {
		return 0, nil, err
	}

	method := util.ToString(&txm, "method")
	params, err := n.getParams(txm)
	if err != nil {
		return 0, nil, err
	}

	usedCost, r, err := ss.CallContract(fromAccount, toAccount, method, params)
	if err != nil {
		return 0, nil, err
	}
	return usedCost, r, nil
}

func (n *Node) signTransaction(rootAccount libcore.Address, txm map[string]interface{}) (string, *block.Transaction, error) {
	as := n.accountService

	accountString := util.ToString(&txm, "account")
	_, account, err := as.NewAccountFromAddress(accountString)
	if err != nil {
		return "", nil, err
	}
	secret := util.ToString(&txm, "secret")
	gas := util.ToUint64(&txm, "gas")

	_, fromKey, err := as.NewKeyFromSecret(secret)
	if err != nil {
		return "", nil, err
	}
	fromAccount, err := fromKey.GetAddress()
	if err != nil {
		return "", nil, err
	}
	fromAddress := fromAccount.String()
	if fromAddress != account.String() {
		return "", nil, util.ErrorOfUnmatched("account", "in secret", fromAddress, account.String())
	}

	seq := n.getNextSequence(rootAccount, account)

	tx := &block.Transaction{
		TransactionType: block.TRANSACTION,

		Account:  account,
		Sequence: seq,
		Gas:      gas,
	}

	payloadInfo := &block.PayloadInfo{}
	if util.Has(&txm, "payload") {
		pxm := util.ToMap(&txm, "payload")
		if util.Has(&pxm, "contract") {
			contractInfo := &pb.ContractInfo{}

			cm := util.ToMap(&pxm, "contract")
			if util.Has(&cm, "inputs") {
				if inputString := util.ToString(&cm, "inputs"); len(inputString) > 0 {
					inputs := strings.Split(inputString, ",")
					list := make([][]byte, 0)
					for _, accountString := range inputs {
						_, account, err := as.NewAccountFromAddress(accountString)
						if err != nil {
							return "", nil, err
						}
						accountData, err := account.MarshalBinary()
						if err != nil {
							return "", nil, err
						}
						list = append(list, accountData)
					}
					contractInfo.Inputs = list
				}
				if inputs := util.ToArray(&cm, "inputs"); inputs != nil {
					list := make([][]byte, 0)
					for _, input := range inputs {
						accountString, ok := input.(string)
						if !ok {
							return "", nil, util.ErrorOfInvalid("input", "account")
						}
						_, account, err := as.NewAccountFromAddress(accountString)
						if err != nil {
							return "", nil, err
						}
						accountData, err := account.MarshalBinary()
						if err != nil {
							return "", nil, err
						}
						list = append(list, accountData)
					}
					contractInfo.Inputs = list
				}
			}
			if util.Has(&cm, "outputs") {
				if outputString := util.ToString(&cm, "outputs"); len(outputString) > 0 {
					outputs := strings.Split(outputString, ",")
					list := make([][]byte, 0)
					for _, accountString := range outputs {
						_, account, err := as.NewAccountFromAddress(accountString)
						if err != nil {
							return "", nil, err
						}
						accountData, err := account.MarshalBinary()
						if err != nil {
							return "", nil, err
						}
						list = append(list, accountData)
					}
					contractInfo.Outputs = list
				}
				if outputs := util.ToArray(&cm, "outputs"); outputs != nil {
					list := make([][]byte, 0)
					for _, output := range outputs {
						accountString, ok := output.(string)
						if !ok {
							return "", nil, util.ErrorOfInvalid("output", "account")
						}
						_, account, err := as.NewAccountFromAddress(accountString)
						if err != nil {
							return "", nil, err
						}
						accountData, err := account.MarshalBinary()
						if err != nil {
							return "", nil, err
						}
						list = append(list, accountData)
					}
					contractInfo.Outputs = list
				}
			}
			if util.Has(&cm, "method") {
				contractInfo.Method = util.ToString(&cm, "method")
			}
			if util.Has(&cm, "params") {
				params, err := n.getParams(cm)
				if err != nil {
					return "", nil, err
				}
				contractInfo.Params = params
			}
			contractData, err := core.Marshal(contractInfo)
			if err != nil {
				return "", nil, err
			}
			dataHash, err := n.cryptoService.Hash(contractData)
			if err != nil {
				return "", nil, err
			}
			payloadInfo.Infos = append(payloadInfo.Infos, &block.DataInfo{
				Hash:    dataHash,
				Content: contractData,
			})
		}
		if util.Has(&pxm, "code") {
			codeInfo := &pb.CodeInfo{}

			cm := util.ToMap(&pxm, "code")
			if util.Has(&cm, "data") {
				codeString := util.ToString(&cm, "data")
				codeData, err := hex.DecodeString(codeString)
				if err != nil {
					return "", nil, err
				}
				codeInfo.Code = codeData
			} else if util.Has(&cm, "file") {
				filePath := util.ToString(&cm, "file")
				f, err := os.Open(filePath)
				if err != nil {
					return "", nil, err
				}
				defer f.Close()
				fileData, err := io.ReadAll(f)
				if err != nil {
					return "", nil, err
				}
				codeInfo.Code = fileData
			}
			if util.Has(&cm, "abi") {
				abiString := util.ToString(&cm, "abi")
				abiData, err := hex.DecodeString(abiString)
				if err != nil {
					return "", nil, err
				}
				codeInfo.Abi = abiData
			} else if util.Has(&cm, "json") {
				jsonPath := util.ToString(&cm, "json")
				f, err := os.Open(jsonPath)
				if err != nil {
					return "", nil, err
				}
				defer f.Close()
				jsonData, err := io.ReadAll(f)
				if err != nil {
					return "", nil, err
				}
				codeInfo.Abi = jsonData
			}

			codeData, err := core.Marshal(codeInfo)
			if err != nil {
				return "", nil, err
			}
			dataHash, err := n.cryptoService.Hash(codeData)
			if err != nil {
				return "", nil, err
			}
			payloadInfo.Infos = append(payloadInfo.Infos, &block.DataInfo{
				Hash:    dataHash,
				Content: codeData,
			})
		}
		if util.Has(&pxm, "page") {
			pageInfo := &pb.PageInfo{}

			pm := util.ToMap(&pxm, "page")
			if util.Has(&pm, "name") {
				name := util.ToString(&pm, "name")
				pageInfo.Name = name
			}
			if util.Has(&pm, "data") {
				pageString := util.ToString(&pm, "data")
				pageData, err := hex.DecodeString(pageString)
				if err != nil {
					return "", nil, err
				}
				pageInfo.Data = pageData
			} else if util.Has(&pm, "file") {
				filePath := util.ToString(&pm, "file")
				f, err := os.Open(filePath)
				if err != nil {
					return "", nil, err
				}
				defer f.Close()
				fileData, err := io.ReadAll(f)
				if err != nil {
					return "", nil, err
				}
				pageInfo.Data = fileData
			}

			pageData, err := core.Marshal(pageInfo)
			if err != nil {
				return "", nil, err
			}
			dataHash, err := n.cryptoService.Hash(pageData)
			if err != nil {
				return "", nil, err
			}
			payloadInfo.Infos = append(payloadInfo.Infos, &block.DataInfo{
				Hash:    dataHash,
				Content: pageData,
			})
		}
		if util.Has(&pxm, "meta") {
			metaInfo := &pb.MetaInfo{}

			pm := util.ToMap(&pxm, "meta")
			metaInfo.Symbol = util.ToString(&pm, "symbol")
			metaInfo.Total = util.ToInt64(&pm, "total")
			metaInfo.Items, err = n.getMetaItems(pm)
			if err != nil {
				return "", nil, err
			}
			metaData, err := core.Marshal(metaInfo)
			if err != nil {
				return "", nil, err
			}
			dataHash, err := n.cryptoService.Hash(metaData)
			if err != nil {
				return "", nil, err
			}
			payloadInfo.Infos = append(payloadInfo.Infos, &block.DataInfo{
				Hash:    dataHash,
				Content: metaData,
			})
		}
		if util.Has(&pxm, "token") {
			tokenInfo := &pb.TokenInfo{}

			pm := util.ToMap(&pxm, "token")
			tokenInfo.Symbol = util.ToString(&pm, "symbol")
			tokenInfo.Index = util.ToUint64(&pm, "index")
			tokenInfo.Items, err = n.getTokenItems(pm)
			if err != nil {
				return "", nil, err
			}
			tokenData, err := core.Marshal(tokenInfo)
			if err != nil {
				return "", nil, err
			}
			dataHash, err := n.cryptoService.Hash(tokenData)
			if err != nil {
				return "", nil, err
			}
			payloadInfo.Infos = append(payloadInfo.Infos, &block.DataInfo{
				Hash:    dataHash,
				Content: tokenData,
			})
		}
	}
	if len(payloadInfo.Infos) > 0 {
		tx.Payload = payloadInfo
	}
	err = n.cryptoService.Sign(fromKey, tx)
	if err != nil {
		return "", nil, err
	}
	data, err := tx.MarshalBinary()
	if err != nil {
		return "", nil, err
	}
	err = n.verifyTransaction(rootAccount, tx)
	if err != nil {
		return "", nil, err
	}
	blob := libcore.Bytes(data).String()
	return blob, tx, nil
}

func (n *Node) getOptions(os string) []string {
	if len(os) == 0 {
		return nil
	}
	return strings.Split(os, ";")
}

func (n *Node) getMetaItems(cm map[string]interface{}) ([]*pb.MetaItem, error) {
	if itemString := util.ToString(&cm, "items"); len(itemString) > 0 {
		list := make([]*pb.MetaItem, 0)
		ps := strings.Split(itemString, ",")
		for i := 0; i < len(ps); i++ {
			items := strings.Split(ps[i], ":")
			if len(items) == 4 {
				name := items[0]
				t := items[1]
				options := n.getOptions(items[2])
				desc := items[3]

				_, err := core.GetDataTypeByName(t)
				if err != nil {
					return nil, err
				}

				list = append(list, &pb.MetaItem{
					Name:    name,
					Type:    t,
					Options: options,
					Desc:    desc,
				})
			} else {
				return nil, util.ErrorOfInvalid("meta items", itemString)
			}
		}
		return list, nil
	}
	if items := util.ToArray(&cm, "items"); items != nil {
		list := make([]*pb.MetaItem, 0)
		for i := 0; i < len(items); i++ {
			item := items[i]
			m, ok := item.(map[string]interface{})
			if !ok {
				return nil, util.ErrorOfInvalid("meta item", "not a map")
			}
			name := util.ToString(&m, "name")
			t := util.ToString(&m, "type")

			_, err := core.GetDataTypeByName(t)
			if err != nil {
				return nil, err
			}

			o := util.ToString(&m, "options")
			options := n.getOptions(o)

			desc := util.ToString(&m, "desc")

			list = append(list, &pb.MetaItem{
				Name:    name,
				Type:    t,
				Options: options,
				Desc:    desc,
			})
		}
		return list, nil
	}
	return nil, util.ErrorOfNotFound("meta items", "map")
}

func (n *Node) getTokenItems(cm map[string]interface{}) ([]*pb.TokenItem, error) {
	if itemString := util.ToString(&cm, "items"); len(itemString) > 0 {
		list := make([]*pb.TokenItem, 0)
		ps := strings.Split(itemString, ",")
		for i := 0; i < len(ps); i++ {
			items := strings.Split(ps[i], ":")
			if len(items) == 2 {
				name := items[0]
				value := items[1]

				list = append(list, &pb.TokenItem{
					Name:  name,
					Value: value,
				})
			} else {
				return nil, util.ErrorOfInvalid("token items", itemString)
			}
		}
		return list, nil
	}
	if items := util.ToArray(&cm, "items"); items != nil {
		list := make([]*pb.TokenItem, 0)
		for i := 0; i < len(items); i++ {
			item := items[i]
			m, ok := item.(map[string]interface{})
			if !ok {
				return nil, util.ErrorOfInvalid("token item", "not a map")
			}
			name := util.ToString(&m, "name")
			value := util.ToString(&m, "value")

			list = append(list, &pb.TokenItem{
				Name:  name,
				Value: value,
			})
		}
		return list, nil
	}
	return nil, util.ErrorOfNotFound("token items", "map")
}

func (n *Node) getParams(cm map[string]interface{}) ([][]byte, error) {
	if params := util.ToString(&cm, "params"); len(params) > 0 {
		list := make([][]byte, 0)
		ps := strings.Split(params, ",")
		for i := 0; i < len(ps); i++ {
			items := strings.Split(ps[i], ":")
			if len(items) == 2 {
				t, err := core.GetDataTypeByName(items[1])
				if err != nil {
					return nil, err
				}
				data, err := t.FromString(items[0])
				if err != nil {
					return nil, err
				}
				list = append(list, data)
			} else {
				return nil, util.ErrorOfInvalid("parameter", params)
			}
		}
		return list, nil
	}
	if params := util.ToArray(&cm, "params"); params != nil {
		list := make([][]byte, 0)
		for i := 0; i < len(params); i++ {
			item := params[i]
			data, err := core.MarshalData(item)
			if err != nil {
				return nil, err
			}
			list = append(list, data)
		}
		return list, nil
	}
	return nil, util.ErrorOfNotFound("params", "in contract")
}

func (n *Node) verifyTransaction(rootAccount libcore.Address, tx libblock.Transaction) error {
	_, _, err := n.cryptoService.Raw(tx, libcrypto.RawBinary)
	if err != nil {
		return err
	}
	_, err = n.consensusService.VerifyTransaction(rootAccount, tx)
	if err != nil {
		return err
	}
	return nil
}

func (n *Node) sendTransaction(rootAccount libcore.Address, tx libblock.Transaction) (libblock.TransactionWithData, error) {
	err := n.verifyTransaction(rootAccount, tx)
	if err != nil {
		return nil, err
	}
	hash := tx.GetHash()
	glog.Infoln("verify transaction", hash.String())

	txWithData, _, err := n.processTransaction(rootAccount, tx)
	if err != nil {
		return nil, err
	}
	return txWithData, nil
}

func (n *Node) processTransaction(rootAccount libcore.Address, tx libblock.Transaction) (libblock.TransactionWithData, libblock.Transaction, error) {
	n.transactionLocker.Lock()
	defer n.transactionLocker.Unlock()

	entry, err := n.GetEntry(rootAccount)
	if err != nil {
		return nil, nil, err
	}
	cs := n.cryptoService
	ss := entry.storage

	h, _, err := cs.Raw(tx, libcrypto.RawBinary)
	if err != nil {
		return nil, nil, err
	}
	ok, err := n.consensusService.VerifyTransaction(rootAccount, tx)
	if err != nil {
		return nil, nil, err
	}
	if !ok {
		return nil, nil, util.ErrorOfInvalid("verify", "transaction")
	}

	err = ss.CreateSandbox()
	if err != nil {
		return nil, nil, err
	}
	defer ss.CancelSandbox()
	txWithData, err := n.consensusService.ProcessTransaction(rootAccount, tx)
	if err != nil {
		return nil, nil, err
	}

	// util.PrintJSON("txWithData", txWithData)
	ok, err = n.addTransaction(rootAccount, txWithData)
	if err != nil {
		return nil, nil, err
	}
	if ok {
		glog.Infoln(">>> receive transaction", h.String())
	} else {
		glog.Infoln(">>> drop exists transaction", h.String())
	}

	return txWithData, tx, nil
}

func (n *Node) getNextSequence(rootAccount libcore.Address, address libcore.Address) uint64 {
	accountEntry, err := n.consensusService.GetAccountInfo(rootAccount, address)
	if err != nil {
		return uint64(1)
	}
	return accountEntry.Sequence + 1
}

func (n *Node) getRootAccount(item map[string]interface{}) (libcore.Address, error) {
	if item != nil && util.Has(&item, "root") {
		root := util.ToString(&item, "root")
		_, r, err := n.accountService.NewAccountFromAddress(root)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
	return nil, nil
}

func (n *Node) _call(params []interface{}, f func(rootAccount libcore.Address, item map[string]interface{}) (interface{}, error)) (interface{}, error) {
	l := len(params)
	if l > 0 {
		list := make([]interface{}, l)
		for i := 0; i < l; i++ {
			item, ok := params[i].(map[string]interface{})
			if !ok {
				return nil, util.ErrorOfInvalid("map", "parameter")
			}
			rootAccount, err := n.getRootAccount(item)
			if err != nil {
				return nil, err
			}
			result, err := f(rootAccount, item)
			if err != nil {
				return nil, err
			}
			list[i] = result
		}
		if l == 1 { // for one, just return the result instead of the array
			return list[0], nil
		}
		return list, nil
	}
	result, err := f(nil, nil)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (n *Node) Call(method string, params []interface{}) (interface{}, error) {
	switch method {
	case "blockNumber":
		result, err := n._call(params, func(rootAccount libcore.Address, item map[string]interface{}) (interface{}, error) {
			entry, err := n.GetEntry(rootAccount)
			if err != nil {
				return nil, err
			}
			return entry.GetBlockNumber(), nil
		})
		if err != nil {
			return nil, err
		}
		return result, nil

	case "createWallet":
		result, err := n._call(params, func(rootAccount libcore.Address, item map[string]interface{}) (interface{}, error) {
			as := n.accountService

			t := util.ToString(&item, "type")
			p := util.ToString(&item, "password")
			tt, k, err := as.GenerateFamilySeed(t + "." + p)
			if err != nil {
				return nil, err
			}
			a, err := k.GetAddress()
			if err != nil {
				return nil, err
			}
			address := a.String()
			pub, err := k.GetPublic()
			if err != nil {
				return nil, err
			}
			publicString, err := pub.MarshalText()
			if err != nil {
				return nil, err
			}
			priv, err := k.GetPrivate()
			if err != nil {
				return nil, err
			}
			privateString, err := priv.MarshalText()
			if err != nil {
				return nil, err
			}
			return &map[string]interface{}{
				"type":    tt.String(),
				"address": address,
				"private": string(privateString),
				"public":  string(publicString),
			}, nil
		})
		if err != nil {
			return nil, err
		}
		return result, nil

	case "getBalance":
		result, err := n._call(params, func(rootAccount libcore.Address, item map[string]interface{}) (interface{}, error) {
			entry, err := n.GetEntry(rootAccount)
			if err != nil {
				return nil, err
			}
			as := n.accountService
			ss := entry.storage

			address := util.ToString(&item, "address")
			_, a, err := as.NewAccountFromAddress(address)
			if err != nil {
				return nil, err
			}
			gas, err := ss.GetGas(a)
			if err != nil {
				return nil, err
			}
			return gas.String(), nil
		})
		if err != nil {
			return nil, err
		}
		return result, nil

	case "getMeta":
		result, err := n._call(params, func(rootAccount libcore.Address, item map[string]interface{}) (interface{}, error) {
			entry, err := n.GetEntry(rootAccount)
			if err != nil {
				return nil, err
			}
			ss := entry.storage

			symbol := util.ToString(&item, "symbol")
			index, account, info, err := ss.ReadMeta(libcore.Symbol(symbol))
			if err != nil {
				return nil, err
			}
			return map[string]interface{}{
				"account": account.String(),
				"info":    info,
				"index":   index,
			}, nil
		})
		if err != nil {
			return nil, err
		}
		return result, nil

	case "getToken":
		result, err := n._call(params, func(rootAccount libcore.Address, item map[string]interface{}) (interface{}, error) {
			entry, err := n.GetEntry(rootAccount)
			if err != nil {
				return nil, err
			}
			ss := entry.storage

			symbol := util.ToString(&item, "symbol")
			index := util.ToUint64(&item, "index")
			account, info, err := ss.ReadToken(libcore.Symbol(symbol), index)
			if err != nil {
				return nil, err
			}
			return map[string]interface{}{
				"account": account.String(),
				"info":    info,
			}, nil
		})
		if err != nil {
			return nil, err
		}
		return result, nil

	case "getTransactionCount":
		result, err := n._call(params, func(rootAccount libcore.Address, item map[string]interface{}) (interface{}, error) {
			as := n.accountService
			address := util.ToString(&item, "address")
			_, a, err := as.NewAccountFromAddress(address)
			if err != nil {
				return nil, err
			}
			seq := n.getNextSequence(rootAccount, a)
			return seq, nil
		})
		if err != nil {
			return nil, err
		}
		return result, nil

	case "getTransactionReceipt":
		result, err := n._call(params, func(rootAccount libcore.Address, item map[string]interface{}) (interface{}, error) {
			entry, err := n.GetEntry(rootAccount)
			if err != nil {
				return nil, err
			}
			ms := entry.merkle

			hashString := util.ToString(&item, "hash")
			h, err := hex.DecodeString(hashString)
			if err != nil {
				return nil, err
			}
			txWithData, err := ms.GetTransactionByHash(libcore.Hash(h))
			if err != nil {
				return nil, err
			}
			receipt := txWithData.GetReceipt()
			_, err = n.consensusService.HashReceipt(receipt)
			if err != nil {
				return nil, err
			}
			return receipt, nil
		})
		if err != nil {
			return nil, err
		}
		return result, nil

	case "getTransactionByHash":
		result, err := n._call(params, func(rootAccount libcore.Address, item map[string]interface{}) (interface{}, error) {
			entry, err := n.GetEntry(rootAccount)
			if err != nil {
				return nil, err
			}
			ms := entry.merkle

			hashString := util.ToString(&item, "hash")
			h, err := hex.DecodeString(hashString)
			if err != nil {
				return nil, err
			}
			txWithData, err := ms.GetTransactionByHash(libcore.Hash(h))
			if err != nil {
				return nil, err
			}
			_, err = n.consensusService.HashTransaction(txWithData)
			if err != nil {
				return nil, err
			}
			return txWithData, nil
		})
		if err != nil {
			return nil, err
		}
		return result, nil

	case "getTransactionByIndex":
		result, err := n._call(params, func(rootAccount libcore.Address, item map[string]interface{}) (interface{}, error) {
			entry, err := n.GetEntry(rootAccount)
			if err != nil {
				return nil, err
			}
			as := n.accountService
			ms := entry.merkle

			address := util.ToString(&item, "address")
			_, a, err := as.NewAccountFromAddress(address)
			if err != nil {
				return nil, err
			}
			index := util.ToUint64(&item, "index")

			txWithData, err := ms.GetTransactionByIndex(a, index)
			if err != nil {
				return nil, err
			}
			_, err = n.consensusService.HashTransaction(txWithData)
			if err != nil {
				return nil, err
			}
			return txWithData, nil
		})
		if err != nil {
			return nil, err
		}
		return result, nil

	case "getBlockByHash":
		result, err := n._call(params, func(rootAccount libcore.Address, item map[string]interface{}) (interface{}, error) {
			entry, err := n.GetEntry(rootAccount)
			if err != nil {
				return nil, err
			}
			ms := entry.merkle

			hashString := util.ToString(&item, "hash")
			h, err := hex.DecodeString(hashString)
			if err != nil {
				return nil, err
			}
			block, err := ms.GetBlockByHash(libcore.Hash(h))
			if err != nil {
				return nil, err
			}
			_, err = n.consensusService.HashBlock(block)
			if err != nil {
				return nil, err
			}
			return block, nil
		})
		if err != nil {
			return nil, err
		}
		return result, nil

	case "getBlockByNumber":
		result, err := n._call(params, func(rootAccount libcore.Address, item map[string]interface{}) (interface{}, error) {
			entry, err := n.GetEntry(rootAccount)
			if err != nil {
				return nil, err
			}
			ms := entry.merkle

			index := util.ToUint64(&item, "num")
			block, err := ms.GetBlockByIndex(index)
			if err != nil {
				return nil, err
			}
			_, err = n.consensusService.HashBlock(block)
			if err != nil {
				return nil, err
			}
			return block, nil
		})
		if err != nil {
			return nil, err
		}
		return result, nil

	case "getNodeInfo":
		address := n.self.GetAddress(n)
		publicData, err := n.self.Key.MarshalText()
		if err != nil {
			return nil, err
		}
		return &map[string]interface{}{
			"address": address,
			"public":  string(publicData),
			"node":    n.config.GetNodeId(),
		}, nil

	case "getData":
		maxCost := int64(1000000)
		result, err := n._call(params, func(rootAccount libcore.Address, item map[string]interface{}) (interface{}, error) {
			entry, err := n.GetEntry(rootAccount)
			if err != nil {
				return nil, err
			}
			ss := entry.storage

			hashString := util.ToString(&item, "hash")
			h, err := hex.DecodeString(hashString)
			if err != nil {
				return nil, err
			}
			format := util.ToString(&item, "format")
			usedCost, r, err := ss.GetData(libcore.Hash(h), format)
			if err != nil {
				return nil, err
			}

			maxCost -= usedCost
			if maxCost < 0 {
				return nil, util.ErrorOf("too long", "call", "contract")
			}

			return r, nil
		})
		if err != nil {
			return nil, err
		}
		return result, nil

	case "getUserData":
		result, err := n._call(params, func(rootAccount libcore.Address, item map[string]interface{}) (interface{}, error) {
			entry, err := n.GetEntry(rootAccount)
			if err != nil {
				return nil, err
			}
			as := n.accountService
			ss := entry.storage

			addressString := util.ToString(&item, "address")
			_, a, err := as.NewAccountFromAddress(addressString)
			if err != nil {
				return nil, err
			}
			accountString := util.ToString(&item, "account")
			_, b, err := as.NewAccountFromAddress(accountString)
			if err != nil {
				return nil, err
			}
			info, err := ss.ReadUser(a, b)
			if err != nil {
				return nil, err
			}
			return &map[string]interface{}{
				"account": b.String(),
				"key":     hex.EncodeToString(info.Key),
				"nonce":   hex.EncodeToString(info.Nonce),
				"hash":    libcore.Hash(info.Data.Hash).String(),
			}, nil
		})
		if err != nil {
			return nil, err
		}
		return result, nil

	case "getContractData":
		maxCost := int64(1000000)
		result, err := n._call(params, func(rootAccount libcore.Address, item map[string]interface{}) (interface{}, error) {
			usedCost, r, err := n.getContractData(rootAccount, item)
			if err != nil {
				return nil, err
			}

			maxCost -= usedCost
			if maxCost < 0 {
				return nil, util.ErrorOf("too long", "call", "contract")
			}

			return r, nil
		})
		if err != nil {
			return nil, err
		}
		return result, nil

	case "callContract":
		maxCost := int64(1000000)
		result, err := n._call(params, func(rootAccount libcore.Address, item map[string]interface{}) (interface{}, error) {
			usedCost, r, err := n.callContract(rootAccount, item)
			if err != nil {
				return nil, err
			}

			maxCost -= usedCost
			if maxCost < 0 {
				return nil, util.ErrorOf("too long", "call", "contract")
			}

			return r, nil
		})
		if err != nil {
			return nil, err
		}
		return result, nil

	case "signTransaction":
		result, err := n._call(params, func(rootAccount libcore.Address, item map[string]interface{}) (interface{}, error) {
			blob, tx, err := n.signTransaction(rootAccount, item)
			if err != nil {
				return nil, err
			}
			hash := tx.GetHash()
			glog.Infoln("sign transaction", hash.String(), len(blob))
			return blob, nil
		})
		if err != nil {
			return nil, err
		}
		return result, nil

	case "sendTransaction":
		result, err := n._call(params, func(rootAccount libcore.Address, item map[string]interface{}) (interface{}, error) {
			_, tx, err := n.signTransaction(rootAccount, item)
			if err != nil {
				return nil, err
			}
			_, err = n.sendTransaction(rootAccount, tx)
			if err != nil {
				return nil, err
			}

			hash := tx.GetHash()
			return hash, nil
		})
		if err != nil {
			return nil, err
		}
		return result, nil

	case "sendRawTransaction":
		result, err := n._call(params, func(rootAccount libcore.Address, item map[string]interface{}) (interface{}, error) {
			blob := util.ToString(&item, "blob")

			data, err := hex.DecodeString(blob)
			if err != nil {
				return nil, err
			}
			tx := &block.Transaction{}
			err = tx.UnmarshalBinary(data)
			if err != nil {
				return nil, err
			}
			_, err = n.sendTransaction(rootAccount, tx)
			if err != nil {
				if util.IsTest() || util.IsDebug() {
					blobDir := "./consensus/data"
					isExists, isDir := util.Exists(blobDir)
					if isExists && isDir {
						blobFile := path.Join(blobDir, "tx.blob")
						e := util.WriteFile(blobFile, []byte(blob))
						if e != nil {
							return nil, e
						}
						glog.Infoln("write blob to ", blobFile)

						rawFile := path.Join(blobDir, "tx.raw")
						e = util.WriteFile(rawFile, []byte(fmt.Sprintf("%v", data)))
						if e != nil {
							return nil, e
						}
						glog.Infoln("write raw to ", rawFile)

						jsonFile := path.Join(blobDir, "tx.json")
						jsonBytes, e := json.MarshalIndent(tx, "", " ")
						if e != nil {
							return nil, e
						}
						e = util.WriteFile(jsonFile, jsonBytes)
						if e != nil {
							return nil, e
						}
						glog.Infoln("write json to ", jsonFile)
					}
				}
				glog.Error(err)
				return nil, err
			}

			hash := tx.GetHash()
			return hash, nil
		})
		if err != nil {
			return nil, err
		}
		return result, nil

	case "getStateByHash":
		result, err := n._call(params, func(rootAccount libcore.Address, item map[string]interface{}) (interface{}, error) {
			entry, err := n.GetEntry(rootAccount)
			if err != nil {
				return nil, err
			}
			ms := entry.merkle

			hashString := util.ToString(&item, "hash")
			h, err := hex.DecodeString(hashString)
			if err != nil {
				return nil, err
			}
			s, err := ms.GetStateByHash(libcore.Hash(h))
			if err != nil {
				return nil, err
			}
			_, err = n.consensusService.HashState(s)
			if err != nil {
				return nil, err
			}
			return s, nil
		})
		if err != nil {
			return nil, err
		}
		return result, nil

	case "getStateByAddress":
		result, err := n._call(params, func(rootAccount libcore.Address, item map[string]interface{}) (interface{}, error) {
			as := n.accountService

			stateTypeName := util.ToString(&item, "type")
			address := util.ToString(&item, "address")
			stateType := libblock.GetStateTypeByName(stateTypeName)
			switch stateType {
			case block.ACCOUNT_STATE:
				_, a, err := as.NewAccountFromAddress(address)
				if err != nil {
					return nil, err
				}
				s, err := n.consensusService.GetAccountInfo(rootAccount, a)
				if err != nil {
					return nil, err
				}
				_, err = n.consensusService.HashState(s)
				if err != nil {
					return nil, err
				}
				return s, nil
			default:
				return nil, util.ErrorOfUnknown("state", fmt.Sprintf("%d(%s)", stateType, stateTypeName))
			}
		})
		if err != nil {
			return nil, err
		}
		return result, nil

	case "dump":
		result, err := n._call(params, func(rootAccount libcore.Address, item map[string]interface{}) (interface{}, error) {
			entry, err := n.GetEntry(rootAccount)
			if err != nil {
				return nil, err
			}
			ss := entry.storage

			if util.IsTest() || util.IsDebug() {
				ss.Dump(storage.LogPrinter{})
			}
			return ss.RootHash(), nil
		})
		if err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, util.ErrorOfNotFound("method", method)
	}
}

func (n *Node) GetEntry(rootAccount libcore.Address) (*Entry, error) {
	n.entryLocker.RLock()
	root := util.GetString(rootAccount)
	entry, ok := n.entryMap[root]
	n.entryLocker.RUnlock()
	if ok {
		return entry, nil
	}

	n.entryLocker.Lock()
	defer n.entryLocker.Unlock()

	entry, ok = n.entryMap[root]
	if ok {
		return entry, nil
	}

	ms, err := n.getMerkleService(rootAccount)
	if err != nil {
		return nil, err
	}
	ss, err := n.getStorageService(rootAccount)
	if err != nil {
		return nil, err
	}
	entry = &Entry{
		merkle:  ms,
		storage: ss,

		txlist: make([]libblock.TransactionWithData, 0),
		txmap:  make(map[string]libblock.TransactionWithData),
	}
	err = n.load(rootAccount, entry)
	if err != nil {
		return nil, err
	}
	n.entryMap[root] = entry
	return entry, nil
}

func (n *Node) getMerkleService(rootAccount libcore.Address) (libstore.MerkleService, error) {
	ms, err := merkle.NewMerkleService(n.config, n.cryptoService, rootAccount)
	if err != nil {
		return nil, err
	}
	return ms, nil
}

func (n *Node) getStorageService(rootAccount libcore.Address) (*storage.StorageService, error) {
	ss, err := storage.NewStorageService(n.storageService, rootAccount)
	if err != nil {
		return nil, err
	}
	return ss, nil
}

func (n *Node) addTransaction(rootAccount libcore.Address, txWithData libblock.TransactionWithData) (bool, error) {
	_, err := n.consensusService.HashTransaction(txWithData)
	if err != nil {
		return false, err
	}

	entry, err := n.GetEntry(rootAccount)
	if err != nil {
		return false, err
	}

	h := txWithData.GetTransaction().GetHash().String()
	_, ok := entry.txmap[h]
	if ok {
		return false, util.ErrorOf("already exists", "transaction", h)
	}

	entry.txlist = append(entry.txlist, txWithData)
	entry.txmap[h] = txWithData
	return true, nil
}

func (n *Node) _generateBlock(rootAccount libcore.Address, list []libblock.TransactionWithData, rb libblock.Block) (libblock.Block, error) {
	config := n.config
	cs := n.cryptoService
	entry, err := n.GetEntry(rootAccount)
	if err != nil {
		return nil, err
	}
	ms := entry.merkle
	ss := entry.storage

	v := entry.ValidatedBlock

	var b *block.Block
	if v == nil { //genesis
		if len(list) > 0 {
			return nil, util.ErrorOfInvalid("transactions", "genesis block")
		}
		var states []libblock.State
		if rootAccount != nil {
			// v, err := util.NewValue("100000000000000")
			// if err != nil {
			// 	return nil, err
			// }
			// err = ss.UpdateGas(rootAccount, *v)
			// if err != nil {
			// 	return nil, err
			// }
			gasAccount := config.GetGasAccount()
			// err = ss.UpdateGas(gasAccount, v.ZeroClone())
			// if err != nil {
			// 	return nil, err
			// }
			states = []libblock.State{
				&block.AccountState{
					State: block.State{
						StateType:  block.ACCOUNT_STATE,
						Account:    rootAccount,
						Sequence:   uint64(0),
						BlockIndex: uint64(0),
					},
				},
				&block.AccountState{
					State: block.State{
						StateType:  block.ACCOUNT_STATE,
						Account:    gasAccount,
						Sequence:   uint64(0),
						BlockIndex: uint64(0),
					},
				},
			}
		} else {
			account := n.config.GetGasAccount()
			// info := &pb.MetaInfo{
			// 	Symbol: block.ACCOUNT_STATE.String(),
			// 	Total:  int64(-1),
			// 	Items: []*pb.MetaItem{
			// 		{
			// 			Name:    "account",
			// 			Type:    core.CORE_DATA_STRING.String(),
			// 			Options: make([]string, 0),
			// 			Desc:    "the account",
			// 		},
			// 	},
			// }
			// dataInfo, err := ss.WriteMeta(account, info)
			// if err != nil {
			// 	return nil, err
			// }
			states = []libblock.State{
				&block.AccountState{
					State: block.State{
						StateType:  block.ACCOUNT_STATE,
						Account:    account,
						Sequence:   uint64(0),
						BlockIndex: uint64(0),
					},
					// Token: dataInfo,
				},
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
			Account:    rootAccount,
			BlockIndex: uint64(0),
			ParentHash: func(rb libblock.Block) libcore.Hash {
				if rb != nil {
					return rb.GetHash()
				}
				return libcrypto.ZeroHash(cs)
			}(rb),
			Transactions:    []libblock.TransactionWithData{},
			TransactionHash: ms.GetTransactionRoot(),

			States:    states,
			StateHash: ms.GetStateRoot(),

			RootHash: ss.RootHash(),

			Timestamp: func(rb libblock.Block) int64 {
				if rb != nil {
					return rb.GetTime()
				}
				return time.Now().UnixNano()
			}(rb),
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
			Account:    rootAccount,
			BlockIndex: v.GetIndex() + 1,
			ParentHash: v.GetHash(),

			Transactions:    list,
			TransactionHash: ms.GetTransactionRoot(),

			States:    states,
			StateHash: ms.GetStateRoot(),

			RootHash: ss.RootHash(),

			Timestamp: time.Now().UnixNano(),
		}

		err := ms.Cancel()
		if err != nil {
			return nil, err
		}
	}

	_, _, err = cs.Raw(b, libcrypto.RawBinary)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (n *Node) AddBlock(rootAccount libcore.Address, b libblock.Block) error {
	n.transactionLocker.Lock()
	defer n.transactionLocker.Unlock()

	entry, err := n.GetEntry(rootAccount)
	if err != nil {
		return err
	}
	ms := entry.merkle
	ss := entry.storage

	err = ss.CommitSandbox()
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
	entry.ValidatedBlock = b

	e := n.self.getPeerEntry(rootAccount)
	e.BlockNumber = int64(b.GetIndex())

	return nil
}

func (n *Node) VerifyParent(rootAccount libcore.Address, b libblock.Block) error {
	if rootAccount != nil {
		if b.GetParentHash().IsZero() {
			return util.ErrorOfInvalid("the parent hash of the genesis block", "block")
		}
		rootEntry, err := n.GetEntry(nil)
		if err != nil {
			return err
		}
		_, err = rootEntry.merkle.GetBlockByHash(b.GetParentHash())
		if err != nil {
			return err
		}
	} else {
		if !b.GetParentHash().IsZero() {
			return util.ErrorOfInvalid("the parent hash of the genesis block", "block")
		}
	}
	return nil
}

func (n *Node) VerifyBlock(rootAccount libcore.Address, b libblock.Block) (ok bool, err error) {
	n.transactionLocker.Lock()
	defer n.transactionLocker.Unlock()

	entry, err := n.GetEntry(rootAccount)
	if err != nil {
		return false, err
	}
	cs := n.cryptoService
	ms := entry.merkle
	ss := entry.storage

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

		ok, err = n.consensusService.VerifyTransaction(rootAccount, tx)
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
		newWithData, e := n.consensusService.ProcessTransaction(rootAccount, tx)
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

	if entry.ValidatedBlock != nil {
		if b.GetIndex() != (entry.ValidatedBlock.GetIndex() + 1) {
			ok = false
			err = util.ErrorOfUnmatched("index", "block", (entry.ValidatedBlock.GetIndex() + 1), b.GetIndex())
			return
		}
		if b.GetParentHash().String() != entry.ValidatedBlock.GetHash().String() {
			ok = false
			err = util.ErrorOfUnmatched("parent hash", "block", entry.ValidatedBlock.GetHash().String(), b.GetParentHash().String())
			return
		}
	} else {
		if b.GetIndex() != 0 {
			ok = false
			err = util.ErrorOfInvalid("index", "block")
			return
		}
		err = n.VerifyParent(rootAccount, b)
		if err != nil {
			ok = false
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

	blk := b.(*block.Block)
	srh := ss.RootHash()
	sbh := blk.GetRootHash()
	if srh.String() != sbh.String() {
		return false, util.ErrorOfUnmatched("root hash", "block", srh.String(), sbh.String())
	}

	return
}

func (n *Node) GenerateBlock(rootAccount libcore.Address, rb libblock.Block) (libblock.Block, error) {
	n.transactionLocker.Lock()
	defer n.transactionLocker.Unlock()

	entry, err := n.GetEntry(rootAccount)
	if err != nil {
		return nil, err
	}
	list := entry.txlist
	entry.txlist = make([]libblock.TransactionWithData, 0)
	entry.txmap = make(map[string]libblock.TransactionWithData)

	return n._generateBlock(rootAccount, list, rb)
}

func (n *Node) generate() {
	for {
		hasTx := false
		for root, entry := range n.entryMap {
			_, rootAccount, _ := n.accountService.NewAccountFromAddress(root)
			if entry.Consensused && len(entry.txlist) > 0 {
				hasTx = hasTx || true

				block, err := n.GenerateBlock(rootAccount, nil)
				if err != nil {
					glog.Error(err)
				} else {
					h, err := n.consensusService.HashBlock(block)
					if err != nil {
						glog.Error(err)
					} else {
						glog.Infof("=== %s: generate block %d, %s, %d\n", util.GetString(rootAccount), block.GetIndex(), h.String(), len(block.GetTransactions()))

						_, err = n.VerifyBlock(rootAccount, block)
						if err != nil {
							glog.Error(err)
						} else {
							err = n.AddBlock(rootAccount, block)
							if err != nil {
								glog.Error(err)
							} else {
								data, err := block.MarshalBinary()
								if err != nil {
									glog.Error(err)
								} else {
									n.broadcast(data)
								}
							}
						}
					}
				}
			}
		}
		if !hasTx {
			for root, entry := range n.entryMap {
				if entry.Consensused {
					glog.Infof("=== %s: block %d, %s, prepare %d, %d\n", root, entry.GetBlockNumber(), entry.GetBlockHash(), entry.GetBlockNumber()+1, len(entry.txlist))
				}
			}

			for i := 0; i < int(n.config.GetBlockDuration()*5); i++ {
				hasTx := false
				for _, entry := range n.entryMap {
					if len(entry.txlist) > 0 {
						hasTx = true
						break
					}
				}
				if !hasTx {
					time.Sleep(200 * time.Millisecond)
				} else {
					break
				}
			}
		}
	}
}

func (n *Node) newId() uint64 {
	return atomic.AddUint64(&n.messageId, 1)
}

func (n *Node) broadcast(data []byte) {
	msg, err := n.CreateMessage(data)
	if err != nil {
		glog.Error(err)
	} else {
		n.outs <- msg
	}
}

func (n *Node) connect() {
	config := n.config
	for {
		for root, entry := range n.entryMap {
			_, rootAccount, _ := n.accountService.NewAccountFromAddress(root)
			list := n.ListPeerBy(rootAccount)
			e := n.self.getPeerEntry(rootAccount)
			e.PeerCount = int64(len(list))

			glog.Infoln("=== :", root, 0, n.self.GetIndex(n), n.self.GetAddress(n), n.self.Id, e.Status, e.BlockNumber, e.PeerCount)
			for i := 0; i < len(list); i++ {
				p := list[i]
				e := p.getPeerEntry(rootAccount)
				glog.Infoln("==> :", root, i+1, p.GetIndex(n), p.GetAddress(n), p.Id, e.Status, e.BlockNumber, e.PeerCount)
			}
			if len(list) == 0 {
				entry.Consensused = n.PrepareConsensus(rootAccount)
			}
		}

		time.Sleep(time.Duration(config.GetBlockDuration()) * time.Second)
	}
}

func (n *Node) ConnectTo(p *Peer) {
	glog.Infoln("connected to", p.GetAddress(n))
}

func (n *Node) SendRequestInfo(p *Peer) {
	glog.Infoln("request to", p.GetAddress(n))
}

func (n *Node) PrepareConsensus(rootAccount libcore.Address) bool {
	l := len(n.config.GetBootstraps())
	if l == 0 {
		return true
	}

	root := util.GetString(rootAccount)
	list := n.ListPeerBy(rootAccount)
	count := int64(0)
	entry, err := n.GetEntry(rootAccount)
	if err != nil {
		glog.Error(err)
		return false
	}
	currentBlock := entry.GetBlockNumber()
	currentCount := int64(len(list))
	for i := 0; i < len(list); i++ {
		p := list[i]
		e, ok := p.peermap[root]
		if ok {
			if e.Status >= PeerKnown && e.BlockNumber == currentBlock && (e.PeerCount == currentCount) {
				count++
			}
		}
	}
	if count > 0 && (count+1) == currentCount {
		for i := 0; i < len(list); i++ {
			p := list[i]
			e, ok := p.peermap[root]
			if ok {
				e.Status = PeerConsensused
			}
		}
		return true
	}
	return false
}

func (n *Node) discoveryPeer(p *Peer) {
	config := n.config

	for {
		balanced := true
		for root, e := range p.peermap {
			_, rootAccount, _ := n.accountService.NewAccountFromAddress(root)
			entry, err := n.GetEntry(rootAccount)
			if err != nil {
				glog.Error(err)
			} else {
				switch e.Status {
				case PeerNone:
					n.ConnectTo(p)
				case PeerConnected:
					n.SendRequestInfo(p)
				case PeerKnown:
					entry.Consensused = n.PrepareConsensus(rootAccount)
				}
				if e.Status >= PeerKnown && entry.GetBlockNumber() > e.BlockNumber {
					balanced = balanced && false

					glog.Infoln("===", entry.GetBlockNumber(), e.lastSendBlock, e.lastSendTime, p.address, p.Id, e.Status, e.BlockNumber, e.PeerCount)
					if e.lastSendBlock >= 0 && e.lastSendBlock == (e.BlockNumber+1) && (time.Since(e.lastSendTime) < (time.Duration(3 * int64(time.Second)))) {
						time.Sleep(3 * time.Second)
					} else {
						block, err := entry.merkle.GetBlockByIndex(uint64(e.BlockNumber + 1))
						if err != nil {
							glog.Error(err)
						} else {
							_, err = n.consensusService.HashBlock(block)
							if err != nil {
								glog.Error(err)
							} else {
								blockData, err := block.MarshalBinary()
								if err != nil {
									glog.Error(err)
								} else {
									mid, msgData, err := n.EncodeMessage(blockData)
									if err != nil {
										glog.Error(err)
									} else {
										err = n.net.SendMsg(n.config.GetChainId(), p.Id, TOPIC_PEER_DATA, msgData)
										if err != nil {
											glog.Error(err)
										} else {
											e.lastSendBlock = int64(block.GetIndex())
											e.lastSendTime = time.Now()
											glog.Infof(">>> send data %d(%s) to %s(%s)\n", mid, core.GetInfo(blockData), p.GetAddress(n), p.Id)
										}
									}
								}
							}
						}
						time.Sleep(time.Second)
					}
				}
			}
		}
		if balanced {
			time.Sleep(time.Duration(config.GetBlockDuration()) * time.Second)
		}
	}
}

func (n *Node) load(rootAccount libcore.Address, entry *Entry) error {
	n.transactionLocker.Lock()
	defer n.transactionLocker.Unlock()

	e := n.self.getPeerEntry(rootAccount)

	current := uint64(0)
	for {
		b, err := entry.merkle.GetBlockByIndex(current)
		if err != nil {
			glog.Error(err)
			break
		}
		_, err = n.consensusService.HashBlock(b)
		if err != nil {
			return err
		}
		if entry.ValidatedBlock != nil {
			if b.GetIndex() != (entry.ValidatedBlock.GetIndex() + 1) {
				return util.ErrorOfInvalid("block index", fmt.Sprintf("%d != %d", b.GetIndex(), (entry.ValidatedBlock.GetIndex()+1)))
			}
			if !b.GetParentHash().Equals(entry.ValidatedBlock.GetHash()) {
				return util.ErrorOfInvalid("block hash", fmt.Sprintf("%s != %s", b.GetParentHash().String(), entry.ValidatedBlock.GetHash().String()))
			}
		} else {
			if b.GetIndex() != 0 {
				return util.ErrorOfInvalid("genesis block", fmt.Sprintf("%d, %s", b.GetIndex(), b.GetHash().String()))
			}
			err := n.VerifyParent(rootAccount, b)
			if err != nil {
				return err
			}
		}
		entry.ValidatedBlock = b

		allData, err := b.Raw(libcrypto.RawBinary)
		if err != nil {
			glog.Error(err)
			break
		}
		ignoreData, err := b.Raw(libcrypto.RawIgnoreVariableFields)
		if err != nil {
			glog.Error(err)
			break
		}
		signData, err := b.Raw(libcrypto.RawIgnoreSigningFields)
		if err != nil {
			glog.Error(err)
			break
		}
		hashData, err := b.Raw(libcrypto.RawIgnoreContent)
		if err != nil {
			glog.Error(err)
			break
		}
		glog.Infoln("load block", current, b.GetHash().String(), len(allData), len(ignoreData), len(signData), len(hashData))

		current++
	}
	e.BlockNumber = entry.GetBlockNumber()
	e.Status = PeerConsensused

	return nil
}

func (n *Node) Dump(rootAccount libcore.Address, printer core.Printer) {
	entry, err := n.GetEntry(rootAccount)
	if err != nil {
		printer.Printf(err.Error())
	} else {
		ss := entry.storage
		ss.Dump(printer)
	}
}

func (n *Node) send() {
	for {
		m := <-n.outs
		msgBytes, err := core.Marshal(m)
		if err != nil {
			glog.Error(err)
		} else {
			data := m.GetData()
			list := n.ListPeer()
			for i := 0; i < len(list); i++ {
				p := list[i]
				err := n.net.SendMsg(n.config.GetChainId(), p.Id, TOPIC_PEER_MESSAGE, msgBytes)
				if err != nil {
					glog.Error(err)
				} else {
					glog.Infof(">>> send message %d(%s) to %s(%s)\n", m.Id, core.GetInfo(data), p.GetAddress(n), p.Id)
				}
			}
		}
	}
}

func (n *Node) receive() {
	for {
		m := <-n.ins
		meta, fromPeer, data := n.ReceiveMessage("message", m)
		fromIndex := m.GetNode()
		glog.Infof("<<< receive message from node %d, length: %d\n", fromIndex, len(data))
		if fromPeer != nil {
			glog.Infof("<<< receive peer message %d(%s) from %s(%s)\n", m.Id, core.GetInfo(data), fromPeer.GetAddress(n), fromPeer.Id)
			switch meta {
			case core.CORE_BLOCK:
				b := &block.Block{}
				err := b.UnmarshalBinary(data)
				if err != nil {
					glog.Error(err)
				} else {
					h, err := n.consensusService.HashBlock(b)
					if err != nil {
						glog.Error(err)
					} else {
						rootAccount := b.GetAccount()
						entry, err := n.GetEntry(rootAccount)
						if err != nil {
							glog.Error(err)
						} else {
							if int64(b.GetIndex()) <= entry.GetBlockNumber() {
								glog.Infoln("<<< drop message block", util.GetString(rootAccount), b.GetIndex(), h.String(), len(b.GetTransactions()))
							} else {
								glog.Infoln("<<< receive message block", util.GetString(rootAccount), b.GetIndex(), h.String(), len(b.GetTransactions()))
								_, err = n.VerifyBlock(rootAccount, b)
								if err != nil {
									glog.Error(err)
								} else {
									glog.Infoln("<<< verify message block", util.GetString(rootAccount), b.GetIndex(), h.String(), len(b.GetTransactions()))
									err = n.AddBlock(rootAccount, b)
									if err != nil {
										glog.Error(err)
									} else {
										glog.Infoln("<<< add message block", util.GetString(rootAccount), b.GetIndex(), h.String(), len(b.GetTransactions()))
									}
								}
							}
							n.SendPeerInfo(rootAccount, fromPeer)
						}
					}
				}

			default:
				glog.Errorln("<<< error message")
			}
		} else {
			glog.Infof("<<< receive unknown message %d(%s)\n", m.Id, core.GetInfo(data))
			if len(data) > 0 {
				meta := data[0]
				switch meta {
				default:
					glog.Warningln("<<< unknown peer", fromIndex)
				}
			}
		}
	}
}

func (n *Node) discoveryHandler(publisher string, msgData []byte) error {
	glog.Infof("[%s][%s] recv discovery from peer[%s], len: %d", n.config.GetChainId(), n.self.GetAddress(n), publisher, len(msgData))
	m, err := n.DecodeMessage(msgData)
	if err != nil {
		return err
	}
	meta, msg, err := core.Unmarshal(m.GetData())
	if err != nil {
		return err
	}
	if meta != core.CORE_PEER_INFO {
		return util.ErrorOfInvalid("format", "peer info")
	}
	peerInfo := msg.(*pb.PeerInfo)
	_, publicKey, err := n.accountService.NewPublicFromBytes(peerInfo.PublicKey)
	if err != nil {
		return err
	} else {
		address, err := publicKey.GenerateAddress()
		if err != nil {
			return err
		}
		index, err := n.GetIndex(address.String())
		if err != nil {
			return err
		}
		rootAccount, err := block.ByteToAddress(peerInfo.Account)
		if err != nil {
			return err
		}
		p := n.GetPeer(index)
		if p == nil {
			p = &Peer{
				Id:      publisher,
				Key:     publicKey,
				peermap: make(map[string]*peerEntry),
			}
			n.AddPeer(p)
			n.UpdateStatus(p, rootAccount, PeerKnown)
		} else {
			n.UpdateSafeStatus(p, rootAccount, PeerKnown)
		}
		n.UpdatePeer(p, peerInfo)
	}
	return nil
}

func (n *Node) GetIndex(address string) (uint64, error) {
	_, a, err := n.accountService.NewAccountFromAddress(address)
	if err != nil {
		return 0, err
	}
	data, err := a.MarshalBinary()
	if err != nil {
		return 0, err
	}
	m := new(big.Int).SetBytes(data)
	return m.Uint64(), nil
}

func (n *Node) messageHandler(id string, data []byte) error {
	glog.Infof("[%s][%s] recv message from peer[%s], len: %d", n.config.GetChainId(), n.self.GetAddress(n), id, len(data))
	msg, err := n.DecodeMessage(data)
	if err != nil {
		return err
	}
	n.ins <- msg
	return nil
}

func (n *Node) dataHandler(id string, msgData []byte) error {
	glog.Infof("[%s][%s] recv data from peer[%s], len: %d", n.config.GetChainId(), n.self.GetAddress(n), id, len(msgData))
	m, err := n.DecodeMessage(msgData)
	if err != nil {
		return err
	}
	meta, fromPeer, data := n.ReceiveMessage("data", m)
	if fromPeer != nil {
		switch meta {
		case core.CORE_PEER_INFO:
			_, msg, err := core.Unmarshal(data)
			if err != nil {
				return err
			}
			peerInfo := msg.(*pb.PeerInfo)
			n.UpdatePeer(fromPeer, peerInfo)

		case core.CORE_BLOCK:
			b := &block.Block{}
			err = b.UnmarshalBinary(data)
			if err != nil {
				glog.Error(err)
			} else {
				h, err := n.consensusService.HashBlock(b)
				if err != nil {
					glog.Error(err)
				} else {
					rootAccount := b.GetAccount()
					entry, err := n.GetEntry(rootAccount)
					if err != nil {
						glog.Error(err)
					} else {
						if int64(b.GetIndex()) <= entry.GetBlockNumber() {
							glog.Infoln("<<< drop data block", util.GetString(rootAccount), b.GetIndex(), h.String(), len(b.GetTransactions()))
						} else {
							glog.Infoln("<<< receive data block", util.GetString(rootAccount), b.GetIndex(), h.String(), len(b.GetTransactions()))

							_, err = n.VerifyBlock(rootAccount, b)
							if err != nil {
								glog.Error(err)
							} else {
								glog.Infoln("<<< verify data block", util.GetString(rootAccount), b.GetIndex(), h.String(), len(b.GetTransactions()))
								err = n.AddBlock(rootAccount, b)
								if err != nil {
									glog.Error(err)
								} else {
									glog.Infoln("<<< add data block", util.GetString(rootAccount), b.GetIndex(), h.String(), len(b.GetTransactions()))
								}
							}
						}
						n.SendPeerInfo(rootAccount, fromPeer)
					}
				}
			}
		}
	}
	return nil
}

func (n *Node) SendPeerInfo(rootAccount libcore.Address, toPeer *Peer) {
	rootData, err := block.AddressToByte(rootAccount)
	if err != nil {
		glog.Error(err)
	} else {
		e := n.self.getPeerEntry(rootAccount)
		peerData, err := core.Marshal(&pb.PeerInfo{
			Account:     rootData,
			BlockNumber: e.BlockNumber,
			PeerCount:   e.PeerCount,
		})
		if err != nil {
			glog.Error(err)
		} else {
			mid, msgData, err := n.EncodeMessage(peerData)
			if err != nil {
				glog.Error(err)
			} else {
				err = n.net.SendMsg(n.config.GetChainId(), toPeer.Id, TOPIC_PEER_DATA, msgData)
				if err != nil {
					glog.Error(err)
				} else {
					glog.Infof(">>> send data %d(%s) to %s(%s)\n", mid, core.GetInfo(peerData), toPeer.GetAddress(n), toPeer.Id)
				}
			}
		}
	}
}

func (n *Node) LoadPageByName(name string) (*zipfs.FileSystem, error) {
	if len(name) == 0 {
		return nil, util.ErrorOfInvalid("name", name)
	}
	_, account, _ := n.accountService.NewAccountFromAddress(name)
	if account != nil {
		return nil, util.ErrorOfInvalid("name", name)
	}
	entry, err := n.GetEntry(nil)
	if err != nil {
		return nil, err
	}
	ss := entry.storage
	data, err := ss.ReadPageByName(name)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(data)
	fs, err := zipfs.NewFromReaderAt(reader, int64(len(data)), nil)
	if err != nil {
		return nil, err
	}
	return fs, nil
}

func (n *Node) LoadPageByNameOrAddress(nameOrAddress string) (*zipfs.FileSystem, error) {
	if len(nameOrAddress) == 0 {
		return nil, util.ErrorOfInvalid("name", nameOrAddress)
	}

	entry, err := n.GetEntry(nil)
	if err != nil {
		return nil, err
	}
	ss := entry.storage

	var data []byte
	_, account, err := n.accountService.NewAccountFromAddress(nameOrAddress)
	if err != nil {
		data, err = ss.ReadPageByName(nameOrAddress)
		if err != nil {
			return nil, err
		}
	} else {
		data, err = ss.ReadPageByAddress(account)
		if err != nil {
			return nil, err
		}
	}
	reader := bytes.NewReader(data)
	fs, err := zipfs.NewFromReaderAt(reader, int64(len(data)), nil)
	if err != nil {
		return nil, err
	}
	return fs, nil
}

func (n *Node) LoadPageByAddress(address string) (*zipfs.FileSystem, error) {
	_, account, err := n.accountService.NewAccountFromAddress(address)
	if err != nil {
		return nil, err
	}

	entry, err := n.GetEntry(nil)
	if err != nil {
		return nil, err
	}
	ss := entry.storage

	data, err := ss.ReadPageByAddress(account)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(data)
	fs, err := zipfs.NewFromReaderAt(reader, int64(len(data)), nil)
	if err != nil {
		return nil, err
	}
	return fs, nil
}

func (n *Node) start() error {
	readyC := make(chan struct{})

	var err error
	n.net, err = p2p.InitNet(n.config, readyC)
	if err != nil {
		return err
	}
	err = n.net.Start()
	if err != nil {
		return err
	}
	err = n.net.InitPubSub(n.config.GetChainId(), 0)
	if err != nil {
		return err
	}

	err = n.net.SubscribeWithChainId(n.config.GetChainId(), TOPIC_PEER_DISCOVERY, n.discoveryHandler)
	if err != nil {
		return err
	}
	err = n.net.DirectMsgHandle(n.config.GetChainId(), TOPIC_PEER_MESSAGE, n.messageHandler)
	if err != nil {
		return err
	}
	err = n.net.DirectMsgHandle(n.config.GetChainId(), TOPIC_PEER_DATA, n.dataHandler)
	if err != nil {
		return err
	}
	close(readyC)

	n.self.Id = n.config.GetNodeId()
	return nil
}

func (n *Node) Start() error {
	if err := n.start(); err != nil {
		return err
	}

	go n.send()
	go n.receive()

	go n.connect()
	go n.generate()
	go n.discovery()

	return nil
}

func (n *Node) Close() error {
	for _, entry := range n.entryMap {
		if entry.merkle != nil {
			if err := entry.merkle.Close(); err != nil {
				return err
			}
		}
		// if entry.storage != nil {
		// 	if err := entry.storage.Close(); err != nil {
		// 		return err
		// 	}
		// }
	}
	if n.storageService != nil {
		if err := n.storageService.Close(); err != nil {
			return err
		}
	}
	if n.net != nil {
		if err := n.net.CancelSubscribeWithChainId(n.config.GetChainId(), TOPIC_PEER_DISCOVERY); err != nil {
			return err
		}
		if err := n.net.CancelDirectMsgHandle(n.config.GetChainId(), TOPIC_PEER_MESSAGE); err != nil {
			return err
		}
		if err := n.net.CancelDirectMsgHandle(n.config.GetChainId(), TOPIC_PEER_DATA); err != nil {
			return err
		}
		if err := n.net.Stop(); err != nil {
			return err
		}
	}
	return nil
}

func (n *Node) discovery() {
	for {
		peers, err := n.net.ChainNodesInfo(n.config.GetChainId())
		if err != nil {
			glog.Error(err)
		} else {
			peerMap := make(map[string]struct{})
			for _, p := range peers {
				peerMap[p.NodeUid] = struct{}{}
			}
			for _, p := range n.peers {
				_, ok := peerMap[p.Id]
				if !ok {
					n.RemovePeer(p)
				}
			}
		}

		pubKeyData, err := n.self.Key.MarshalBinary()
		if err != nil {
			glog.Error(err)
		} else {
			for root := range n.entryMap {
				_, rootAccount, _ := n.accountService.NewAccountFromAddress(root)
				rootData, err := block.AddressToByte(rootAccount)
				if err != nil {
					glog.Error(err)
				} else {
					peers := n.ListPeerBy(rootAccount)
					e := n.self.getPeerEntry(rootAccount)
					e.PeerCount = int64(len(peers))

					entry, err := n.GetEntry(rootAccount)
					if err != nil {
						glog.Error(err)
					} else {
						peerData, err := core.Marshal(&pb.PeerInfo{
							Account:     rootData,
							PublicKey:   pubKeyData,
							BlockNumber: entry.GetBlockNumber(),
							PeerCount:   int64(len(peers)),
						})
						if err != nil {
							glog.Error(err)
						} else {
							_, msgData, err := n.EncodeMessage(peerData)
							if err != nil {
								glog.Error(err)
							} else {
								err = n.net.BroadcastWithChainId(n.config.GetChainId(), TOPIC_PEER_DISCOVERY, msgData)
								if err != nil {
									glog.Error(err)
								}
							}
						}
					}
				}
			}
		}

		time.Sleep(10 * time.Second)
	}
}

func (n *Node) GetNodeKey() libaccount.Key {
	return n.key
}

func (n *Node) GetPeer(index uint64) *Peer {
	n.peerLocker.Lock()
	defer n.peerLocker.Unlock()

	return n.peers[index]
}

func (n *Node) AddPeer(p *Peer) error {
	n.peerLocker.Lock()
	defer n.peerLocker.Unlock()

	_, ok := n.peers[p.GetIndex(n)]
	if !ok {
		n.peers[p.GetIndex(n)] = p
		go n.discoveryPeer(p)
	}
	return nil
}

func (n *Node) RemovePeer(p *Peer) error {
	n.peerLocker.Lock()
	defer n.peerLocker.Unlock()

	delete(n.peers, p.GetIndex(n))
	return nil
}

func (n *Node) UpdatePeer(p *Peer, info *pb.PeerInfo) error {
	n.peerLocker.Lock()
	defer n.peerLocker.Unlock()

	a, err := block.ByteToAddress(info.Account)
	if err != nil {
		return err
	}
	entry := p.getPeerEntry(a)
	entry.BlockNumber = info.BlockNumber
	entry.PeerCount = info.PeerCount

	return nil
}

func (n *Node) UpdateStatus(p *Peer, rootAccount libcore.Address, s Status) error {
	n.peerLocker.Lock()
	defer n.peerLocker.Unlock()

	entry := p.getPeerEntry(rootAccount)
	entry.Status = s

	return nil
}

func (n *Node) UpdateSafeStatus(p *Peer, rootAccount libcore.Address, s Status) error {
	n.peerLocker.Lock()
	defer n.peerLocker.Unlock()

	entry := p.getPeerEntry(rootAccount)
	if entry.Status < s {
		entry.Status = s
	}
	return nil
}

func (n *Node) ListPeer() []*Peer {
	n.peerLocker.RLock()
	defer n.peerLocker.RUnlock()

	list := make([]*Peer, 0)
	for _, p := range n.peers {
		list = append(list, p)
	}
	return list
}

func (n *Node) ListPeerBy(rootAccount libcore.Address) []*Peer {
	n.peerLocker.RLock()
	defer n.peerLocker.RUnlock()

	root := util.GetString(rootAccount)
	list := make([]*Peer, 0)
	for _, p := range n.peers {
		_, ok := p.peermap[root]
		if ok {
			list = append(list, p)
		}
	}
	return list
}

func (n *Node) CreateMessage(data []byte) (*pb.Message, error) {
	return &pb.Message{
		Id:   n.newId(),
		Data: data,
		Node: n.self.GetIndex(n),
	}, nil
}

func (n *Node) EncodeMessage(data []byte) (uint64, []byte, error) {
	m, err := n.CreateMessage(data)
	if err != nil {
		return 0, nil, err
	}
	msgData, err := core.Marshal(m)
	if err != nil {
		return 0, nil, err
	}
	return m.Id, msgData, nil
}

func (n *Node) DecodeMessage(data []byte) (*pb.Message, error) {
	meta, m, err := core.Unmarshal(data)
	if err != nil {
		return nil, err
	}
	if meta != core.CORE_MESSAGE {
		return nil, util.ErrorOfInvalid("format", "message")
	}
	return m.(*pb.Message), nil
}

func (n *Node) ReceiveMessage(t string, msg *pb.Message) (core.DataType, *Peer, []byte) {
	fromIndex := msg.GetNode()
	data := msg.GetData()
	glog.Infof("<<< receive %s from node %d, length: %d\n", t, fromIndex, len(data))
	fromPeer := n.GetPeer(fromIndex)
	if fromPeer != nil {
		glog.Infof("<<< receive peer %s %d(%s) from %s(%s)\n", t, msg.Id, core.GetInfo(data), fromPeer.GetAddress(n), fromPeer.Id)
		if len(data) > 0 {
			meta := core.DataType(data[0])
			return meta, fromPeer, data
		}
	}
	return 0, nil, nil
}

func (n *Node) HashBlock(b libblock.Block) (libcore.Hash, error) {
	return n.consensusService.HashBlock(b)
}
