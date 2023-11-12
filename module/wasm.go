package module

import (
    "archive/zip"
    "errors"
    "fmt"
    "github.com/hashicorp/hcl/v2/hclsimple"
    "github.com/labstack/gommon/log"
    "github.com/second-state/WasmEdge-go/wasmedge"
    "io"
    "io/fs"
    "math"
    "math/rand"
    "net"
    "os"
    "path/filepath"
    "slices"
    "strconv"
    "strings"
    "sync"
    "time"
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

/*func Init() {
    wasmedge.SetLogErrorLevel()

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
}*/

const ManifestFileName = "manifest.hcl"
const ModulesFolderLocation = "/etc/encedeus/modules"

// const ModulesFolderLocation = "/etc/encedeus/modules"

var (
    ErrInvalidManifest = errors.New("invalid manifest file")
)

type SemVerVersion struct {
    MinorVersion uint32
    MajorVersion uint32
    PatchVersion uint32
    Suffix       string
}

func (v SemVerVersion) String() string {
    return fmt.Sprintf("%v.%v.%v-%s", v.MajorVersion, v.MinorVersion, v.PatchVersion, v.Suffix)
}

type Store struct {
    mu                sync.Mutex
    Modules           []*Module
    ModulesFolderPath string
}

func NewStore() *Store {
    store := new(Store)
    store.Modules = make([]*Module, 5)

    return store
}

func (ms *Store) GetAvailablePort() Port {
    maxPort := math.MaxUint16
    minPort := 100

    isNotAvailable := func(port Port) func(m *Module) bool {
        return func(m *Module) bool {
            if m.Port == port {
                return true
            }

            timeout := time.Second
            conn, err := net.DialTimeout("tcp", net.JoinHostPort("localhost", strconv.Itoa(int(port))), timeout)
            if err == nil && conn != nil {
                defer conn.Close()
                return true
            }

            return false
        }
    }

    var port Port
    for slices.ContainsFunc(ms.Modules, isNotAvailable(port)) {
        port = Port(rand.Intn(maxPort-minPort) + minPort)
    }

    return port
}

func (ms *Store) LoadOne(emaPath string) (*Module, error) {
    zipReader, err := zip.OpenReader(emaPath)
    if err != nil {
        return nil, err
    }
    defer zipReader.Close()

    manifest, err := NewManifestFromEma(zipReader)
    if err != nil {
        return nil, err
    }

    if len(manifest.FrontendMainFile) == 0 && len(manifest.BackendMainFile) == 0 {
        return nil, ErrInvalidManifest
    }

    unzipLocation := filepath.Join(ms.ModulesFolderPath, fmt.Sprintf("/cache/%s", manifest.Name))

    _, err = os.Open(unzipLocation)
    if errors.Is(err, fs.ErrNotExist) {
        err = Unzip(zipReader, unzipLocation)
        if err != nil {
            return nil, err
        }
    }

    port := ms.GetAvailablePort()
    if len(manifest.FrontendMainFile) != 0 {
        frontendMain, err := os.Open(filepath.Join(unzipLocation, manifest.FrontendMainFile))
        if err != nil {
            return nil, err
        }

        err = ExecuteModuleWasmFile(frontendMain, port)
        if err != nil {
            return nil, err
        }
    }

    module := new(Module)
    module.Manifest = *manifest
    module.Port = port

    ms.mu.Lock()
    defer ms.mu.Unlock()
    ms.Modules = append(ms.Modules, module)

    return module, nil
}

func (ms *Store) LoadAll() error {
    modulesFolder, err := os.ReadDir(ms.ModulesFolderPath)
    if err != nil {
        return err
    }

    for _, ema := range modulesFolder {
        if filepath.Ext(ema.Name()) != "ema" {
            continue
        }

        go func(name string) {
            _, err := ms.LoadOne(filepath.Join(ms.ModulesFolderPath, name))
            if err != nil {
                log.Errorf("Failed loading module %s: %e", name, err)
            }
        }(ema.Name())
    }

    return nil
}

type Port uint16
type Module struct {
    Manifest Manifest
    Port     Port
}

type Manifest struct {
    Name             string
    Authors          []string
    Verison          SemVerVersion
    FrontendMainFile string
    BackendMainFile  string
}

func NewManifestFromEma(reader *zip.ReadCloser) (*Manifest, error) {
    manifestFile, err := reader.Open(ManifestFileName)
    if err != nil {
        return nil, err
    }
    defer manifestFile.Close()

    stat, _ := manifestFile.Stat()
    b := make([]byte, stat.Size())
    _, err = manifestFile.Read(b)
    if err != nil {
        return nil, err
    }

    manifest := new(Manifest)
    err = hclsimple.Decode(ManifestFileName, b, nil, manifest)
    if err != nil {
        return nil, err
    }

    return manifest, nil
}

func Unzip(r *zip.ReadCloser, dest string) error {
    // r, err := zip.OpenReader(src)
    // if err != nil {
    //     return err
    // }
    // defer func() {
    //     if err := r.Close(); err != nil {
    //         panic(err)
    //     }
    // }()

    os.MkdirAll(dest, 0755)

    // Closure to address file descriptors issue with all the deferred .Close() methods
    extractAndWriteFile := func(f *zip.File) error {
        rc, err := f.Open()
        if err != nil {
            return err
        }
        defer func() {
            if err := rc.Close(); err != nil {
                panic(err)
            }
        }()

        path := filepath.Join(dest, f.Name)

        // Check for ZipSlip (Directory traversal)
        if !strings.HasPrefix(path, filepath.Clean(dest)+string(os.PathSeparator)) {
            return fmt.Errorf("illegal file path: %s", path)
        }

        if f.FileInfo().IsDir() {
            os.MkdirAll(path, f.Mode())
        } else {
            os.MkdirAll(filepath.Dir(path), f.Mode())
            f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
            if err != nil {
                return err
            }
            defer func() {
                if err := f.Close(); err != nil {
                    panic(err)
                }
            }()

            _, err = io.Copy(f, rc)
            if err != nil {
                return err
            }
        }
        return nil
    }

    for _, f := range r.File {
        err := extractAndWriteFile(f)
        if err != nil {
            return err
        }
    }

    return nil
}

func ExecuteWasmFileWithDefaults(buf []byte, args []string) error {
    conf := wasmedge.NewConfigure(wasmedge.REFERENCE_TYPES)
    conf.AddConfig(wasmedge.WASI)
    vm := wasmedge.NewVMWithConfig(conf)
    wasi := vm.GetImportModule(wasmedge.WASI)
    wasi.InitWasi(
        args,
        os.Environ(),
        []string{".:."},
    )

    err := vm.LoadWasmBuffer(buf)
    if err != nil {
        return err
    }

    defer vm.Release()
    defer conf.Release()

    return nil
}

func ExecuteModuleWasmFile(f fs.File, port Port) error {
    stat, _ := f.Stat()
    buf := make([]byte, stat.Size())
    _, err := f.Read(buf)
    if err != nil {
        return err
    }

    err = ExecuteWasmFileWithDefaults(buf, []string{fmt.Sprintf("-p %v", port)})
    if err != nil {
        return err
    }

    return nil
}
