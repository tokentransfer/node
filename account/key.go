package account

import (
	"fmt"
	"strings"

	"github.com/tokentransfer/node/util"

	libaccount "github.com/tokentransfer/interfaces/account"
	libcore "github.com/tokentransfer/interfaces/core"
)

type Key struct {
	privateKey *Private
	publicKey  *Public
	address    *Address
}

func GenerateFamilySeed(t libaccount.KeyType, password string) (libaccount.Key, error) {
	k := t.CreateWith([]byte(password))
	p, err := k.GetPrivate()
	if err != nil {
		return nil, err
	}
	return &Key{
		privateKey: &Private{
			keyType:    t,
			PrivateKey: p,
		},
	}, nil
}

func (p *Key) GetKeyType() libaccount.KeyType {
	return p.privateKey.GetKeyType()
}

func (p *Key) UnmarshalBinary(data []byte) error {
	priv := &Private{}
	err := priv.UnmarshalBinary(data)
	if err != nil {
		return err
	}
	p.privateKey = priv
	return nil
}

func (p *Key) MarshalBinary() ([]byte, error) {
	return p.privateKey.MarshalBinary()
}

func (p *Key) UnmarshalText(b []byte) error {
	priv := &Private{}
	err := priv.UnmarshalText(b)
	if err != nil {
		return err
	}
	p.privateKey = priv
	return nil
}

func (p *Key) MarshalText() ([]byte, error) {
	return p.privateKey.MarshalText()
}

func (p *Key) Sign(hash libcore.Hash, msg []byte) (libcore.Signature, error) {
	return p.privateKey.Sign(hash, msg)
}

func (p *Key) Verify(hash libcore.Hash, msg []byte, signature libcore.Signature) (bool, error) {
	publicKey, err := p.GetPublic()
	if err != nil {
		return false, err
	}
	return publicKey.Verify(hash, msg, signature)
}

func (p *Key) GetPrivate() (libaccount.PrivateKey, error) {
	return p.privateKey, nil
}

func (p *Key) GetPublic() (libaccount.PublicKey, error) {
	if p.publicKey == nil {
		pk, err := p.privateKey.GeneratePublic()
		if err != nil {
			return nil, err
		}
		p.publicKey = pk.(*Public)
	}
	return p.publicKey, nil
}

func (p *Key) GetAddress() (libcore.Address, error) {
	_, err := p.GetPublic()
	if err != nil {
		return nil, err
	}
	if p.address == nil {
		a, err := p.publicKey.GenerateAddress()
		if err != nil {
			return nil, err
		}
		p.address = a.(*Address)
	}
	return p.address, nil
}

type Private struct {
	keyType libaccount.KeyType
	libaccount.PrivateKey
}

func (p *Private) GetKeyType() libaccount.KeyType {
	return p.keyType
}

func (p *Private) UnmarshalBinary(data []byte) error {
	if len(data) < 2 {
		return util.ErrorOfInvalid("data length", fmt.Sprintf("%d", len(data)))
	}
	t := libaccount.KeyType(util.BYTE_ORDER.Uint16(data))
	k := t.Create()
	err := k.UnmarshalBinary(data[2:])
	if err != nil {
		return err
	}
	priv, err := k.GetPrivate()
	if err != nil {
		return err
	}
	p.keyType = t
	p.PrivateKey = priv
	return nil
}

func (p *Private) MarshalBinary() ([]byte, error) {
	data, err := p.PrivateKey.MarshalBinary()
	if err != nil {
		return nil, err
	}
	data = append([]byte{0x00, 0x00}, data...)
	util.BYTE_ORDER.PutUint16(data, uint16(p.keyType))
	return data, nil
}

func (p *Private) UnmarshalText(b []byte) error {
	s := string(b)
	list := strings.Split(s, ".")
	if len(list) != 2 {
		return util.ErrorOfInvalid("private key", s)
	}
	keyTypeName := list[0]
	t, err := libaccount.GetKeyTypeByName(keyTypeName)
	if err != nil {
		return err
	}
	k := t.Create()
	err = k.UnmarshalText([]byte(list[1]))
	if err != nil {
		return err
	}
	priv, err := k.GetPrivate()
	if err != nil {
		return err
	}
	p.keyType = t
	p.PrivateKey = priv
	return nil
}

func (p *Private) MarshalText() ([]byte, error) {
	keyTypeName := strings.ToLower(p.keyType.String())
	secret, err := p.PrivateKey.MarshalText()
	if err != nil {
		return nil, err
	}
	text := strings.Join([]string{keyTypeName, string(secret)}, ".")
	return []byte(text), nil
}

func (p *Private) Sign(hash libcore.Hash, msg []byte) (libcore.Signature, error) {
	hashBytes := []byte(hash)
	sig, err := p.PrivateKey.Sign(hashBytes, msg)
	if err != nil {
		return nil, err
	}
	return sig, nil
}

