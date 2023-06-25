package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"panel/dto"
	"panel/ent"
	"panel/ent/user"
)

// GetUserAuthDataAndHashByUsername returns the user's uuid and hashed password provided the username of the user
func GetUserAuthDataAndHashByUsername(username string) (string, dto.AccessTokenDTO, error) {
	userData, err := Db.User.Query().Where(user.Name(username)).First(context.Background())

	if err != nil {
		if !ent.IsNotFound(err) {
			log.Errorf("error querying db on user login (username): %v", err)
		}

		return "", dto.AccessTokenDTO{}, err
	}

	fmt.Println(userData)

	roleData, _ := Db.Role.Get(context.Background(), userData.RoleID)

	return userData.Password, dto.AccessTokenDTO{
		UserId:      userData.UUID,
		Permissions: roleData.Permissions,
	}, nil
}

// GetUserAuthDataAndHashByEmail returns the user's uuid and hashed password provided the email of the user
func GetUserAuthDataAndHashByEmail(email string) (string, dto.AccessTokenDTO, error) {
	userData, err := Db.User.Query().Where(user.Email(email)).Select("uuid", "password", "user_role").First(context.Background())

	if err != nil {
		if ent.IsNotFound(err) {
			log.Errorf("error querying db on user login (email): %v", err)
		}

		return "", dto.AccessTokenDTO{}, err
	}

	roleData, _ := Db.Role.Get(context.Background(), userData.RoleID)

	return userData.Password, dto.AccessTokenDTO{
		UserId:      userData.UUID,
		Permissions: roleData.Permissions,
	}, nil
}

// GetUserAuthDataAndHashByEmail returns the user's uuid and hashed password provided the email of the user
func GetUserAuthDataAndHashByUUID(userId uuid.UUID) (dto.AccessTokenDTO, error) {
	userData, err := Db.User.Query().Where(user.UUID(userId)).Select("role_id").First(context.Background())

	if err != nil {
		if ent.IsNotFound(err) {
			return dto.AccessTokenDTO{}, errors.New("user not found")
		}
		log.Errorf("error querying db on user login (email): %v", err)

		return dto.AccessTokenDTO{}, err
	}

	roleData, _ := Db.Role.Get(context.Background(), userData.RoleID)

	return dto.AccessTokenDTO{
		UserId:      userId,
		Permissions: roleData.Permissions,
	}, nil
}
