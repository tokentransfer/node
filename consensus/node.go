package consensus

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
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
	"github.com/tokentransfer/node/merkle"
	"github.com/tokentransfer/node/p2p"
	"github.com/tokentransfer/node/util"

	libaccount "github.com/tokentransfer/interfaces/account"
	libblock "github.com/tokentransfer/interfaces/block"
	libcore "github.com/tokentransfer/interfaces/core"
	libcrypto "github.com/tokentransfer/interfaces/crypto"
	libstore "github.com/tokentransfer/interfaces/store"
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

func (p *Peer) GetIndex() uint64 {
	if p.index == 0 {
		address, err := p.Key.GenerateAddress()
		if err != nil {
			return 0
		}
		index, err := core.GetIndex(address.String())
		if err != nil {
			return 0
		}

		p.address = address.String()
		p.index = index
	}
	return p.index
}

func (p *Peer) GetAddress() string {
	if len(p.address) == 0 {
		address, err := p.Key.GenerateAddress()
		if err != nil {
			return ""
		}
		index, err := core.GetIndex(address.String())
		if err != nil {
			return ""
		}

		p.address = address.String()
		p.index = index
	}
	return p.address
}

type Node struct {
	Consensused bool

	messageId uint64

	cryptoService    libcrypto.CryptoService
	accountService   libaccount.AccountService
	merkleService    libstore.MerkleService
	consensusService *ConsensusService
	storageService   *StorageService

	transactions      []libblock.TransactionWithData
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

		transactionLocker: &sync.Mutex{},

		accountService: account.NewAccountService(),
		cryptoService:  &crypto.CryptoService{},
	}
}

func (n *Node) Init(c libcore.Config) error {
	n.config = c.(*config.Config)
	core.Init(n.config)
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

	merkleService, err := merkle.NewMerkleService(n.config, n.cryptoService)
	if err != nil {
		return err
	}
	storageService, err := NewStorageService(n.config)
	if err != nil {
		return err
	}
	consensusService := &ConsensusService{
		CryptoService:  n.cryptoService,
		MerkleService:  merkleService,
		Config:         n.config,
		AccountService: n.accountService,
		StorageService: storageService,
	}
	n.merkleService = merkleService
	n.consensusService = consensusService
	n.storageService = storageService

	return nil
}

func (n *Node) signTransaction(txm map[string]interface{}) (string, *block.Transaction, error) {
	as := n.accountService

	from := util.ToString(&txm, "from")
	secret := util.ToString(&txm, "secret")
	to := util.ToString(&txm, "to")
	value := util.ToString(&txm, "value")
	gas := util.ToInt64(&txm, "gas")

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
		return "", nil, fmt.Errorf("error account: %s != %s", fromAddress, from)
	}
	_, toAccount, err := as.NewAccountFromAddress(to)
	if err != nil {
		return "", nil, err
	}
	amount, err := core.NewAmount(value)
	if err != nil {
		return "", nil, err
	}
	seq := n.getNextSequence(fromAccount)
	tx := &block.Transaction{
		TransactionType: libblock.TransactionType(1),

		Account:     fromAccount,
		Sequence:    seq,
		Amount:      *amount,
		Gas:         gas,
		Destination: toAccount,
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
				if params := util.ToString(&cm, "params"); len(params) > 0 {
					list := make([][]byte, 0)
					ps := strings.Split(params, ",")
					for i := 0; i < len(ps); i++ {
						items := strings.Split(ps[i], ":")
						if len(items) == 2 {
							t, err := core.GetDataTypeByName(items[1])
							if err != nil {
								return "", nil, err
							}
							data, err := t.FromString(items[0])
							if err != nil {
								return "", nil, err
							}
							list = append(list, data)
						} else {
							return "", nil, util.ErrorOfInvalid("parameter", params)
						}
					}
					contractInfo.Params = list
				}
				if params := util.ToArray(&cm, "params"); params != nil {
					list := make([][]byte, 0)
					for i := 0; i < len(params); i++ {
						item := params[i]
						data, err := core.MarshalData(item)
						if err != nil {
							return "", nil, err
						}
						list = append(list, data)
					}
					contractInfo.Params = list
				}
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
	err = n.verifyTransaction(tx)
	if err != nil {
		return "", nil, err
	}
	blob := libcore.Bytes(data).String()
	return blob, tx, nil
}

