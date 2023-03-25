package vm

import (
	"fmt"
	"io"
	"os"
	"testing"

	. "github.com/tokentransfer/check"
)

type WasmSuite struct{}

func Test_Config(t *testing.T) {
	s := Suite(&WasmSuite{})
	TestingRun(t, s)
}

func (suite *WasmSuite) TestVMLib(c *C) {
	cost := int64(100)
	f, err := os.Open("./testdata/lib.wasm")
	c.Assert(err, IsNil)
	wasmBytes, err := io.ReadAll(f)
	c.Assert(err, IsNil)
	remainCost, results, err := RunWasm(100, wasmBytes, "add", [][]byte{
		[]byte("4"),
		[]byte("3"),
	})
	c.Assert(err, IsNil)
	c.Assert(len(results), Equals, 1)
	c.Assert(results, DeepEquals, []byte("7"))
	fmt.Println(cost, remainCost, cost-remainCost)
}

func (suite *WasmSuite) TestVMDemo(c *C) {
	cost := int64(100)
	f, err := os.Open("./testdata/demo.wasm")
	c.Assert(err, IsNil)
	wasmBytes, err := io.ReadAll(f)
	c.Assert(err, IsNil)
	remainCost, results, err := RunWasm(cost, wasmBytes, "add", [][]byte{
		[]byte("2"),
		[]byte("3"),
	})
	c.Assert(err, IsNil)
	c.Assert(len(results), Equals, 1)
	c.Assert(results, DeepEquals, []byte("5"))
	fmt.Println(cost, remainCost, cost-remainCost)
}
