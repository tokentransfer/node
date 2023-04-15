package consensus

import (
	"encoding/hex"
	"fmt"
	"testing"

	libcrypto "github.com/tokentransfer/interfaces/crypto"
	"github.com/tokentransfer/node/block"
	"github.com/tokentransfer/node/chunk"
	"github.com/tokentransfer/node/config"
	"github.com/tokentransfer/node/core"
	"github.com/tokentransfer/node/util"

	. "github.com/tokentransfer/check"
)

type NodeSuite struct{}

func Test_Node(t *testing.T) {
	s := Suite(&NodeSuite{})
	TestingRun(t, s)
}

func dump(c *C, s core.Storage, name string) {
	fmt.Println("=================" + name + "=================")
	s.DumpLog(chunk.LogPrinter{})

	root, err := s.Group("/")
	c.Assert(err, IsNil)
	dumpGroup(c, root, chunk.LogPrinter{})
}

func dumpGroup(c *C, g core.Group, log core.Printer) {
	g.Dump(log)
	for _, item := range g.List() {
		ok, err := g.IsGroup(item)
		c.Assert(err, IsNil)
		if ok {
			sub, err := g.Group(item)
			c.Assert(err, IsNil)
			dumpGroup(c, sub, log)
		}
	}
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
		_, err = n.consensusService.HashBlock(block)
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

	dump(c, n.storageService.storage, "before")

	blob := "650801121602006DA68A0C5DAAE0715AE6B62F00F548A2C6981C2F180322083130302F54455354280A32160200396C520DD96C0DFBAE874713732C701C0F4D9EEF3A3A0A380A20F63546DCF74C70142978051086B94FC47F22D2E1B844AA149D2B71357910D93B12148E1A0361646422053603000000220536040000004243020004D33A199D322FAFD28867E3B9FB2F5AC081D56CFF1AE803635730E1D01B77D837D9EE578346DD88B68D21B9A61B8F1EFE9B2574F08B4A471F864FA7EA7A29185C4A9C013130303930343635393332303234373139343333333637303039333238343834383632323836373136393239303330393033373031353631343234343331383333353733323032393732353134373A3638343630333239353536343837373039343332333236313432383731303039363031363037393137313131393836353338313535343937393732393630323632333936323835323839393131"
	data, err := hex.DecodeString(blob)
	c.Assert(err, IsNil)

	tx := &block.Transaction{}
	err = tx.UnmarshalBinary(data)
	c.Assert(err, IsNil)

	h, _, err := n.cryptoService.Raw(tx, libcrypto.RawBinary)
	c.Assert(err, IsNil)
	ok, err := n.consensusService.VerifyTransaction(tx)
	c.Assert(err, IsNil)
	c.Assert(ok, Equals, true)

	err = n.storageService.CreateSandbox()
	c.Assert(err, IsNil)
	_, err = n.consensusService.ProcessTransaction(tx)
	c.Assert(err, IsNil)
	err = n.storageService.CommitSandbox()
	c.Assert(err, IsNil)

	util.PrintJSON(">> hash", h)
	// util.PrintJSON(">>> tx", tx)
	// util.PrintJSON(">>> txWithData", txWithData)

	dump(c, n.storageService.storage, "after")
}
