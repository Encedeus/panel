package util

import (
    "errors"
    "github.com/Encedeus/panel/config"
    "github.com/Encedeus/panel/proto"
    protoapi "github.com/Encedeus/panel/proto/go"
    "github.com/golang-jwt/jwt"
    "github.com/labstack/echo/v4"
    "strings"
    "time"
)

const (
    // RefreshTokenExpireTime 1 week
    RefreshTokenExpireTime = 168 * time.Hour
    // AccessTokenExpireTime 15 minutes
    AccessTokenExpireTime = 15 * time.Minute
)

type AccountAPIKeyClaims struct {
    jwt.StandardClaims
    *protoapi.AccountAPIKeyToken
}

type TokenClaims struct {
    jwt.StandardClaims
    *protoapi.Token
}

// GenerateAccessToken generates an access token containing the uuid of a user that expires in 15 minutes
func GenerateAccessToken(userData *protoapi.AccessToken) (string, error) {
    tokenClaims := TokenClaims{
        Token: &protoapi.Token{
            Type:   protoapi.TokenType_ACCESS_TOKEN,
            UserId: userData.Token.UserId,
        },
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

func GenerateAPIKey(keyData *protoapi.AccountAPIKeyToken) (string, error) {
    keyClaims := AccountAPIKeyClaims{
        AccountAPIKeyToken: keyData,
    }

    keyClaims.IssuedAt = time.Now().Unix()
    keyClaims.Token.Type = protoapi.TokenType_ACCOUNT_API_KEY

    key := jwt.NewWithClaims(jwt.SigningMethodHS256, keyClaims)
    keyString, err := key.SignedString([]byte(config.Config.Auth.JWTSecretAccess))
    if err != nil {
        return "", err
    }

    return keyString, nil
}

// GenerateRefreshToken generates a refresh token containing the uuid of a user that expires in a week
func GenerateRefreshToken(keyData *protoapi.RefreshToken) (string, error) {
    refreshTokenClaims := TokenClaims{
        Token: &protoapi.Token{
            Type:   protoapi.TokenType_REFRESH_TOKEN,
            UserId: keyData.Token.UserId,
        },
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
func GetTokenPair(keyData *protoapi.Token) (string, string, error) {
    accessToken, err := GenerateAccessToken(proto.ProtoTokenToAccessToken(keyData))
    if err != nil {
        // log.Errorf("error generating access token %v", err1)
        return "", "", err
    }

    refreshToken, err := GenerateRefreshToken(proto.ProtoTokenToRefreshToken(keyData))
    if err != nil {
        // log.Errorf("error generating refresh token %v", err2)
        return "", "", err
    }

    return accessToken, refreshToken, nil
}

// GetTokenFromHeader extracts the token from the auth header
func GetTokenFromHeader(ctx echo.Context) string {
    // removes "Bearer" in front of the token and returns the token
    return strings.TrimPrefix(ctx.Request().Header.Get("Authorization"), "Bearer ")
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

    if strings.TrimSpace(tokenString) == "" {
        return false, claims, nil
    }

    _, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
        if token.Method != jwt.SigningMethodHS256 {
            return nil, errors.New("unexpected jwt signing method")
        }
        return []byte(config.Config.Auth.JWTSecretAccess), nil
    })

    if claims.Token.Type != protoapi.TokenType_ACCESS_TOKEN {
        return false, TokenClaims{}, ErrInvalidTokenType
    }

    if err != nil {
        return false, claims, err
    }

    // isUpdated, err := services.IsUserUpdated(ctx, db, claims.UserID, claims.IssuedAt)
    // if err != nil || isUpdated {
    //     return false, claims, err
    // }

    return true, claims, nil
}

func ValidateAccountAPIKey(tokenString string) (bool, AccountAPIKeyClaims, error) {
    // parse the JWT and check the signing method
    claims := new(AccountAPIKeyClaims)

    if strings.TrimSpace(tokenString) == "" {
        return false, *claims, nil
    }

    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        if token.Method != jwt.SigningMethodHS256 {
            return nil, errors.New("unexpected jwt signing method")
        }

        if err := token.Claims.Valid(); err != nil {
            return nil, err
        }

        if _, ok := token.Claims.(AccountAPIKeyClaims); !ok {
            return nil, ErrInvalidTokenType
        }

        return []byte(config.Config.Auth.JWTSecretAccess), nil
    })

    if _, ok := token.Claims.(AccountAPIKeyClaims); !ok {
        return false, AccountAPIKeyClaims{}, ErrInvalidTokenType
    }

    if err != nil {
        return false, AccountAPIKeyClaims{}, err
    }

    // isUpdated, err := services.IsUserUpdated(ctx, db, claims.UserID, claims.IssuedAt)
    // if err != nil || isUpdated {
    //     return false, claims, err
    // }

    return true, *claims, nil
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

    // isUpdated, err := services.IsUserUpdated(claims.UserID, claims.IssuedAt)
    // if err != nil || isUpdated {
    //     return false, claims, err
    // }

    return true, claims, nil
}
