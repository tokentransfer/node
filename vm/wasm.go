package vm

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tokentransfer/node/core"
	"github.com/tokentransfer/node/util"
	"github.com/tokentransfer/node/vm/env"
	"github.com/tokentransfer/node/vm/wasi"
)

const (
	TYPE_FUNCTION = "function"
)

type WasmModule struct {
	cost          *api.Cost
	ctx           context.Context
	runtimeConfig wazero.RuntimeConfig
	moduleConfig  wazero.ModuleConfig
	runtime       wazero.Runtime

	mallocFn api.Function
	freeFn   api.Function
	moveFn   api.Function
	funcDef  map[string]interface{}
}

func (wm *WasmModule) Close() error {
	return wm.runtime.Close(wm.ctx)
}

func (wm *WasmModule) Load(cost *api.Cost, wasmCode []byte) (api.Module, error) {
	ctx := context.WithValue(context.Background(), "cost", cost)
	runtimeConfig := wazero.NewRuntimeConfigInterpreter()
	runtime := wazero.NewRuntimeWithConfig(ctx, runtimeConfig)

	// wasi_snapshot_preview1.MustInstantiate(ctx, runtime)
	moduleConfig := wazero.NewModuleConfig().WithStartFunctions()

	//load dependency modules
	err := env.LoadModule(ctx, runtime)
	if err != nil {
		return nil, err
	}
	err = wasi.LoadModule(ctx, runtime)
	if err != nil {
		return nil, err
	}

	wm.cost = cost
	wm.ctx = ctx
	wm.runtimeConfig = runtimeConfig
	wm.moduleConfig = moduleConfig
	wm.runtime = runtime

	mod, err := runtime.InstantiateWithConfig(ctx, wasmCode, moduleConfig)
	if err != nil {
		return nil, err
	}
	return mod, nil
}

func (wm *WasmModule) Verify(mod api.Module, abiCode []byte, method string, params [][]byte) (api.Function, error) {
	memoryMap, defineMap, err := getModuleMap(abiCode)
	if err != nil {
		return nil, err
	}

	mallocName := util.ToString(&memoryMap, "malloc")
	mallocFn := mod.ExportedFunction(mallocName)
	if mallocFn == nil {
		return nil, util.ErrorOfNonexists("function", "malloc")
	} else {
		wm.mallocFn = mallocFn
	}
	freeName := util.ToString(&memoryMap, "free")
	freeFn := mod.ExportedFunction(freeName)
	if freeFn == nil {
		return nil, util.ErrorOfNonexists("function", "free")
	} else {
		wm.freeFn = freeFn
	}
	moveName := util.ToString(&memoryMap, "move")
	moveFn := mod.ExportedFunction(moveName)
	if moveFn == nil {
		return nil, util.ErrorOfNonexists("function", "free")
	} else {
		wm.moveFn = moveFn
	}

	m := util.ToMap(&defineMap, method)
	if m == nil {
		return nil, util.ErrorOfNonexists("method in abi", method)
	} else {
		wm.funcDef = m
	}
	f := mod.ExportedFunction(method)
	if f == nil {
		return nil, util.ErrorOfNonexists("method in wasm module", method)
	}
	types := f.Definition().ParamTypes()
	tLen := len(types)
	pLen, _, err := getInterface(method, m)
	if err != nil {
		return nil, err
	}
	if tLen != pLen {
		return nil, util.ErrorOfInvalid("parameter", fmt.Sprintf("%d != %d", tLen, pLen))
	}
	return f, nil
}

func getBigNumberValue(data []byte) (int64, error) {
	meta, value, err := core.UnmarshalData(data)
	if err != nil {
		return 0, err
	}
	var v int64
	var e error
	switch meta {
	case core.CORE_DATA_INT8:
		v = int64(value.(int8))
		e = nil
	case core.CORE_DATA_UINT8:
		v = int64(value.(uint8))
		e = nil
	case core.CORE_DATA_INT16:
		v = int64(value.(int16))
		e = nil
	case core.CORE_DATA_UINT16:
		v = int64(value.(uint16))
		e = nil
	case core.CORE_DATA_INT32:
		v = int64(value.(int32))
		e = nil
	case core.CORE_DATA_UINT32:
		v = int64(value.(uint32))
		e = nil
	case core.CORE_DATA_INT64:
		v = value.(int64)
		e = nil
	case core.CORE_DATA_UINT64:
		v = int64(value.(uint64))
		e = nil
	case core.CORE_DATA_FLOAT32:
		v = int64(api.EncodeF32(value.(float32)))
		e = nil
	case core.CORE_DATA_FLOAT64:
		v = int64(api.EncodeF64(value.(float64)))
		e = nil
	case core.CORE_DATA_STRING:
		v, e = strconv.ParseInt(value.(string), 10, 32)
	default:
		v = 0
		e = util.ErrorOfInvalid("type", "data")
	}
	if e != nil {
		return 0, e
	}
	return v, nil
}

