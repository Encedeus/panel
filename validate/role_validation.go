package validate

import (
    "context"
    "github.com/Encedeus/panel/ent"
    "github.com/Encedeus/panel/ent/role"
    protoapi "github.com/Encedeus/panel/proto/go"
    "github.com/google/uuid"
    "github.com/microcosm-cc/bluemonday"
    "strings"
)

func IsRoleID(ctx context.Context, db *ent.Client, id *protoapi.UUID) bool {
    if strings.TrimSpace(id.Value) == "" {
        return false
    }

    _, err := db.Role.Get(ctx, uuid.MustParse(id.Value))

    return err == nil
}

func IsRoleName(ctx context.Context, db *ent.Client, name string) bool {
    if strings.TrimSpace(name) == "" || len(name) < 3 {
        return false
    }

    _, err := db.Role.Query().Where(role.NameEQ(name)).First(ctx)

    return err == nil
}

func IsPermission(permission string) bool {
    if len(permission) > 64 || len(permission) < 3 {
        return false
    }

    p := bluemonday.StrictPolicy()
    if s := p.Sanitize(permission); s != permission {
        return false
    }

    return true
}

func IsPermissionList(permissions []string) bool {
    for _, permission := range permissions {
        if !IsPermission(permission) {
            return false
        }
    }

    return true
}
