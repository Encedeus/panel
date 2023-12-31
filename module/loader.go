package module

import (
    "archive/zip"
    "context"
    "errors"
    "fmt"
    "github.com/filecoin-project/go-jsonrpc"
    "github.com/hashicorp/hcl/v2/hclsimple"
    "github.com/labstack/gommon/log"
    "github.com/second-state/WasmEdge-go/wasmedge"
    "io"
    "io/fs"
    "math"
    "math/rand"
    "net"
    "net/http"
    "os"
    "path/filepath"
    "slices"
    "strconv"
    "strings"
    "sync"
    "time"
)

const ManifestFileName = "manifest.hcl"

// const ModulesFolderLocation = "/home/optimuseprime/Projects/Encedeus/test/panel/modules"

// const ModulesFolderLocation = "/etc/encedeus/modules"

var (
    ErrInvalidManifest = errors.New("invalid manifest file")
)

type SemVerVersion struct {
    MajorVersion uint32
    MinorVersion uint32
    PatchVersion uint32
    Suffix       string
}

func SemVerFromString(s string) SemVerVersion {
    semVer := SemVerVersion{}

    v := strings.Split(s, ".")
    major, _ := strconv.Atoi(v[0])
    minor, _ := strconv.Atoi(v[1])

    v1 := strings.Split(v[2], "-")
    patch, _ := strconv.Atoi(v1[0])
    suffix := v1[1]

    semVer.MajorVersion = uint32(major)
    semVer.MinorVersion = uint32(minor)
    semVer.PatchVersion = uint32(patch)
    semVer.Suffix = suffix

    return semVer
}

func (v SemVerVersion) String() string {
    return fmt.Sprintf("%v.%v.%v-%s", v.MajorVersion, v.MinorVersion, v.PatchVersion, v.Suffix)
}

type Store struct {
    mu                sync.Mutex
    Modules           []*Module
    ModulesFolderPath string
    RPCPort           Port
}

func NewStore(modulesPath string) *Store {
    store := new(Store)
    store.ModulesFolderPath = modulesPath
    store.Modules = make([]*Module, 0, 5)
    store.RPCPort = store.GetAvailablePort()

    return store
}

func (ms *Store) GetAvailablePort() Port {
    maxPort := math.MaxUint16
    minPort := 1024

    isAvailable := func(port Port) func(m *Module) bool {
        return func(m *Module) bool {
            if m != nil {
                if m.RPCPort == port {
                    return true
                }
            }

            timeout := time.Second
            conn, err := net.DialTimeout("tcp", net.JoinHostPort("localhost", strconv.Itoa(int(port))), timeout)
            defer conn.Close()
            if err == nil && conn != nil {
                return true
            }

            return false
        }
    }

    var port = Port(rand.Intn(maxPort-minPort) + minPort)
    for slices.ContainsFunc(ms.Modules, isAvailable(port)) {
        port = Port(rand.Intn(maxPort-minPort) + minPort)
    }
    fmt.Printf("Available port: %v\n", port)

    return port
}

type Configuration struct {
    Port     Port
    HostPort Port
    Manifest Manifest
}

type HandshakeResponse struct {
    RegisteredCommands []*Command
}

func (ms *Store) LoadOne(emaPath string) (*Module, error) {
    zipReader, err := zip.OpenReader(filepath.Join(ms.ModulesFolderPath, emaPath))
    if err != nil {
        log.Errorf("%e", err)
        return nil, err
    }
    defer zipReader.Close()

    manifest, err := NewManifestFromEma(zipReader)
    if err != nil {
        log.Errorf("%e", err)
        return nil, err
    }

    if len(manifest.BackendMainFile) == 0 && len(manifest.Frontend.TabName) == 0 {
        log.Errorf("%e", ErrInvalidManifest)
        return nil, ErrInvalidManifest
    }

    unzipLocation := filepath.Join(ms.ModulesFolderPath, fmt.Sprintf("/cache/%s", manifest.Name))

    _, err = os.Open(unzipLocation)
    if errors.Is(err, fs.ErrNotExist) {
        err = Unzip(zipReader, unzipLocation)
        if err != nil {
            log.Errorf("%e", err)
            return nil, err
        }
    }

    // Backend loading
    backendPort := ms.GetAvailablePort()
    rpcPort := ms.GetAvailablePort()
    if len(manifest.BackendMainFile) != 0 {
        backendMain, err := os.Open(filepath.Join(unzipLocation, "backend", manifest.BackendMainFile))
        if err != nil {
            log.Errorf("%e", err)
            return nil, err
        }

        err = ExecuteModuleWasmFile(ms.ModulesFolderPath, manifest.Name, backendMain, rpcPort, backendPort)
        if err != nil {
            log.Errorf("%e", err)
            return nil, err
        }
    }

    var frontendServer *FrontendServer
    // Frontend loading
    if len(manifest.Frontend.TabName) != 0 {
        port := ms.GetAvailablePort()
        frontendServer = NewFrontendServer(
            Platform(manifest.Frontend.Platform),
            filepath.Join(unzipLocation, "frontend"),
            port)

        err = frontendServer.Start()
        if err != nil {
            log.Errorf("%e", err)
            return nil, err
        }
    }

    fmt.Printf("Frontend port: %v\n", frontendServer.Port)

    module := new(Module)
    module.Manifest = *manifest
    module.BackendPort = backendPort
    module.RPCPort = rpcPort
    module.Store = ms
    module.FrontendServer = frontendServer

    err = module.beginHandshake()
    if err != nil {
        log.Errorf("%e", err)
        return nil, err
    }
    // fmt.Println("Hadnshake done")

    ms.mu.Lock()
    defer ms.mu.Unlock()

    ms.Modules = append(ms.Modules, module)
    log.Infof("Modules: %+v\n", ms.Modules)

    return module, nil
}

