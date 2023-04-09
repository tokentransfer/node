package vm

import (
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
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
	// []byte("410a0d0a02763012070a0536d20400000a0d0a02763212070a05362e1600000a1d0a046c69737412150a13400a070a0536d20400000a070a05362e160000"),
	// []byte("410a0d0a02763012070a0536d20400000a0d0a02763212070a05362f1600000a1d0a046c69737412150a13400a070a0536d20400000a070a05362f160000"),
	data, err := hex.DecodeString("410a0d0a02763012070a0536d20400000a0d0a02763212070a05362e1600000a1d0a046c69737412150a13400a070a0536d20400000a070a05362e160000")
	c.Assert(err, IsNil)
	return data
}

func loadWasm(c *C, prefix string) ([]byte, []byte) {
	wf, err := os.Open(fmt.Sprintf("%s.wasm", prefix))
	c.Assert(err, IsNil)
	wasmCode, err := io.ReadAll(wf)
	c.Assert(err, IsNil)

	af, err := os.Open(fmt.Sprintf("%s.json", prefix))
	c.Assert(err, IsNil)
	abiData, err := io.ReadAll(af)
	c.Assert(err, IsNil)

	return wasmCode, abiData
}

func getParams(c *C, params []interface{}) ([]string, [][]byte) {
	list := make([][]byte, 0)
	rets := make([]string, 0)
	for _, item := range params {
		data, err := core.MarshalData(item)
		c.Assert(err, IsNil)

		list = append(list, data)
		rets = append(rets, fmt.Sprintf("%v", item))
	}
	return rets, list
}

func getResult(c *C, retData []byte) (interface{}, string) {
	_, value, err := core.UnmarshalData(retData)
	c.Assert(err, IsNil)
	return value, fmt.Sprintf("%s:%v(%d)", core.GetInfo(retData), value, len(retData))
}

func (suite *WasmSuite) testVMMain(c *C) {
	cost := int64(100)
	wasmData := getWasmData(c)

	wasmCode, abiCode := loadWasm(c, "./testdata/main_wasm")
	rets, paramData := getParams(c, []interface{}{
		uint32(5),
		uint32(6),
	})
	fmt.Println("params: ", strings.Join(rets, ","))
	usedCost, newWasmData, results, err := RunWasm(cost, wasmCode, abiCode, wasmData, "add", paramData)
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

func (suite *WasmSuite) testVMLib(c *C) {
	cost := int64(100)
	wasmCode, abiData := loadWasm(c, "./testdata/lib")
	wasmData := getWasmData(c)
	rets, paramData := getParams(c, []interface{}{
		uint32(4),
		uint32(3),
	})
	fmt.Println("params: ", strings.Join(rets, ","))
	usedCost, newWasmData, results, err := RunWasm(cost, wasmCode, abiData, wasmData, "add", paramData)
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

func (suite *WasmSuite) testVMDemo(c *C) {
	cost := int64(100)
	wasmCode, abiData := loadWasm(c, "./testdata/demo")
	wasmData := getWasmData(c)
	rets, paramData := getParams(c, []interface{}{
		uint32(4),
		uint32(3),
	})
	fmt.Println("params: ", strings.Join(rets, ","))
	usedCost, newWasmData, results, err := RunWasm(cost, wasmCode, abiData, wasmData, "add", paramData)
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

func (suite *WasmSuite) TestVerify(c *C) {
	wasmCode, abiData := loadWasm(c, "./testdata/wasm_lib_bg")
	err := VerifyWasm(wasmCode, abiData)
	c.Assert(err, IsNil)
}

func (suite *WasmSuite) TestVMBg(c *C) {
	cost := int64(10000)
	wasmCode, abiData := loadWasm(c, "./testdata/wasm_lib_bg")

	cases := []map[string]interface{}{
		{
			"method": "add",
			"params": []interface{}{
				int32(4),
				int32(3),
			},
			"result": int32(7),
		},
		{
			"method": "test_string_i32_ret_string",
			"params": []interface{}{
				"hello",
				int32(3),
			},
			"result": "hello-3",
		},
		{
			"method": "test_i32_string",
			"params": []interface{}{
				int32(3),
				"hello",
			},
			"result": int32(8),
		},
		{
			"method": "test_string_i32",
			"params": []interface{}{
				"hello",
				int32(3),
			},
			"result": int32(8),
		},
		{
			"method": "test_i64_string_i32",
			"params": []interface{}{
				int64(4),
				"hello",
				int32(3),
			},
			"result": int64(12),
		},
		{
			"method": "test_string_i64_i32",
			"params": []interface{}{
				"hello",
				int64(4),
				int32(3),
			},
			"result": int32(12),
		},
		{
			"method": "test_string_i64_i32_ret_string",
			"params": []interface{}{
				"hello",
				int64(4),
				int32(3),
			},
			"result": "5-4-3",
		},
		{
			"method": "test_bytes_i64_i32_ret_string",
			"params": []interface{}{
				[]byte("hello"),
				int64(4),
				int32(3),
			},
			"result": "5-4-3",
		},
		{
			"method": "test_string_string_ret_string",
			"params": []interface{}{
				"hello",
				"world",
			},
			"result": "hello-world",
		},
		{
			"method": "test_f32_f64_ret_f64",
			"params": []interface{}{
				float32(2.5),
				float64(3.4),
			},
			"result": float64(5.9),
		},
	}
	for index, testCase := range cases {
		fmt.Println("======", index, "======")
		method := testCase["method"].(string)
		params := testCase["params"].([]interface{})
		result := testCase["result"]
		c.Assert(method, NotNil)
		c.Assert(params, NotNil)
		c.Assert(result, NotNil)

		wasmData := getWasmData(c)
		rets, paramData := getParams(c, params)
		fmt.Println("params: ", strings.Join(rets, ","))
		usedCost, newWasmData, retData, err := RunWasm(cost, wasmCode, abiData, wasmData, method, paramData)
		c.Assert(err, IsNil)
		retValue, retString := getResult(c, retData)
		c.Assert(retValue, Equals, result)
		fmt.Println(cost, usedCost, cost-usedCost)
		fmt.Println("return: ", len(wasmData), len(newWasmData), retString)
		if len(newWasmData) > 0 {
			meta, msg, err := core.Unmarshal(newWasmData)
			c.Assert(err, IsNil)
			c.Assert(meta, Equals, core.CORE_DATA_MAP)
			m := msg.(*pb.DataMap)
			c.Assert(m, NotNil)
		}
	}
}
