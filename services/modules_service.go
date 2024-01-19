package services

import (
    "context"
    "fmt"
    "github.com/Encedeus/panel/module"
    "github.com/Encedeus/panel/proto"
    protoapi "github.com/Encedeus/panel/proto/go"
)

func FindAllModules(_ context.Context, store *module.Store, req *protoapi.FindAllModulesRequest) *protoapi.FindAllModulesResponse {
    modules := make([]*protoapi.Module, len(store.Modules))

    for i, m := range store.Modules {
        fmt.Printf("What is nil 1: %+v", modules)
        fmt.Printf("What is nil 3: %v", store.RPCPort)
        if req.FrontendOnly {
            if len(m.Manifest.Frontend.Platform) != 0 {
                modules[i] = proto.ModuleToProtoModule(*m)
                continue
            }
        }
        if req.BackendOnly {
            if len(m.Manifest.Backend.MainFile) != 0 {
                modules[i] = proto.ModuleToProtoModule(*m)
                continue
            }
        }

        fmt.Printf("What is nil 1: %v", modules[i])
        fmt.Printf("What is nil 2: %v", m)
        modules[i] = proto.ModuleToProtoModule(*m)
    }

    return &protoapi.FindAllModulesResponse{
        Modules: modules,
    }
}
