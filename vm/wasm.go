package vm

import (
	"context"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tokentransfer/node/core"
)

func CreateWasm(wasmCode []byte, method string, params ...uint64) ([]uint64, error) {
	return nil, nil
}

func RunWasm(cost int64, wasmCode []byte, method string, params ...uint64) (int64, []uint64, error) {
	apiCost := api.NewCost()
	ctx := context.WithValue(context.Background(), "cost", apiCost)
	c := wazero.NewRuntimeConfigInterpreter().WithCost(cost)
	r := wazero.NewRuntimeWithConfig(ctx, c)
	defer r.Close(ctx)
	// wasi_snapshot_preview1.MustInstantiate(ctx, r)
	mc := wazero.NewModuleConfig()
	mod, err := r.InstantiateWithConfig(ctx, wasmCode, mc)
	if err != nil {
		return -1, nil, err
	}
	f := mod.ExportedFunction(method)
	if f == nil {
		return -1, nil, core.ErrorOfNonexists("method in wasm module", method)
	}
	results, err := f.Call(ctx, params...)
	if err != nil {
		return -1, nil, err
	}
	return apiCost.GetCost(), results, nil
}
