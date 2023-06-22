package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"panel/ent"
	"panel/ent/user"
)

// GetUserPasswordHashAndUUIDByUsername returns the user's uuid and hashed password provided the username of the user
func GetUserPasswordHashAndUUIDByUsername(username string) (string, uuid.UUID, error) {
	userData, err := Db.User.Query().Where(user.Name(username)).Select("uuid", "password").First(context.Background())

	if err != nil {
		if ent.IsNotFound(err) {
			return "", uuid.UUID{}, errors.New("user not found")
		}
		log.Errorf("error querying db on user login (username): %v", err)

		return "", uuid.UUID{}, err
	}

	return userData.Password, userData.UUID, nil
}

// GetUserPasswordHashAndUUIDByEmail returns the user's uuid and hashed password provided the email of the user
func GetUserPasswordHashAndUUIDByEmail(email string) (string, uuid.UUID, error) {
	userData, err := Db.User.Query().Where(user.Name(email)).Select("uuid", "password").First(context.Background())

	if err != nil {
		if ent.IsNotFound(err) {
			return "", uuid.UUID{}, errors.New("user not found")
		}
		log.Errorf("error querying db on user login (email): %v", err)
		return "", uuid.UUID{}, err
	}

	return userData.Password, userData.UUID, nil
}
