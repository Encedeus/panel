package service

import (
	"context"
	"github.com/labstack/gommon/log"
	"panel/ent"
	"panel/ent/role"
)

func CreateUserRoleId(name string, email string, passwordHash string, roleId int) error {

	roleData, err := Db.Role.Get(context.Background(), roleId)

	if err != nil {
		log.Fatalf("error querying role")
		return err
	}

	return CreateUser(name, email, passwordHash, roleData)
}

func CreateUserRoleName(name string, email string, passwordHash string, roleName string) error {
	roleData, err := Db.Role.Query().Where(role.Name(roleName)).First(context.Background())

	if err != nil {
		log.Fatalf("error querying role")
		return err
	}

	return CreateUser(name, email, passwordHash, roleData)
}

func CreateUser(name string, email string, passwordHash string, role *ent.Role) error {
	_, err := Db.User.Create().
		SetName(name).
		SetEmail(email).
		SetPfp([]byte{1, 2, 3}). //TODO: add a profile picture generator (github style)
		SetPassword(passwordHash).
		SetRole(role).
		Save(context.Background())
	return err
}