func (n *Node) verifyTransaction(tx libblock.Transaction) error {
	_, _, err := n.cryptoService.Raw(tx, libcrypto.RawBinary)
	if err != nil {
		return err
	}
	_, err = n.consensusService.VerifyTransaction(tx)
	if err != nil {
		return err
	}
	return nil
}

func (n *Node) sendTransaction(tx libblock.Transaction) (libblock.TransactionWithData, error) {
	err := n.verifyTransaction(tx)
	if err != nil {
		return nil, err
	}
	hash := tx.GetHash()
	glog.Infoln("verify transaction", hash.String())

	txWithData, _, err := n.processTransaction(tx)
	if err != nil {
		return nil, err
	}
	return txWithData, nil
}

func (n *Node) processTransaction(tx libblock.Transaction) (libblock.TransactionWithData, libblock.Transaction, error) {
	n.transactionLocker.Lock()
	defer n.transactionLocker.Unlock()

	h, _, err := n.cryptoService.Raw(tx, libcrypto.RawBinary)
	if err != nil {
		return nil, nil, err
	}
	ok, err := n.consensusService.VerifyTransaction(tx)
	if err != nil {
		return nil, nil, err
	}
	if !ok {
		return nil, nil, errors.New("error transaction")
	}

	err = n.storageService.CreateSandbox()
	if err != nil {
		return nil, nil, err
	}
	defer n.storageService.CancelSandbox()
	txWithData, err := n.consensusService.ProcessTransaction(tx)
	if err != nil {
		return nil, nil, err
	}

	// util.PrintJSON("txWithData", txWithData)
	_, ok = n.AddTransaction(txWithData)
	if ok {
		glog.Infoln(">>> receive transaction", h.String())
	} else {
		glog.Infoln(">>> drop transaction", h.String())
	}

	return txWithData, tx, nil
}

func (n *Node) getNextSequence(address libcore.Address) uint64 {
	accountEntry, err := n.consensusService.GetAccountInfo(address)
	if err != nil {
		return uint64(1)
	}
	return accountEntry.Sequence + 1
}

