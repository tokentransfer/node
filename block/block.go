package block

import (
	"github.com/tokentransfer/node/core"
	"github.com/tokentransfer/node/core/pb"
	"github.com/tokentransfer/node/util"

	libblock "github.com/tokentransfer/interfaces/block"
	libcore "github.com/tokentransfer/interfaces/core"
	libcrypto "github.com/tokentransfer/interfaces/crypto"
)

// reference from github.com/tokentransfer/chain/block/block.go
type Block struct {
	Hash libcore.Hash

	Account         libcore.Address
	BlockIndex      uint64
	ParentHash      libcore.Hash
	RootHash        libcore.Hash
	TransactionHash libcore.Hash
	StateHash       libcore.Hash
	Timestamp       int64

	Transactions []libblock.TransactionWithData
	States       []libblock.State
}

func (b *Block) GetAccount() libcore.Address {
	return b.Account
}

func (b *Block) GetIndex() uint64 {
	return b.BlockIndex
}

func (b *Block) GetHash() libcore.Hash {
	return b.Hash
}

func (b *Block) SetHash(h libcore.Hash) {
	b.Hash = h
}

func (b *Block) UnmarshalBinary(data []byte) error {
	meta, msg, err := core.Unmarshal(data)
	if err != nil {
		return err
	}
	if meta != core.CORE_BLOCK {
		return util.ErrorOfInvalid("data", "block")
	}
	block := msg.(*pb.Block)

	a, err := ByteToAddress(block.Account)
	if err != nil {
		return err
	}
	b.Account = a
	b.BlockIndex = block.BlockIndex
	b.ParentHash = libcore.Hash(block.ParentHash)
	b.RootHash = libcore.Hash(block.RootHash)
	b.TransactionHash = libcore.Hash(block.TransactionHash)
	b.StateHash = libcore.Hash(block.StateHash)
	b.Timestamp = block.Timestamp

	l := len(block.Transactions)
	transactions := make([]libblock.TransactionWithData, l)
	for i := 0; i < l; i++ {
		data, err := core.Marshal(block.Transactions[i])
		if err != nil {
			return err
		}
		tx := &TransactionWithData{}
		err = tx.UnmarshalBinary(data)
		if err != nil {
			return err
		}
		transactions[i] = tx
	}
	b.Transactions = transactions

	l = len(block.States)
	states := make([]libblock.State, l)
	for i := 0; i < l; i++ {
		data := block.States[i]
		state, err := ReadState(data)
		if err != nil {
			return err
		}
		states[i] = state
	}
	b.States = states

	return nil
}

func (b *Block) MarshalBinary() ([]byte, error) {
	return b.Raw(libcrypto.RawBinary)
}

func (b *Block) Raw(rt libcrypto.RawType) ([]byte, error) {
	a, err := AddressToByte(b.Account)
	if err != nil {
		return nil, err
	}
	block := &pb.Block{
		Account:         a,
		BlockIndex:      b.BlockIndex,
		ParentHash:      []byte(b.ParentHash),
		RootHash:        []byte(b.RootHash),
		TransactionHash: []byte(b.TransactionHash),
		StateHash:       []byte(b.StateHash),
		Timestamp:       b.Timestamp,
	}

	l := len(b.Transactions)
	transactions := make([]*pb.TransactionWithData, l)
	for i := 0; i < l; i++ {
		tx := b.Transactions[i]
		data, err := tx.Raw(rt)
		if err != nil {
			return nil, err
		}
		_, msg, err := core.Unmarshal(data)
		if err != nil {
			return nil, err
		}
		transactions[i] = msg.(*pb.TransactionWithData)
	}
	block.Transactions = transactions

	l = len(b.States)
	states := make([][]byte, l)
	for i := 0; i < l; i++ {
		state := b.States[i]
		data, err := state.Raw(rt)
		if err != nil {
			return nil, err
		}
		states[i] = data
	}
	block.States = states

	return core.Marshal(block)
}

func (b *Block) GetParentHash() libcore.Hash {
	return libcore.Hash(b.ParentHash)
}

func (b *Block) GetRootHash() libcore.Hash {
	return libcore.Hash(b.RootHash)
}

func (b *Block) GetTransactionHash() libcore.Hash {
	return libcore.Hash(b.TransactionHash)
}

func (b *Block) GetStateHash() libcore.Hash {
	return libcore.Hash(b.StateHash)
}

func (b *Block) GetTransactions() []libblock.TransactionWithData {
	l := len(b.Transactions)
	ret := make([]libblock.TransactionWithData, l)
	for i := 0; i < l; i++ {
		ret[i] = b.Transactions[i]
	}
	return ret
}

func (b *Block) GetStates() []libblock.State {
	l := len(b.States)
	ret := make([]libblock.State, l)
	for i := 0; i < l; i++ {
		ret[i] = b.States[i]
	}
	return ret
}

func (b *Block) GetTime() int64 {
	return b.Timestamp
}
