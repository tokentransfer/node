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

	dump(n.storageService.storage, "before")

	blob := "650801121602006DA68A0C5DAAE0715AE6B62F00F548A2C6981C2F180422083130302F54455354280A32160200396C520DD96C0DFBAE874713732C701C0F4D9EEF3A410A3F0A20B8043A3EE92BEA2B433A123432AA438F2B48B298B005541160632CC273F7CB99121B8E120770726F636573731A0536D20400001A013E1A05362E1600004243020004D33A199D322FAFD28867E3B9FB2F5AC081D56CFF1AE803635730E1D01B77D837D9EE578346DD88B68D21B9A61B8F1EFE9B2574F08B4A471F864FA7EA7A29185C4A9B0134373035303438313530363134373637393036323239323430383737373332303136333531313933323733303931393034353735343839303630363033353437383738343336333333313236303A3234373839333530303938343830363438373738333331373231313230303337343238383931363830393031333639363034333431343139323530353836373539333931303830363039353231"
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

	dump(n.storageService.storage, "after")
}
