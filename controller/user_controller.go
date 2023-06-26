package controller

import (
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
		userEndpoint := server.Group("/user")

		userEndpoint.Use(middleware.AccessJWTAuth)

		userEndpoint.POST("/create", handleCreateUser)

	})
}

func handleCreateUser(ctx echo.Context) error {

	if !util.DoesTokenContainPermission("create_user", ctx) {
		return ctx.JSON(http.StatusUnauthorized, echo.Map{
			"message": "unauthorised",
		})
	}

	userInfo := dto.CreateUserDTO{}
	ctx.Bind(&userInfo)

	if userInfo.Name == "" || userInfo.Password == "" || userInfo.Email == "" {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "bad request",
		})
	}

	var err error

	if userInfo.RoleName != "" {
		err = service.CreateUserRoleName(userInfo.Name, userInfo.Email, util.HashPassword(userInfo.Password), userInfo.RoleName)
	} else if userInfo.RoleId != 0 {
		err = service.CreateUserRoleId(userInfo.Name, userInfo.Email, util.HashPassword(userInfo.Password), userInfo.RoleId)
	} else {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "either role name or id must be specified",
		})
	}

	if err != nil {
		if ent.IsNotFound(err) {
			return ctx.JSON(http.StatusNotFound, echo.Map{
				"message": err.Error(),
			})
		}

		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(http.StatusCreated, echo.Map{"message": "ok"})
}
