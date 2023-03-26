package consensus

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"chainmaker.org/chainmaker/protocol/v2"
	"github.com/caivega/glog"
	"github.com/tokentransfer/node/account"
	"github.com/tokentransfer/node/block"
	"github.com/tokentransfer/node/conf"
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

	transactions      []libblock.TransactionWithData
	transactionLocker *sync.Mutex

	net protocol.Net

	outs chan *pb.Message
	ins  chan *pb.Message

	ready *sync.WaitGroup
	timer *time.Ticker

	config *conf.Config
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

		accountService: &account.AccountService{},
		cryptoService:  &crypto.CryptoService{},
	}
}

func (n *Node) Init(c libcore.Config) error {
	n.config = c.(*conf.Config)
	core.Init(n.config)

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
		panic(err)
	}
	storageService, err := NewStorageService(n.config)
	if err != nil {
		panic(err)
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

	return nil
}

func (n *Node) signTransaction(txm map[string]interface{}) (string, *block.Transaction, error) {
	as := n.accountService

	from := util.ToString(&txm, "from")
	secret := util.ToString(&txm, "secret")
	to := util.ToString(&txm, "to")
	value := util.ToString(&txm, "value")

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
		Gas:         int64(10),
		Destination: toAccount,
	}

	payloadInfo := &pb.PayloadInfo{}
	if util.Has(&txm, "contract") {
		contractInfo := &pb.ContractInfo{}

		cm := util.ToMap(&txm, "contract")
		if util.Has(&cm, "account") {
			accountString := util.ToString(&cm, "account")
			_, account, err := as.NewAccountFromAddress(accountString)
			if err != nil {
				return "", nil, err
			}
			accountData, err := account.MarshalBinary()
			if err != nil {
				return "", nil, err
			}
			contractInfo.Account = accountData
		}
		if util.Has(&cm, "method") {
			contractInfo.Method = util.ToString(&cm, "method")
		}
		if util.Has(&cm, "params") {
			params := util.ToArray(&cm, "params")
			list := make([][]byte, 0)
			for i := 0; i < len(params); i++ {
				s := params[i].(string)
				list = append(list, []byte(s))
			}
			contractInfo.Params = list
		}
		if util.Has(&cm, "code") {
			codeString := util.ToString(&cm, "code")
			codeData, err := hex.DecodeString(codeString)
			if err != nil {
				return "", nil, err
			}
			contractInfo.Code = codeData
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
			contractInfo.Code = fileData
		}

		contractData, err := core.Marshal(contractInfo)
		if err != nil {
			return "", nil, err
		}

		payloadInfo.Payload = append(payloadInfo.Payload, contractData)
	}
	if len(payloadInfo.Payload) > 0 {
		payloadData, err := core.Marshal(payloadInfo)
		if err != nil {
			return "", nil, err
		}
		tx.Payload = libcore.Bytes(payloadData)
	}

	err = n.cryptoService.Sign(fromKey, tx)
	if err != nil {
		return "", nil, err
	}
	data, err := tx.MarshalBinary()
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
	fmt.Println("verify transaction", hash.String())

	txWithData, _, err := n.processTransaction(tx)
	if err != nil {
		fmt.Println("error", err) // ignore the error after verified transaction
	} else {
		data, err := tx.MarshalBinary()
		if err != nil {
			return nil, err
		}
		n.broadcast(data)
		fmt.Println("send transaction", hash.String())
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
	txWithData, err := n.consensusService.ProcessTransaction(tx)
	if err != nil {
		return nil, nil, err
	}

	// util.PrintJSON("txWithData", txWithData)
	_, ok = n.AddTransaction(txWithData)
	if ok {
		fmt.Println(">>> receive transaction", h.String())
	} else {
		fmt.Println(">>> drop transaction", h.String())
	}

	return txWithData, tx, nil
}