func (n *Node) Call(method string, params []interface{}) (interface{}, error) {
	as := n.accountService

	switch method {
	case "blockNumber":
		result := n.consensusService.GetBlockNumber()
		return result, nil

	case "createWallet":
		item := params[0].(map[string]interface{})
		t := util.ToString(&item, "type")
		p := util.ToString(&item, "password")
		tt, k, err := n.accountService.GenerateFamilySeed(t + "." + p)
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

	case "getBalance":
		item := params[0].(map[string]interface{})
		address := util.ToString(&item, "address")
		_, a, err := as.NewAccountFromAddress(address)
		if err != nil {
			return nil, err
		}
		accountEntry, err := n.consensusService.GetAccountInfo(a)
		if err != nil {
			return nil, err
		}
		return accountEntry.Amount, nil

	case "getTransactionCount":
		item := params[0].(map[string]interface{})
		address := util.ToString(&item, "address")
		_, a, err := as.NewAccountFromAddress(address)
		if err != nil {
			return nil, err
		}
		seq := n.getNextSequence(a)
		return seq, nil

	case "getTransactionReceipt":
		item := params[0].(map[string]interface{})
		hashString := util.ToString(&item, "hash")
		h, err := hex.DecodeString(hashString)
		if err != nil {
			return nil, err
		}
		txWithData, err := n.merkleService.GetTransactionByHash(libcore.Hash(h))
		if err != nil {
			return nil, err
		}
		receipt := txWithData.GetReceipt()
		_, err = n.consensusService.HashReceipt(receipt)
		if err != nil {
			return nil, err
		}
		return receipt, nil

	case "getTransactionByHash":
		item := params[0].(map[string]interface{})
		hashString := util.ToString(&item, "hash")
		h, err := hex.DecodeString(hashString)
		if err != nil {
			return nil, err
		}
		txWithData, err := n.merkleService.GetTransactionByHash(libcore.Hash(h))
		if err != nil {
			return nil, err
		}
		_, err = n.consensusService.HashTransaction(txWithData)
		if err != nil {
			return nil, err
		}
		return txWithData, nil

	case "getTransactionByIndex":
		item := params[0].(map[string]interface{})
		address := util.ToString(&item, "address")
		_, a, err := as.NewAccountFromAddress(address)
		if err != nil {
			return nil, err
		}
		index := util.AsUint64(&item, "index")

		txWithData, err := n.merkleService.GetTransactionByIndex(a, index)
		if err != nil {
			return nil, err
		}
		_, err = n.consensusService.HashTransaction(txWithData)
		if err != nil {
			return nil, err
		}
		return txWithData, nil

	case "getBlockByHash":
		item := params[0].(map[string]interface{})
		hashString := util.ToString(&item, "hash")
		h, err := hex.DecodeString(hashString)
		if err != nil {
			return nil, err
		}
		block, err := n.merkleService.GetBlockByHash(libcore.Hash(h))
		if err != nil {
			return nil, err
		}
		_, err = n.consensusService.HashBlock(block)
		if err != nil {
			return nil, err
		}
		return block, nil

	case "getBlockByNumber":
		item := params[0].(map[string]interface{})
		index := util.AsUint64(&item, "num")
		block, err := n.merkleService.GetBlockByIndex(index)
		if err != nil {
			return nil, err
		}
		_, err = n.consensusService.HashBlock(block)
		if err != nil {
			return nil, err
		}
		return block, nil

	case "signTransaction":
		l := len(params)
		list := make([]string, l)
		for i := 0; i < l; i++ {
			item := params[i].(map[string]interface{})
			blob, tx, err := n.signTransaction(item)
			if err != nil {
				return nil, err
			}
			hash := tx.GetHash()
			list[i] = blob
			glog.Infoln("sign transaction", hash.String(), blob)
		}
		return list, nil

	case "sendTransaction":
		l := len(params)
		list := make([]string, l)
		for i := 0; i < l; i++ {
			item := params[i].(map[string]interface{})
			_, tx, err := n.signTransaction(item)
			if err != nil {
				return nil, err
			}
			_, err = n.sendTransaction(tx)
			if err != nil {
				return nil, err
			}

			hash := tx.GetHash()
			list[i] = hash.String()
		}
		return list, nil

	case "sendRawTransaction":
		l := len(params)
		list := make([]string, l)
		for i := 0; i < l; i++ {
			blob := params[i].(string)

			data, err := hex.DecodeString(blob)
			if err != nil {
				return nil, err
			}
			tx := &block.Transaction{}
			err = tx.UnmarshalBinary(data)
			if err != nil {
				return nil, err
			}
			_, err = n.sendTransaction(tx)
			if err != nil {
				return nil, err
			}

			hash := tx.GetHash()
			list[i] = hash.String()
		}
		return list, nil

	case "getStateByHash":
		item := params[0].(map[string]interface{})
		hashString := util.ToString(&item, "hash")
		h, err := hex.DecodeString(hashString)
		if err != nil {
			return nil, err
		}
		s, err := n.merkleService.GetStateByHash(libcore.Hash(h))
		if err != nil {
			return nil, err
		}
		_, err = n.consensusService.HashState(s)
		if err != nil {
			return nil, err
		}
		return s, nil

	case "getStateByAddress":
		item := params[0].(map[string]interface{})
		stateTypeName := util.ToString(&item, "type")
		address := util.ToString(&item, "address")
		stateType := libblock.GetStateTypeByName(stateTypeName)
		switch stateType {
		case block.ACCOUNT_STATE:
			_, a, err := as.NewAccountFromAddress(address)
			if err != nil {
				return nil, err
			}
			s, err := n.consensusService.GetAccountInfo(a)
			if err != nil {
				return nil, err
			}
			_, err = n.consensusService.HashState(s)
			if err != nil {
				return nil, err
			}
			return s, nil
		}
		return nil, errors.New("error")

	case "dump":
		n.storageService.Dump(chunk.LogPrinter{})
		return n.storageService.storage.Root().String(), nil

	default:
		return nil, fmt.Errorf("no such method %s", method)
	}
}

