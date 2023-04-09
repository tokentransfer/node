package wasi

import (
	"context"

	"github.com/tetratelabs/wazero"
)

const ModuleName = "wasi_snapshot_preview1"

func LoadModule(ctx context.Context, runtime wazero.Runtime) error {
	builder := runtime.NewHostModuleBuilder(ModuleName)
	exportFunctions(builder)
	_, err := builder.Instantiate(ctx)
	if err != nil {
		return err
	}
	return nil
}

func exportFunctions(builder wazero.HostModuleBuilder) {
	fdWrite.Export(builder)
}