func (n *Node) getNextSequence(address libcore.Address) uint64 {
	accountEntry, err := n.consensusService.GetAccountInfo(address, nil, nil)
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
	case "getBalance":
		address := params[0].(string)
		_, a, err := as.NewAccountFromAddress(address)
		if err != nil {
			return nil, err
		}
		accountEntry, err := n.consensusService.GetAccountInfo(a, nil, nil)
		if err != nil {
			return nil, err
		}
		return accountEntry.Amount, nil
	case "getTransactionCount":
		address := params[0].(string)
		_, a, err := as.NewAccountFromAddress(address)
		if err != nil {
			return nil, err
		}
		seq := n.getNextSequence(a)
		return seq, nil
	case "getTransactionReceipt":
		hashString := params[0].(string)
		h, err := hex.DecodeString(hashString)
		if err != nil {
			return nil, err
		}
		txWithData, err := n.merkleService.GetTransactionByHash(libcore.Hash(h))
		if err != nil {
			return nil, err
		}
		receipt := txWithData.GetReceipt()
		_, err = n.HashReceipt(receipt)
		if err != nil {
			return nil, err
		}
		return receipt, nil
	case "getTransactionByHash":
		hashString := params[0].(string)
		h, err := hex.DecodeString(hashString)
		if err != nil {
			return nil, err
		}
		txWithData, err := n.merkleService.GetTransactionByHash(libcore.Hash(h))
		if err != nil {
			return nil, err
		}
		_, err = n.HashTransaction(txWithData)
		if err != nil {
			return nil, err
		}
		return txWithData, nil
	case "getTransactionByIndex":
		address := params[0].(string)
		_, a, err := as.NewAccountFromAddress(address)
		if err != nil {
			return nil, err
		}
		index := uint64(params[1].(float64))

		txWithData, err := n.merkleService.GetTransactionByIndex(a, index)
		if err != nil {
			return nil, err
		}
		_, err = n.HashTransaction(txWithData)
		if err != nil {
			return nil, err
		}
		return txWithData, nil
	case "getBlockByHash":
		hashString := params[0].(string)
		h, err := hex.DecodeString(hashString)
		if err != nil {
			return nil, err
		}
		block, err := n.merkleService.GetBlockByHash(libcore.Hash(h))
		if err != nil {
			return nil, err
		}
		_, err = n.HashBlock(block)
		if err != nil {
			return nil, err
		}
		return block, nil
	case "getBlockByNumber":
		index := uint64(params[0].(float64))
		block, err := n.merkleService.GetBlockByIndex(index)
		if err != nil {
			return nil, err
		}
		_, err = n.HashBlock(block)
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
			fmt.Println("sign transaction", hash.String(), blob)
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
		data, err := hex.DecodeString(params[0].(string))
		if err != nil {
			return nil, err
		}
		s, err := n.merkleService.GetStateByHash(libcore.Hash(data))
		if err != nil {
			return nil, err
		}
		_, err = n.HashState(s)
		if err != nil {
			return nil, err
		}
		return s, nil
	case "getStateByName":
		stateType := libblock.GetStateTypeByName(params[0].(string))
		switch stateType {
		case block.ACCOUNT_STATE:
			address := params[1].(string)
			_, a, err := as.NewAccountFromAddress(address)
			if err != nil {
				return nil, err
			}
			s, err := n.consensusService.GetAccountInfo(a, nil, nil)
			if err != nil {
				return nil, err
			}
			_, err = n.HashState(s)
			if err != nil {
				return nil, err
			}
			return s, nil
		}
		return nil, errors.New("error")
	case "receiveStateByName":
		stateType := libblock.GetStateTypeByName(params[0].(string))
		switch stateType {
		case block.ACCOUNT_STATE:
			address := params[1].(string)
			_, a, err := as.NewAccountFromAddress(address)
			if err != nil {
				return nil, err
			}
			s, err := n.consensusService.GetAccountInfo(a, nil, nil)
			if err != nil {
				return nil, err
			}
			_, err = n.HashState(s)
			if err != nil {
				return nil, err
			}
			data, err := s.MarshalBinary()
			if err != nil {
				return nil, err
			}
			return hex.EncodeToString(data), nil
		}
		return nil, errors.New("error")
	default:
		return nil, fmt.Errorf("no such method %s", method)
	}
}

