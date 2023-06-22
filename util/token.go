package util

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"panel/config"
	"strings"
	"time"
)

// GenerateAccessToken generates an access token containing the uuid of a user that expires in 15 minutes
func GenerateAccessToken(userId uuid.UUID) (string, error) {
	// generate a token containing the user's uuid
	accessToken := jwt.New(jwt.SigningMethodHS256)
	accessClaims := accessToken.Claims.(jwt.MapClaims)
	accessClaims["userId"] = userId
	accessClaims["exp"] = time.Now().Add(15 * time.Minute).Unix() // expires in 15 minutes

	// sign the token
	accessTokenString, err := accessToken.SignedString([]byte(config.Config.Auth.JWTSecretAccess))
	if err != nil {
		return "", err
	}

	return accessTokenString, nil
}

// GenerateRefreshToken generates a refresh token containing the uuid of a user that expires in a week
func GenerateRefreshToken(userId uuid.UUID) (string, error) {
	// generate a token containing the user's uuid
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshClaims["userId"] = userId
	refreshClaims["exp"] = time.Now().Add(168 * time.Hour).Unix() // expires in a  week

	// sign the token
	refreshTokenString, err := refreshToken.SignedString([]byte(config.Config.Auth.JWTSecretRefresh))
	if err != nil {
		return "", err
	}

	return refreshTokenString, nil
}

// GetTokenPair returns an access and a refresh token
func GetTokenPair(userId uuid.UUID) (string, string, error) {
	accessToken, err1 := GenerateAccessToken(userId)
	refreshToken, err2 := GenerateRefreshToken(userId)

	if err1 != nil {
		log.Errorf("error generating access token %v", err1)
		return "", "", err1
	}
	if err2 != nil {
		log.Errorf("error generating refresh token %v", err2)
		return "", "", err2
	}

	return accessToken, refreshToken, nil
}

// GetTokenFromHeader extracts the token from the auth header
func GetTokenFromHeader(ctx echo.Context) string {
	// remove "Bearer " in front of the token
	return strings.Split(ctx.Request().Header.Get("Authorization"), " ")[1]
}

// ValidateJWT returns data about the validity of a JWT as well as the uuid encoded inside the JWT
func ValidateJWT(tokenString string, secretKey string) (bool, *uuid.UUID, error) {

	// parse the JWT and check the signing method
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("unexpected jwt signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return false, nil, err
	}

	// extract the payload (claims) of the JWT
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return false, nil, err
	}

	// get the uuid of the user
	id, err := uuid.Parse(claims["userId"].(string))

	if err != nil {
		return false, nil, err
	}

	return true, &id, nil
}

func ValidateAccessJWT(token string) (bool, *uuid.UUID, error) {
	return ValidateJWT(token, config.Config.Auth.JWTSecretAccess)
}
func ValidateRefreshJWT(token string) (bool, *uuid.UUID, error) {
	return ValidateJWT(token, config.Config.Auth.JWTSecretRefresh)
}
