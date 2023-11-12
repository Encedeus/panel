package controllers

import (
    "github.com/Encedeus/panel/ent"
    "github.com/Encedeus/panel/middleware"
    protoapi "github.com/Encedeus/panel/proto/go"
    "github.com/Encedeus/panel/security"
    "github.com/Encedeus/panel/services"
    "github.com/labstack/echo/v4"
    "github.com/labstack/gommon/log"
    "net/http"
    "net/mail"
    "strings"
    "time"
)

type AuthController struct {
    Controller
}

func (ac AuthController) registerRoutes(srv *Server) {
    authEndpoint := srv.Group("auth")
    {
        authEndpoint.POST("/signin", func(c echo.Context) error {
            return ac.handleUserSignIn(c, srv.DB)
        })

        authEndpoint.Use(middleware.RefreshJWTAuth)

        authEndpoint.GET("/refresh", func(c echo.Context) error {
            return ac.handleRefreshToken(c, srv.DB)
        })
        authEndpoint.DELETE("/signout", func(c echo.Context) error {
            return ac.handleSignOut(c, srv.DB)
        })
    }
}

func (AuthController) handleUserSignIn(c echo.Context, db *ent.Client) error {
    ctx := c.Request().Context()
    signInReq := new(protoapi.UserSignInRequest)
    // error safe because of the json syntax middleware
    err := c.Bind(signInReq)
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "message": err.Error(),
        })
    }

    var (
        passwordHash string
        tokenData    *protoapi.Token
    )

    // check which method was used for log in
    if _, err := mail.ParseAddress(signInReq.Uid); err != nil {
        passwordHash, tokenData, err = services.GetUserAuthDataAndHashByUsername(ctx, db, signInReq.Uid)
    } else if strings.TrimSpace(signInReq.Uid) != "" {
        passwordHash, tokenData, err = services.GetUserAuthDataAndHashByEmail(ctx, db, signInReq.Uid)
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
    auth := security.VerifyHash(signInReq.Password, passwordHash)

    if !auth {
        return c.JSON(http.StatusUnauthorized, echo.Map{
            "message": "unauthorised",
        })
    }

    // generate access and refresh tokens
    accessToken, refreshToken, err := services.GetTokenPair(tokenData)

    // set refresh token cookie
    c.SetCookie(&http.Cookie{
        Name:     "encedeus_refreshToken",
        Value:    refreshToken,
        Secure:   true,
        Expires:  time.Now().Add(services.RefreshTokenExpireTime),
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
    token, _ := services.GetRefreshTokenFromCookie(c)
    _, userData, _ := services.ValidateRefreshJWT(token)

    // generate access token
    accessToken, _ := services.GenerateAccessToken(&protoapi.AccessToken{
        Token: &protoapi.Token{
            UserId: userData.Token.UserId,
            Type:   protoapi.TokenType_ACCESS_TOKEN,
        },
    })

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
