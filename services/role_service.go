package services

import (
    "context"
    "errors"
    "github.com/Encedeus/panel/dto"
    "github.com/Encedeus/panel/ent"
    "github.com/Encedeus/panel/ent/role"
    "github.com/google/uuid"
    "time"
)

/*func CreateRole(name string, permissions []string) (uuid.UUID, error) {
    roleData, err := db.Role.Create().
        SetName(name).
        SetPermissions(permissions).
        Save(ctx)

    return roleData.ID, err
}*/

func CreateRole(ctx context.Context, db *ent.Client, roleInfo dto.CreateRoleDTO) (uuid.UUID, error) {
    roleData, err := db.Role.Create().
        SetName(roleInfo.Name).
        SetPermissions(roleInfo.Permissions).
        Save(ctx)

    return roleData.ID, err
}

func UpdateRole(ctx context.Context, db *ent.Client, roleInfo dto.UpdateRoleDTO) error {
    roleData, err := db.Role.Get(ctx, roleInfo.ID)
    if err != nil {
        return err
    }

    if IsRoleDeleted(roleData) {
        return errors.New("role deleted")
    }

    if roleInfo.Name != "" {
        _, err = roleData.Update().SetName(roleInfo.Name).Save(ctx)
        if err != nil {
            return err
        }
    }

    if s := roleInfo.ID.String(); len(roleInfo.Permissions) != 0 && s != "" {
        _, err = roleData.Update().SetPermissions(roleInfo.Permissions).Save(ctx)
        if err != nil {
            return err
        }
    }
    return err
}

func DeleteRole(ctx context.Context, db *ent.Client, roleId uuid.UUID) error {

    if s := roleId.String(); s == "" {
        return nil
    }

    roleData, err := db.Role.Get(ctx, roleId)
    if err != nil {
        return err
    }

    if IsRoleDeleted(roleData) {
        return errors.New("already deleted")
    }

    _, err = roleData.Update().SetDeletedAt(time.Now()).Save(ctx)
    return err
}

func FindRole(ctx context.Context, db *ent.Client, roleId uuid.UUID) (*ent.Role, error) {
    roleData, err := db.Role.Query().
        Where(role.IDEQ(roleId)).
        Select("name", "created_at", "updated_at", "deleted_at", "permissions").
        First(ctx)
    if err != nil {
        return nil, err
    }

    if IsRoleDeleted(roleData) {
        return nil, errors.New("role deleted")
    }

    return roleData, err
}

func IsRoleDeleted(role *ent.Role) bool {
    return role.DeletedAt.Unix() != -62135596800
}