func (n *Node) AddTransaction(txWithData libblock.TransactionWithData) (libcore.Hash, bool) {
	h, err := n.HashTransaction(txWithData)
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
			fmt.Println(">>> drop block transaction", h.String())
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

	now := time.Now()
	nowTime := now.UnixNano()
	nextTime := (nowTime - nowTime%int64(BLOCK_DURATION*time.Second)) + int64(BLOCK_DURATION*time.Second)
	duration := (nextTime - nowTime)
	if duration > 0 {
		time.Sleep(time.Duration(nextTime - nowTime))
	}

	n.timer = time.NewTicker(BLOCK_DURATION * time.Second)
	for {
		t := <-n.timer.C

		if len(n.transactions) > 0 {
			block, err := n.GenerateBlock()
			if err != nil {
				panic(err)
			}
			_, err = n.HashBlock(block)
			if err != nil {
				panic(err)
			}
			fmt.Printf("=== generate block %d, %s, %d\n", block.GetIndex(), t.String(), len(block.GetTransactions()))

			_, err = n.consensusService.VerifyBlock(block)
			if err != nil {
				panic(err)
			}

			err = n.consensusService.AddBlock(block)
			if err != nil {
				panic(err)
			}

			data, err := block.MarshalBinary()
			if err != nil {
				panic(err)
			}
			n.broadcast(data)
		} else {
			fmt.Printf("=== prepare block %d, %s, %d\n", n.GetBlockNumber(), n.GetBlockHash(), len(n.transactions))
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

		fmt.Println(">>>", 0, n.self.GetIndex(), n.self.GetAddress(), n.self.Status, n.GetBlockNumber(), n.self.PeerCount)
		for i := 0; i < len(list); i++ {
			p := list[i]
			fmt.Println(">>>", i+1, p.GetIndex(), p.GetAddress(), p.Status, p.BlockNumber, p.PeerCount)
		}

		if !lastConsensused && n.Consensused {
			lastConsensused = n.Consensused
			n.ready.Done()
		}

		time.Sleep(time.Duration(config.GetBlockDuration()) * time.Second)
	}
}

func (n *Node) ConnectTo(p *Peer) {
	fmt.Println("connected to", p.GetAddress())
}

func (n *Node) SendRequestInfo(p *Peer) {
	fmt.Println("request to", p.GetAddress())
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

	for {
		switch p.Status {
		case PeerNone:
			n.ConnectTo(p)
		case PeerConnected:
			n.SendRequestInfo(p)
		case PeerKnown:
			n.Consensused = n.PrepareConsensus()
		}

		fmt.Println("=== discovery peer", p.GetAddress(), p.Id, p.Status, p.BlockNumber, p.PeerCount, n.GetBlockNumber())
		if p.Status >= PeerKnown && n.GetBlockNumber() > p.BlockNumber {
			for i := p.BlockNumber + 1; i <= n.GetBlockNumber() && n.GetBlockNumber() > p.BlockNumber; i++ {
				block, err := n.merkleService.GetBlockByIndex(uint64(i))
				if err != nil {
					glog.Error(err)
				} else {
					_, err = n.HashBlock(block)
					if err != nil {
						glog.Error(err)
					} else {
						blockData, err := block.MarshalBinary()
						if err != nil {
							glog.Error(err)
						} else {
							mid, msgData, err := n.SendMessage(blockData)
							if err != nil {
								glog.Error(err)
							} else {
								err = n.net.SendMsg(n.config.GetChainId(), p.Id, TOPIC_PEER_DATA, msgData)
								if err != nil {
									glog.Error(err)
								} else {
									fmt.Printf(">>> send data %d(%s) to %s(%s)\n", mid, core.GetInfo(blockData), p.GetAddress(), p.Id)
								}
							}
						}
					}
				}

				time.Sleep(2 * time.Second)
			}
			time.Sleep(time.Second)
		} else {
			time.Sleep(time.Duration(config.GetBlockDuration()) * time.Second)
		}
	}
}

func (n *Node) Load() {
	current := uint64(0)
	for {
		b, err := n.merkleService.GetBlockByIndex(current)
		if err != nil {
			break
		}
		_, err = n.HashBlock(b)
		if err != nil {
			break
		}

		if n.consensusService.ValidatedBlock != nil {
			if b.GetIndex() != (n.consensusService.ValidatedBlock.GetIndex() + 1) {
				break
			}
			if !b.GetParentHash().Equals(n.consensusService.ValidatedBlock.GetHash()) {
				break
			}
		} else {
			if b.GetIndex() != 0 {
				break
			}
			if !b.GetParentHash().IsZero() {
				break
			}
		}

		n.consensusService.ValidatedBlock = b
		fmt.Println("load block", current, b.GetHash().String())

		current++
	}
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
			log.Println(err)
		} else {
			data := m.GetData()
			list := n.ListPeer()
			for i := 0; i < len(list); i++ {
				p := list[i]
				err := n.net.SendMsg(n.config.GetChainId(), p.Id, TOPIC_PEER_MESSAGE, msgBytes)
				if err != nil {
					fmt.Println(">>>", core.GetInfo(data), libcore.Bytes(data).String())
					log.Println(err)
				} else {
					fmt.Printf(">>> send message %d(%s) to %s(%s)\n", m.Id, core.GetInfo(data), p.GetAddress(), p.Id)
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
		fmt.Printf("<<< receive message from node %d, length: %d\n", fromIndex, len(data))
		if fromPeer != nil {
			fmt.Printf("<<< receive peer message %d(%s) from %s(%s)\n", m.Id, core.GetInfo(data), fromPeer.GetAddress(), fromPeer.Id)
			switch meta {
			case core.CORE_BLOCK:
				b := &block.Block{}
				err := b.UnmarshalBinary(data)
				if err != nil {
					glog.Error(err)
				} else {
					h, err := n.HashBlock(b)
					if err != nil {
						glog.Error(err)
					} else {
						if int64(b.GetIndex()) <= n.consensusService.GetBlockNumber() {
							fmt.Println("<<< drop block", b.GetIndex(), h.String(), len(b.GetTransactions()))
						} else {
							fmt.Println("<<< receive block", b.GetIndex(), h.String(), len(b.GetTransactions()))
							_, err = n.consensusService.VerifyBlock(b)
							if err != nil {
								glog.Error(err)
							} else {
								fmt.Println("<<< verify block", b.GetIndex(), h.String(), len(b.GetTransactions()))
								err = n.consensusService.AddBlock(b)
								if err != nil {
									glog.Error(err)
								}
							}
						}

					}
				}

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
							fmt.Println(">>> transaction", hash)
						} else {
							fmt.Println(">>> drop transaction", hash)
						}
						fmt.Println("<<< receive transaction", hash)
					}
				}

			default:
				fmt.Println("<<< error message")
			}
		} else {
			fmt.Printf("<<< receive unknown message %d(%s)\n", m.Id, core.GetInfo(data))
			if len(data) > 0 {
				meta := data[0]
				switch meta {
				default:
					fmt.Println("<<< unknown peer", fromIndex)
				}
			}
		}
	}
}

