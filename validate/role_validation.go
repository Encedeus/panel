package validate

import (
    "context"
    "github.com/Encedeus/panel/ent"
    "github.com/Encedeus/panel/ent/role"
    protoapi "github.com/Encedeus/panel/proto/go"
    "github.com/google/uuid"
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
    if strings.TrimSpace(name) == "" {
        return false
    }

    _, err := db.Role.Query().Where(role.NameEQ(name)).First(ctx)

    return err == nil
}
