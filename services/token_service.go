package services

import (
    "github.com/Encedeus/panel/config"
    "github.com/Encedeus/panel/proto"
    protoapi "github.com/Encedeus/panel/proto/go"
    "github.com/golang-jwt/jwt/v5"
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
    jwt.RegisteredClaims
    *protoapi.AccountAPIKeyToken
}

type TokenClaims struct {
    jwt.RegisteredClaims
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

    tokenClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(AccessTokenExpireTime))
    tokenClaims.IssuedAt = jwt.NewNumericDate(time.Now())

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

    keyClaims.IssuedAt = jwt.NewNumericDate(time.Now())
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
    refreshTokenClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(RefreshTokenExpireTime))
    refreshTokenClaims.IssuedAt = jwt.NewNumericDate(time.Now())

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

/*type AccessTokenClaims interface {
    AccountAPIKeyClaims | TokenClaims
}*/

func ValidateAccessJWT(tokenString string) (isValid bool, token *protoapi.Token, err error) {
    // parse the JWT and check the signing method
    tAcc, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(config.Config.Auth.JWTSecretAccess), nil
    })
    tcl, ok := tAcc.Claims.(*TokenClaims)
    if ok && tAcc.Valid {
        if tcl.Token != nil {
            if tcl.Token.Type == protoapi.TokenType_ACCESS_TOKEN {
                token = tcl.Token
            }
        }
    }
    if token != nil {
        return true, token, nil
    }

    tApi, err := jwt.ParseWithClaims(tokenString, &AccountAPIKeyClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(config.Config.Auth.JWTSecretAccess), nil
    })
    acl, ok := tApi.Claims.(*AccountAPIKeyClaims)
    if ok && tApi.Valid {
        if acl.Token != nil {
            if acl.Token.Type == protoapi.TokenType_ACCOUNT_API_KEY {
                token = acl.Token
            }
        }
    }
    if token != nil {
        return true, token, nil
    }

    // isUpdated, err := services.IsUserUpdated(ctx, db, claims.UserID, claims.IssuedAt)
    // if err != nil || isUpdated {
    //     return false, claims, err
    // }

    return false, nil, err
}

func ValidateRefreshJWT(tokenString string) (bool, TokenClaims, error) {
    // parse the JWT and check the signing method

    claims := TokenClaims{}

    _, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
        if token.Method != jwt.SigningMethodHS256 {
            return nil, ErrUnexpectedJWTSigningMethod
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
