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

func CreateRole(name string, permissions []string) (uuid.UUID, error) {
    roleData, err := Db.Role.Create().
        SetName(name).
        SetPermissions(permissions).
        Save(context.Background())

    return roleData.ID, err
}

func UpdateRole(roleInfo dto.UpdateRoleDTO) error {
    roleData, err := Db.Role.Get(context.Background(), roleInfo.ID)
    if err != nil {
        return err
    }

    if IsRoleDeleted(roleData) {
        return errors.New("role deleted")
    }

    if roleInfo.Name != "" {
        _, err = roleData.Update().SetName(roleInfo.Name).Save(context.Background())
        if err != nil {
            return err
        }
    }

    if s := roleInfo.ID.String(); len(roleInfo.Permissions) != 0 && s != "" {
        _, err = roleData.Update().SetPermissions(roleInfo.Permissions).Save(context.Background())
        if err != nil {
            return err
        }
    }
    return err
}

func DeleteRole(roleId uuid.UUID) error {

    if s := roleId.String(); s == "" {
        return nil
    }

    roleData, err := Db.Role.Get(context.Background(), roleId)
    if err != nil {
        return err
    }

    if IsRoleDeleted(roleData) {
        return errors.New("already deleted")
    }

    _, err = roleData.Update().SetDeletedAt(time.Now()).Save(context.Background())
    return err
}

func FindRole(roleId uuid.UUID) (*ent.Role, error) {
    roleData, err := Db.Role.Query().
        Where(role.IDEQ(roleId)).
        Select("name", "created_at", "updated_at", "deleted_at", "permissions").
        First(context.Background())
    if err != nil {
        return nil, err
    }

    if IsRoleDeleted(roleData) {
        return nil, errors.New("role deleted")
    }

    return roleData, err
}

func IsRoleDeleted(roleData *ent.Role) bool {
    return roleData.DeletedAt.Unix() != -62135596800
}
