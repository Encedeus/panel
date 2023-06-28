package controller

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
	"panel/config"
	"panel/dto"
	"panel/ent"
	"panel/middleware"
	"panel/service"
	"panel/util"
)

func init() {
	addController(func(server *echo.Echo, db *ent.Client) {
		userEndpoint := server.Group("/user")

		userEndpoint.Static("/pfp", "./pfp")

		userEndpoint.Use(middleware.AccessJWTAuth)

		userEndpoint.POST("/create", handleCreateUser)
		userEndpoint.POST("/setpfp", SetPfp)
	})
}

func handleCreateUser(ctx echo.Context) error {
	userID, _ := uuid.Parse(ctx.Request().Header.Get("UUID"))

	if service.DoesUserHavePermission("create_user", userID) {
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

func SetPfp(ctx echo.Context) error {

	file, err := ctx.FormFile("file")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": "bad request"})
	}

	src, err := file.Open()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"message": "internal server error"})
	}
	defer src.Close()

	dst, err := os.Create(fmt.Sprintf("%s/%s", config.Config.CDN.Directory, ctx.Request().Header.Get("UUID")))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"message": "internal server error"})
	}

	if _, err = io.Copy(dst, src); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"message": "internal server error"})
	}

	return ctx.JSON(http.StatusOK, echo.Map{"message": "ok"})
}
