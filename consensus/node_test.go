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

	blob := "65080112160200229B160CF3D28F5EFF890EA1A8645F2622C4B33B18062206302F5445535428C09A0C32160200229B160CF3D28F5EFF890EA1A8645F2622C4B33B3AAD010AAA010A20194CF652D48DF6C77488CE476999E11CEBC0BD277384D93006CEBF2035BE0F7E1285018E120770726F636573731A0536D20400001A6B3E343130613332306130343663363937333734313232613061323834303061306230613039333730313030303030303030303030303030306130623061303933373032303030303030303030303030303030613062306130393337303630303030303030303030303030301A05362E16000042430200048DC56BD40C6DDACB8BA5972AE220F980F4DA7ECBA45D6BABBC65FC3A04CE2FEAA08FA11B26C74D9AB8AB0CC92E6DDE8ABD27835BDD35DA549E739640518F40924A9A01323038383931363430353833323132383739313836303834333931303734363731353037363235393133333230373139363031343632313437363534373234373039303734383331353830313A3835333635333436363538313033303536363832393835373534323938383236333330363332393039383937353035393039313433323534323238363237313139353935313030383938323237"
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