func (n *Node) discoveryHandler(publisher string, msgData []byte) error {
	glog.Infof("[%s][%s] recv discovery from peer[%s], len: %d", n.config.GetChainId(), n.self.GetAddress(), publisher, len(msgData))
	m, err := n.GetMessage(msgData)
	if err != nil {
		return err
	}
	meta, msg, err := core.Unmarshal(m.GetData())
	if err != nil {
		return err
	}
	if meta != core.CORE_PEER_INFO {
		return core.ErrorOfInvalid("format", "peer info")
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
		p := n.peers[index]
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
	msg, err := n.GetMessage(data)
	if err != nil {
		return err
	}
	n.ins <- msg
	return nil
}

func (n *Node) dataHandler(id string, msgData []byte) error {
	glog.Infof("[%s][%s] recv data from peer[%s], len: %d", n.config.GetChainId(), n.self.GetAddress(), id, len(msgData))
	m, err := n.GetMessage(msgData)
	if err != nil {
		return err
	}
	meta, fromPeer, data := n.ReceiveMessage("message", m)
	fromIndex := m.GetNode()
	fmt.Printf("<<< receive data from node %d, length: %d\n", fromIndex, len(data))
	if fromPeer != nil {
		fmt.Printf("<<< receive peer data %d(%s) from %s(%s)\n", m.Id, core.GetInfo(data), fromPeer.GetAddress(), fromPeer.Id)
		switch meta {
		case core.CORE_PEER_INFO:
			_, msg, err := core.Unmarshal(data)
			if err != nil {
				return err
			}
			peerInfo := msg.(*pb.PeerInfo)
			index, err := core.GetIndex(id)
			if err != nil {
				return err
			}
			p := n.peers[index]
			if p != nil {
				p.BlockNumber = peerInfo.BlockNumber
				p.PeerCount = peerInfo.PeerCount
			}

		case core.CORE_BLOCK:
			b := &block.Block{}
			err = b.UnmarshalBinary(data)
			if err != nil {
				glog.Error(err)
			} else {
				h, err := n.HashBlock(b)
				if err != nil {
					glog.Error(err)
				} else {
					if int64(b.GetIndex()) <= n.consensusService.GetBlockNumber() {
						fmt.Println("<<< drop block", b.GetIndex(), h.String(), len(b.GetTransactions()))
					} else {
						fmt.Println("<<< receive block", b.GetIndex(), h.String(), len(b.GetTransactions()))

						_, err = n.consensusService.VerifyBlock(b)
						if err != nil {
							glog.Error(err)
						} else {
							fmt.Println("<<< verify block", b.GetIndex(), h.String(), len(b.GetTransactions()))
							err = n.consensusService.AddBlock(b)
							if err != nil {
								glog.Error(err)
							}
						}
					}
				}
			}

			peers, _ := n.net.ChainNodesInfo(n.config.GetChainId())
			peerData, err := core.Marshal(&pb.PeerInfo{
				BlockNumber: n.GetBlockNumber(),
				PeerCount:   int64(len(peers)),
			})
			if err != nil {
				glog.Error(err)
			} else {
				mid, msgData, err := n.SendMessage(peerData)
				if err != nil {
					glog.Error(err)
				} else {
					err = n.net.SendMsg(n.config.GetChainId(), fromPeer.Id, TOPIC_PEER_DATA, msgData)
					if err != nil {
						glog.Error(err)
					} else {
						fmt.Printf(">>> send data %d(%s) to %s\n", mid, core.GetInfo(peerData), id)
					}
				}
			}
		}
	}
	return nil
}

func (n *Node) start() {
	readyC := make(chan struct{})

	var err error
	n.net, err = p2p.InitNet(n.config, readyC)
	if err != nil {
		panic(err)
	}
	err = n.net.Start()
	if err != nil {
		panic(err)
	}
	err = n.net.InitPubSub(n.config.GetChainId(), 0)
	if err != nil {
		panic(err)
	}

	err = n.net.SubscribeWithChainId(n.config.GetChainId(), TOPIC_PEER_DISCOVERY, n.discoveryHandler)
	if err != nil {
		panic(err)
	}
	err = n.net.DirectMsgHandle(n.config.GetChainId(), TOPIC_PEER_MESSAGE, n.messageHandler)
	if err != nil {
		panic(err)
	}
	err = n.net.DirectMsgHandle(n.config.GetChainId(), TOPIC_PEER_DATA, n.dataHandler)
	if err != nil {
		panic(err)
	}
	close(readyC)

	n.self.Id = n.config.GetNodeId()
}

func (n *Node) Start() error {
	n.Load()
	n.start()

	go n.send()
	go n.receive()

	go n.connect()
	go n.generate()
	go n.discovery()

	return nil
}

func (n *Node) Stop() error {
	n.net.CancelSubscribeWithChainId(n.config.GetChainId(), TOPIC_PEER_DISCOVERY)
	n.net.CancelDirectMsgHandle(n.config.GetChainId(), TOPIC_PEER_MESSAGE)
	n.net.CancelDirectMsgHandle(n.config.GetChainId(), TOPIC_PEER_DATA)
	n.net.Stop()

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
				_, msgData, err := n.SendMessage(peerData)
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
				fmt.Println("> discovery", p.NodeUid)
				for index, addr := range p.NodeAddress {
					fmt.Printf("  %d: %s\n", index, addr)
				}
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
			time.Sleep(60 * 10 * time.Second)
		} else {
			time.Sleep(10 * time.Second)
		}
	}
}

