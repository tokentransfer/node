package block

import (
	"errors"

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

	Name   string
	Amount string

	User  *DataInfo
	Code  *DataInfo
	Page  *DataInfo
	Token *DataInfo
	Data  *DataInfo
	File  *DataInfo
}

func (s *AccountState) GetName() string {
	return s.Name
}

func (s *AccountState) GetStateKey() []string {
	list := make([]string, 0)
	name := s.GetName()
	if len(name) > 0 {
		list = append(list, name)
		list = append(list, core.GetVersionKey(name, s.Version, "-"))
	}
	list = append(list, s.Account.String())
	list = append(list, core.GetVersionKey(s.Account.String(), s.Version, "-"))
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
	s.Amount = state.Amount

	s.User = fromDataInfo(state.User)
	s.Code = fromDataInfo(state.Code)
	s.Page = fromDataInfo(state.Page)
	s.Token = fromDataInfo(state.Token)
	s.Data = fromDataInfo(state.Data)
	s.File = fromDataInfo(state.File)
	return nil
}

func (s *AccountState) MarshalBinary() ([]byte, error) {
	a, err := s.Account.MarshalBinary()
	if err != nil {
		return nil, err
	}

	return core.Marshal(&pb.AccountState{
		State: &pb.State{
			StateType:  uint32(core.CORE_ACCOUNT_STATE),
			BlockIndex: s.BlockIndex,
			Account:    a,
			Sequence:   s.Sequence,
			Previous:   []byte(s.Previous),
			Version:    s.Version,
		},
		Amount: s.Amount,

		Name:  s.Name,
		User:  toDataInfo(s.User, libcrypto.RawBinary),
		Code:  toDataInfo(s.Code, libcrypto.RawBinary),
		Page:  toDataInfo(s.Page, libcrypto.RawBinary),
		Token: toDataInfo(s.Token, libcrypto.RawBinary),
		Data:  toDataInfo(s.Data, libcrypto.RawBinary),
		File:  toDataInfo(s.File, libcrypto.RawBinary),
	})
}

func (s *AccountState) Raw(ignoreSigningFields bool) ([]byte, error) {
	a, err := s.Account.MarshalBinary()
	if err != nil {
		return nil, err
	}

	if ignoreSigningFields {
		return core.Marshal(&pb.AccountState{
			State: &pb.State{
				StateType: uint32(core.CORE_ACCOUNT_STATE),
				Account:   a,
				Sequence:  s.Sequence,
				Previous:  []byte(s.Previous),
				Version:   s.Version,
			},
			Amount: s.Amount,

			Name:  s.Name,
			User:  toDataInfo(s.User, libcrypto.RawIgnoreSigningFields),
			Code:  toDataInfo(s.Code, libcrypto.RawIgnoreSigningFields),
			Page:  toDataInfo(s.Page, libcrypto.RawIgnoreSigningFields),
			Token: toDataInfo(s.Token, libcrypto.RawIgnoreSigningFields),
			Data:  toDataInfo(s.Data, libcrypto.RawIgnoreSigningFields),
			File:  toDataInfo(s.File, libcrypto.RawIgnoreSigningFields),
		})
	} else { // ignore variable fields
		return core.Marshal(&pb.AccountState{
			State: &pb.State{
				StateType: uint32(core.CORE_ACCOUNT_STATE),
				Account:   a,
				Sequence:  s.Sequence,
				Previous:  []byte(s.Previous),
				Version:   s.Version,
			},
			Amount: s.Amount,

			Name:  s.Name,
			User:  toDataInfo(s.User, libcrypto.RawIgnoreVariableFields),
			Code:  toDataInfo(s.Code, libcrypto.RawIgnoreVariableFields),
			Page:  toDataInfo(s.Page, libcrypto.RawIgnoreVariableFields),
			Token: toDataInfo(s.Token, libcrypto.RawIgnoreVariableFields),
			Data:  toDataInfo(s.Data, libcrypto.RawIgnoreVariableFields),
			File:  toDataInfo(s.File, libcrypto.RawIgnoreVariableFields),
		})
	}

}

func ReadState(data []byte) (libblock.State, error) {
	if len(data) == 0 {
		return nil, errors.New("error entry")
	}
	meta := core.GetMeta(data)
	switch meta {
	case core.CORE_ACCOUNT_STATE:
		s := &AccountState{}
		err := s.UnmarshalBinary(data)
		if err != nil {
			return nil, err
		}
		return s, nil
	default:
		return nil, errors.New("error data")
	}
}

func CloneState(state libblock.State) (libblock.State, error) {
	data, err := state.MarshalBinary()
	if err != nil {
		return nil, err
	}
	s, err := ReadState(data)
	if err != nil {
		return nil, err
	}
	return s, nil
}
