package service

import (
	"context"
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
	"panel/ent"
	"panel/ent/role"
	"panel/ent/user"
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
