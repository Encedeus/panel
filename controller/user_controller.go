package controller

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
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

		userEndpoint.Static("/pfp", config.Config.CDN.Directory)

		userEndpoint.Use(middleware.AccessJWTAuth)

		userEndpoint.POST("/create", handleCreateUser)
		userEndpoint.POST("/setpfp", setPfp)
		userEndpoint.POST("/update", handleUserUpdate)
	})
}

func handleCreateUser(ctx echo.Context) error {
	// get uuid from header provided by the middleware
	userId, _ := uuid.Parse(ctx.Request().Header.Get("UUID"))

	// check permissions
	if !service.DoesUserHavePermission("create_user", userId) {
		return ctx.JSON(http.StatusUnauthorized, echo.Map{
			"message": "unauthorised",
		})
	}

	userInfo := dto.CreateUserDTO{}
	ctx.Bind(&userInfo)

	// check if all the fields are provided
	if userInfo.Name == "" || userInfo.Password == "" || userInfo.Email == "" {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "bad request",
		})
	}

	var err error
	// check which method was used for role assignment
	if userInfo.RoleName != "" {
		err = service.CreateUserRoleName(userInfo.Name, userInfo.Email, util.HashPassword(userInfo.Password), userInfo.RoleName)
	} else if userInfo.RoleId != 0 {
		err = service.CreateUserRoleId(userInfo.Name, userInfo.Email, util.HashPassword(userInfo.Password), userInfo.RoleId)
	} else {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "either role name or id must be specified",
		})
	}

	// error checking
	if err != nil {
		if ent.IsNotFound(err) {
			return ctx.JSON(http.StatusNotFound, echo.Map{
				"message": "role not found",
			})
		}

		if ent.IsConstraintError(err) {
			return ctx.JSON(http.StatusConflict, echo.Map{
				"message": "username taken",
			})
		}

		// log any uncaught errors
		log.Errorf("uncaught error querying role: %v", err)

		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(http.StatusCreated, echo.Map{"message": "ok"})
}

func setPfp(ctx echo.Context) error {
	// get file from multipart
	file, err := ctx.FormFile("file")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": "bad request"})
	}

	// open file
	src, err := file.Open()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"message": "internal server error"})
	}
	defer src.Close()

	// create file
	dst, err := os.Create(fmt.Sprintf("%s/%s", config.Config.CDN.Directory, ctx.Request().Header.Get("UUID")))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"message": "internal server error"})
	}

	// write to file
	if _, err = io.Copy(dst, src); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"message": "internal server error"})
	}

	return ctx.JSON(http.StatusOK, echo.Map{"message": "ok"})
}

func handleUserUpdate(ctx echo.Context) error {
	userId, _ := uuid.Parse(ctx.Request().Header.Get("UUID"))

	// check permissions
	if !service.DoesUserHavePermission("update_user", userId) {
		return ctx.JSON(http.StatusUnauthorized, echo.Map{
			"message": "unauthorised",
		})
	}

	updateInfo := dto.UpdateUserDTO{}
	ctx.Bind(&updateInfo)

	// check if no fields are provided
	if updateInfo.Name == "" && updateInfo.Email == "" && updateInfo.Password == "" && updateInfo.RoleName == "" && updateInfo.RoleId == 0 {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "bad request",
		})
	}

	err := service.UpdateUser(updateInfo, userId)

	// error checking
	if err != nil {

		if ent.IsNotFound(err) {
			return ctx.JSON(http.StatusNotFound, echo.Map{
				"message": "role not found",
			})
		}
		if ent.IsValidationError(err) {
			return ctx.JSON(http.StatusNotFound, echo.Map{
				"message": "validation error",
			})
		}
		if ent.IsConstraintError(err) {
			return ctx.JSON(http.StatusNotFound, echo.Map{
				"message": "constraint error",
			})
		}

		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{"message": "ok"})
}
