package wbg

import (
	"context"
	"fmt"

	"github.com/tetratelabs/wazero/api"

	"github.com/tokentransfer/node/vm/wasm"
)

var wbgThrow = wasm.NewFunc(
	"__wbindgen_throw", api.GoModuleFunc(wbgThrowFn),
	[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32},
	[]string{"msg_ptr", "msg_len"},
	[]api.ValueType{},
	[]string{},
)

func wbgThrowFn(_ context.Context, mod api.Module, params []uint64) {
	ptr := params[0]
	len := params[1]
	data, ok := mod.Memory().Read(uint32(ptr), uint32(len))
	if !ok {
		panic(fmt.Sprintf("error read memory: %d, %d", ptr, len))
	}
	panic(string(data))
}