func getF32Value(data []byte) (float32, error) {
	meta, value, err := core.UnmarshalData(data)
	if err != nil {
		return 0, err
	}
	var v float32
	var e error
	switch meta {
	case core.CORE_DATA_INT8:
		v = float32(value.(int8))
		e = nil
	case core.CORE_DATA_UINT8:
		v = float32(value.(uint8))
		e = nil
	case core.CORE_DATA_INT16:
		v = float32(value.(int16))
		e = nil
	case core.CORE_DATA_UINT16:
		v = float32(value.(uint16))
		e = nil
	case core.CORE_DATA_INT32:
		v = float32(value.(int32))
		e = nil
	case core.CORE_DATA_UINT32:
		v = float32(value.(uint32))
		e = nil
	case core.CORE_DATA_INT64:
		v = float32(value.(int64))
		e = nil
	case core.CORE_DATA_UINT64:
		v = float32(value.(uint64))
		e = nil
	case core.CORE_DATA_FLOAT32:
		v = value.(float32)
		e = nil
	case core.CORE_DATA_STRING:
		f32, err := strconv.ParseFloat(value.(string), 32)
		if err != nil {
			return 0, err
		}
		v = float32(f32)
		e = nil
	default:
		v = 0
		e = util.ErrorOfInvalid("type", "data")
	}
	if e != nil {
		return 0, e
	}
	return v, nil
}

func getF64Value(data []byte) (float64, error) {
	meta, value, err := core.UnmarshalData(data)
	if err != nil {
		return 0, err
	}
	var v float64
	var e error
	switch meta {
	case core.CORE_DATA_INT8:
		v = float64(value.(int8))
		e = nil
	case core.CORE_DATA_UINT8:
		v = float64(value.(uint8))
		e = nil
	case core.CORE_DATA_INT16:
		v = float64(value.(int16))
		e = nil
	case core.CORE_DATA_UINT16:
		v = float64(value.(uint16))
		e = nil
	case core.CORE_DATA_INT32:
		v = float64(value.(int32))
		e = nil
	case core.CORE_DATA_UINT32:
		v = float64(value.(uint32))
		e = nil
	case core.CORE_DATA_INT64:
		v = float64(value.(int64))
		e = nil
	case core.CORE_DATA_UINT64:
		v = float64(value.(uint64))
		e = nil
	case core.CORE_DATA_FLOAT32:
		v = float64(value.(float32))
		e = nil
	case core.CORE_DATA_FLOAT64:
		v = value.(float64)
		e = nil
	case core.CORE_DATA_STRING:
		v, e = strconv.ParseFloat(value.(string), 32)
	default:
		v = 0
		e = util.ErrorOfInvalid("type", "data")
	}
	if e != nil {
		return 0, e
	}
	return v, nil
}

