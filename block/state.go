package block

import (
	"github.com/tokentransfer/node/core"
	"github.com/tokentransfer/node/core/pb"
	"github.com/tokentransfer/node/util"

	libblock "github.com/tokentransfer/interfaces/block"
	libcore "github.com/tokentransfer/interfaces/core"
	libcrypto "github.com/tokentransfer/interfaces/crypto"
)

type State struct {
	Hash libcore.Hash

	Account    libcore.Address
	Sequence   uint64
	BlockIndex uint64
	Previous   libcore.Hash
	Version    uint64

	StateType libblock.StateType
}

func (s *State) GetHash() libcore.Hash {
	return s.Hash
}

func (s *State) SetHash(h libcore.Hash) {
	s.Hash = h
}

func (s *State) GetStateType() libblock.StateType {
	return s.StateType
}

func (s *State) GetBlockIndex() uint64 {
	return s.BlockIndex
}

func (s *State) SetBlockIndex(index uint64) {
	s.BlockIndex = index
}

func (s *State) GetPrevious() libcore.Hash {
	return s.Previous
}

func (s *State) GetAccount() libcore.Address {
	return s.Account
}

func (s *State) GetIndex() uint64 {
	return s.Sequence
}

func (s *State) GetVersion() uint64 {
	return s.Version
}

type AccountState struct {
	State

	Name string

	User     *DataInfo
	Code     *DataInfo
	Page     *DataInfo
	Token    *DataInfo
	Contract *DataInfo
	Data     *DataInfo

	PublicKey libcore.PublicKey
	RootHash  libcore.Hash
}

func (s *AccountState) GetName() string {
	return s.Name
}

func (s *AccountState) GetStateKey() []string {
	list := make([]string, 0)
	name := s.GetName()
	if len(name) > 0 {
		list = append(list, name)
		list = append(list, util.GetVersionKey(name, s.Version))
	}
	list = append(list, s.Account.String())
	list = append(list, util.GetVersionKey(s.Account.String(), s.Version))
	return list
}

func (s *AccountState) UnmarshalBinary(data []byte) error {
	meta, msg, err := core.Unmarshal(data)
	if err != nil {
		return err
	}
	if meta != core.CORE_ACCOUNT_STATE {
		return util.ErrorOfInvalid("state", "data")
	}

	state := msg.(*pb.AccountState)

	_, account, err := as.NewAccountFromBytes(state.State.Account[:])
	if err != nil {
		return err
	}

	s.State.StateType = libblock.StateType(core.CORE_ACCOUNT_STATE)
	s.State.BlockIndex = state.State.BlockIndex
	s.State.Account = account
	s.State.Sequence = state.State.Sequence
	s.State.Previous = state.State.Previous
	s.State.Version = state.State.Version

	s.Name = state.Name

	s.User = FromDataInfo(state.User)
	s.Code = FromDataInfo(state.Code)
	s.Page = FromDataInfo(state.Page)
	s.Token = FromDataInfo(state.Token)
	s.Contract = FromDataInfo(state.Contract)
	s.Data = FromDataInfo(state.Data)

	s.PublicKey = libcore.PublicKey(state.PublicKey)
	s.RootHash = libcore.Hash(state.RootHash)
	return nil
}

func (s *AccountState) MarshalBinary() ([]byte, error) {
	return s.Raw(libcrypto.RawBinary)
}

func (s *AccountState) Raw(rt libcrypto.RawType) ([]byte, error) {
	a, err := s.Account.MarshalBinary()
	if err != nil {
		return nil, err
	}

	switch rt {
	case libcrypto.RawBinary:
		return core.Marshal(&pb.AccountState{
			State: &pb.State{
				StateType:  uint32(core.CORE_ACCOUNT_STATE),
				BlockIndex: s.BlockIndex,
				Account:    a,
				Sequence:   s.Sequence,
				Previous:   []byte(s.Previous),
				Version:    s.Version,
			},

			Name:     s.Name,
			User:     ToDataInfo(s.User, rt),
			Code:     ToDataInfo(s.Code, rt),
			Page:     ToDataInfo(s.Page, rt),
			Token:    ToDataInfo(s.Token, rt),
			Contract: ToDataInfo(s.Contract, rt),
			Data:     ToDataInfo(s.Data, rt),

			PublicKey: []byte(s.PublicKey),
			RootHash:  []byte(s.RootHash),
		})
	}
	return core.Marshal(&pb.AccountState{
		State: &pb.State{
			StateType: uint32(core.CORE_ACCOUNT_STATE),
			Account:   a,
			Sequence:  s.Sequence,
			Previous:  []byte(s.Previous),
			Version:   s.Version,
		},

		Name:     s.Name,
		User:     ToDataInfo(s.User, rt),
		Code:     ToDataInfo(s.Code, rt),
		Page:     ToDataInfo(s.Page, rt),
		Token:    ToDataInfo(s.Token, rt),
		Contract: ToDataInfo(s.Contract, rt),
		Data:     ToDataInfo(s.Data, rt),

		PublicKey: []byte(s.PublicKey),
		RootHash:  []byte(s.RootHash),
	})
}
