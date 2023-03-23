package block

import (
	"errors"

	"github.com/tokentransfer/node/core"
	"github.com/tokentransfer/node/core/pb"

	libblock "github.com/tokentransfer/interfaces/block"
	libcore "github.com/tokentransfer/interfaces/core"
)

type State struct {
	Hash libcore.Hash

	Account    libcore.Address
	Sequence   uint64
	BlockIndex uint64
	Previous   libcore.Hash

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

type AccountState struct {
	State

	Amount core.Amount
}

func (s *AccountState) GetStateKey() string {
	return core.GetAccountKey(s.Account, s.Amount.Currency, s.Amount.Issuer, "-")
}

func (s *AccountState) UnmarshalBinary(data []byte) error {
	meta, msg, err := core.Unmarshal(data)
	if err != nil {
		return err
	}
	if meta != core.CORE_ACCOUNT_STATE {
		return errors.New("error State data")
	}

	state := msg.(*pb.AccountState)

	_, account, err := as.NewAccountFromBytes(state.State.Account[:])
	if err != nil {
		return err
	}

	a, err := core.NewAmount(state.Amount)
	if err != nil {
		return err
	}

	s.State.StateType = libblock.StateType(core.CORE_ACCOUNT_STATE)
	s.State.BlockIndex = state.State.BlockIndex
	s.State.Account = account
	s.State.Sequence = state.State.Sequence
	s.State.Previous = state.State.Previous
	s.Amount = *a
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
		},
		Amount: s.Amount.String(),
	})
}

func (s *AccountState) Raw(ignoreSigningFields bool) ([]byte, error) {
	a, err := s.Account.MarshalBinary()
	if err != nil {
		return nil, err
	}

	return core.Marshal(&pb.AccountState{
		State: &pb.State{
			StateType: uint32(core.CORE_ACCOUNT_STATE),
			Account:   a,
			Sequence:  s.Sequence,
			Previous:  []byte(s.Previous),
		},
		Amount: s.Amount.String(),
	})
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