func (wm *WasmModule) Run(mod api.Module, f api.Function, wasmData []byte, method string, params [][]byte) ([]byte, []byte, error) {
	if wm.funcDef != nil {
		inputs := util.ToArray(&wm.funcDef, "inputs")
		outputs := util.ToArray(&wm.funcDef, "outputs")
		hasOutput := false
		outputType := core.CORE_DATA_INT64
		outputData := false
		if len(outputs) > 0 {
			output := outputs[0]
			m, ok := output.(map[string]interface{})
			if ok {
				t, err := core.GetDataTypeByName(util.ToString(&m, "type"))
				if err != nil {
					return nil, nil, err
				}
				hasOutput = true
				outputType = t
				outputData = util.ToBoolean(&m, "data")
			}
		}

		list := make([]uint64, 0)
		for index, item := range inputs {
			m, ok := item.(map[string]interface{})
			if ok {
				data := params[index]

				t, err := core.GetDataTypeByName(util.ToString(&m, "type"))
				if err != nil {
					return nil, nil, err
				}
				if t.Primary() {
					ui64, err := core.AsUint64(data)
					if err != nil {
						return nil, nil, err
					}
					list = append(list, ui64)
					fmt.Println(index, t, ui64)
				} else {
					inputData := util.ToBoolean(&m, "data")
					if inputData && t == core.CORE_DATA_STRING {
						s := hex.EncodeToString(wasmData)
						retData, err := core.MarshalData(s)
						if err != nil {
							return nil, nil, err
						}
						data = retData // use wasm data instead
					}
					d, err := core.AsBytes(data)
					if err != nil {
						return nil, nil, err
					}
					dataSize := uint64(len(d))
					dataPtr, err := wm.malloc(mod, dataSize)
					if err != nil {
						return nil, nil, err
					}
					ok := mod.Memory().Write(uint32(dataPtr), d)
					if !ok {
						return nil, nil, util.ErrorOfInvalid("write memory", fmt.Sprintf("%d, %d", dataPtr, dataSize))
					}
					fmt.Println(index, t, d, dataPtr, dataSize)
					list = append(list, dataPtr, dataSize)
					defer wm.free(mod, dataPtr, dataSize)
				}
			}
		}
		if hasOutput && !outputType.Primary() {
			moveData, err := wm.move(mod, -16)
			if err != nil {
				return nil, nil, err
			}
			list = append([]uint64{moveData}, list...)
			defer wm.move(mod, 16)
		}
		results, err := f.Call(wm.ctx, list...)
		if err != nil {
			return nil, nil, err
		}
		if hasOutput {
			if outputType.Primary() {
				returnData, err := outputType.FromUint64(results[0])
				if err != nil {
					return nil, nil, err
				}
				return nil, returnData, nil
			} else {
				moveData := list[0]
				fmt.Println(moveData, results)
				retPtr, _ := mod.Memory().ReadUint32Le(uint32(moveData))
				retSize, _ := mod.Memory().ReadUint32Le(uint32(moveData + 4))
				retData, ok := mod.Memory().Read(uint32(retPtr), uint32(retSize))
				if !ok {
					return nil, nil, util.ErrorOfInvalid("read memory", fmt.Sprintf("%d, %d", retPtr, retSize))
				}
				defer wm.free(mod, uint64(retPtr), uint64(retSize))

				if outputData {
					newWasmData, err := hex.DecodeString(string(retData))
					if err != nil {
						return nil, nil, err
					}
					return newWasmData, nil, nil // return as wasm data
				} else {
					returnData, err := outputType.FromBytes(retData)
					if err != nil {
						return nil, nil, err
					}
					return nil, returnData, nil
				}
			}
		} else {
			return nil, nil, nil
		}
	} else {
		types := f.Definition().ParamTypes()
		tLen := len(types)
		pLen := len(params)
		if tLen != pLen {
			return nil, nil, util.ErrorOfInvalid("parameter", fmt.Sprintf("%d != %d", tLen, pLen))
		}
		list := make([]uint64, 0)
		for i := 0; i < len(types); i++ {
			t := types[i]
			p := params[i]
			switch t {
			case api.ValueTypeI32:
				i32, err := core.AsUint64(p)
				if err != nil {
					return nil, nil, err
				}
				list = append(list, api.EncodeI32(int32(i32)))
			case api.ValueTypeI64:
				i64, err := core.AsUint64(p)
				if err != nil {
					return nil, nil, err
				}
				list = append(list, api.EncodeI64(int64(i64)))
			case api.ValueTypeF32:
				f32, err := getF32Value(p)
				if err != nil {
					return nil, nil, err
				}
				list = append(list, api.EncodeF32(float32(f32)))
			case api.ValueTypeF64:
				f64, err := getF64Value(p)
				if err != nil {
					return nil, nil, err
				}
				list = append(list, api.EncodeF64(f64))
			default:
				return nil, nil, util.ErrorOfInvalid("parameter type", api.ValueTypeName(t))
			}
		}

		results, err := f.Call(wm.ctx, list...)
		if err != nil {
			return nil, nil, err
		}

		types = f.Definition().ResultTypes()
		tLen = len(types)
		if tLen > 1 {
			return nil, nil, util.ErrorOfInvalid("return", "> 1")
		}
		rLen := len(results)
		if tLen != rLen {
			return nil, nil, util.ErrorOfInvalid("return", fmt.Sprintf("%d != %d", tLen, rLen))
		}

		resultData := []byte(nil)
		if tLen == 1 {
			t := types[0]
			r := results[0]
			switch t {
			case api.ValueTypeI32:
				i32 := api.DecodeI32(r)
				resultData = []byte(fmt.Sprintf("%d", i32))
			case api.ValueTypeI64:
				resultData = []byte(fmt.Sprintf("%d", r))
			case api.ValueTypeF32:
				f32 := api.DecodeF32(r)
				resultData = []byte(fmt.Sprintf("%f", f32))
			case api.ValueTypeF64:
				f64 := api.DecodeF64(r)
				resultData = []byte(fmt.Sprintf("%f", f64))
			default:
				return nil, nil, util.ErrorOfInvalid("result type", api.ValueTypeName(t))
			}
		}
		return nil, resultData, nil
	}
}

