package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"panel/dto"
	"panel/ent"
	"panel/middleware"
	"panel/service"
	"panel/util"
)

func init() {
	addController(func(server *echo.Echo, db *ent.Client) {
		usersEndpoint := server.Group("/auth")
		{
			usersEndpoint.POST("/login", userLoginHandler)

			usersEndpoint.Use(middleware.RefreshJWTAuth)

			usersEndpoint.GET("/refresh", tokenRefreshHandler)
		}
	})
}

func userLoginHandler(ctx echo.Context) error {
	var loginInfo dto.UserLoginDTO
	// error safe because of the json syntax middleware
	ctx.Bind(&loginInfo)

	var (
		err          error
		passwordHash string
		tokenData    dto.AccessTokenDTO
	)

	// check which method was used for log in
	if loginInfo.Username != "" {
		passwordHash, tokenData, err = service.GetUserAuthDataAndHashByUsername(loginInfo.Username)
	} else if loginInfo.Email != "" {
		passwordHash, tokenData, err = service.GetUserAuthDataAndHashByEmail(loginInfo.Email)
	} else {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "either username or email must be specified",
		})
	}

	// handle errors
	if err != nil {
		if ent.IsNotFound(err) {
			return ctx.JSON(http.StatusNotFound, echo.Map{
				"message": err.Error(),
			})
		}

		log.Errorf("uncaught error querying user: %v", err)

		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": "internal server error",
		})
	}

	// check if the password hash is a match
	auth := util.VerifyHash(loginInfo.Password, passwordHash)

	if !auth {
		return ctx.JSON(http.StatusUnauthorized, echo.Map{
			"message": "unauthorised",
		})
	}

	// generate access and refresh tokens
	accessToken, refreshToken, err := util.GetTokenPair(tokenData)

	return ctx.JSON(http.StatusCreated, echo.Map{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}

func tokenRefreshHandler(ctx echo.Context) error {
	// error safe because of the RefreshJWTAuth middleware
	_, userData, _ := util.ValidateRefreshJWT(util.GetTokenFromHeader(ctx))

	// generate access token
	accessToken, _ := util.GenerateAccessToken(dto.AccessTokenDTO{UserId: userData.UserId})

	return ctx.JSON(http.StatusOK, echo.Map{
		"accessToken": accessToken,
	})
}