func (p *Private) GeneratePublic() (libaccount.PublicKey, error) {
	pk, err := p.PrivateKey.GeneratePublic()
	if err != nil {
		return nil, err
	}
	return &Public{
		keyType:   p.keyType,
		PublicKey: pk,
	}, nil
}

func (p *Private) GetSecret() (string, error) {
	b, err := p.MarshalText()
	return string(b), err
}

type Public struct {
	keyType libaccount.KeyType
	libaccount.PublicKey
}

func (p *Public) GetKeyType() libaccount.KeyType {
	return p.keyType
}

func (p *Public) UnmarshalBinary(data []byte) error {
	if len(data) < 2 {
		return util.ErrorOfInvalid("data length", fmt.Sprintf("%d", len(data)))
	}
	t := libaccount.KeyType(util.BYTE_ORDER.Uint16(data))
	k := t.Create()
	pk, err := k.GetPublic()
	if err != nil {
		return err
	}
	err = pk.UnmarshalBinary(data[2:])
	if err != nil {
		return err
	}
	p.keyType = t
	p.PublicKey = pk
	return nil
}

func (p *Public) MarshalBinary() ([]byte, error) {
	data, err := p.PublicKey.MarshalBinary()
	if err != nil {
		return nil, err
	}
	data = append([]byte{0x00, 0x00}, data...)
	util.BYTE_ORDER.PutUint16(data, uint16(p.keyType))
	return data, nil
}

func (p *Public) UnmarshalText(b []byte) error {
	s := string(b)
	list := strings.Split(s, ".")
	if len(list) != 2 {
		return util.ErrorOfInvalid("public key", s)
	}
	keyTypeName := list[0]
	t, err := libaccount.GetKeyTypeByName(keyTypeName)
	if err != nil {
		return err
	}
	k := t.Create()
	pk, err := k.GetPublic()
	if err != nil {
		return err
	}
	err = pk.UnmarshalText(b[4:])
	if err != nil {
		return err
	}
	p.keyType = t
	p.PublicKey = pk
	return nil
}

func (p *Public) MarshalText() ([]byte, error) {
	keyTypeName := strings.ToLower(p.keyType.String())
	publicKeyString, err := p.PublicKey.MarshalText()
	if err != nil {
		return nil, err
	}
	text := strings.Join([]string{keyTypeName, string(publicKeyString)}, ".")
	return []byte(text), nil
}

func (p *Public) Verify(hash libcore.Hash, msg []byte, signature libcore.Signature) (bool, error) {
	return p.PublicKey.Verify(hash, msg, signature)
}

func (p *Public) GenerateAddress() (libcore.Address, error) {
	pka, err := p.PublicKey.GenerateAddress()
	if err != nil {
		return nil, err
	}
	data, err := pka.MarshalBinary()
	if err != nil {
		return nil, err
	}
	data = append([]byte{0x00, 0x00}, data...)
	util.BYTE_ORDER.PutUint16(data, uint16(p.keyType))
	a := Address{}
	err = a.UnmarshalBinary(data)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

type Address [22]byte

func (a *Address) GetKeyType() libaccount.KeyType {
	data := a[:]
	t := libaccount.KeyType(util.BYTE_ORDER.Uint16(data))
	return t
}

func (a *Address) UnmarshalText(b []byte) error {
	s := string(b)
	list := strings.Split(s, ".")
	if len(list) != 2 {
		return util.ErrorOfInvalid("address", s)
	}
	keyTypeName := list[0]
	t, err := libaccount.GetKeyTypeByName(keyTypeName)
	if err != nil {
		return err
	}
	k := t.Create()
	pka, err := k.GetAddress()
	if err != nil {
		return err
	}
	err = pka.UnmarshalText(b[4:])
	if err != nil {
		return err
	}
	data, err := pka.MarshalBinary()
	if err != nil {
		return err
	}
	data = append([]byte{0x00, 0x00}, data...)
	util.BYTE_ORDER.PutUint16(data, uint16(t))
	copy(a[:], data)
	return nil
}

func (a Address) MarshalText() ([]byte, error) {
	data := a[:]
	t := libaccount.KeyType(util.BYTE_ORDER.Uint16(data))
	keyTypeName := strings.ToLower(t.String())

	k := t.Create()
	addr, err := k.GetAddress()
	if err != nil {
		return nil, err
	}
	err = addr.UnmarshalBinary(data[2:])
	if err != nil {
		return nil, err
	}
	address, err := addr.MarshalText()
	if err != nil {
		return nil, err
	}
	text := strings.Join([]string{keyTypeName, string(address)}, ".")

	return []byte(text), nil
}

func (a *Address) UnmarshalBinary(data []byte) error {
	copy(a[:], data)
	return nil
}

func (a Address) MarshalBinary() ([]byte, error) {
	return a[:], nil
}

func (a Address) String() string {
	b, _ := a.MarshalText()
	return string(b)
}
