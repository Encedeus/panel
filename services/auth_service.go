package services

import (
    "context"
    "github.com/Encedeus/panel/ent"
    "github.com/Encedeus/panel/ent/user"
    "github.com/Encedeus/panel/proto"
    protoapi "github.com/Encedeus/panel/proto/go"
    "github.com/labstack/gommon/log"
)

// GetUserAuthDataAndHashByUsername returns the user's uuid and hashed password provided the username of the user
func GetUserAuthDataAndHashByUsername(ctx context.Context, db *ent.Client, username string) (string, *protoapi.Token, error) {
    userData, err := db.User.Query().Where(user.Name(username)).Select(user.FieldID, user.FieldPassword).First(ctx)

    if err != nil {
        if !ent.IsNotFound(err) {
            log.Errorf("error querying db on user login (username): %v", err)
        }

        return "", nil, err
    }

    return userData.Password, &protoapi.Token{
        UserId: proto.UUIDToProtoUUID(userData.ID),
    }, nil
}

// GetUserAuthDataAndHashByEmail returns the user's uuid and hashed password provided the email of the user
func GetUserAuthDataAndHashByEmail(ctx context.Context, db *ent.Client, email string) (string, *protoapi.Token, error) {
    userData, err := db.User.Query().Where(user.Email(email)).Select(user.FieldID, user.FieldPassword).First(ctx)

    if err != nil {
        if ent.IsNotFound(err) {
            log.Errorf("error querying db on user login (email): %v", err)
        }

        return "", nil, err
    }

    return userData.Password, &protoapi.Token{
        UserId: proto.UUIDToProtoUUID(userData.ID),
    }, nil
}
