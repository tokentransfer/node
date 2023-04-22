package btc

import (
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/base58"

	libaccount "github.com/tokentransfer/interfaces/account"
	libcore "github.com/tokentransfer/interfaces/core"
	"github.com/tokentransfer/node/util"
)

func GenerateFamilySeed(password string) (*Key, error) {
	b := util.Sha512Half([]byte(password))
	k := &Key{}
	err := k.UnmarshalBinary(b)
	if err != nil {
		return nil, err
	}
	return k, nil
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
	*btcec.PrivateKey
}

func (p *Private) UnmarshalBinary(data []byte) error {
	privKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), data)
	p.PrivateKey = privKey
	return nil
}

func (p *Private) MarshalBinary() ([]byte, error) {
	return p.PrivateKey.Serialize(), nil
}

func (p *Private) UnmarshalText(b []byte) error {
	data := base58.Decode(string(b))
	return p.UnmarshalBinary(data)
}

func (p *Private) MarshalText() ([]byte, error) {
	data, err := p.MarshalBinary()
	if err != nil {
		return nil, err
	}
	s := base58.Encode(data)
	return []byte(s), nil
}

func (p *Private) Sign(hash libcore.Hash, msg []byte) (libcore.Signature, error) {
	hashBytes := []byte(hash)
	sig, err := p.PrivateKey.Sign(hashBytes)
	if err != nil {
		return nil, err
	}
	return libcore.Signature(sig.Serialize()), nil
}

func (p *Private) GeneratePublic() (libaccount.PublicKey, error) {
	return &Public{p.PrivateKey.PubKey()}, nil
}

func (p *Private) GetSecret() (string, error) {
	b, err := p.MarshalText()
	return string(b), err
}

type Public struct {
	*btcec.PublicKey
}

func (p *Public) UnmarshalBinary(data []byte) error {
	pk, err := btcec.ParsePubKey(data, btcec.S256())
	if err != nil {
		return err
	}
	p.PublicKey = pk
	return nil
}

func (p *Public) MarshalBinary() ([]byte, error) {
	return p.SerializeCompressed(), nil
}

func (p *Public) UnmarshalText(b []byte) error {
	data := base58.Decode(string(b))
	return p.UnmarshalBinary(data)
}

func (p *Public) MarshalText() ([]byte, error) {
	data, err := p.MarshalBinary()
	if err != nil {
		return nil, err
	}
	s := base58.Encode(data)
	return []byte(s), nil
}

func (p *Public) Verify(hash libcore.Hash, msg []byte, signature libcore.Signature) (bool, error) {
	signatureBytes := []byte(signature)
	sig, err := btcec.ParseDERSignature(signatureBytes, btcec.S256())
	if err != nil {
		return false, err
	}
	hashBytes := []byte(hash)
	return sig.Verify(hashBytes, p.PublicKey), nil
}

func (p *Public) GenerateAddress() (libcore.Address, error) {
	data, err := p.MarshalBinary()
	if err != nil {
		return nil, err
	}
	h := btcutil.Hash160(data)
	a := Address{}
	err = a.UnmarshalBinary(h[:])
	if err != nil {
		return nil, err
	}
	return &a, nil
}

type Address [20]byte

func (a *Address) UnmarshalText(b []byte) error {
	addr, err := btcutil.DecodeAddress(string(b), &chaincfg.MainNetParams)
	if err != nil {
		return err
	}
	address, ok := addr.(*btcutil.AddressPubKeyHash)
	if !ok {
		return util.ErrorOf("unsupported", "address format", addr.String())
	}
	copy(a[:], address.Hash160()[:])
	return nil
}

func (a Address) MarshalText() ([]byte, error) {
	addr, err := btcutil.NewAddressPubKeyHash(a[:], &chaincfg.MainNetParams)
	if err != nil {
		return nil, err
	}
	s := addr.EncodeAddress()
	return []byte(s), nil
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
