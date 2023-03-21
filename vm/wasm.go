package vm

import (
	"context"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
	"github.com/tokentransfer/node/core"
)

func RunWasm(wasmCode []byte, method string, params ...uint64) ([]uint64, error) {
	ctx := context.Background()
	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx)
	wasi_snapshot_preview1.MustInstantiate(ctx, r)
	mod, err := r.Instantiate(ctx, wasmCode)
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
