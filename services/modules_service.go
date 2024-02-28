package services

import (
	"context"
	"fmt"
	"github.com/Encedeus/panel/config"
	"github.com/Encedeus/panel/module"
	"github.com/Encedeus/panel/proto"
	protoapi "github.com/Encedeus/panel/proto/go"
	"io"
	"net/http"
	"os"
	"path/filepath"
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

func InstallModule(fileName string, uri string) error {

	path := fmt.Sprintf("%s/%s.ema",
		filepath.Join(config.Config.StorageLocationPath, config.Config.Modules.ModulesDirectory),
		fileName,
	)

	out, err := os.Create(path)
	defer out.Close()

	if err != nil {
		return err
	}

	resp, err := http.Get(uri)
	defer resp.Body.Close()

	if err != nil {
		return err
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
