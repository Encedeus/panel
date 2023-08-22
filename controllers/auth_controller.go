package controllers

import (
    "github.com/Encedeus/panel/dto"
    "github.com/Encedeus/panel/ent"
    "github.com/Encedeus/panel/hashing"
    "github.com/Encedeus/panel/middleware"
    "github.com/Encedeus/panel/services"
    "github.com/Encedeus/panel/util"
    "github.com/labstack/echo/v4"
    "github.com/labstack/gommon/log"
    "net/http"
    "time"
)

func init() {
    addController(func(server *echo.Echo, db *ent.Client) {
        usersEndpoint := server.Group("auth")
        {
            usersEndpoint.POST("/login", userLoginHandler)

            usersEndpoint.Use(middleware.RefreshJWTAuth)

            usersEndpoint.GET("/refresh", tokenRefreshHandler)

            usersEndpoint.DELETE("/logout", logoutHandler)
        }
    })
}

func userLoginHandler(ctx echo.Context) error {
    var loginInfo dto.UserLoginDTO
    // error safe because of the json syntax middleware
    _ = ctx.Bind(&loginInfo)

    var (
        err          error
        passwordHash string
        tokenData    dto.TokenDTO
    )

    // check which method was used for log in
    if loginInfo.Username != "" {
        passwordHash, tokenData, err = services.GetUserAuthDataAndHashByUsername(loginInfo.Username)
    } else if loginInfo.Email != "" {
        passwordHash, tokenData, err = services.GetUserAuthDataAndHashByEmail(loginInfo.Email)
    } else {
        return ctx.JSON(http.StatusBadRequest, echo.Map{
            "message": "either username or email must be specified",
        })
    }

    // handle errors
    if err != nil {
        if ent.IsNotFound(err) {
            return ctx.JSON(http.StatusNotFound, echo.Map{
                "message": "user not found",
            })
        }

        log.Errorf("uncaught error querying user: %v", err)

        return ctx.JSON(http.StatusInternalServerError, echo.Map{
            "message": "internal server error",
        })
    }

    // check if the password hash is a match
    auth := hashing.VerifyHash(loginInfo.Password, passwordHash)

    if !auth {
        return ctx.JSON(http.StatusUnauthorized, echo.Map{
            "message": "unauthorised",
        })
    }

    // generate access and refresh tokens
    accessToken, refreshToken, err := util.GetTokenPair(tokenData)

    // set refresh token cookie
    ctx.SetCookie(&http.Cookie{
        Name:     "encedeus_refreshToken",
        Value:    refreshToken,
        Secure:   true,
        Expires:  time.Now().Add(util.RefreshTokenExpireTime),
        SameSite: http.SameSiteStrictMode,
        HttpOnly: true,
        Path:     "/",
    })

    return ctx.JSON(http.StatusCreated, echo.Map{
        "accessToken": accessToken,
    })
}

func tokenRefreshHandler(ctx echo.Context) error {
    // error safe because of the RefreshJWTAuth middleware
    token, _ := util.GetRefreshTokenFromCookie(ctx)
    _, userData, _ := util.ValidateRefreshJWT(token)

    // generate access token
    accessToken, _ := util.GenerateAccessToken(dto.AccessTokenDTO{UserId: userData.UserId})

    return ctx.JSON(http.StatusOK, echo.Map{
        "accessToken": accessToken,
    })
}

func logoutHandler(ctx echo.Context) error {
    ctx.SetCookie(&http.Cookie{
        Name:     "encedeus_refreshToken",
        HttpOnly: true,
        SameSite: http.SameSiteStrictMode,
        Expires:  time.UnixMilli(0),
        Secure:   true,
        Path:     "/",
    })

    return ctx.NoContent(http.StatusOK)
}
