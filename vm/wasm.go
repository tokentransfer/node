package vm

import (
	"context"
	"fmt"
	"strconv"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tokentransfer/node/util"
)

func RunWasm(cost int64, wasmCode []byte, wasmData []byte, method string, params [][]byte) (int64, []byte, []byte, error) {
	apiCost := api.NewCost()
	ctx := context.WithValue(context.Background(), "cost", apiCost)
	config := wazero.NewRuntimeConfigInterpreter().WithCost(cost)
	runtime := wazero.NewRuntimeWithConfig(ctx, config)
	defer runtime.Close(ctx)
	// wasi_snapshot_preview1.MustInstantiate(ctx, r)
	mc := wazero.NewModuleConfig()
	mod, err := runtime.InstantiateWithConfig(ctx, wasmCode, mc)
	if err != nil {
		return -1, nil, nil, err
	}
	f := mod.ExportedFunction(method)
	if f == nil {
		return -1, nil, nil, util.ErrorOfNonexists("method in wasm module", method)
	}
	types := f.Definition().ParamTypes()
	tLen := len(types)
	pLen := len(params)
	if tLen != pLen {
		return -1, nil, nil, util.ErrorOfInvalid("parameter", fmt.Sprintf("%d != %d", tLen, pLen))
	}
	list := make([]uint64, 0)
	for i := 0; i < len(types); i++ {
		t := types[i]
		p := params[i]
		s := string(p)
		switch t {
		case api.ValueTypeI32:
			i32, err := strconv.ParseInt(s, 10, 32)
			if err != nil {
				return -1, nil, nil, err
			}
			list = append(list, api.EncodeI32(int32(i32)))
		case api.ValueTypeI64:
			i64, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				return -1, nil, nil, err
			}
			list = append(list, api.EncodeI64(i64))
		case api.ValueTypeF32:
			f32, err := strconv.ParseFloat(s, 32)
			if err != nil {
				return -1, nil, nil, err
			}
			list = append(list, api.EncodeF32(float32(f32)))
		case api.ValueTypeF64:
			f64, err := strconv.ParseFloat(s, 64)
			if err != nil {
				return -1, nil, nil, err
			}
			list = append(list, api.EncodeF64(f64))
		}
	}

	results, err := f.Call(ctx, list...)
	if err != nil {
		return -1, nil, nil, err
	}

	types = f.Definition().ResultTypes()
	tLen = len(types)
	if tLen > 1 {
		return -1, nil, nil, util.ErrorOfInvalid("return", "> 1")
	}
	rLen := len(results)
	if tLen != rLen {
		return -1, nil, nil, util.ErrorOfInvalid("return", fmt.Sprintf("%d != %d", tLen, rLen))
	}

	remainCost := apiCost.GetCost()
	result := []byte(nil)
	if tLen == 1 {
		t := types[0]
		r := results[0]
		switch t {
		case api.ValueTypeI32:
			i32 := api.DecodeI32(r)
			result = []byte(fmt.Sprintf("%d", i32))
		case api.ValueTypeI64:
			result = []byte(fmt.Sprintf("%d", r))
		case api.ValueTypeF32:
			f32 := api.DecodeF32(r)
			result = []byte(fmt.Sprintf("%f", f32))
		case api.ValueTypeF64:
			f64 := api.DecodeF64(r)
			result = []byte(fmt.Sprintf("%f", f64))
		}
	}

	return (cost - remainCost), nil, result, nil
}
