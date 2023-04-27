package wbg

import (
	"context"
	"fmt"

	"github.com/tetratelabs/wazero/api"

	"github.com/tokentransfer/node/vm/wasm"
)

var wbgThrow = wasm.NewHostFunc(
	"__wbindgen_throw", wasm.WasmFunc(wbgThrowFn),
	[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32},
	"msg_ptr", "msg_len",
)

func wbgThrowFn(_ context.Context, mod api.Module, params []uint64) uint32 {
	ptr := params[0]
	len := params[1]
	data, ok := mod.Memory().Read(uint32(ptr), uint32(len))
	if !ok {
		panic(fmt.Sprintf("error read memory: %d, %d", ptr, len))
	}
	panic(string(data))
}
