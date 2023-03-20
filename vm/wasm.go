package vm

import (
	"context"
	"fmt"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
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
		return nil, fmt.Errorf("no such method: %s", method)
	}
	results, err := f.Call(ctx, params...)
	if err != nil {
		return nil, err
	}
	return results, nil
}
