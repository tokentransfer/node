package vm

import (
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"testing"

	. "github.com/tokentransfer/check"
	"github.com/tokentransfer/node/core"
	"github.com/tokentransfer/node/core/pb"
)

type WasmSuite struct{}

func Test_Config(t *testing.T) {
	s := Suite(&WasmSuite{})
	TestingRun(t, s)
}

func getWasmData(c *C) []byte {
	data, err := hex.DecodeString("410a0d0a02763012070a0536d20400000a0d0a02763212070a05362e1600000a1d0a046c69737412150a13400a070a0536d20400000a070a05362e160000")
	c.Assert(err, IsNil)
	return data
}

func (suite *WasmSuite) TestVMMain(c *C) {
	cost := int64(1000000)
	f, err := os.Open("./testdata/main.wasm")
	c.Assert(err, IsNil)
	wasmCode, err := io.ReadAll(f)
	c.Assert(err, IsNil)
	wasmData := getWasmData(c)
	usedCost, newWasmData, results, err := RunWasm(cost, wasmCode, wasmData, "add", [][]byte{
		[]byte("5"),
		[]byte("6"),
	})
	c.Assert(err, IsNil)
	c.Assert(len(results), Equals, 2)
	c.Assert(results, DeepEquals, []byte("11"))
	meta, msg, err := core.Unmarshal(newWasmData)
	c.Assert(err, IsNil)
	c.Assert(meta, Equals, core.CORE_DATA_MAP)
	m := msg.(*pb.DataMap)
	c.Assert(m, NotNil)
	fmt.Println(cost, usedCost, cost-usedCost)
}

func (suite *WasmSuite) TestVMLib(c *C) {
	cost := int64(100)
	f, err := os.Open("./testdata/lib.wasm")
	c.Assert(err, IsNil)
	wasmCode, err := io.ReadAll(f)
	c.Assert(err, IsNil)
	wasmData := getWasmData(c)
	usedCost, newWasmData, results, err := RunWasm(cost, wasmCode, wasmData, "add", [][]byte{
		[]byte("4"),
		[]byte("3"),
	})
	c.Assert(err, IsNil)
	c.Assert(len(results), Equals, 1)
	c.Assert(results, DeepEquals, []byte("7"))
	meta, msg, err := core.Unmarshal(newWasmData)
	c.Assert(err, IsNil)
	c.Assert(meta, Equals, core.CORE_DATA_MAP)
	m := msg.(*pb.DataMap)
	c.Assert(m, NotNil)
	fmt.Println(cost, usedCost, cost-usedCost)
}

func (suite *WasmSuite) TestVMDemo(c *C) {
	cost := int64(100)
	f, err := os.Open("./testdata/demo.wasm")
	c.Assert(err, IsNil)
	wasmCode, err := io.ReadAll(f)
	c.Assert(err, IsNil)
	wasmData := getWasmData(c)
	usedCost, newWasmData, results, err := RunWasm(cost, wasmCode, wasmData, "add", [][]byte{
		[]byte("2"),
		[]byte("3"),
	})
	c.Assert(err, IsNil)
	c.Assert(len(results), Equals, 1)
	c.Assert(results, DeepEquals, []byte("5"))
	meta, msg, err := core.Unmarshal(newWasmData)
	c.Assert(err, IsNil)
	c.Assert(meta, Equals, core.CORE_DATA_MAP)
	m := msg.(*pb.DataMap)
	c.Assert(m, NotNil)
	fmt.Println(cost, usedCost, cost-usedCost)
}

func (suite *WasmSuite) TestVMBg(c *C) {
	cost := int64(100)
	f, err := os.Open("./testdata/wasm_lib_bg.wasm")
	c.Assert(err, IsNil)
	wasmCode, err := io.ReadAll(f)
	c.Assert(err, IsNil)
	wasmData := getWasmData(c)
	usedCost, newWasmData, results, err := RunWasm(cost, wasmCode, wasmData, "process", [][]byte{
		[]byte("410a0d0a02763012070a0536d20400000a0d0a02763212070a05362e1600000a1d0a046c69737412150a13400a070a0536d20400000a070a05362e160000"),
		// []byte("410a0d0a02763012070a0536d20400000a0d0a02763212070a05362f1600000a1d0a046c69737412150a13400a070a0536d20400000a070a05362f160000"),
	})
	c.Assert(err, IsNil)
	c.Assert(len(results), Equals, 1)
	c.Assert(results, DeepEquals, []byte("5"))
	meta, msg, err := core.Unmarshal(newWasmData)
	c.Assert(err, IsNil)
	c.Assert(meta, Equals, core.CORE_DATA_MAP)
	m := msg.(*pb.DataMap)
	c.Assert(m, NotNil)
	fmt.Println(cost, usedCost, cost-usedCost)
}
