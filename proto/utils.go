package proto

import (
	"github.com/Encedeus/panel/ent"
	"github.com/Encedeus/panel/module"
	protoapi "github.com/Encedeus/panel/proto/go"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"
)

func ProtoUUIDToUUID(id *protoapi.UUID) uuid.UUID {
	return uuid.MustParse(id.Value)
}

func UUIDToProtoUUID(id uuid.UUID) *protoapi.UUID {
	return &protoapi.UUID{
		Value: id.String(),
	}
}

func EntUserEntityToProtoUser(user *ent.User) *protoapi.User {
	return &protoapi.User{
		Id:        UUIDToProtoUUID(user.ID),
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		RoleId:    UUIDToProtoUUID(user.RoleID),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}

func ProtoUserToEntUserEntity(user *protoapi.User) *ent.User {
	return &ent.User{
		ID:        ProtoUUIDToUUID(user.Id),
		CreatedAt: user.CreatedAt.AsTime(),
		UpdatedAt: user.UpdatedAt.AsTime(),
		DeletedAt: user.DeletedAt.AsTime(),
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		RoleID:    ProtoUUIDToUUID(user.RoleId),
	}
}

func EntRoleEntityToProtoRole(role *ent.Role) *protoapi.Role {
	return &protoapi.Role{
		Id:          UUIDToProtoUUID(role.ID),
		Name:        role.Name,
		CreatedAt:   timestamppb.New(role.CreatedAt),
		UpdatedAt:   timestamppb.New(role.UpdatedAt),
		DeletedAt:   timestamppb.New(role.DeletedAt),
		Permissions: role.Permissions,
	}
}

func ProtoRoleToEntRoleEntity(role *protoapi.Role) *ent.Role {
	return &ent.Role{
		ID:          ProtoUUIDToUUID(role.Id),
		Name:        role.Name,
		CreatedAt:   role.CreatedAt.AsTime(),
		UpdatedAt:   role.UpdatedAt.AsTime(),
		DeletedAt:   role.DeletedAt.AsTime(),
		Permissions: role.Permissions,
	}
}

func ProtoTokenToAccessToken(token *protoapi.Token) *protoapi.AccessToken {
	return &protoapi.AccessToken{
		Token: token,
	}
}

func ProtoTokenToRefreshToken(token *protoapi.Token) *protoapi.RefreshToken {
	return &protoapi.RefreshToken{
		Token: token,
	}
}

func ProtoAccountAPIKeyCreateRequestToToken(req *protoapi.AccountAPIKeyCreateRequest) *protoapi.AccountAPIKeyToken {
	return &protoapi.AccountAPIKeyToken{
		Token: &protoapi.Token{
			UserId: req.UserId,
			Type:   protoapi.TokenType_ACCOUNT_API_KEY,
		},
		IpAddresses: req.IpAddresses,
		Description: req.Description,
	}
}

func EntAccountAPIKeyToProtoKey(key *ent.ApiKey) *protoapi.AccountAPIKey {
	return &protoapi.AccountAPIKey{
		Id:          UUIDToProtoUUID(key.ID),
		CreatedAt:   timestamppb.New(key.CreatedAt),
		UpdatedAt:   timestamppb.New(key.UpdatedAt),
		Description: key.Description,
		IpAddresses: key.IPAddresses,
		UserId:      UUIDToProtoUUID(key.UserID),
		Key:         key.Key,
	}
}

func MarshalControllerProtoResponseToJSON(c *echo.Context, okStatus int, message proto.Message) (err error) {
	json, err := protojson.Marshal(message)
	if err != nil {
		return (*c).JSON(http.StatusInternalServerError, echo.Map{
			"message": "internal server error",
		})
	}

	return (*c).JSONBlob(okStatus, json)
}

func ModuleToProtoModule(m module.Module) *protoapi.Module {
	pm := &protoapi.Module{
		Store: &protoapi.ModuleStore{
			ModulesFolderPath: m.Store.ModulesFolderPath,
			RpcPort: &protoapi.Port{
				Value: uint32(m.Store.RPCPort),
			},
		},
		Manifest: &protoapi.ModuleManifest{
			Name:    m.Manifest.Name,
			Authors: m.Manifest.Authors,
			Version: m.Manifest.Version,
			Frontend: &protoapi.ModuleManifestFrontend{
				TabName: m.Manifest.Frontend.TabName,
				// TabIconPath: m.Manifest.Frontend.TabIconPath,
				Platform: &protoapi.ModulePlatform{
					Value: m.Manifest.Frontend.Platform,
				},
			},
			Backend: &protoapi.ModuleManifestBackend{
				Main:               m.Manifest.Backend.MainFile,
				RegisteredCommands: m.Manifest.Backend.RegisteredCommands,
			},
		},
		FrontendServer: &protoapi.ModuleFrontendServer{
			Platform: &protoapi.ModulePlatform{
				Value: string(m.FrontendServer.Platform),
			},
			Environment: m.FrontendServer.Environment,
			EntryPoint:  m.FrontendServer.EntryPoint,
			AssetsPath:  m.FrontendServer.AssetsPath,
			Port: &protoapi.Port{
				Value: uint32(m.FrontendServer.Port),
			},
		},
		BackendPort: &protoapi.Port{
			Value: uint32(m.Backend.BackendPort),
		},
		RpcPort: &protoapi.Port{
			Value: uint32(m.Backend.RPCPort),
		},
	}

	return pm
}
