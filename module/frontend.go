package module

import (
    "embed"
    "fmt"
    vite "github.com/torenware/vite-go"
    "html/template"
    "mime"
    "net/http"
    "os"
    "path/filepath"
)

type Platform string

const (
    PLATFORM_SVELTE  = "svelte"
    PLATFORM_VANILLA = "vanilla"
    PLATFORM_REACT   = "react"
    PLATFORM_VUE     = "vue"
)

//go:embed "frontend-template.gohtml"
var frontendTemplate embed.FS

type FrontendServer struct {
    Platform    Platform
    Environment string
    EntryPoint  string
    AssetsPath  string
    Port        Port
}

func NewFrontendServer(platform Platform, assetsPath string, port Port) *FrontendServer {
    srv := &FrontendServer{
        Platform:    platform,
        AssetsPath:  assetsPath,
        Port:        port,
        EntryPoint:  "src/main.js",
        Environment: "production",
    }

    return srv
}

func (s *FrontendServer) Glue() (*vite.VueGlue, error) {
    config := &vite.ViteConfig{
        Platform:    string(s.Platform),
        Environment: "production",
        EntryPoint:  s.EntryPoint,
        FS:          os.DirFS(s.AssetsPath),
    }

    glue, err := vite.NewVueGlue(config)
    if err != nil {
        return nil, err
    }

    return glue, nil
}

func (s *FrontendServer) Start() error {
    config := &vite.ViteConfig{
        Platform:    string(s.Platform),
        Environment: "production",
        EntryPoint:  s.EntryPoint,
        FS:          os.DirFS(s.AssetsPath),
    }

    glue, err := vite.NewVueGlue(config)
    if err != nil {
        return err
    }

    mux := http.NewServeMux()

    fsHandler, err := glue.FileServer()
    if err != nil {
        return err
    }

    mux.Handle(config.URLPrefix, fsHandler)
    mux.Handle("/", http.HandlerFunc(s.servePage))

    go func() {
        http.ListenAndServe(fmt.Sprintf(":%v", s.Port), mux)
    }()

    return nil
}

func (s *FrontendServer) serveOneFile(w http.ResponseWriter, _ *http.Request, uri, contentType string) {
    strippedURI := uri[:]
    // path := filepath.Join(s.AssetsPath, "dist", strippedURI)
    // fmt.Printf("Path: %v\n", path)
    buf, err := os.ReadFile(filepath.Join(s.AssetsPath, "dist", strippedURI))
    // fmt.Printf("Buf: %v\n", buf)

    if err != nil {
        w.WriteHeader(http.StatusNotFound)
    }

    w.Header().Add("Content-Type", contentType)
    _, err = w.Write(buf)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
    }
}

func (s *FrontendServer) servePage(w http.ResponseWriter, r *http.Request) {
    ext := filepath.Ext(r.RequestURI)
    contentType := mime.TypeByExtension(ext)
    // fmt.Printf("Content type: %v, Ext: %v\n", contentType, ext)
    if contentType != "" {
        s.serveOneFile(w, r, r.RequestURI, contentType)

        return
    }

    t, err := template.ParseFS(frontendTemplate, "frontend-template.gohtml")
    if err != nil || t == nil {
        w.WriteHeader(http.StatusNotFound)
    }

    glue, err := s.Glue()
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
    }

    t.Execute(w, glue)
}
