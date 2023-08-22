package services

import (
    "context"
    "github.com/Encedeus/panel/dto"
    "github.com/Encedeus/panel/ent"
    "github.com/Encedeus/panel/ent/user"
    "github.com/labstack/gommon/log"
)

// GetUserAuthDataAndHashByUsername returns the user's uuid and hashed password provided the username of the user
func GetUserAuthDataAndHashByUsername(username string) (string, dto.TokenDTO, error) {
    userData, err := Db.User.Query().Where(user.Name(username)).Select("id", "password").First(context.Background())

    if err != nil {
        if !ent.IsNotFound(err) {
            log.Errorf("error querying db on user login (username): %v", err)
        }

        return "", dto.TokenDTO{}, err
    }

    return userData.Password, dto.TokenDTO{
        UserId: userData.ID,
    }, nil
}

// GetUserAuthDataAndHashByEmail returns the user's uuid and hashed password provided the email of the user
func GetUserAuthDataAndHashByEmail(email string) (string, dto.TokenDTO, error) {
    userData, err := Db.User.Query().Where(user.Email(email)).Select("id", "password").First(context.Background())

    if err != nil {
        if ent.IsNotFound(err) {
            log.Errorf("error querying db on user login (email): %v", err)
        }

        return "", dto.TokenDTO{}, err
    }

    return userData.Password, dto.TokenDTO{
        UserId: userData.ID,
    }, nil
}