func (wm *WasmModule) malloc(m api.Module, size uint64) (uint64, error) {
	if wm.mallocFn != nil {
		results, err := wm.mallocFn.Call(wm.ctx, size)
		if err != nil {
			return 0, err
		}
		return results[0], nil
	}
	return 0, util.ErrorOfNonexists("function", "malloc")
}

func (wm *WasmModule) free(m api.Module, ptr uint64, size uint64) error {
	if wm.freeFn != nil {
		_, err := wm.freeFn.Call(wm.ctx, ptr, size)
		if err != nil {
			return err
		}
		return nil
	}
	return util.ErrorOfNonexists("function", "malloc")
}

func (wm *WasmModule) move(m api.Module, delta int32) (uint64, error) {
	if wm.moveFn != nil {
		results, err := wm.moveFn.Call(wm.ctx, uint64(delta))
		if err != nil {
			return 0, err
		}
		return results[0], nil
	}
	return 0, util.ErrorOfNonexists("function", "move")
}

func verifyType(t core.DataType, pt api.ValueType) bool {
	return t.Primary() && (pt == api.ValueTypeI32 || pt == api.ValueTypeI64 || pt == api.ValueTypeF32 || pt == api.ValueTypeF64)
}

func getInterface(fn string, m map[string]interface{}) (int, []string, error) {
	inputs := util.ToArray(&m, "inputs")
	outputs := util.ToArray(&m, "outputs")

	wasmMap := util.ToMap(&m, "wasm")
	wasmOutputs := util.ToArray(&wasmMap, "outputs")

	paramSize := 0
	paramTypes := make([]string, 0)
	if len(outputs) > 1 || len(wasmOutputs) > 1 {
		return 0, nil, util.ErrorOfInvalid("return count", fmt.Sprintf("%d, > 1", len(outputs)))
	} else {
		if len(outputs) == 1 {
			o := outputs[0]
			m, ok := o.(map[string]interface{})
			if !ok {
				return 0, nil, util.ErrorOfInvalid("return type", fmt.Sprintf("%d", 0))
			} else {
				t, err := core.GetDataTypeByName(util.ToString(&m, "type"))
				if err != nil {
					return 0, nil, err
				}
				if !t.Primary() {
					if len(wasmOutputs) != 0 {
						return 0, nil, util.ErrorOfInvalid("wasm return count", fmt.Sprintf("%d", len(wasmOutputs)))
					}
					paramSize++
					paramTypes = append(paramTypes, "int32")
				} else {
					if len(wasmOutputs) == 1 {
						wo := wasmOutputs[0]
						wm, ok := wo.(map[string]interface{})
						if !ok {
							pn := util.ToString(&wm, "name")
							pt, err := core.GetDataTypeByName(util.ToString(&wm, "type"))
							if err != nil {
								return 0, nil, err
							}
							if t != pt {
								return 0, nil, util.ErrorOfInvalid("return type", fmt.Sprintf("%s != %s in %s(%s)", t, pt, fn, pn))
							}
						}
					} else {
						return 0, nil, util.ErrorOfInvalid("wasm return count", fmt.Sprintf("%d", len(wasmOutputs)))
					}
				}
			}
		}
	}
	for index, item := range inputs {
		m, ok := item.(map[string]interface{})
		if !ok {
			return 0, nil, util.ErrorOfInvalid("parameter", fmt.Sprintf("%d", index))
		} else {
			t, err := core.GetDataTypeByName(util.ToString(&m, "type"))
			if err != nil {
				return 0, nil, err
			}
			if t.Primary() {
				paramSize++
				paramTypes = append(paramTypes, t.String())
			} else {
				paramSize += 2
				paramTypes = append(paramTypes, "int32", "int32")
			}
		}
	}
	return paramSize, paramTypes, nil
}

