package account

import (
	"strings"

	"github.com/tokentransfer/node/account/btc"
	"github.com/tokentransfer/node/account/eth"
	"github.com/tokentransfer/node/account/gmt"

	libaccount "github.com/tokentransfer/interfaces/account"
	libcore "github.com/tokentransfer/interfaces/core"
)

const (
	BTC libaccount.KeyType = 1
	ETH libaccount.KeyType = 2
	GMT libaccount.KeyType = 3
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
	GMT.Register("GMT", func(t libaccount.KeyType, seed []byte) libaccount.Key {
		key, _ := gmt.GenerateFamilySeed(string(seed))
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
	list := strings.Split(password, ".")
	var t libaccount.KeyType
	var e error
	var p string
	if len(list) != 2 {
		t = service.defaultType
		e = nil
		p = password
	} else {
		keyTypeName := list[0]
		p = list[1]
		t, e = libaccount.GetKeyTypeByName(keyTypeName)
	}
	if e != nil {
		return libaccount.KeyType(0), nil, e
	}
	k, err := GenerateFamilySeed(t, p)
	if err != nil {
		return libaccount.KeyType(0), nil, err
	}
	return t, k, nil
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
