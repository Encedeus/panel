package util

import (
    "errors"
    "github.com/Encedeus/panel/config"
    "github.com/Encedeus/panel/dto"
    "github.com/Encedeus/panel/services"
    "github.com/golang-jwt/jwt"
    "github.com/google/uuid"
    "github.com/labstack/echo/v4"
    "github.com/labstack/gommon/log"
    "strings"
    "time"
)

const (
    RefreshTokenExpireTime = 168 * time.Hour
    AccessTokenExpireTime  = 15 * time.Minute
)

// GenerateAccessToken generates an access token containing the uuid of a user that expires in 15 minutes
func GenerateAccessToken(userData dto.AccessTokenDTO) (string, error) {
    userData.ExpiresAt = time.Now().Add(AccessTokenExpireTime).Unix()
    userData.IssuedAt = time.Now().Unix()

    accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, userData)
    accessTokenString, err := accessToken.SignedString([]byte(config.Config.Auth.JWTSecretAccess))
    if err != nil {
        return "", err
    }

    return accessTokenString, nil
}

// GenerateRefreshToken generates a refresh token containing the uuid of a user that expires in a week
func GenerateRefreshToken(userData dto.RefreshTokenDTO) (string, error) {
    // generate a token containing the user's uuid
    userData.ExpiresAt = time.Now().Add(RefreshTokenExpireTime).Unix()
    userData.IssuedAt = time.Now().Unix()

    accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, userData)
    accessTokenString, err := accessToken.SignedString([]byte(config.Config.Auth.JWTSecretRefresh))
    if err != nil {
        return "", err
    }

    return accessTokenString, nil
}

// GetTokenPair returns an access and a refresh token
func GetTokenPair(userData dto.AccessTokenDTO) (string, string, error) {
    accessToken, err1 := GenerateAccessToken(userData)
    refreshToken, err2 := GenerateRefreshToken(dto.RefreshTokenDTO{UserId: userData.UserId})

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
    // removes "Bearer" in front of the token and returns the token
    return strings.Split(ctx.Request().Header.Get("Authorization"), " ")[1]
}

// GetRefreshTokenFromCookie extracts the refresh token from a browser cookie
func GetRefreshTokenFromCookie(ctx echo.Context) (string, error) {
    cookie, err := ctx.Cookie("encedeus_refreshToken")
    if err != nil {
        return "", err
    }

    return cookie.Value, nil
}

func ValidateAccessJWT(tokenString string) (bool, dto.AccessTokenDTO, error) {
    // parse the JWT and check the signing method

    claims := dto.AccessTokenDTO{}

    _, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
        if token.Method != jwt.SigningMethodHS256 {
            return nil, errors.New("unexpected jwt signing method")
        }
        return []byte(config.Config.Auth.JWTSecretAccess), nil
    })

    if err != nil {
        return false, claims, err
    }

    isUpdated, err := IsUserUpdated(claims.UserId, claims.IssuedAt)
    if err != nil || isUpdated {
        return false, claims, err
    }

    return true, claims, err
}
func ValidateRefreshJWT(tokenString string) (bool, dto.RefreshTokenDTO, error) {
    // parse the JWT and check the signing method

    claims := dto.RefreshTokenDTO{}

    _, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
        if token.Method != jwt.SigningMethodHS256 {
            return nil, errors.New("unexpected jwt signing method")
        }
        return []byte(config.Config.Auth.JWTSecretRefresh), nil
    })

    if err != nil {
        return false, claims, err
    }

    isUpdated, err := IsUserUpdated(claims.UserId, claims.IssuedAt)
    if err != nil || isUpdated {
        return false, claims, err
    }

    return true, claims, err
}

func IsUserUpdated(userId uuid.UUID, issuedAt int64) (bool, error) {
    lastUpdate, err := services.GetLastUpdate(userId)
    if err != nil {
        return true, err
    }

    if lastUpdate > issuedAt {
        return true, nil
    }

    return false, nil
}
