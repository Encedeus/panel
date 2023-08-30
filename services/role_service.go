package services

import (
    "context"
    "errors"
    "github.com/Encedeus/panel/ent"
    "github.com/Encedeus/panel/ent/role"
    "github.com/Encedeus/panel/proto"
    protoapi "github.com/Encedeus/panel/proto/go"
    "github.com/Encedeus/panel/validate"
    "google.golang.org/protobuf/types/known/timestamppb"
    "strings"
    "time"
)

func CreateRole(ctx context.Context, db *ent.Client, req *protoapi.RoleCreateRequest) (*protoapi.RoleCreateResponse, error) {
    if !validate.IsRoleName(ctx, db, req.Name) {
        return nil, ErrInvalidRoleName
    }
    if !validate.IsPermissionList(req.Permissions) {
        return nil, ErrInvalidPermission
    }

    roleData, err := db.Role.Create().
        SetName(req.Name).
        SetPermissions(req.Permissions).
        Save(ctx)

    resp := &protoapi.RoleCreateResponse{
        Role: &protoapi.Role{
            Id:          proto.UUIDToProtoUUID(roleData.ID),
            CreatedAt:   timestamppb.New(roleData.CreatedAt),
            UpdatedAt:   timestamppb.New(roleData.UpdatedAt),
            DeletedAt:   timestamppb.New(roleData.DeletedAt),
            Name:        roleData.Name,
            Permissions: roleData.Permissions,
        },
    }

    return resp, err
}

func UpdateRole(ctx context.Context, db *ent.Client, req *protoapi.RoleUpdateRequest) (*protoapi.RoleUpdateResponse, error) {
    if !validate.IsRoleName(ctx, db, req.Name) {
        return nil, ErrInvalidRoleName
    }
    if !validate.IsRoleID(ctx, db, req.Id) {
        return nil, ErrInvalidRoleID
    }
    if !validate.IsPermissionList(req.Permissions) {
        return nil, ErrInvalidPermission
    }

    roleData, err := db.Role.Get(ctx, proto.ProtoUUIDToUUID(req.Id))
    if err != nil {
        return nil, err
    }

    if IsRoleDeleted(roleData) {
        return nil, errors.New("role deleted")
    }

    if req.Name != "" {
        _, err = roleData.Update().SetName(req.Name).Save(ctx)
        if err != nil {
            return nil, err
        }
    }

    if s := req.Id.Value; len(req.Permissions) != 0 && s != "" {
        _, err = roleData.Update().SetPermissions(req.Permissions).Save(ctx)
        if err != nil {
            return nil, err
        }
    }

    roleData, err = db.Role.Get(ctx, proto.ProtoUUIDToUUID(req.Id))
    if err != nil {
        return nil, err
    }

    resp := &protoapi.RoleUpdateResponse{
        Role: &protoapi.Role{
            Id:          proto.UUIDToProtoUUID(roleData.ID),
            CreatedAt:   timestamppb.New(roleData.CreatedAt),
            UpdatedAt:   timestamppb.New(roleData.UpdatedAt),
            DeletedAt:   timestamppb.New(roleData.DeletedAt),
            Name:        roleData.Name,
            Permissions: roleData.Permissions,
        },
    }

    return resp, err
}

func DeleteRole(ctx context.Context, db *ent.Client, req *protoapi.RoleDeleteRequest) (*protoapi.RoleDeleteResponse, error) {
    if s := req.Id.Value; strings.TrimSpace(s) == "" {
        return nil, nil
    }

    roleData, err := db.Role.Get(ctx, proto.ProtoUUIDToUUID(req.Id))
    if err != nil {
        return nil, err
    }

    if IsRoleDeleted(roleData) {
        return nil, errors.New("already deleted")
    }

    _, err = roleData.Update().SetDeletedAt(time.Now()).Save(ctx)
    if err != nil {
        return nil, err
    }

    resp := &protoapi.RoleDeleteResponse{}

    return resp, nil
}

func FindRole(ctx context.Context, db *ent.Client, req *protoapi.RoleFindOneRequest) (*protoapi.RoleFindOneResponse, error) {
    roleData, err := db.Role.Query().
        Where(role.IDEQ(proto.ProtoUUIDToUUID(req.Id))).
        First(ctx)
    if err != nil {
        return nil, err
    }

    if IsRoleDeleted(roleData) {
        return nil, errors.New("role deleted")
    }

    resp := &protoapi.RoleFindOneResponse{
        Role: proto.EntRoleEntityToProtoRole(roleData),
    }

    return resp, err
}

func IsRoleDeleted(role *ent.Role) bool {
    return role.DeletedAt.Unix() != -62135596800
}
