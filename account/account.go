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
	BTC.Register("BTC", func(t libaccount.KeyType, seed []byte) libaccount.Key {
		key, _ := btc.GenerateFamilySeed(string(seed))
		return key
	})
	ETH.Register("ETH", func(t libaccount.KeyType, seed []byte) libaccount.Key {
		key, _ := eth.GenerateFamilySeed(string(seed))
		return key
	})
}

type accountService struct {
	defaultType libaccount.KeyType
}

func NewAccountService() *accountService {
	return &accountService{defaultType: ETH}
}

func NewAccountServiceWith(defaultType libaccount.KeyType) *accountService {
	return &accountService{defaultType: defaultType}
}

func (service *accountService) GenerateFamilySeed(password string) (libaccount.KeyType, libaccount.Key, error) {
	k, err := GenerateFamilySeed(service.defaultType, password)
	if err != nil {
		return libaccount.KeyType(0), nil, err
	}
	return service.defaultType, k, nil
}

func (service *accountService) NewKeyFromSecret(secret string) (libaccount.KeyType, libaccount.Key, error) {
	key := &Key{}
	err := key.UnmarshalText([]byte(secret))
	if err != nil {
		return 0, nil, err
	}
	return key.GetKeyType(), key, nil
}

func (service *accountService) NewKeyFromBytes(data []byte) (libaccount.KeyType, libaccount.Key, error) {
	key := &Key{}
	err := key.UnmarshalBinary(data)
	if err != nil {
		return 0, nil, err
	}
	return key.GetKeyType(), key, nil
}

func (service *accountService) NewPublicFromHex(s string) (libaccount.KeyType, libaccount.PublicKey, error) {
	key := &Public{}
	err := key.UnmarshalText([]byte(s))
	if err != nil {
		return 0, nil, err
	}
	return key.GetKeyType(), key, nil
}

func (service *accountService) NewPublicFromBytes(data []byte) (libaccount.KeyType, libaccount.PublicKey, error) {
	key := &Public{}
	err := key.UnmarshalBinary(data)
	if err != nil {
		return 0, nil, err
	}
	return key.GetKeyType(), key, nil
}

func (service *accountService) NewAccountFromAddress(address string) (libaccount.KeyType, libcore.Address, error) {
	a := &Address{}
	err := a.UnmarshalText([]byte(address))
	if err != nil {
		return 0, nil, err
	}
	return a.GetKeyType(), a, nil
}

func (service *accountService) NewAccountFromBytes(data []byte) (libaccount.KeyType, libcore.Address, error) {
	a := &Address{}
	err := a.UnmarshalBinary(data)
	if err != nil {
		return 0, nil, err
	}
	return a.GetKeyType(), a, nil
}
