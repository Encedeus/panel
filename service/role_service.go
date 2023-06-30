package service

import (
	"context"
	"panel/dto"
	"time"
)

func CreateRole(name string, permissions []string) (int, error) {
	roleData, err := Db.Role.Create().
		SetName(name).
		SetPermissions(permissions).
		Save(context.Background())
	return roleData.ID, err
}

func UpdateRole(roleInfo dto.UpdateRoleDTO) error {
	roleData, err := Db.Role.Get(context.Background(), roleInfo.Id)
	if err != nil {
		return err
	}

	if roleInfo.Name != "" {
		_, err = roleData.Update().SetName(roleInfo.Name).Save(context.Background())
		if err != nil {
			return err
		}
	}

	if len(roleInfo.Permissions) != 0 && roleInfo.Id != 1 {
		_, err = roleData.Update().SetPermissions(roleInfo.Permissions).Save(context.Background())
		if err != nil {
			return err
		}
	}
	return err
}

func DeleteRole(roleId int) error {

	if roleId == 1 {
		return nil
	}

	roleData, err := Db.Role.Get(context.Background(), roleId)
	if err != nil {
		return err
	}

	_, err = roleData.Update().SetDeletedAt(time.Now()).Save(context.Background())
	return err
}
