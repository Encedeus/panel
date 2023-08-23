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

type AuthController struct {
    Controller
}

func (ac AuthController) registerRoutes(srv *Server) {
    authEndpoint := srv.Group("auth")
    {
        authEndpoint.POST("/login", func(c echo.Context) error {
            return ac.handleUserSignIn(c, srv.DB)
        })

        authEndpoint.Use(middleware.RefreshJWTAuth)

        authEndpoint.GET("/refresh", func(c echo.Context) error {
            return ac.handleRefreshToken(c, srv.DB)
        })
        authEndpoint.DELETE("/logout", func(c echo.Context) error {
            return ac.handleSignOut(c, srv.DB)
        })
    }
}

func (AuthController) handleUserSignIn(c echo.Context, db *ent.Client) error {
    ctx := c.Request().Context()
    var loginInfo dto.UserLoginDTO
    // error safe because of the json syntax middleware
    _ = c.Bind(&loginInfo)

    var (
        err          error
        passwordHash string
        tokenData    dto.TokenDTO
    )

    // check which method was used for log in
    if loginInfo.Username != "" {
        passwordHash, tokenData, err = services.GetUserAuthDataAndHashByUsername(ctx, db, loginInfo.Username)
    } else if loginInfo.Email != "" {
        passwordHash, tokenData, err = services.GetUserAuthDataAndHashByEmail(ctx, db, loginInfo.Email)
    } else {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "message": "either username or email must be specified",
        })
    }

    // handle errors
    if err != nil {
        if ent.IsNotFound(err) {
            return c.JSON(http.StatusNotFound, echo.Map{
                "message": "user not found",
            })
        }

        log.Errorf("uncaught error querying user: %v", err)

        return c.JSON(http.StatusInternalServerError, echo.Map{
            "message": "internal server error",
        })
    }

    // check if the password hash is a match
    auth := hashing.VerifyHash(loginInfo.Password, passwordHash)

    if !auth {
        return c.JSON(http.StatusUnauthorized, echo.Map{
            "message": "unauthorised",
        })
    }

    // generate access and refresh tokens
    accessToken, refreshToken, err := util.GetTokenPair(tokenData)

    // set refresh token cookie
    c.SetCookie(&http.Cookie{
        Name:     "encedeus_refreshToken",
        Value:    refreshToken,
        Secure:   true,
        Expires:  time.Now().Add(util.RefreshTokenExpireTime),
        SameSite: http.SameSiteStrictMode,
        HttpOnly: true,
        Path:     "/",
    })

    return c.JSON(http.StatusCreated, echo.Map{
        "accessToken": accessToken,
    })
}

func (AuthController) handleRefreshToken(c echo.Context, _ *ent.Client) error {
    // error safe because of the RefreshJWTAuth middleware
    token, _ := util.GetRefreshTokenFromCookie(c)
    _, userData, _ := util.ValidateRefreshJWT(token)

    // generate access token
    accessToken, _ := util.GenerateAccessToken(dto.AccessTokenDTO{UserID: userData.UserID})

    return c.JSON(http.StatusOK, echo.Map{
        "accessToken": accessToken,
    })
}

func (AuthController) handleSignOut(c echo.Context, _ *ent.Client) error {
    c.SetCookie(&http.Cookie{
        Name:     "encedeus_refreshToken",
        HttpOnly: true,
        SameSite: http.SameSiteStrictMode,
        Expires:  time.UnixMilli(0),
        Secure:   true,
        Path:     "/",
    })

    return c.NoContent(http.StatusOK)
}
