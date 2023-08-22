package util

import (
    "errors"
    "github.com/Encedeus/panel/config"
    "github.com/Encedeus/panel/dto"
    "github.com/golang-jwt/jwt"
    "github.com/labstack/echo/v4"
    "strings"
    "time"
)

const (
    RefreshTokenExpireTime = 168 * time.Hour
    AccessTokenExpireTime  = 15 * time.Minute
)

type AccountAPIKeyClaims struct {
    jwt.StandardClaims
    dto.AccountAPIKeyDTO
}

type TokenClaims struct {
    jwt.StandardClaims
    dto.TokenDTO
}

// GenerateAccessToken generates an access token containing the uuid of a user that expires in 15 minutes
func GenerateAccessToken(userData dto.AccessTokenDTO) (string, error) {
    tokenClaims := TokenClaims{
        TokenDTO: dto.TokenDTO(userData),
    }

    tokenClaims.ExpiresAt = time.Now().Add(AccessTokenExpireTime).Unix()
    tokenClaims.IssuedAt = time.Now().Unix()

    accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
    accessTokenString, err := accessToken.SignedString([]byte(config.Config.Auth.JWTSecretAccess))
    if err != nil {
        return "", err
    }

    return accessTokenString, nil
}

func GenerateAPIKey(keyData dto.AccountAPIKeyDTO) (string, error) {
    keyClaims := AccountAPIKeyClaims{
        AccountAPIKeyDTO: keyData,
    }

    keyClaims.IssuedAt = time.Now().Unix()

    key := jwt.NewWithClaims(jwt.SigningMethodHS256, keyClaims)
    keyString, err := key.SignedString([]byte(config.Config.Auth.JWTSecretAccess))
    if err != nil {
        return "", err
    }

    return keyString, nil
}

// GenerateRefreshToken generates a refresh token containing the uuid of a user that expires in a week
func GenerateRefreshToken(userData dto.RefreshTokenDTO) (string, error) {
    refreshTokenClaims := TokenClaims{
        TokenDTO: dto.TokenDTO(userData),
    }

    // generate a token containing the user's uuid
    refreshTokenClaims.ExpiresAt = time.Now().Add(RefreshTokenExpireTime).Unix()
    refreshTokenClaims.IssuedAt = time.Now().Unix()

    accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
    accessTokenString, err := accessToken.SignedString([]byte(config.Config.Auth.JWTSecretRefresh))
    if err != nil {
        return "", err
    }

    return accessTokenString, nil
}

// GetTokenPair returns an access and a refresh token
func GetTokenPair(userData dto.TokenDTO) (string, string, error) {
    accessToken, err := GenerateAccessToken(dto.AccessTokenDTO(userData))
    if err != nil {
        // log.Errorf("error generating access token %v", err1)
        return "", "", err
    }

    refreshToken, err := GenerateRefreshToken(dto.RefreshTokenDTO(userData))
    if err != nil {
        // log.Errorf("error generating refresh token %v", err2)
        return "", "", err
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

func ValidateAccessJWT(tokenString string) (bool, TokenClaims, error) {
    // parse the JWT and check the signing method

    claims := TokenClaims{}

    _, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
        if token.Method != jwt.SigningMethodHS256 {
            return nil, errors.New("unexpected jwt signing method")
        }
        return []byte(config.Config.Auth.JWTSecretAccess), nil
    })

    if err != nil {
        return false, claims, err
    }

    // isUpdated, err := services.IsUserUpdated(claims.UserId, claims.IssuedAt)
    // if err != nil || isUpdated {
    //     return false, claims, err
    // }

    return true, claims, err
}
func ValidateRefreshJWT(tokenString string) (bool, TokenClaims, error) {
    // parse the JWT and check the signing method

    claims := TokenClaims{}

    _, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
        if token.Method != jwt.SigningMethodHS256 {
            return nil, errors.New("unexpected jwt signing method")
        }
        return []byte(config.Config.Auth.JWTSecretRefresh), nil
    })

    if err != nil {
        return false, claims, err
    }

    // isUpdated, err := services.IsUserUpdated(claims.UserId, claims.IssuedAt)
    // if err != nil || isUpdated {
    //     return false, claims, err
    // }

    return true, claims, err
}
