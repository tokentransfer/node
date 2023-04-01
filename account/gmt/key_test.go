package gmt

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
	s := "sc1zvylljq2kjv6wgctf5zrrgw2dz7chcemec3rse4jq4285qw4xxe8svqkuqy" //bc1qszykvnt4s3das7dex456m5euxff5eexetnxkrj
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
