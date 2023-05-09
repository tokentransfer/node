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
	"github.com/tokentransfer/node/chunk"
	"github.com/tokentransfer/node/config"
	"github.com/tokentransfer/node/core"
	"github.com/tokentransfer/node/core/pb"
	"github.com/tokentransfer/node/crypto"
	"github.com/tokentransfer/node/p2p"
	"github.com/tokentransfer/node/util"

	libaccount "github.com/tokentransfer/interfaces/account"
	libblock "github.com/tokentransfer/interfaces/block"
	libcore "github.com/tokentransfer/interfaces/core"
	libcrypto "github.com/tokentransfer/interfaces/crypto"
)

type Status int

const (
	PeerNone Status = iota
	PeerConnected
	PeerKnown
	PeerConsensused
)

func (s Status) String() string {
	switch s {
	case PeerNone:
		return "none"
	case PeerConnected:
		return "connected"
	case PeerKnown:
		return "known"
	case PeerConsensused:
		return "consensused"
	}
	return "unknown status"
}

const (
	BLOCK_DURATION       = 10
	TOPIC_PEER_DISCOVERY = "peer_discovery"
	TOPIC_PEER_DATA      = "peer_data"
	TOPIC_PEER_MESSAGE   = "peer_message"
)

type Peer struct {
	Id          string
	Key         libaccount.PublicKey
	BlockNumber int64
	PeerCount   int64
	Status      Status

	address string
	index   uint64
}

func (p *Peer) GetPublicKey() libaccount.PublicKey {
	return p.Key
}

func (p *Peer) GetIndex(n *Node) uint64 {
	if p.index == 0 {
		address, err := p.Key.GenerateAddress()
		if err != nil {
			return 0
		}
		index, err := n.GetIndex(address.String())
		if err != nil {
			return 0
		}

		p.address = address.String()
		p.index = index
	}
	return p.index
}

func (p *Peer) GetAddress(n *Node) string {
	if len(p.address) == 0 {
		address, err := p.Key.GenerateAddress()
		if err != nil {
			return ""
		}
		index, err := n.GetIndex(address.String())
		if err != nil {
			return ""
		}

		p.address = address.String()
		p.index = index
	}
	return p.address
}

type transactionEntry struct {
	ValidatedBlock libblock.Block
	txlist         []libblock.TransactionWithData
	txmap          map[string]libblock.TransactionWithData
}

type Node struct {
	Consensused bool

	messageId uint64

	cryptoService    libcrypto.CryptoService
	accountService   libaccount.AccountService
	consensusService *ConsensusService
	storageService   *StorageService

	transactionMap    map[libcore.Address]*transactionEntry
	transactionLocker *sync.Mutex

	net protocol.Net

	outs chan *pb.Message
	ins  chan *pb.Message

	ready *sync.WaitGroup

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

		ready: ready,

		transactionMap:    make(map[libcore.Address]*transactionEntry),
		transactionLocker: &sync.Mutex{},

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
		Key:    pubKey,
		Status: PeerConsensused,
	}

	storageService, err := NewStorageService(n.config, n.cryptoService, n.accountService)
	if err != nil {
		return err
	}
	consensusService := &ConsensusService{
		CryptoService:  n.cryptoService,
		Config:         n.config,
		AccountService: n.accountService,
		StorageService: storageService,
	}
	n.consensusService = consensusService
	n.storageService = storageService

	return nil
}

func (n *Node) getContractData(rootAccount libcore.Address, txm map[string]interface{}) (int64, interface{}, error) {
	as := n.accountService
	ss := n.storageService

	from := util.ToString(&txm, "from")
	to := util.ToString(&txm, "to")

	cc := ss.GetChunkService(rootAccount)

	_, fromAccount, err := as.NewAccountFromAddress(from)
	if err != nil {
		return 0, nil, err
	}

	_, toAccount, err := as.NewAccountFromAddress(to)
	if err != nil {
		return 0, nil, err
	}

	format := util.ToString(&txm, "format")
	usedCost, r, err := cc.GetContractData(fromAccount, fromAccount, toAccount, format)
	if err != nil {
		return 0, nil, err
	}
	return usedCost, r, nil
}

