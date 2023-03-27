package consensus

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/tokentransfer/node/block"
	"github.com/tokentransfer/node/chunk"
	"github.com/tokentransfer/node/config"
	"github.com/tokentransfer/node/core"

	. "github.com/tokentransfer/check"
)

type NodeSuite struct{}

func Test_Node(t *testing.T) {
	s := Suite(&NodeSuite{})
	TestingRun(t, s)
}

func dump(s core.Storage, name string) {
	fmt.Println("=================" + name + "=================")
	s.DumpLog(chunk.LogPrinter{})
}

func (suite *NodeSuite) TestProcess(c *C) {
	config, err := config.NewConfig("../config.json")
	c.Assert(err, IsNil)

	n := NewNode()
	err = n.Init(config)
	c.Assert(err, IsNil)

	n.Load()

	if n.GetBlockNumber() < 0 {
		block, err := n.GenerateBlock()
		if err != nil {
			panic(err)
		}
		_, err = n.HashBlock(block)
		if err != nil {
			panic(err)
		}
		_, err = n.VerifyBlock(block)
		if err != nil {
			panic(err)
		}
		err = n.AddBlock(block)
		if err != nil {
			panic(err)
		}
		fmt.Printf("=== generate block %d, %s, %d\n", block.GetIndex(), block.GetHash().String(), len(block.GetTransactions()))
	}

	dump(n.storageService.storage, "before")

	blob := "65080112146DA68A0C5DAAE0715AE6B62F00F548A2C6981C2F180122083130302F54455354280A321442F32B004DA1093D51AE40A58F38E33BA4F463973ADA018D0AD6018E22D2010061736D01000000010D0260027F7F017F60027D7D017D03030200010503010011071603066D656D6F727902000361646400000373756D00010A11020700200020016A0B070020002001920B0083010970726F64756365727302086C616E6775616765010452757374000C70726F6365737365642D62790305727573746325312E37302E302D6E696768746C79202833396632363537643120323032332D30332D3039290677616C72757306302E31392E300C7761736D2D62696E6467656E12302E322E3834202863656138636333643229424104D33A199D322FAFD28867E3B9FB2F5AC081D56CFF1AE803635730E1D01B77D837D9EE578346DD88B68D21B9A61B8F1EFE9B2574F08B4A471F864FA7EA7A29185C4A9C013130343534383630353634383138363432383833313136383033323331353139343133373539303231303939363636333931373839353730333130363634373834313133393830363737363635383A3337313239313831393338323832353235353339383431393232303839383636373830363933383334363033313238363934303839373236363237383134373235313837363936353535313633"
	data, err := hex.DecodeString(blob)
	c.Assert(err, IsNil)

	tx := &block.Transaction{}
	err = tx.UnmarshalBinary(data)
	c.Assert(err, IsNil)
	_, _, err = n.processTransaction(tx)
	c.Assert(err, IsNil)

	// util.PrintJSON(">>> tx", tx)
	// util.PrintJSON(">>> txWithData", txWithData)

	dump(n.storageService.storage, "after")
}
