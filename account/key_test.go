package account

import (
	"fmt"
	"math/big"
	"testing"

	. "github.com/tokentransfer/check"
)

type KeySuite struct{}

func Test_Key(t *testing.T) {
	s := Suite(&KeySuite{})
	TestingRun(t, s)
}

func (suite *KeySuite) TestAccount(c *C) {
	// s := "btc.FzxtNfYryNN11BnoRGsmfpT1JfeNVwC62vYi5RbCsTZj" //btc.1CeCLCBm3isypgFCh3GgV8v2yyvP8t1XuS
	s := "eth.beec9ec61c17b04cb9e4a9b7017e749f92835e2743e95f94cde218d667b14109" //eth.0x6da68a0c5dAAE0715AE6b62F00f548A2C6981c2f
	k := &Key{}
	err := k.UnmarshalText([]byte(s))
	// k, err := GenerateFamilySeed(BTC, "masterpassphrase")
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
