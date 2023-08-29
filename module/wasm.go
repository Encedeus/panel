package module

import (
    "os"
    "sync"

    "github.com/second-state/WasmEdge-go/wasmedge"
)

// func InitWasmtime() {
//     dir, err := os.MkdirTemp("", "out")
//     check(err)
//     defer os.RemoveAll(dir)
//     stdoutPath := filepath.Join(dir, "stdout")
//
//     engine := wasmtime.NewEngine()
//
//     module, err := wasmtime.NewModuleFromFile(engine, "/mnt/h/Programming/Web/Workspace/Projects/Encedeus/panel/module_test/go/main.wasm")
//     check(err)
//
//     linker := wasmtime.NewLinker(engine)
//     err = linker.DefineWasi()
//     check(err)
//     wasiConfig := wasmtime.NewWasiConfig()
//     wasiConfig.SetStdoutFile(stdoutPath)
//
//     store := wasmtime.NewStore(engine)
//     store.SetWasi(wasiConfig)
//
//     // exposeFuncs(store, linker)
//
//     err = executeModule(store, module, linker)
//     check(err)
//
//     out, err := os.ReadFile(stdoutPath)
//     check(err)
//     fmt.Print(string(out))
// }
//
// func check(e error) {
//     if e != nil {
//         panic(e)
//     }
// }
//
// func executeModule(store wasmtime.Storelike, module *wasmtime.Module, linker *wasmtime.Linker) (err error) {
//     instance, err := linker.Instantiate(store, module)
//     if err != nil {
//         return err
//     }
//
//     run := instance.GetFunc(store, "_start")
//     if run == nil {
//         return errors.New("module is invalid")
//     }
//
//     _, err = run.Call(store)
//     if err != nil {
//         return err
//     }
//
//     return nil
// }

// func exposeFuncs(store wasmtime.Storelike, linker *wasmtime.Linker) {
// }

func Init() {
    // wasmedge.SetLogErrorLevel()

    // bytes, err := os.ReadFile("/mnt/h/Programming/Web/Workspace/Projects/Encedeus/wasmedge-quickjs/wasmedge_quickjs.wasm")
    // if err != nil {
    //     panic(err)
    // }
    //
    // _, err = vm.RunWasmBuffer(bytes, "_start")
    // if err != nil {
    //     panic(err)
    // }

    var wg sync.WaitGroup
    wg.Add(2)
    go func() {
        defer wg.Done()
        conf := wasmedge.NewConfigure(wasmedge.REFERENCE_TYPES)
        conf.AddConfig(wasmedge.WASI)
        vm := wasmedge.NewVMWithConfig(conf)
        wasi := vm.GetImportModule(wasmedge.WASI)
        wasi.InitWasi(
            os.Args[1:],
            os.Environ(),
            []string{".:."},
        )

        defer vm.Release()
        defer conf.Release()

        vm.RunWasmFile("/mnt/h/Programming/Web/Workspace/Projects/Encedeus/test/js/main.wasm", "_start")
    }()

    go func() {
        defer wg.Done()
        conf := wasmedge.NewConfigure(wasmedge.REFERENCE_TYPES)
        conf.AddConfig(wasmedge.WASI)
        vm := wasmedge.NewVMWithConfig(conf)
        wasi := vm.GetImportModule(wasmedge.WASI)
        wasi.InitWasi(
            os.Args[1:],
            os.Environ(),
            []string{".:."},
        )

        defer vm.Release()
        defer conf.Release()

        vm.RunWasmFile("/mnt/h/Programming/Web/Workspace/Projects/Encedeus/test/go/main.wasm", "_start")
    }()

    wg.Wait()
}
