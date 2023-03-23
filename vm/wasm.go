package vm

import (
	"context"

	"github.com/tetratelabs/wazero"
	"github.com/tokentransfer/node/core"
)

func CreateWasm(wasmCode []byte, method string, params ...uint64) ([]uint64, error) {
	return nil, nil
}

func RunWasm(wasmCode []byte, method string, params ...uint64) ([]uint64, error) {
	ctx := context.Background()
	c := wazero.NewRuntimeConfigInterpreter().WithCost(10)
	r := wazero.NewRuntimeWithConfig(ctx, c)
	defer r.Close(ctx)
	// wasi_snapshot_preview1.MustInstantiate(ctx, r)
	mc := wazero.NewModuleConfig()
	mod, err := r.InstantiateWithConfig(ctx, wasmCode, mc)
	if err != nil {
		return nil, err
	}
	f := mod.ExportedFunction(method)
	if f == nil {
		return nil, core.ErrorOfNonexists("method in wasm module", method)
	}
	results, err := f.Call(ctx, params...)
	if err != nil {
		return nil, err
	}
	return results, nil
}
