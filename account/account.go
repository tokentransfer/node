package account

import (
	"github.com/tokentransfer/node/account/btc"
	"github.com/tokentransfer/node/account/eth"

	libaccount "github.com/tokentransfer/interfaces/account"
	libcore "github.com/tokentransfer/interfaces/core"
)

const (
	BTC libaccount.KeyType = 1
	ETH libaccount.KeyType = 2
)

func init() {
	// KeyType
	ETH.Register("ETH", func(t libaccount.KeyType) libaccount.Key {
		key, _ := eth.GenerateFamilySeed("masterpassphrase")
		return key
	})
	BTC.Register("BTC", func(t libaccount.KeyType) libaccount.Key {
		key, _ := btc.GenerateFamilySeed("masterpassphrase")
		return key
	})
}

type AccountService struct {
}

func (service *AccountService) GenerateFamilySeed(password string) (libaccount.KeyType, libaccount.Key, error) {
	key, err := eth.GenerateFamilySeed(password)
	if err != nil {
		return 0, nil, err
	}
	return ETH, key, nil
}

func (service *AccountService) NewKeyFromSecret(secret string) (libaccount.KeyType, libaccount.Key, error) {
	key := &eth.Key{}
	err := key.UnmarshalText([]byte(secret))
	if err != nil {
		return 0, nil, err
	}
	return ETH, key, nil
}

func (service *AccountService) NewKeyFromBytes(data []byte) (libaccount.KeyType, libaccount.Key, error) {
	key := &eth.Key{}
	err := key.UnmarshalBinary(data)
	if err != nil {
		return 0, nil, err
	}
	return ETH, key, nil
}

func (service *AccountService) NewPublicFromHex(s string) (libaccount.KeyType, libaccount.PublicKey, error) {
	key := &eth.Public{}
	err := key.UnmarshalText([]byte(s))
	if err != nil {
		return 0, nil, err
	}
	return ETH, key, nil
}

func (service *AccountService) NewPublicFromBytes(data []byte) (libaccount.KeyType, libaccount.PublicKey, error) {
	key := &eth.Public{}
	err := key.UnmarshalBinary(data)
	if err != nil {
		return 0, nil, err
	}
	return ETH, key, nil
}

func (service *AccountService) NewAccountFromAddress(address string) (libaccount.KeyType, libcore.Address, error) {
	a := &eth.Address{}
	err := a.UnmarshalText([]byte(address))
	if err != nil {
		return 0, nil, err
	}
	return ETH, a, nil
}

func (service *AccountService) NewAccountFromBytes(data []byte) (libaccount.KeyType, libcore.Address, error) {
	a := &eth.Address{}
	err := a.UnmarshalBinary(data)
	if err != nil {
		return 0, nil, err
	}
	return ETH, a, nil
}