func (n *Node) AddTransaction(txWithData libblock.TransactionWithData) (libcore.Hash, bool) {
	h, err := n.consensusService.HashTransaction(txWithData)
	if err != nil {
		return nil, false
	}

	for i := 0; i < len(n.transactions); i++ {
		txWithData := n.transactions[i]
		hi := txWithData.GetTransaction().GetHash()
		if bytes.Equal(h, hi) {
			return nil, false
		}
	}

	n.transactions = append(n.transactions, txWithData)
	return h, true
}

func (n *Node) ClearTransaction(b libblock.Block) {
	n.transactionLocker.Lock()
	defer n.transactionLocker.Unlock()

	m := make(map[string]bool)
	list := b.GetTransactions()
	for i := 0; i < len(list); i++ {
		txWithData := list[i]
		h := txWithData.GetTransaction().GetHash()
		m[h.String()] = true
	}

	transactions := make([]libblock.TransactionWithData, 0)
	for i := 0; i < len(n.transactions); i++ {
		txWithData := n.transactions[i]
		h := txWithData.GetTransaction().GetHash()
		exists, ok := m[h.String()]
		if ok && exists {
			glog.Infoln(">>> drop block transaction", h.String())
		} else {
			transactions = append(transactions, txWithData)
		}
	}
	n.transactions = transactions
}

func (n *Node) GenerateBlock() (libblock.Block, error) {
	n.transactionLocker.Lock()
	defer n.transactionLocker.Unlock()

	list := n.transactions
	n.transactions = make([]libblock.TransactionWithData, 0)

	return n.consensusService.GenerateBlock(list)
}

func (n *Node) VerifyBlock(b libblock.Block) (ok bool, err error) {
	return n.consensusService.VerifyBlock(b)
}

func (n *Node) AddBlock(b libblock.Block) error {
	return n.consensusService.AddBlock(b)
}