func (ms *Store) InitRPCServer() {
    rpcSrv := jsonrpc.NewServer()

    invHandle := new(ModuleInvokeHandler)
    invHandle.ModuleStore = ms
    rpcSrv.Register("ModuleInvokeHandler", invHandle)

    srv := http.Server{
        Addr:         fmt.Sprintf(":%v", ms.RPCPort),
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 5 * time.Second,
        Handler:      rpcSrv,
    }

    fmt.Println("RPC server started")
    log.Fatal(srv.ListenAndServe())
}

func (ms *Store) HasRegisteredCommand(command string) (bool, *Module, string) {
    for _, mod := range ms.Modules {
        for _, cmd := range mod.Manifest.RegisteredCommands {
            if cmd == command {
                return true, mod, cmd
            }
        }
    }

    return false, nil, ""
}

func (ms *Store) LoadAll() error {
    modulesFolder, err := os.ReadDir(ms.ModulesFolderPath)
    if err != nil {
        return err
    }

    // ms.InitRPCServer()

    for _, ema := range modulesFolder {
        if filepath.Ext(ema.Name()) != ".ema" {
            continue
        }

        go func(name string) {
            _, err := ms.LoadOne(name)
            if err != nil {
                log.Errorf("Failed loading module %s: %e", name, err)
            }
        }(ema.Name())
    }

    return nil
}

type Port uint16
type Module struct {
    Store          *Store
    Manifest       Manifest
    BackendPort    Port
    RPCPort        Port
    FrontendServer *FrontendServer
    // RegisteredCommands []*Command
}

func (m *Module) beginHandshake() error {
    var client struct {
        OnHandshake func(config Configuration) HandshakeResponse
    }
    // fmt.Println("Handshake")
    // time.Sleep(2 * time.Second)/**/

    fmt.Printf("Actual RPC port: %v\n", m.RPCPort)
    closer, err := jsonrpc.NewClient(context.Background(), fmt.Sprintf("http://localhost:%v", m.RPCPort), "HandshakeHandler", &client, nil)
    if err != nil {
        return err
    }
    defer closer()

    _ = client.OnHandshake(Configuration{
        Manifest: m.Manifest,
        Port:     m.BackendPort,
        HostPort: m.Store.RPCPort,
    })
    // fmt.Printf("%+v\n", resp)
    // m.RegisteredCommands = resp.RegisteredCommands

    return nil
}

type Manifest struct {
    Name               string   `hcl:"name"`
    Authors            []string `hcl:"authors"`
    Verison            string   `hcl:"version"`
    BackendMainFile    string   `hcl:"backend_main"`
    RegisteredCommands []string `hcl:"commands"`
    Frontend           struct {
        TabName  string `hcl:"tab_name"`
        Platform string `hcl:"platform"`
    } `hcl:"frontend,block"`
}

func (m *Manifest) SemVer() SemVerVersion {
    return SemVerFromString(m.Verison)
}

func NewManifestFromEma(reader *zip.ReadCloser) (*Manifest, error) {
    manifestFile, err := reader.Open(ManifestFileName)
    if err != nil {
        log.Errorf("%e", err)
        return nil, err
    }
    defer manifestFile.Close()

    stat, _ := manifestFile.Stat()
    b := make([]byte, stat.Size())
    _, err = manifestFile.Read(b)
    if err != nil && !errors.Is(err, io.EOF) {
        log.Errorf("%e", err)
        return nil, err
    }

    manifest := new(Manifest)
    err = hclsimple.Decode(ManifestFileName, b, nil, manifest)
    if err != nil {
        log.Errorf("%e", err)
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

func ExecuteWasmFileWithDefaults(dirMapping string, buf []byte, environ []string) error {
    wasmedge.SetLogDebugLevel()

    conf := wasmedge.NewConfigure(wasmedge.REFERENCE_TYPES)
    conf.AddConfig(wasmedge.WASI)
    vm := wasmedge.NewVMWithConfig(conf)
    wasi := vm.GetImportModule(wasmedge.WASI)
    wasi.InitWasi(
        os.Args[1:],
        environ,
        []string{dirMapping},
    )

    _, err := vm.RunWasmBuffer(buf, "_start")
    if err != nil {
        log.Errorf("%e", err)
    }

    vm.Release()
    conf.Release()

    return nil
}

func ExecuteModuleWasmFile(modulesPath, moduleName string, f fs.File, rpcPort Port, backendPort Port) error {
    stat, _ := f.Stat()
    buf := make([]byte, stat.Size())
    _, err := f.Read(buf)
    if err != nil {
        return err
    }

    path := filepath.Join(modulesPath, "cache", moduleName)
    go func() {
        err = ExecuteWasmFileWithDefaults(fmt.Sprintf("/:%s", path), buf,
            []string{
                fmt.Sprintf("MODULE_RPC_PORT=%v", strconv.Itoa(int(rpcPort))),
                fmt.Sprintf("MODULE_MAIN_PORT=%v", strconv.Itoa(int(backendPort))),
            })
        if err != nil {
            log.Errorf("%e", err)
        }
    }()

    return nil
}
