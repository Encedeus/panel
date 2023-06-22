package controller

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
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
		userId       uuid.UUID
	)

	// check which method was used for log in
	if loginInfo.Username != "" {
		passwordHash, userId, err = service.GetUserPasswordHashAndUUIDByUsername(loginInfo.Username)
	} else if loginInfo.Email != "" {
		passwordHash, userId, err = service.GetUserPasswordHashAndUUIDByEmail(loginInfo.Email)
	} else {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "either username or email must be specified",
		})
	}

	// handle errors
	if err != nil {
		if err.Error() == "user not found" {
			return ctx.JSON(http.StatusNotFound, echo.Map{
				"message": err.Error(),
			})
		}

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
	accessToken, refreshToken, err := util.GetTokenPair(userId)

	return ctx.JSON(http.StatusOK, echo.Map{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}

func tokenRefreshHandler(ctx echo.Context) error {
	// error safe because of the RefreshJWTAuth middleware
	_, userId, _ := util.ValidateRefreshJWT(util.GetTokenFromHeader(ctx))

	// generate access token
	accessToken, _ := util.GenerateAccessToken(*userId)

	return ctx.JSON(http.StatusOK, echo.Map{
		"accessToken": accessToken,
	})
}
