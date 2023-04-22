package eth

import (
	"crypto/ecdsa"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	libaccount "github.com/tokentransfer/interfaces/account"
	libcore "github.com/tokentransfer/interfaces/core"
)

func GenerateFamilySeed(password string) (*Key, error) {
	h := crypto.Keccak256([]byte(password))
	key := &Key{}
	err := key.UnmarshalBinary(h)
	if err != nil {
		return nil, err
	}
	return key, nil
}

type Key struct {
	privateKey Private
	publicKey  *Public
	address    *Address
}

func (p *Key) UnmarshalBinary(data []byte) error {
	return p.privateKey.UnmarshalBinary(data)
}

func (p *Key) MarshalBinary() ([]byte, error) {
	return p.privateKey.MarshalBinary()
}

func (p *Key) UnmarshalText(b []byte) error {
	return p.privateKey.UnmarshalText(b)
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
	return &p.privateKey, nil
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
	*ecdsa.PrivateKey
}

func (p *Private) UnmarshalBinary(data []byte) error {
	privateKey, err := crypto.ToECDSA(data)
	if err != nil {
		return err
	}
	p.PrivateKey = privateKey
	return nil
}

func (p *Private) MarshalBinary() ([]byte, error) {
	data := crypto.FromECDSA(p.PrivateKey)
	return data, nil
}

func (p *Private) UnmarshalText(b []byte) error {
	data := common.FromHex(string(b))
	return p.UnmarshalBinary(data)
}

func (p *Private) MarshalText() ([]byte, error) {
	data, err := p.MarshalBinary()
	if err != nil {
		return nil, err
	}
	s := hex.EncodeToString(data)
	return []byte(s), nil
}

func (p *Private) Sign(hash libcore.Hash, msg []byte) (libcore.Signature, error) {
	hashBytes := []byte(hash)

	sig, err := crypto.Sign(hashBytes, p.PrivateKey)
	if err != nil {
		return nil, err
	}
	return libcore.Signature(sig[:64]), nil
}

func (p *Private) GeneratePublic() (libaccount.PublicKey, error) {
	publicKey := p.Public()
	return &Public{publicKey.(*ecdsa.PublicKey)}, nil
}

func (p *Private) GetSecret() (string, error) {
	b, err := p.MarshalText()
	return string(b), err
}

type Public struct {
	*ecdsa.PublicKey
}

func (p *Public) UnmarshalBinary(data []byte) error {
	pk, err := crypto.UnmarshalPubkey(data)
	if err != nil {
		return err
	}
	p.PublicKey = pk
	return nil
}

func (p *Public) MarshalBinary() ([]byte, error) {
	return crypto.FromECDSAPub(p.PublicKey), nil
}

func (p *Public) UnmarshalText(b []byte) error {
	data := common.FromHex(string(b))
	return p.UnmarshalBinary(data)
}

func (p *Public) MarshalText() ([]byte, error) {
	data, err := p.MarshalBinary()
	if err != nil {
		return nil, err
	}
	s := hex.EncodeToString(data)
	return []byte(s), nil
}

func (p *Public) Verify(hash libcore.Hash, msg []byte, signature libcore.Signature) (bool, error) {
	pubk := crypto.FromECDSAPub(p.PublicKey)
	ok := crypto.VerifySignature(pubk, hash, []byte(signature))
	return ok, nil
}

func (p *Public) GenerateAddress() (libcore.Address, error) {
	a := crypto.PubkeyToAddress(*p.PublicKey)

	ra := Address{}
	err := ra.UnmarshalBinary(a[:])
	if err != nil {
		return nil, err
	}
	return &ra, nil
}

type Address [20]byte

func (a *Address) UnmarshalText(b []byte) error {
	address := common.HexToAddress(string(b))
	copy(a[:], address[:])
	return nil
}

func (a Address) MarshalText() ([]byte, error) {
	address := common.BytesToAddress(a[:])
	return []byte(address.Hex()), nil
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
