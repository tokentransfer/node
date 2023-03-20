package vm

import (
	"io"
	"os"
	"testing"

	. "github.com/tokentransfer/check"
)

type ConfigSuite struct{}

func Test_Config(t *testing.T) {
	s := Suite(&ConfigSuite{})
	TestingRun(t, s)
}

func (suite *ConfigSuite) TestVMLib(c *C) {
	f, err := os.Open("./testdata/lib.wasm")
	c.Assert(err, IsNil)
	wasmBytes, err := io.ReadAll(f)
	c.Assert(err, IsNil)
	results, err := RunWasm(wasmBytes, "add", 4, 3)
	c.Assert(err, IsNil)
	c.Assert(len(results), Equals, 1)
	c.Assert(results[0], Equals, uint64(7))
}

func (suite *ConfigSuite) TestVMDemo(c *C) {
	f, err := os.Open("./testdata/demo.wasm")
	c.Assert(err, IsNil)
	wasmBytes, err := io.ReadAll(f)
	c.Assert(err, IsNil)
	results, err := RunWasm(wasmBytes, "add", 2, 3)
	c.Assert(err, IsNil)
	c.Assert(len(results), Equals, 1)
	c.Assert(results[0], Equals, uint64(5))
}
