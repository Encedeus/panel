package services

import (
	"context"
	"github.com/Encedeus/panel/module"
	"github.com/Encedeus/panel/proto"
	protoapi "github.com/Encedeus/panel/proto/go"
)

func FindAllModules(_ context.Context, store *module.Store, req *protoapi.ModulesFindAllRequest) *protoapi.ModulesFindAllResponse {
	modules := make([]*protoapi.Module, len(store.Modules))

	for i, m := range store.Modules {
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

		modules[i] = proto.ModuleToProtoModule(*m)
	}

	return &protoapi.ModulesFindAllResponse{
		Modules: modules,
	}
}

func FindOneModule(_ context.Context, store *module.Store, req *protoapi.ModulesFindOneRequest) (*protoapi.ModulesFindOneResponse, error) {
	var mod *protoapi.Module

	for _, m := range store.Modules {
		if m.ID == proto.ProtoUUIDToUUID(req.Id) {
			mod = proto.ModuleToProtoModule(*m)
		}
	}

	if mod == nil {
		return nil, ErrModuleNotFound
	}

	return &protoapi.ModulesFindOneResponse{
		Modules: []*protoapi.Module{
			mod,
		},
	}, nil
}
