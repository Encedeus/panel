package service

import (
	"context"
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
	"panel/dto"
	"panel/ent"
	"panel/ent/role"
	"panel/ent/user"
	"panel/util"
)

func CreateUserRoleId(name string, email string, passwordHash string, roleId int) error {

	roleData, err := Db.Role.Get(context.Background(), roleId)

	if err != nil {
		return err
	}

	return CreateUser(name, email, passwordHash, roleData)
}

func CreateUserRoleName(name string, email string, passwordHash string, roleName string) error {
	roleData, err := Db.Role.Query().Where(role.Name(roleName)).First(context.Background())

	if err != nil {
		return err
	}

	return CreateUser(name, email, passwordHash, roleData)
}

func CreateUser(name string, email string, passwordHash string, role *ent.Role) error {
	_, err := Db.User.Create().
		SetName(name).
		SetEmail(email).
		SetPassword(passwordHash).
		SetRole(role).
		Save(context.Background())
	return err
}

// DoesUserHavePermission checks if user's role have a permission
func DoesUserHavePermission(permission string, userID uuid.UUID) bool {
	first, err := Db.User.Query().Where(user.UUID(userID)).Select("role_id").First(context.Background())
	if err != nil {
		return false
	}

	roleData, err := Db.Role.Query().Where(role.ID(first.RoleID)).Select("permissions").First(context.Background())
	if err != nil {
		return false
	}

	return slices.Contains(roleData.Permissions, permission) || slices.Contains(roleData.Permissions, "*")
}

// UpdateUser updates the user given an updateInfo dto
func UpdateUser(updateInfo dto.UpdateUserDTO, userId uuid.UUID) error {
	userData, err := Db.User.Query().Where(user.UUID(userId)).First(context.Background())
	if err != nil {
		return err
	}

	if updateInfo.Name != "" {
		_, err = userData.Update().SetName(updateInfo.Name).Save(context.Background())
	}

	if updateInfo.Password != "" {
		_, err = userData.Update().SetPassword(util.HashPassword(updateInfo.Password)).Save(context.Background())
	}

	if updateInfo.Email != "" {
		_, err = userData.Update().SetEmail(updateInfo.Email).Save(context.Background())
	}

	if updateInfo.RoleName != "" {
		roleData, roleErr := Db.Role.Query().Where(role.Name(updateInfo.RoleName)).First(context.Background())
		if roleErr != nil {
			return roleErr
		}
		_, err = userData.Update().SetRole(roleData).Save(context.Background())

	} else if updateInfo.RoleId != 0 {
		roleData, roleErr := Db.Role.Query().Where(role.ID(updateInfo.RoleId)).First(context.Background())
		if roleErr != nil {
			return roleErr
		}
		_, err = userData.Update().SetRole(roleData).Save(context.Background())
	}

	return err
}