func (n *Node) GetNodeKey() libaccount.Key {
	return n.key
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

func (n *Node) HashBlock(b libblock.Block) (libcore.Hash, error) {
	h, _, err := n.cryptoService.Raw(b, libcrypto.RawBinary)
	if err != nil {
		return nil, err
	}
	transactions := b.GetTransactions()
	for i := 0; i < len(transactions); i++ {
		txWithData := transactions[i]
		_, err := n.HashTransaction(txWithData)
		if err != nil {
			return nil, err
		}
	}
	states := b.GetStates()
	for i := 0; i < len(states); i++ {
		state := states[i]
		_, err := n.HashState(state)
		if err != nil {
			return nil, err
		}
	}
	return h, nil
}

func (n *Node) HashTransaction(txWithData libblock.TransactionWithData) (libcore.Hash, error) {
	_, _, err := n.cryptoService.Raw(txWithData, libcrypto.RawBinary)
	if err != nil {
		return nil, err
	}

	h, _, err := n.cryptoService.Raw(txWithData.GetTransaction(), libcrypto.RawBinary)
	if err != nil {
		return nil, err
	}

	_, _, err = n.cryptoService.Raw(txWithData.GetReceipt(), libcrypto.RawBinary)
	if err != nil {
		return nil, err
	}

	_, err = n.HashReceipt(txWithData.GetReceipt())
	if err != nil {
		return nil, err
	}
	return h, nil
}

func (n *Node) HashReceipt(r libblock.Receipt) (libcore.Hash, error) {
	h, _, err := n.cryptoService.Raw(r, libcrypto.RawBinary)
	if err != nil {
		return nil, err
	}

	states := r.GetStates()
	for i := 0; i < len(states); i++ {
		state := states[i]
		_, err := n.HashState(state)
		if err != nil {
			return nil, err
		}
	}
	return h, nil
}

func (n *Node) HashState(s libblock.State) (libcore.Hash, error) {
	h, _, err := n.cryptoService.Raw(s, libcrypto.RawBinary)
	if err != nil {
		return nil, err
	}
	return h, nil
}

func (n *Node) CreateMessage(data []byte) (*pb.Message, error) {
	return &pb.Message{
		Id:   n.newId(),
		Data: data,
		Node: n.self.GetIndex(),
	}, nil
}

func (n *Node) SendMessage(data []byte) (uint64, []byte, error) {
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

func (n *Node) GetMessage(data []byte) (*pb.Message, error) {
	meta, m, err := core.Unmarshal(data)
	if err != nil {
		return nil, err
	}
	if meta != core.CORE_MESSAGE {
		return nil, core.ErrorOfInvalid("format", "message")
	}
	return m.(*pb.Message), nil
}

func (n *Node) ReceiveMessage(t string, msg *pb.Message) (byte, *Peer, []byte) {
	fromIndex := msg.GetNode()
	data := msg.GetData()
	fmt.Printf("<<< receive %s from node %d, length: %d\n", t, fromIndex, len(data))
	fromPeer, ok := n.peers[fromIndex]
	if ok {
		fmt.Printf("<<< receive peer %s %d(%s) from %s(%s)\n", t, msg.Id, core.GetInfo(data), fromPeer.GetAddress(), fromPeer.Id)
		if len(data) > 0 {
			meta := data[0]
			return meta, fromPeer, data
		}
	}
	return 0, nil, nil
}
