package consensus

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
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
				wasmData, err := core.ReadBytes(reader)
				c.Assert(err, IsNil)
				abiData, err := core.ReadBytes(reader)
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

				fmt.Println(">", index, len(dataBytes), string(dataBytes))
				if len(dataBytes) > 0 {
					bs, err := hex.DecodeString(string(dataBytes))
					c.Assert(err, IsNil)
					o, err := core.AsData(bs)
					c.Assert(err, IsNil)
					util.PrintJSON(item, o)
				}
			}
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

	blob := "650801121602006DA68A0C5DAAE0715AE6B62F00F548A2C6981C2F180822083130302F5445535428C09A0C3216020002834B2A40474D0AAB28EF195531F9F4354EE4523A5D0A5B0A201910E5D7D6386C7C5F5367358DDA50FD916F5F5E8B702EDDF5EC2D8BC710D88C12378E0A160200B5A1D2E92521249AB079341095CDAE44CD4998770A1602006B39BA86CAB636A780C9DD68D49B0F493793AA831A046C6973744243020004D33A199D322FAFD28867E3B9FB2F5AC081D56CFF1AE803635730E1D01B77D837D9EE578346DD88B68D21B9A61B8F1EFE9B2574F08B4A471F864FA7EA7A29185C4A9B0132313134333233313433363636383138303132353132373931373038363932323132303333383237363634353232393837343935353431303634373136363539343438393235383332393934343A3131353636353132333831393236333938303432343039353039353536383236313539313739363730393039393739383435323539383031323231303636373632323134353431353935373435"
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
