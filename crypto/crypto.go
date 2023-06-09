package crypto

import (
	"crypto/sha256"
	"hash"

	"github.com/tokentransfer/node/account"
	"github.com/tokentransfer/node/util"

	libaccount "github.com/tokentransfer/interfaces/account"
	libcore "github.com/tokentransfer/interfaces/core"
	libcrypto "github.com/tokentransfer/interfaces/crypto"
)

var as = account.NewAccountService()

type CryptoService struct {
}

func (service *CryptoService) GetSize() int {
	return 32
}

func (service *CryptoService) NewHasher() hash.Hash {
	return sha256.New()
}

func (service *CryptoService) Hash(msg []byte) (libcore.Hash, error) {
	h := service.NewHasher()
	h.Write(msg)
	b := h.Sum(nil)
	return libcore.Hash(b), nil
}

func (service *CryptoService) Raw(h libcrypto.Hashable, rt libcrypto.RawType) (libcore.Hash, []byte, error) {
	rawData, err := h.Raw(rt)
	if err != nil {
		return nil, nil, err
	}
	hashData, err := h.Raw(libcrypto.RawIgnoreVariableFields)
	if err != nil {
		return nil, nil, err
	}
	hash, err := service.Hash(hashData)
	if err != nil {
		return nil, nil, err
	}
	h.SetHash(hash)
	return hash, rawData, nil
}

func (service *CryptoService) Sign(p libaccount.Key, s libcrypto.Signable) error {
	publicKey, err := p.GetPublic()
	if err != nil {
		return err
	}
	publicBytes, err := publicKey.MarshalBinary()
	if err != nil {
		return err
	}
	s.SetPublicKey(libcore.PublicKey(publicBytes))

	data, err := s.Raw(libcrypto.RawIgnoreSigningFields)
	if err != nil {
		return err
	}
	hash, err := service.Hash(data)
	if err != nil {
		return err
	}
	signature, err := p.Sign(hash, data)
	if err != nil {
		return err
	}
	s.SetSignature(signature)

	h, _, err := service.Raw(s, libcrypto.RawIgnoreSigningFields)
	if err != nil {
		return err
	}
	s.SetHash(h)

	return nil
}

func (service *CryptoService) Verify(s libcrypto.Signable) (bool, error) {
	publicBytes := []byte(s.GetPublicKey())

	_, p, err := as.NewPublicFromBytes(publicBytes)
	if err != nil {
		return false, err
	}
	a, err := p.GenerateAddress()
	if err != nil {
		return false, err
	}
	if !libcore.Equals(a, s.GetAccount()) {
		return false, util.ErrorOfUnmatched("address", "in public key", s.GetAccount().String(), a.String())
	}
	data, err := s.Raw(libcrypto.RawIgnoreSigningFields)
	if err != nil {
		return false, err
	}
	hash, err := service.Hash(data)
	if err != nil {
		return false, err
	}
	signature := s.GetSignature()
	return p.Verify(hash, data, signature)
}
