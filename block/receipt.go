package block

import (
	"github.com/tokentransfer/node/core"
	"github.com/tokentransfer/node/core/pb"
	"github.com/tokentransfer/node/util"

	libblock "github.com/tokentransfer/interfaces/block"
	libcore "github.com/tokentransfer/interfaces/core"
	libcrypto "github.com/tokentransfer/interfaces/crypto"
)

type Receipt struct {
	Hash              libcore.Hash
	BlockIndex        uint64
	TransactionIndex  uint32
	TransactionResult libblock.TransactionResult

	States []libblock.State
	Datas  []*DataInfo
}

func (r *Receipt) GetHash() libcore.Hash {
	return r.Hash
}

func (r *Receipt) SetHash(h libcore.Hash) {
	r.Hash = h
}

func (r *Receipt) UnmarshalBinary(data []byte) error {
	meta, msg, err := core.Unmarshal(data)
	if err != nil {
		return err
	}
	if meta != core.CORE_RECEIPT {
		return util.ErrorOfInvalid("data", "receipt")
	}
	receipt := msg.(*pb.Receipt)
	r.TransactionResult = libblock.TransactionResult(receipt.TransactionResult)
	r.TransactionIndex = receipt.TransactionIndex
	r.BlockIndex = receipt.BlockIndex

	list := receipt.GetStates()
	l := len(list)
	states := make([]libblock.State, l)
	for i := 0; i < l; i++ {
		data := list[i]
		s, err := ReadState(data)
		if err != nil {
			return err
		}
		states[i] = s
	}
	r.States = states

	list = receipt.GetDatas()
	l = len(list)
	datas := make([]*DataInfo, l)
	for i := 0; i < l; i++ {
		data := list[i]
		t, msg, err := core.Unmarshal(data)
		if err != nil {
			return err
		}
		if t != core.CORE_DATA_INFO {
			return util.ErrorOfInvalid("data", "info")
		}
		info := msg.(*pb.DataInfo)
		datas[i] = fromDataInfo(info)
	}
	r.Datas = datas

	return nil
}

func (r *Receipt) MarshalBinary() ([]byte, error) {
	return r.Raw(libcrypto.RawBinary)
}

func (r *Receipt) Raw(rt libcrypto.RawType) ([]byte, error) {
	l := len(r.States)
	states := make([][]byte, l)
	for i := 0; i < l; i++ {
		s := r.States[i]
		data, err := s.Raw(rt)
		if err != nil {
			return nil, err
		}
		states[i] = data
	}
	l = len(r.Datas)
	datas := make([][]byte, l)
	for i := 0; i < l; i++ {
		info := r.Datas[i]
		if info != nil {
			pbInfo := toDataInfo(info, rt)
			data, err := core.Marshal(pbInfo)
			if err != nil {
				return nil, err
			}
			datas[i] = data
		}
	}
	receipt := &pb.Receipt{
		TransactionResult: uint32(r.TransactionResult),
		States:            states,
		Datas:             datas,
	}
	return core.Marshal(receipt)
}

func (r *Receipt) GetTransactionIndex() uint32 {
	return r.TransactionIndex
}

func (r *Receipt) SetTransactionIndex(index uint32) {
	r.TransactionIndex = index
}

func (r *Receipt) GetBlockIndex() uint64 {
	return r.BlockIndex
}

func (r *Receipt) SetBlockIndex(index uint64) {
	r.BlockIndex = index
}

func (r *Receipt) GetTransactionResult() libblock.TransactionResult {
	return r.TransactionResult
}

func (r *Receipt) GetStates() []libblock.State {
	return r.States
}
