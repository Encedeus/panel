package module

import (
	"archive/zip"
	"context"
	"errors"
	"fmt"
	"github.com/Encedeus/panel/config"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/google/uuid"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/labstack/gommon/log"
	"github.com/second-state/WasmEdge-go/wasmedge"
	"io"
	"io/fs"
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

/*type ModuleLoadError struct {
	a   *Module
	b   *Module
	err string
}

func (e ModuleLoadError) Error() string {
	return err
}*/

var (
	ErrInvalidManifest          = errors.New("invalid manifest file")
	ErrModuleDependencyNotFound = errors.New("module dependency not found")
	ErrCircularDependency       = errors.New("circular dependency")
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
	Craters           []*Crater
}

func NewStore(modulesPath string) *Store {
	store := new(Store)
	store.ModulesFolderPath = modulesPath
	store.Modules = make([]*Module, 0, 5)
	store.RPCPort = store.GetAvailablePort()

	return store
}

func (ms *Store) FindCraterByName(id string) *Crater {
	for _, m := range ms.Modules {
		for _, c := range m.Manifest.Backend.Craters {
			if c.Id == id || c.Name == id {
				return c
			}
		}
	}
	/*	fmt.Printf("Craters 1: %+v\n", ms.Craters)
		fmt.Printf("Craters 2: %+v\n", ms.Modules[0].Backend.Craters)*/
	/*	for _, c := range ms.Craters {
		if c.Id == id || c.Name == id {
			return c
		}
	}*/

	return nil
}

func (ms *Store) FindCraterVariantByName(id string, c *Crater) *Variant {
	for _, v := range c.Variants {
		if v.Id == id || v.Name == id {
			return v
		}
	}

	return nil
}

func (ms *Store) FindModuleByName(name string) (bool, *Module) {
	for _, m := range ms.Modules {
		if m.Manifest.Name == name {
			return true, m
		}
	}

	return false, nil
}

func (ms *Store) buildDependencyGraph() {
	for _, m := range ms.Modules {
		for _, depName := range m.Manifest.Dependencies {
			isFound, dep := ms.FindModuleByName(depName)
			if !isFound {
				log.Printf("Module %v depends on module %v, but module %v doesn't exist", m.Manifest.Name, depName, depName)
				return
			}
			m.Dependencies = append(m.Dependencies, dep)
		}
	}
}

func (ms *Store) resolveDependencies() []*Module {
	ms.buildDependencyGraph()

	if len(ms.Modules) == 0 {
		return nil
	}

	start := ms.Modules[0]
	resolved := make([]*Module, 0)
	unresolved := make([]*Module, 0)
	err := resolveDependenciesRecurse(start, &resolved, &unresolved)
	if err != nil {
		log.Printf("%v\n", err)
	}

	return resolved
}

func resolveDependenciesRecurse(m *Module, resolved, unresolved *[]*Module) error {
	*unresolved = append(*unresolved, m)
	for _, dep := range m.Dependencies {
		if !slices.Contains(*resolved, dep) {
			if slices.Contains(*unresolved, dep) {
				return ErrCircularDependency
			}
			return resolveDependenciesRecurse(dep, resolved, unresolved)
		}
	}
	*resolved = append(*resolved, m)
	*unresolved = slices.DeleteFunc(*unresolved, func(module *Module) bool {
		return module.Manifest.Name == m.Manifest.Name
	})

	return nil
}

func (ms *Store) GetAvailablePort() Port {
	/*	maxPort := math.MaxUint16
		minPort := 1024

		isAvailable := func(port AddPort) func(m *Module) bool {
			return func(m *Module) bool {
				if m != nil {
					if m.Backend.RPCPort == port {
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

		var port = AddPort(rand.Intn(maxPort-minPort) + minPort)
		for slices.ContainsFunc(ms.Modules, isAvailable(port)) {
			port = AddPort(rand.Intn(maxPort-minPort) + minPort)
		}
		fmt.Printf("Available port: %v\n", port)*/

	server, err := net.Listen("tcp", ":0")

	// If there's an error it likely means no ports
	// are available or something else prevented finding
	// an open port
	if err != nil {
		return 0
	}

	// Defer the closing of the server so it closes
	defer server.Close()

	// Get the host string in the format "127.0.0.1:4444"
	hostString := server.Addr().String()

	// Split the host from the port
	_, portString, err := net.SplitHostPort(hostString)
	if err != nil {
		return 0
	}

	// Return the port as an int
	port, _ := strconv.Atoi(portString)

	return Port(port)
}

type Config struct {
	Port       Port
	HostPort   Port
	Manifest   Manifest
	HostConfig config.Configuration
}

type HandshakeResponse struct {
	//RegisteredCommands []*Command
	//RegisteredCraters []*Crater
}

func (ms *Store) LoadOne(emaPath string, doHandshake bool) (*Module, error) {
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

	if len(manifest.Backend.MainFile) == 0 && len(manifest.Frontend.TabName) == 0 {
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
	fmt.Printf("Backend port: %v\n", backendPort)
	fmt.Printf("RPC port: %v\n", rpcPort)
	if len(manifest.Backend.MainFile) != 0 {
		backendMain, err := os.Open(filepath.Join(unzipLocation, "backend", manifest.Backend.MainFile))
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

	frontendServer := &FrontendServer{}
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

	backend := &Backend{
		BackendPort: backendPort,
		RPCPort:     rpcPort,
	}

	module := new(Module)
	module.Manifest = *manifest
	module.Store = ms
	module.FrontendServer = frontendServer
	module.Backend = backend

	if doHandshake {
		err = module.beginHandshake()
		if err != nil {
			log.Errorf("%e", err)
			return nil, err
		}
	}

	ms.mu.Lock()
	defer ms.mu.Unlock()

	ms.Modules = append(ms.Modules, module)

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

	log.Fatal(srv.ListenAndServe())
}

func (ms *Store) HasRegisteredCommand(command string) (bool, *Module, string) {
	for _, mod := range ms.Modules {
		for _, cmd := range mod.Manifest.Backend.RegisteredCommands {
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

	go ms.InitRPCServer()

	wg := sync.WaitGroup{}
	for _, ema := range modulesFolder {
		if filepath.Ext(ema.Name()) != ".ema" {
			continue
		}
		wg.Add(1)

		go func(name string) {
			_, err := ms.LoadOne(name, false)
			if err != nil {
				log.Errorf("Failed loading module %s: %e", name, err)
			}
			wg.Done()
		}(ema.Name())
	}
	wg.Wait()

	loadOrder := ms.resolveDependencies()
	if len(loadOrder) <= 0 {
		return nil
	}

	wg.Add(len(loadOrder))
	for _, m := range loadOrder {
		go func(module *Module) {
			err = module.beginHandshake()
			if err != nil {
				log.Errorf("%e", err)
			}
			wg.Done()
		}(m)
	}
	wg.Wait()

	return nil
}

type Port uint16
type Module struct {
	ID       uuid.UUID
	Store    *Store
	Manifest Manifest
	Backend  *Backend
	/*	BackendPort    Port
		RPCPort        AddPort*/
	FrontendServer *FrontendServer
	Dependencies   []*Module
	// RegisteredCommands []*Command
}

type Backend struct {
	BackendPort Port
	RPCPort     Port
	Craters     []*Crater
}

func (m *Module) beginHandshake() error {
	var client struct {
		OnHandshake func(config Config) HandshakeResponse
	}

	start := time.Now()
	for {
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("127.0.0.1:%v", m.Backend.RPCPort), 5*time.Second)
		if err == nil {
			_ = conn.Close()
			break
		}
		if time.Now().Sub(start) > 5*time.Second {
			log.Printf("Connection to module %v RPC server at port %v refused: %v", m.ID, m.Backend.RPCPort, err)
			return nil
		}
	}
	closer, err := jsonrpc.NewClient(context.Background(), fmt.Sprintf("http://localhost:%v", m.Backend.RPCPort), "HandshakeHandler", &client, nil)
	if err != nil {
		return err
	}
	defer closer()

	_ = client.OnHandshake(Config{
		Manifest:   m.Manifest,
		Port:       m.Backend.BackendPort,
		HostPort:   m.Store.RPCPort,
		HostConfig: config.Config,
	})
	//fmt.Printf("handshake done\n%+v\n", resp)
	//m.Backend.Craters = resp.RegisteredCraters
	//fmt.Printf("handshake done\n%+v\n", resp)
	// m.RegisteredCommands = resp.RegisteredCommands

	return nil
}

type Manifest struct {
	Name    string   `hcl:"name"`
	Authors []string `hcl:"authors"`
	Version string   `hcl:"version"`
	Backend struct {
		MainFile           string    `hcl:"main"`
		RegisteredCommands []string  `hcl:"commands"`
		Craters            []*Crater `hcl:"crater,block"`
	} `hcl:"backend,block"`
	Frontend struct {
		TabName string `hcl:"tab_name"`
		// TabIconPath string `hcl:"tab_icon"`
		Platform string `hcl:"platform"`
	} `hcl:"frontend,block"`
	Dependencies []string `hcl:"dependencies"`
}

func (m *Manifest) SemVer() SemVerVersion {
	return SemVerFromString(m.Version)
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
	fmt.Printf("%+v\n", manifest)

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