func (n *Node) callContract(rootAccount libcore.Address, txm map[string]interface{}) (int64, interface{}, error) {
	as := n.accountService
	ss := n.storageService

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

	usedCost, r, err := ss.GetChunkService(rootAccount).CallContract(fromAccount, fromAccount, toAccount, method, params)
	if err != nil {
		return 0, nil, err
	}
	return usedCost, r, nil
}

func (n *Node) signTransaction(rootAccount libcore.Address, txm map[string]interface{}) (string, *block.Transaction, error) {
	as := n.accountService

	from := util.ToString(&txm, "from")
	secret := util.ToString(&txm, "secret")
	to := util.ToString(&txm, "to")
	gas := util.ToUint64(&txm, "gas")

	var account libcore.Address
	if util.Has(&txm, "account") {
		accountString := util.ToString(&txm, "account")
		_, a, err := as.NewAccountFromAddress(accountString)
		if err != nil {
			return "", nil, err
		}
		account = a
	} else {
		account = nil
	}

	_, fromKey, err := as.NewKeyFromSecret(secret)
	if err != nil {
		return "", nil, err
	}
	fromAccount, err := fromKey.GetAddress()
	if err != nil {
		return "", nil, err
	}
	fromAddress := fromAccount.String()
	if fromAddress != from {
		return "", nil, util.ErrorOfUnmatched("account", "in secret", from, fromAddress)
	}
	_, toAccount, err := as.NewAccountFromAddress(to)
	if err != nil {
		return "", nil, err
	}
	seq := n.getNextSequence(rootAccount, fromAccount)

	tx := &block.Transaction{
		TransactionType: block.TRANSACTION,

		From:     fromAccount,
		Sequence: seq,
		To:       toAccount,

		Account: account,
		Gas:     gas,
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
	return nil, util.ErrorOfNotFound("params", "map")
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

	ss := n.storageService
	cc := ss.GetChunkService(rootAccount)

	h, _, err := n.cryptoService.Raw(tx, libcrypto.RawBinary)
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

	err = cc.CreateSandbox()
	if err != nil {
		return nil, nil, err
	}
	defer cc.CancelSandbox()
	txWithData, err := n.consensusService.ProcessTransaction(rootAccount, tx)
	if err != nil {
		return nil, nil, err
	}

	// util.PrintJSON("txWithData", txWithData)
	_, ok = n.AddTransaction(rootAccount, txWithData)
	if ok {
		glog.Infoln(">>> receive transaction", h.String())
	} else {
		glog.Infoln(">>> drop transaction", h.String())
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
	as := n.accountService
	ss := n.storageService

	switch method {
	case "blockNumber":
		result, err := n._call(params, func(rootAccount libcore.Address, item map[string]interface{}) (interface{}, error) {
			return n.GetBlockNumber(rootAccount), nil
		})
		if err != nil {
			return nil, err
		}
		return result, nil

	case "createWallet":
		result, err := n._call(params, func(rootAccount libcore.Address, item map[string]interface{}) (interface{}, error) {
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
			address := util.ToString(&item, "address")
			_, a, err := as.NewAccountFromAddress(address)
			if err != nil {
				return nil, err
			}
			gas, err := ss.GetChunkService(rootAccount).GetGas(rootAccount, a)
			if err != nil {
				return nil, err
			}
			return gas.String(), nil
		})
		if err != nil {
			return nil, err
		}
		return result, nil

	case "getTransactionCount":
		result, err := n._call(params, func(rootAccount libcore.Address, item map[string]interface{}) (interface{}, error) {
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
			hashString := util.ToString(&item, "hash")
			h, err := hex.DecodeString(hashString)
			if err != nil {
				return nil, err
			}
			txWithData, err := ss.GetMerkleService(rootAccount).GetTransactionByHash(libcore.Hash(h))
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
			hashString := util.ToString(&item, "hash")
			h, err := hex.DecodeString(hashString)
			if err != nil {
				return nil, err
			}
			txWithData, err := ss.GetMerkleService(rootAccount).GetTransactionByHash(libcore.Hash(h))
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
			address := util.ToString(&item, "address")
			_, a, err := as.NewAccountFromAddress(address)
			if err != nil {
				return nil, err
			}
			index := util.AsUint64(&item, "index")

			txWithData, err := ss.GetMerkleService(rootAccount).GetTransactionByIndex(a, index)
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
			hashString := util.ToString(&item, "hash")
			h, err := hex.DecodeString(hashString)
			if err != nil {
				return nil, err
			}
			block, err := ss.GetMerkleService(rootAccount).GetBlockByHash(libcore.Hash(h))
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
			index := util.AsUint64(&item, "num")
			block, err := ss.GetMerkleService(rootAccount).GetBlockByIndex(index)
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
			hashString := util.ToString(&item, "hash")
			h, err := hex.DecodeString(hashString)
			if err != nil {
				return nil, err
			}
			format := util.ToString(&item, "format")
			usedCost, r, err := ss.GetChunkService(rootAccount).GetData(libcore.Hash(h), format)
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
			info, err := ss.GetChunkService(rootAccount).ReadUser(a, a, b)
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
			return hash, nil
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
			hashString := util.ToString(&item, "hash")
			h, err := hex.DecodeString(hashString)
			if err != nil {
				return nil, err
			}
			s, err := ss.GetMerkleService(rootAccount).GetStateByHash(libcore.Hash(h))
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
			if util.IsTest() || util.IsDebug() {
				ss.GetChunkService(rootAccount).Dump(chunk.LogPrinter{})
			}
			return ss.GetChunkService(rootAccount).RootHash(), nil
		})
		if err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, util.ErrorOfNotFound("method", method)
	}
}

func (n *Node) getTransactionEntry(rootAccount libcore.Address) *transactionEntry {
	entry, ok := n.transactionMap[rootAccount]
	if !ok {
		entry = &transactionEntry{
			txlist: make([]libblock.TransactionWithData, 0),
			txmap:  make(map[string]libblock.TransactionWithData),
		}
		n.transactionMap[rootAccount] = entry
	}
	return entry
}

func (n *Node) AddTransaction(rootAccount libcore.Address, txWithData libblock.TransactionWithData) (libcore.Hash, bool) {
	n.transactionLocker.Lock()
	defer n.transactionLocker.Unlock()

	_, err := n.consensusService.HashTransaction(txWithData)
	if err != nil {
		return nil, false
	}

	entry := n.getTransactionEntry(rootAccount)

	h := txWithData.GetTransaction().GetHash()
	_, ok := entry.txmap[h.String()]
	if ok {
		return nil, false
	}

	entry.txlist = append(entry.txlist, txWithData)
	entry.txmap[h.String()] = txWithData
	return h, true
}

func (n *Node) ClearTransaction(rootAccount libcore.Address, b libblock.Block) {
	n.transactionLocker.Lock()
	defer n.transactionLocker.Unlock()

	entry := n.getTransactionEntry(rootAccount)

	txlist := make([]libblock.TransactionWithData, 0)
	txmap := make(map[string]libblock.TransactionWithData)
	for i := 0; i < len(entry.txlist); i++ {
		txWithData := entry.txlist[i]
		h := txWithData.GetTransaction().GetHash()
		_, ok := entry.txmap[h.String()]
		if ok {
			glog.Infoln(">>> drop block transaction", h.String())
		} else {
			txlist = append(txlist, txWithData)
			txmap[h.String()] = txWithData
		}
	}
	entry.txlist = txlist
	entry.txmap = txmap
}

func (n *Node) _generateBlock(rootAccount libcore.Address, list []libblock.TransactionWithData, entry *transactionEntry) (libblock.Block, error) {
	config := n.config
	cs := n.cryptoService
	ss := n.storageService
	cc := ss.GetChunkService(rootAccount)
	ms := ss.GetMerkleService(rootAccount)

	v := entry.ValidatedBlock

	var b *block.Block
	if v == nil { //genesis
		if len(list) > 0 {
			return nil, util.ErrorOfInvalid("transactions", "genesis block")
		}
		var states []libblock.State
		if rootAccount != nil {
			v, err := util.NewValue("100000000000000")
			if err != nil {
				return nil, err
			}
			err = cc.UpdateGas(rootAccount, rootAccount, *v)
			if err != nil {
				return nil, err
			}
			gasAccount := config.GetGasAccount()
			err = cc.UpdateGas(rootAccount, gasAccount, v.ZeroClone())
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
			Account:    rootAccount,
			BlockIndex: uint64(0),
			ParentHash: libcrypto.ZeroHash(cs),

			Transactions:    []libblock.TransactionWithData{},
			TransactionHash: ms.GetTransactionRoot(),

			States:    states,
			StateHash: ms.GetStateRoot(),

			RootHash: cc.RootHash(),

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
			Account:    rootAccount,
			BlockIndex: v.GetIndex() + 1,
			ParentHash: v.GetHash(),

			Transactions:    list,
			TransactionHash: ms.GetTransactionRoot(),

			States:    states,
			StateHash: ms.GetStateRoot(),

			RootHash: cc.RootHash(),

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

func (n *Node) AddBlock(rootAccount libcore.Address, b libblock.Block) error {
	n.transactionLocker.Lock()
	defer n.transactionLocker.Unlock()

	entry := n.getTransactionEntry(rootAccount)

	ss := n.storageService
	cc := ss.GetChunkService(rootAccount)
	ms := ss.GetMerkleService(rootAccount)

	err := cc.CommitSandbox()
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

	return nil
}

func (n *Node) VerifyBlock(rootAccount libcore.Address, b libblock.Block) (ok bool, err error) {
	n.transactionLocker.Lock()
	defer n.transactionLocker.Unlock()

	entry := n.getTransactionEntry(rootAccount)

	cs := n.cryptoService
	ss := n.storageService
	cc := ss.GetChunkService(rootAccount)
	ms := ss.GetMerkleService(rootAccount)

	ok = true
	err = nil

	defer func() {
		if !ok || err != nil {
			cc.CancelSandbox()
			ms.Cancel()
		}
	}()

	err = cc.CreateSandbox()
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

		err = cc.CreateSandbox()
		if err != nil {
			return
		}
		newWithData, e := n.consensusService.ProcessTransaction(rootAccount, tx)
		if e != nil {
			err = cc.CancelSandbox()
			if err != nil {
				glog.Error(err)
			}
			ok = false
			err = e
			return
		}
		err = cc.CommitSandbox()
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

func (n *Node) GenerateBlock(rootAccount libcore.Address) (libblock.Block, error) {
	n.transactionLocker.Lock()
	defer n.transactionLocker.Unlock()

	entry := n.getTransactionEntry(rootAccount)

	list := entry.txlist
	entry.txlist = make([]libblock.TransactionWithData, 0)
	entry.txmap = make(map[string]libblock.TransactionWithData)

	return n._generateBlock(rootAccount, list, entry)
}

func (n *Node) generate() {
	n.ready.Wait()

	for {
		hasTx := false
		for rootAccount, entry := range n.transactionMap {
			if len(entry.txlist) > 0 {
				hasTx = hasTx || true

				block, err := n.GenerateBlock(rootAccount)
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
			entry := n.getTransactionEntry(nil)
			glog.Infof("=== block %d, %s, prepare %d, %d\n", n.GetBlockNumber(nil), n.GetBlockHash(nil), n.GetBlockNumber(nil)+1, len(entry.txlist))

			for i := 0; i < int(n.config.GetBlockDuration()*5); i++ {
				hasTx := false
				for _, entry := range n.transactionMap {
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
	lastConsensused := n.Consensused
	for {
		list := n.ListPeer()

		glog.Infoln("===", 0, n.self.GetIndex(n), n.self.GetAddress(n), n.self.Id, n.self.Status, n.GetBlockNumber(nil), n.self.PeerCount)
		for i := 0; i < len(list); i++ {
			p := list[i]
			glog.Infoln("==>", i+1, p.GetIndex(n), p.GetAddress(n), p.Id, p.Status, p.BlockNumber, p.PeerCount)
		}
		if len(list) == 0 {
			n.Consensused = n.PrepareConsensus()
		}
		if !lastConsensused && n.Consensused {
			lastConsensused = n.Consensused
			n.ready.Done()
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

func (n *Node) PrepareConsensus() bool {
	list := n.ListPeer()
	count := int64(0)
	currentBlock := n.GetBlockNumber(nil)
	currentCount := n.self.PeerCount
	for i := 0; i < len(list); i++ {
		p := list[i]
		if p.Status >= PeerKnown && p.BlockNumber == currentBlock && (p.PeerCount == currentCount) {
			count++
		}
	}
	l := len(n.config.GetBootstraps())
	if l == 0 {
		return true
	}
	if count > 0 && (count+1) == currentCount {
		for i := 0; i < len(list); i++ {
			p := list[i]
			p.Status = PeerConsensused
		}
		return true
	}
	return false
}

func (n *Node) discoveryPeer(p *Peer) {
	config := n.config

	lastSendBlock := int64(-1)
	lastSendTime := time.Now()
	for {
		switch p.Status {
		case PeerNone:
			n.ConnectTo(p)
		case PeerConnected:
			n.SendRequestInfo(p)
		case PeerKnown:
			n.Consensused = n.PrepareConsensus()
		}
		if p.Status >= PeerKnown && n.GetBlockNumber(nil) > p.BlockNumber {
			glog.Infoln("===", n.GetBlockNumber(nil), lastSendBlock, lastSendTime, p.address, p.Id, p.Status, p.BlockNumber, p.PeerCount)
			if lastSendBlock >= 0 && lastSendBlock == (p.BlockNumber+1) && (time.Since(lastSendTime) < (time.Duration(int64(config.GetBlockDuration()) * int64(time.Second)))) {
				time.Sleep(3 * time.Second)
			} else {
				block, err := n.storageService.GetMerkleService(nil).GetBlockByIndex(uint64(p.BlockNumber + 1))
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
									lastSendBlock = int64(block.GetIndex())
									lastSendTime = time.Now()
									glog.Infof(">>> send data %d(%s) to %s(%s)\n", mid, core.GetInfo(blockData), p.GetAddress(n), p.Id)
								}
							}
						}
					}
				}
				time.Sleep(time.Second)
			}
		} else {
			time.Sleep(time.Duration(config.GetBlockDuration()) * time.Second)
		}
	}
}

func (n *Node) Load(rootAccount libcore.Address) error {
	n.transactionLocker.Lock()
	defer n.transactionLocker.Unlock()

	entry := n.getTransactionEntry(rootAccount)

	ss := n.storageService
	ms := ss.GetMerkleService(rootAccount)

	current := uint64(0)
	for {
		b, err := ms.GetBlockByIndex(current)
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
			if !b.GetParentHash().IsZero() {
				return util.ErrorOfInvalid("the parent hash of the genesis block", b.GetParentHash().String())
			}
		}
		entry.ValidatedBlock = b

		allData, err := b.MarshalBinary()
		if err != nil {
			glog.Error(err)
			break
		}
		ignoreData, err := b.Raw(false)
		if err != nil {
			glog.Error(err)
			break
		}
		signData, err := b.Raw(true)
		if err != nil {
			glog.Error(err)
			break
		}
		glog.Infoln("load block", current, b.GetHash().String(), len(allData), len(ignoreData), len(signData))

		current++
	}
	return nil
}

func (n *Node) Dump(rootAccount libcore.Address, printer core.Printer) {
	ss := n.storageService
	cc := ss.GetChunkService(rootAccount)

	cc.Dump(printer)
}

func (n *Node) GetBlockNumber(rootAccount libcore.Address) int64 {
	entry := n.getTransactionEntry(rootAccount)
	if entry.ValidatedBlock != nil {
		return int64(entry.ValidatedBlock.GetIndex())
	}
	return -1
}

func (n *Node) GetBlockHash(rootAccount libcore.Address) string {
	entry := n.getTransactionEntry(rootAccount)
	if entry.ValidatedBlock != nil {
		return entry.ValidatedBlock.GetHash().String()
	}
	return ""
}

func (n *Node) GetBlock(rootAccount libcore.Address) libblock.Block {
	entry := n.getTransactionEntry(rootAccount)
	if entry.ValidatedBlock != nil {
		return entry.ValidatedBlock
	}
	return nil
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
						if int64(b.GetIndex()) <= n.GetBlockNumber(rootAccount) {
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
					}
				}
				n.SendPeerInfo(fromPeer)

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
		p := n.GetPeer(index)
		if p == nil {
			p = &Peer{
				Id:          publisher,
				Key:         publicKey,
				Status:      PeerKnown,
				BlockNumber: peerInfo.BlockNumber,
				PeerCount:   peerInfo.PeerCount,
			}
			n.AddPeer(p)
		} else {
			if p.Status < PeerKnown {
				p.Status = PeerKnown
			}
			p.BlockNumber = peerInfo.BlockNumber
			p.PeerCount = peerInfo.PeerCount
		}
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
			fromPeer.BlockNumber = peerInfo.BlockNumber
			fromPeer.PeerCount = peerInfo.PeerCount

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
					if int64(b.GetIndex()) <= n.GetBlockNumber(rootAccount) {
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
				}
			}
			n.SendPeerInfo(fromPeer)
		}
	}
	return nil
}

func (n *Node) SendPeerInfo(toPeer *Peer) {
	peers, _ := n.net.ChainNodesInfo(n.config.GetChainId())
	peerData, err := core.Marshal(&pb.PeerInfo{
		BlockNumber: n.GetBlockNumber(nil),
		PeerCount:   int64(len(peers)),
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

func (n *Node) LoadPageByName(name string) (*zipfs.FileSystem, error) {
	if len(name) == 0 {
		return nil, util.ErrorOfInvalid("name", name)
	}
	_, account, _ := n.accountService.NewAccountFromAddress(name)
	if account != nil {
		return nil, util.ErrorOfInvalid("name", name)
	}
	data, err := n.storageService.GetChunkService(nil).ReadPageByName(name)
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
	var data []byte
	_, account, err := n.accountService.NewAccountFromAddress(nameOrAddress)
	if err != nil {
		data, err = n.storageService.GetChunkService(nil).ReadPageByAddress(account)
		if err != nil {
			return nil, err
		}
	} else {
		data, err = n.storageService.GetChunkService(nil).ReadPageByName(nameOrAddress)
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
	data, err := n.storageService.GetChunkService(nil).ReadPageByAddress(account)
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
	if err := n.Load(nil); err != nil {
		return err
	}
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

func (n *Node) Stop() error {
	if err := n.storageService.Close(); err != nil {
		return err
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
		pubKeyData, err := n.self.Key.MarshalBinary()
		if err != nil {
			glog.Error(err)
		} else {
			peers, _ := n.net.ChainNodesInfo(n.config.GetChainId())
			n.self.PeerCount = int64(len(peers))

			peerData, err := core.Marshal(&pb.PeerInfo{
				PublicKey:   pubKeyData,
				BlockNumber: n.GetBlockNumber(nil),
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

		if n.Consensused {
			time.Sleep(60 * time.Second)
		} else {
			time.Sleep(10 * time.Second)
		}
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

func (n *Node) ListPeer() []*Peer {
	n.peerLocker.RLock()
	defer n.peerLocker.RUnlock()

	list := make([]*Peer, 0)
	for _, p := range n.peers {
		list = append(list, p)
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