func verifyInterface(fn string, m map[string]interface{}, fd api.FunctionDefinition) error {
	n := util.ToString(&m, "name")
	if fn != n {
		return util.ErrorOfInvalid("function", fmt.Sprintf("%s != %s", fn, n))
	}
	t := util.ToString(&m, "type")
	if t != TYPE_FUNCTION {
		return util.ErrorOfInvalid("function", fmt.Sprintf("%s != %s in %s", t, TYPE_FUNCTION, n))
	}

	wasmMap := util.ToMap(&m, "wasm")
	wasmInputs := util.ToArray(&wasmMap, "inputs")
	wasmOutputs := util.ToArray(&wasmMap, "outputs")

	wasmParams := fd.ParamTypes()
	wasmResults := fd.ResultTypes()
	if len(wasmInputs) != len(wasmParams) {
		return util.ErrorOfInvalid("wasm parameters", fmt.Sprintf("%d != %d", len(wasmInputs), len(wasmParams)))
	}
	if len(wasmOutputs) != len(wasmResults) {
		return util.ErrorOfInvalid("wasm results", fmt.Sprintf("%d != %d", len(wasmOutputs), len(wasmResults)))
	}

	paramSize, paramTypes, err := getInterface(fn, m)
	if err != nil {
		return err
	}
	if paramSize != len(wasmInputs) {
		return util.ErrorOfInvalid("return type", fmt.Sprintf("%d != %d", paramSize, len(wasmInputs)))
	}
	for index, t := range paramTypes {
		o := wasmInputs[index]
		m, ok := o.(map[string]interface{})
		if !ok {
			return util.ErrorOfInvalid("parameter", fmt.Sprintf("%d", index))
		} else {
			wn := util.ToString(&m, "name")
			wt := util.ToString(&m, "type")
			if t != wt {
				return util.ErrorOfInvalid("return type", fmt.Sprintf("%s != %s in %s(%s)", t, wt, fn, wn))
			}
		}
	}
	for index, pt := range fd.ParamTypes() {
		input := wasmInputs[index]
		inputMap, ok := input.(map[string]interface{})
		if !ok {
			return util.ErrorOfInvalid("parameter type", fmt.Sprintf("%d", index))
		} else {
			n := util.ToString(&inputMap, "name")
			t, err := core.GetDataTypeByName(util.ToString(&inputMap, "type"))
			if err != nil {
				return err
			}
			if !verifyType(t, pt) {
				return util.ErrorOfInvalid("parameter", fmt.Sprintf("%s, %s", n, t))
			}
		}
	}
	return nil
}

func getModuleMap(abiCode []byte) (map[string]interface{}, map[string]interface{}, error) {
	abiMap := make(map[string]interface{})
	err := json.Unmarshal(abiCode, &abiMap)
	if err != nil {
		return nil, nil, err
	}
	memoryMap := util.ToMap(&abiMap, "memory")

	list := util.ToArray(&abiMap, "interfaces")
	interfaceMap := make(map[string]interface{})
	for index, item := range list {
		m, ok := item.(map[string]interface{})
		if !ok {
			return nil, nil, util.ErrorOfInvalid("interface", fmt.Sprintf("index = %d", index))
		}
		name := util.ToString(&m, "name")
		interfaceMap[name] = m
	}
	return memoryMap, interfaceMap, nil
}

func VerifyWasm(wasmCode []byte, abiCode []byte) error {
	apiCost := api.NewCost()
	wm := &WasmModule{}
	mod, err := wm.Load(apiCost, wasmCode)
	if err != nil {
		return err
	}
	defer wm.Close()

	if len(abiCode) > 0 {
		memoryMap, defineMap, err := getModuleMap(abiCode)
		if err != nil {
			return err
		}

		mallocName := util.ToString(&memoryMap, "malloc")
		if mod.ExportedFunction(mallocName) == nil {
			return util.ErrorOfNonexists("function", "malloc")
		}
		freeName := util.ToString(&memoryMap, "free")
		if mod.ExportedFunction(freeName) == nil {
			return util.ErrorOfNonexists("function", "free")
		}
		moveName := util.ToString(&memoryMap, "move")
		if mod.ExportedFunction(moveName) == nil {
			return util.ErrorOfNonexists("function", "move")
		}

		for name, fd := range mod.ExportedFunctionDefinitions() {
			item, ok := defineMap[name]
			if ok {
				m, ok := item.(map[string]interface{})
				if !ok {
					return util.ErrorOfInvalid("function", name)
				}
				err = verifyInterface(name, m, fd)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func RunWasm(cost int64, wasmCode []byte, abiCode []byte, wasmData []byte, method string, params [][]byte) (int64, []byte, []byte, error) {
	apiCost := api.NewCostWith(cost)

	wm := &WasmModule{}
	mod, err := wm.Load(apiCost, wasmCode)
	if err != nil {
		return -1, nil, nil, err
	}
	defer wm.Close()

	f, err := wm.Verify(mod, abiCode, method, params)
	if err != nil {
		return -1, nil, nil, err
	}
	newWasmData, retData, err := wm.Run(mod, f, wasmData, method, params)
	if err != nil {
		return -1, nil, nil, err
	}
	remainCost := apiCost.GetCost()

	return (cost - remainCost), newWasmData, retData, nil
}
