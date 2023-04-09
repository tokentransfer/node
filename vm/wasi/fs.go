package wasi

import (
	"context"

	"github.com/tetratelabs/wazero/api"

	"github.com/tokentransfer/node/vm/wasm"
)

var fdWrite = wasm.NewHostFunc(
	"fd_write", wasm.WasmFunc(fdWriteFn),
	[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
	"fd", "iovs", "iovs_len", "result.nwritten",
)

func fdWriteFn(_ context.Context, mod api.Module, params []uint64) uint32 {
	return 0
}
