package btc

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"

	. "github.com/tokentransfer/check"
	libcore "github.com/tokentransfer/interfaces/core"
)

type KeySuite struct{}

func Test_Key(t *testing.T) {
	s := Suite(&KeySuite{})
	TestingRun(t, s)
}

func (suite *KeySuite) TestAccount(c *C) {
	s := "FzxtNfYryNN11BnoRGsmfpT1JfeNVwC62vYi5RbCsTZj" //1CeCLCBm3isypgFCh3GgV8v2yyvP8t1XuS
	k := &Key{}
	err := k.UnmarshalText([]byte(s))
	// k, err := GenerateFamilySeed("masterpassphrase")
	if err != nil {
		panic(err)
	}
	a, err := k.GetAddress()
	if err != nil {
		panic(err)
	}
	address := a.String()
	p, err := k.GetPrivate()
	if err != nil {
		panic(err)
	}
	secret, err := p.GetSecret()
	if err != nil {
		panic(err)
	}
	fmt.Println("address", address)
	fmt.Println("secret", secret)

	data, err := a.MarshalBinary()
	if err != nil {
		panic(err)
	}
	i := new(big.Int)
	i.SetBytes(data)
	fmt.Println(address, i.Uint64())
}

func (suite *KeySuite) TestSignAndVerify(c *C) {
	msgBlob := "65080112146DA68A0C5DAAE0715AE6B62F00F548A2C6981C2F1801280A321442F32B004DA1093D51AE40A58F38E33BA4F46397424104D33A199D322FAFD28867E3B9FB2F5AC081D56CFF1AE803635730E1D01B77D837D9EE578346DD88B68D21B9A61B8F1EFE9B2574F08B4A471F864FA7EA7A29185C4A413C76F5EBE03A949D0776192F50C2284A65ECED1EB2DCA1C37719A30C650598D57DA6E63598223A881545D612F9B20B1CAC507AA8EAC458A39A347F0CE8A9720B00"
	msg, err := hex.DecodeString(msgBlob)
	c.Assert(err, IsNil)

	h := make([]byte, 32)
	n, err := rand.Read(h)
	c.Assert(err, IsNil)
	c.Assert(n, Equals, 32)

	k, err := GenerateFamilySeed("masterpassphrase")
	c.Assert(err, IsNil)

	pk, err := k.GetPublic()
	c.Assert(err, IsNil)

	pkdata, err := pk.MarshalBinary()
	c.Assert(err, IsNil)

	pk2 := &Public{}
	err = pk2.UnmarshalBinary(pkdata)
	c.Assert(err, IsNil)

	pk2data, err := pk2.MarshalBinary()
	c.Assert(err, IsNil)

	c.Assert(pkdata, DeepEquals, pk2data)

	fmt.Println("hash", libcore.Hash(h).String())
	sig, err := k.Sign(h, msg)
	c.Assert(err, IsNil)

	ok, err := pk.Verify(h, msg, sig)
	c.Assert(err, IsNil)
	c.Assert(ok, Equals, true)
}