func (n *Node) generate() {
	n.ready.Wait()

	for {
		if len(n.transactions) > 0 {
			block, err := n.GenerateBlock()
			if err != nil {
				glog.Error(err)
			} else {
				h, err := n.consensusService.HashBlock(block)
				if err != nil {
					glog.Error(err)
				} else {
					glog.Infof("=== generate block %d, %s, %d\n", block.GetIndex(), h.String(), len(block.GetTransactions()))

					_, err = n.consensusService.VerifyBlock(block)
					if err != nil {
						glog.Error(err)
					} else {
						err = n.consensusService.AddBlock(block)
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
		} else {
			glog.Infof("=== block %d, %s, prepare %d, %d\n", n.GetBlockNumber(), n.GetBlockHash(), n.GetBlockNumber()+1, len(n.transactions))
			for i := 0; i < int(n.config.GetBlockDuration()*5); i++ {
				if len(n.transactions) == 0 {
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

		glog.Infoln("===", 0, n.self.GetIndex(), n.self.GetAddress(), n.self.Id, n.self.Status, n.GetBlockNumber(), n.self.PeerCount)
		for i := 0; i < len(list); i++ {
			p := list[i]
			glog.Infoln("==>", i+1, p.GetIndex(), p.GetAddress(), p.Id, p.Status, p.BlockNumber, p.PeerCount)
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
	glog.Infoln("connected to", p.GetAddress())
}

func (n *Node) SendRequestInfo(p *Peer) {
	glog.Infoln("request to", p.GetAddress())
}

func (n *Node) PrepareConsensus() bool {
	list := n.ListPeer()
	count := int64(0)
	currentBlock := n.GetBlockNumber()
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
		if p.Status >= PeerKnown && n.GetBlockNumber() > p.BlockNumber {
			glog.Infoln("===", n.GetBlockNumber(), lastSendBlock, lastSendTime, p.address, p.Id, p.Status, p.BlockNumber, p.PeerCount)
			if lastSendBlock >= 0 && lastSendBlock == (p.BlockNumber+1) && (time.Since(lastSendTime) < (time.Duration(int64(config.GetBlockDuration()) * int64(time.Second)))) {
				time.Sleep(3 * time.Second)
			} else {
				block, err := n.merkleService.GetBlockByIndex(uint64(p.BlockNumber + 1))
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
									glog.Infof(">>> send data %d(%s) to %s(%s)\n", mid, core.GetInfo(blockData), p.GetAddress(), p.Id)
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

func (n *Node) Load() error {
	current := uint64(0)
	for {
		b, err := n.merkleService.GetBlockByIndex(current)
		if err != nil {
			glog.Error(err)
			break
		}
		_, err = n.consensusService.HashBlock(b)
		if err != nil {
			glog.Error(err)
			break
		}

		if n.consensusService.ValidatedBlock != nil {
			if b.GetIndex() != (n.consensusService.ValidatedBlock.GetIndex() + 1) {
				glog.Error(util.ErrorOfInvalid("block index", fmt.Sprintf("%d != %d", b.GetIndex(), (n.consensusService.ValidatedBlock.GetIndex()+1))))
				break
			}
			if !b.GetParentHash().Equals(n.consensusService.ValidatedBlock.GetHash()) {
				glog.Error(util.ErrorOfInvalid("block hash", fmt.Sprintf("%s != %s", b.GetParentHash().String(), n.consensusService.ValidatedBlock.GetHash().String())))
				break
			}
		} else {
			if b.GetIndex() != 0 {
				glog.Error(util.ErrorOfInvalid("genesis block", fmt.Sprintf("%d, %s", b.GetIndex(), b.GetHash().String())))
				break
			}
			if !b.GetParentHash().IsZero() {
				glog.Error(util.ErrorOfInvalid("the parent hash of the genesis block", b.GetParentHash().String()))
				break
			}
		}

		n.consensusService.ValidatedBlock = b

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

func (n *Node) Dump(printer core.Printer) {
	n.storageService.Dump(printer)
}

func (n *Node) GetBlockNumber() int64 {
	if n.consensusService != nil {
		return n.consensusService.GetBlockNumber()
	}
	return -1
}

func (n *Node) GetBlockHash() string {
	if n.consensusService != nil {
		return n.consensusService.GetBlockHash()
	}
	return ""
}

func (n *Node) GetBlock() libblock.Block {
	if n.consensusService != nil {
		return n.consensusService.GetBlock()
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
					glog.Infof(">>> send message %d(%s) to %s(%s)\n", m.Id, core.GetInfo(data), p.GetAddress(), p.Id)
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
			glog.Infof("<<< receive peer message %d(%s) from %s(%s)\n", m.Id, core.GetInfo(data), fromPeer.GetAddress(), fromPeer.Id)
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
						if int64(b.GetIndex()) <= n.consensusService.GetBlockNumber() {
							glog.Infoln("<<< drop message block", b.GetIndex(), h.String(), len(b.GetTransactions()))
						} else {
							glog.Infoln("<<< receive message block", b.GetIndex(), h.String(), len(b.GetTransactions()))
							_, err = n.consensusService.VerifyBlock(b)
							if err != nil {
								glog.Error(err)
							} else {
								glog.Infoln("<<< verify message block", b.GetIndex(), h.String(), len(b.GetTransactions()))
								err = n.consensusService.AddBlock(b)
								if err != nil {
									glog.Error(err)
								} else {
									glog.Infoln("<<< add message block", b.GetIndex(), h.String(), len(b.GetTransactions()))
								}
							}
						}
					}
				}
				n.SendPeerInfo(fromPeer)

			case core.CORE_TRANSACTION:
				tx := &block.Transaction{}
				err := tx.UnmarshalBinary(data)
				if err != nil {
					log.Println(err)
				} else {
					txWithData, _, err := n.processTransaction(tx)
					if err != nil {
						glog.Error(err)
					} else {
						hash := tx.GetHash().String()
						// util.PrintJSON("txWithData", txWithData)
						_, ok := n.AddTransaction(txWithData)
						if ok {
							glog.Infoln(">>> transaction", hash)
						} else {
							glog.Infoln(">>> drop transaction", hash)
						}
						glog.Infoln("<<< receive transaction", hash)
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
	glog.Infof("[%s][%s] recv discovery from peer[%s], len: %d", n.config.GetChainId(), n.self.GetAddress(), publisher, len(msgData))
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
		index, err := core.GetIndex(address.String())
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

func (n *Node) messageHandler(id string, data []byte) error {
	glog.Infof("[%s][%s] recv message from peer[%s], len: %d", n.config.GetChainId(), n.self.GetAddress(), id, len(data))
	msg, err := n.DecodeMessage(data)
	if err != nil {
		return err
	}
	n.ins <- msg
	return nil
}

func (n *Node) dataHandler(id string, msgData []byte) error {
	glog.Infof("[%s][%s] recv data from peer[%s], len: %d", n.config.GetChainId(), n.self.GetAddress(), id, len(msgData))
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
					if int64(b.GetIndex()) <= n.consensusService.GetBlockNumber() {
						glog.Infoln("<<< drop data block", b.GetIndex(), h.String(), len(b.GetTransactions()))
					} else {
						glog.Infoln("<<< receive data block", b.GetIndex(), h.String(), len(b.GetTransactions()))

						_, err = n.consensusService.VerifyBlock(b)
						if err != nil {
							glog.Error(err)
						} else {
							glog.Infoln("<<< verify data block", b.GetIndex(), h.String(), len(b.GetTransactions()))
							err = n.consensusService.AddBlock(b)
							if err != nil {
								glog.Error(err)
							} else {
								glog.Infoln("<<< add data block", b.GetIndex(), h.String(), len(b.GetTransactions()))
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
		BlockNumber: n.GetBlockNumber(),
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
				glog.Infof(">>> send data %d(%s) to %s(%s)\n", mid, core.GetInfo(peerData), toPeer.GetAddress(), toPeer.Id)
			}
		}
	}
}

func (n *Node) LoadPageByName(name string) (*zipfs.FileSystem, error) {
	if len(name) == 0 {
		return nil, util.ErrorOfInvalid("name", name)
	}
	data, err := n.storageService.ReadPageByName(name)
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

func (n *Node) LoadPageByAddress(address string) (*zipfs.FileSystem, error) {
	_, account, err := n.accountService.NewAccountFromAddress(address)
	if err != nil {
		return nil, err
	}
	data, err := n.storageService.ReadPageByAddress(account)
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
	if err := n.Load(); err != nil {
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
	if err := n.merkleService.Close(); err != nil {
		return err
	}
	if err := n.storageService.Close(); err != nil {
		return err
	}
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
				BlockNumber: n.GetBlockNumber(),
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

	_, ok := n.peers[p.GetIndex()]
	if !ok {
		n.peers[p.GetIndex()] = p
		go n.discoveryPeer(p)
	}
	return nil
}

func (n *Node) RemovePeer(p *Peer) error {
	n.peerLocker.Lock()
	defer n.peerLocker.Unlock()

	delete(n.peers, p.GetIndex())
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
		Node: n.self.GetIndex(),
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
		glog.Infof("<<< receive peer %s %d(%s) from %s(%s)\n", t, msg.Id, core.GetInfo(data), fromPeer.GetAddress(), fromPeer.Id)
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
