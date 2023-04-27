package gmt

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/btcsuite/btcutil/bech32"

	"github.com/piligo/gmsm/sm2"

	libaccount "github.com/tokentransfer/interfaces/account"
	libcore "github.com/tokentransfer/interfaces/core"
	"github.com/tokentransfer/node/util"
)

func GenerateFamilySeed(password string) (*Key, error) {
	b := util.Sm3Half([]byte(password))
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
	*sm2.PrivateKey
}

func (p *Private) UnmarshalBinary(data []byte) error {
	d := new(big.Int).SetBytes(data)
	c := sm2.P256Sm2()
	privateKey := new(sm2.PrivateKey)
	privateKey.PublicKey.Curve = c
	privateKey.D = d
	privateKey.PublicKey.X, privateKey.PublicKey.Y = c.ScalarBaseMult(d.Bytes())
	p.PrivateKey = privateKey
	return nil
}

func (p *Private) MarshalBinary() ([]byte, error) {
	return p.PrivateKey.D.Bytes(), nil
}

func (p *Private) UnmarshalText(b []byte) error {
	_, _, data, err := decodeSegWitAddress(string(b))
	if err != nil {
		return err
	}
	return p.UnmarshalBinary(data)
}

func (p *Private) MarshalText() ([]byte, error) {
	data, err := p.MarshalBinary()
	if err != nil {
		return nil, err
	}
	s, err := encodeSegWitAddress("sc", byte(0x02), data)
	if err != nil {
		return nil, err
	}
	return []byte(s), nil
}

func (p *Private) Sign(hash libcore.Hash, msg []byte) (libcore.Signature, error) {
	hashBytes := []byte(hash)
	sig, err := p.PrivateKey.Sign(nil, hashBytes, nil)
	if err != nil {
		return nil, err
	}
	return libcore.Signature(sig), nil
}

func (p *Private) GeneratePublic() (libaccount.PublicKey, error) {
	pk := p.PrivateKey.Public().(*sm2.PublicKey)
	return &Public{pk}, nil
}

func (p *Private) GetSecret() (string, error) {
	b, err := p.MarshalText()
	return string(b), err
}

type Public struct {
	*sm2.PublicKey
}

func (p *Public) UnmarshalBinary(data []byte) error {
	pk, err := sm2.ParseSm2PublicKey(data)
	if err != nil {
		return err
	}
	p.PublicKey = pk
	return nil
}

func (p *Public) MarshalBinary() ([]byte, error) {
	data, err := sm2.MarshalSm2PublicKey(p.PublicKey)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (p *Public) UnmarshalText(b []byte) error {
	_, _, data, err := decodeSegWitAddress(string(b))
	if err != nil {
		return err
	}
	return p.UnmarshalBinary(data)
}

func (p *Public) MarshalText() ([]byte, error) {
	data, err := p.MarshalBinary()
	if err != nil {
		return nil, err
	}
	s, err := encodeSegWitAddress("pc", byte(0x01), data)
	if err != nil {
		return nil, err
	}
	return []byte(s), nil
}

func (p *Public) Verify(hash libcore.Hash, msg []byte, signature libcore.Signature) (bool, error) {
	signatureBytes := []byte(signature)
	ok := p.PublicKey.Verify(msg, signatureBytes)
	if !ok {
		return false, util.ErrorOfInvalid("signature", "verify failed")
	}
	return true, nil
}

func (p *Public) GenerateAddress() (libcore.Address, error) {
	data, err := p.MarshalBinary()
	if err != nil {
		return nil, err
	}
	h := util.Sm3Half(data)[:20]
	a := Address{}
	err = a.UnmarshalBinary(h[:])
	if err != nil {
		return nil, err
	}
	return &a, nil
}

type Address [20]byte

func (a *Address) UnmarshalText(b []byte) error {
	_, _, data, err := decodeSegWitAddress(string(b))
	if err != nil {
		return err
	}
	copy(a[:], data[:])
	return nil
}

func (a Address) MarshalText() ([]byte, error) {
	data, err := a.MarshalBinary()
	if err != nil {
		return nil, err
	}
	s, err := encodeSegWitAddress("bc", byte(0x00), data)
	if err != nil {
		return nil, err
	}
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

// reference from AddressWitnessPubKeyHash(github.com/btcsuite/btcutil/address.go)
// encodeSegWitAddress creates a bech32 encoded address string representation
// from witness version and witness program.
func encodeSegWitAddress(witnessHrp string, witnessVersion byte, witnessProgram []byte) (string, error) {
	// Group the address bytes into 5 bit groups, as this is what is used to
	// encode each character in the address string.
	converted, err := bech32.ConvertBits(witnessProgram, 8, 5, true)
	if err != nil {
		return "", err
	}

	// Concatenate the witness version and program, and encode the resulting
	// bytes using bech32 encoding.
	combined := make([]byte, len(converted)+1)
	combined[0] = witnessVersion
	copy(combined[1:], converted)
	bech, err := bech32.Encode(witnessHrp, combined)
	if err != nil {
		return "", err
	}

	// Check validity by decoding the created address.
	hrp, version, program, err := decodeSegWitAddress(bech)
	if err != nil {
		return "", util.ErrorOfInvalid("segwit address", err.Error())
	}

	if hrp != witnessHrp || version != witnessVersion || !bytes.Equal(program, witnessProgram) {
		return "", util.ErrorOfInvalid("segwit address", "unmatched")
	}

	return bech, nil
}

// decodeSegWitAddress parses a bech32 encoded segwit address string and
// returns the witness version and witness program byte representation.
func decodeSegWitAddress(address string) (string, byte, []byte, error) {
	// Decode the bech32 encoded address.
	hrp, data, err := bech32.Decode(address)
	if err != nil {
		return "", 0, nil, err
	}

	// The first byte of the decoded address is the witness version, it must
	// exist.
	if len(data) < 1 {
		return "", 0, nil, util.ErrorOfEmpty("witness", "version")
	}

	// ...and be <= 16.
	version := data[0]
	if version > 16 {
		return "", 0, nil, util.ErrorOfInvalid("witness version", fmt.Sprintf("%v", version))
	}

	// The remaining characters of the address returned are grouped into
	// words of 5 bits. In order to restore the original witness program
	// bytes, we'll need to regroup into 8 bit words.
	regrouped, err := bech32.ConvertBits(data[1:], 5, 8, false)
	if err != nil {
		return "", 0, nil, err
	}

	// The regrouped data must be between 2 and 40 bytes.
	if len(regrouped) < 2 || len(regrouped) > 40 {
		return "", 0, nil, util.ErrorOfInvalid("data length", fmt.Sprintf("%d", len(regrouped)))
	}

	// For witness version 0, address MUST be exactly 20 or 32 bytes.
	if version == 0 && len(regrouped) != 20 && len(regrouped) != 32 {
		return "", 0, nil, util.ErrorOfInvalid("data length for witness version 0", fmt.Sprintf("%d", len(regrouped)))
	}

	return hrp, version, regrouped, nil
}
