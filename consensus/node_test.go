package consensus

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"testing"

	"github.com/caivega/glog"
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
	dumpGroup(c, s, root, chunk.LogPrinter{})
}

func dumpGroup(c *C, s core.Storage, g core.Group, log core.Printer) {
	g.Dump(log)
	for index, item := range g.List() {
		ok, err := g.IsGroup(item)
		c.Assert(err, IsNil)
		if ok {
			sub, err := g.Group(item)
			c.Assert(err, IsNil)
			dumpGroup(c, s, sub, log)
		} else {
			switch g.Name() {
			case "code":
				dataKey, err := g.GetKey(item)
				c.Assert(err, IsNil)

				data, err := s.Get(dataKey)
				c.Assert(err, IsNil)
				reader := data.Open()
				wasmData, err := util.ReadBytes(reader)
				c.Assert(err, IsNil)
				abiData, err := util.ReadBytes(reader)
				c.Assert(err, IsNil)
				reader.Close()
				data.Dispose()

				fmt.Println(">", index, len(wasmData), len(abiData))

			default:
				dataKey, err := g.GetKey(item)
				c.Assert(err, IsNil)

				data, err := s.Get(dataKey)
				c.Assert(err, IsNil)
				reader := data.Open()
				buf := new(bytes.Buffer)
				_, err = io.Copy(buf, reader)
				c.Assert(err, IsNil)
				reader.Close()
				data.Dispose()
				dataBytes := buf.Bytes()

				fmt.Println(">", index, len(dataBytes))
				if len(dataBytes) > 0 {
					c.Assert(err, IsNil)
					o, err := core.AsData(dataBytes)
					if err != nil {
						fmt.Println(item, err, string(dataBytes))
					} else {
						util.PrintJSON(item, o)
					}
				}
			}
		}
	}
}

func (suite *NodeSuite) TestTransaction(c *C) {
	blobData, err := util.ReadFile("./data/tx.blob")
	c.Assert(err, IsNil)
	fmt.Println(len(blobData), string(blobData))
	data, err := hex.DecodeString(string(blobData))
	c.Assert(err, IsNil)

	tx := &block.Transaction{}
	err = tx.UnmarshalBinary(data)
	c.Assert(err, IsNil)
}

func (suite *NodeSuite) TestProcess(c *C) {
	config, err := config.NewConfig("../config.json")
	c.Assert(err, IsNil)
	// config.SetMode("debug")

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

	blobData, err := util.ReadFile("./data/tx.blob")
	c.Assert(err, IsNil)
	fmt.Println(len(blobData), string(blobData))
	data, err := hex.DecodeString(string(blobData))
	c.Assert(err, IsNil)

	tx := &block.Transaction{}
	err = tx.UnmarshalBinary(data)
	c.Assert(err, IsNil)

	err = n.verifyTransaction(tx)
	c.Assert(err, IsNil)
	_, _, err = n.processTransaction(tx)
	c.Assert(err, IsNil)

	block, err := n.GenerateBlock()
	c.Assert(err, IsNil)
	h, err := n.consensusService.HashBlock(block)
	c.Assert(err, IsNil)
	_, err = n.consensusService.VerifyBlock(block)
	c.Assert(err, IsNil)
	glog.Infof("=== generate block %d, %s, %d\n", block.GetIndex(), h.String(), len(block.GetTransactions()))
	util.PrintJSON(">> hash", h)
	// util.PrintJSON(">>> tx", tx)
	// util.PrintJSON(">>> txWithData", txWithData)
	// err = n.consensusService.AddBlock(block)
	// c.Assert(err, IsNil)

	dump(c, n.storageService.storage, "after")
}
