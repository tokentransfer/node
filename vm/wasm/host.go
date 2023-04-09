package wasm

import (
	"context"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
)

type HostFunc struct {
	// Name is equivalent to the same method on api.FunctionDefinition.
	Name string

	// Function
	Func api.GoModuleFunction

	// ParamTypes is equivalent to the same method on api.FunctionDefinition.
	ParamTypes []api.ValueType

	// ParamNames is equivalent to the same method on api.FunctionDefinition.
	ParamNames []string

	// ResultTypes is equivalent to the same method on api.FunctionDefinition.
	ResultTypes []api.ValueType

	// ResultNames is equivalent to the same method on api.FunctionDefinition.
	ResultNames []string
}

func NewHostFunc(
	name string,
	f api.GoModuleFunction,
	paramTypes []api.ValueType,
	paramNames ...string,
) *HostFunc {
	return &HostFunc{
		Name:        name,
		Func:        f,
		ParamTypes:  paramTypes,
		ParamNames:  paramNames,
		ResultTypes: []api.ValueType{api.ValueTypeI32},
		ResultNames: []string{"errno"},
	}
}

func (f *HostFunc) Export(builder wazero.HostModuleBuilder) {
	builder.NewFunctionBuilder().WithGoModuleFunction(f.Func, f.ParamTypes, f.ResultTypes).
		WithParameterNames(f.ParamNames...).
		WithResultNames(f.ResultNames...).
		Export(f.Name)
}

// wasiFunc special cases that all WASI functions return a single Errno
// result. The returned value will be written back to the stack at index zero.
type WasmFunc func(ctx context.Context, mod api.Module, stack []uint64) uint32

// Call implements the same method as documented on api.GoModuleFunction.
func (f WasmFunc) Call(ctx context.Context, mod api.Module, stack []uint64) {
	// Write the result back onto the stack
	stack[0] = uint64(f(ctx, mod, stack))
}
