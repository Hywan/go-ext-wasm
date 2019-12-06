package wasmertest

import (
	"github.com/stretchr/testify/assert"
	wasm "github.com/wasmerio/go-ext-wasm/wasmer"
	"path"
	"runtime"
	"testing"
)

func getBytes(wasmFile ...string) []byte {
	_, filename, _, _ := runtime.Caller(0)
	modulePath := path.Join(path.Dir(filename), "testdata", path.Join(wasmFile...))

	bytes, _ := wasm.ReadBytes(modulePath)

	return bytes
}

func TestWasiC(t *testing.T) {
	module, err := wasm.Compile(getBytes("wasi_c.wasm"))
	assert.NoError(t, err)

	wasiVersion := wasm.WasiGetVersion(module)

	assert.Equal(t, wasiVersion, wasm.Snapshot0)

	importObject := wasm.NewDefaultWasiImportObjectForVersion(wasiVersion)

	instance, err := module.InstantiateWithImportObject(importObject)
	assert.NoError(t, err)

	defer instance.Close()

	// `_starts` calls the `main` function. It prints `Hello, WASI!`.
	start, exists := instance.Exports["_start"]
	assert.Equal(t, true, exists)

	_, err = start()
	assert.NoError(t, err)

	// `sum` is a regular exported function.
	sum, exists := instance.Exports["sum"]
	assert.Equal(t, true, exists)

	output, err := sum(1, 2)
	assert.NoError(t, err)
	assert.Equal(t, wasm.I32(3), output)

}
