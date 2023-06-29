package service

import (
	"context"
	"github.com/labstack/gommon/log"
	"panel/dto"
	"panel/ent"
	"panel/ent/user"
)

// GetUserAuthDataAndHashByUsername returns the user's uuid and hashed password provided the username of the user
func GetUserAuthDataAndHashByUsername(username string) (string, dto.AccessTokenDTO, error) {
	userData, err := Db.User.Query().Where(user.Name(username)).Select("uuid", "password").First(context.Background())

	if err != nil {
		if !ent.IsNotFound(err) {
			log.Errorf("error querying db on user login (username): %v", err)
		}

		return "", dto.AccessTokenDTO{}, err
	}

	return userData.Password, dto.AccessTokenDTO{
		UserId: userData.UUID,
	}, nil
}

// GetUserAuthDataAndHashByEmail returns the user's uuid and hashed password provided the email of the user
func GetUserAuthDataAndHashByEmail(email string) (string, dto.AccessTokenDTO, error) {
	userData, err := Db.User.Query().Where(user.Email(email)).Select("uuid", "password").First(context.Background())

	if err != nil {
		if ent.IsNotFound(err) {
			log.Errorf("error querying db on user login (email): %v", err)
		}

		return "", dto.AccessTokenDTO{}, err
	}

	return userData.Password, dto.AccessTokenDTO{
		UserId: userData.UUID,
	}, nil
}
