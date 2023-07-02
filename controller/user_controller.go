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
		userEndpoint.GET("/", getUser)
		userEndpoint.Static("/pfp", config.Config.CDN.Directory)

		userEndpoint.Use(middleware.AccessJWTAuth)

		userEndpoint.POST("/create", handleCreateUser)
		userEndpoint.PUT("/setpfp", setPfp)
		userEndpoint.PATCH("/update", handleUpdateUser)
		userEndpoint.DELETE("/delete", handleDeleteUser)
	})
}

func getUser(ctx echo.Context) error {
	userInfo := dto.GetUserDTO{}
	ctx.Bind(&userInfo)

	if userInfo.UserId.ID() == 0 {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": "bad request"})
	}

	userData, err := service.GetUser(userInfo.UserId)

	if err != nil {
		if ent.IsNotFound(err) {
			return ctx.JSON(http.StatusNotFound, echo.Map{"message": "user not found"})
		}
		if err.Error() == "user deleted" {
			return ctx.JSON(http.StatusGone, echo.Map{"message": "user deleted"})
		}

		log.Errorf("error querying user: %v", err)

		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"name":       userData.Name,
		"email":      userData.Email,
		"created_at": userData.CreatedAt,
		"updated_at": userData.UpdatedAt,
		"role_id":    userData.RoleID,
	})
}

func handleCreateUser(ctx echo.Context) error {
	// get uuid from header provided by the middleware
	authUUID, _ := uuid.Parse(ctx.Request().Header.Get("UUID"))

	// check permissions
	if !service.DoesUserHavePermission("create_user", authUUID) {
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

	var (
		err    error
		userId *uuid.UUID
	)
	// check which method was used for role assignment
	if userInfo.RoleName != "" {
		userId, err = service.CreateUserRoleName(userInfo.Name, userInfo.Email, util.HashPassword(userInfo.Password), userInfo.RoleName)
	} else if userInfo.RoleId != 0 {
		userId, err = service.CreateUserRoleId(userInfo.Name, userInfo.Email, util.HashPassword(userInfo.Password), userInfo.RoleId)
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

	return ctx.JSON(http.StatusCreated, echo.Map{"uuid": userId.String()})
}

func handleUpdateUser(ctx echo.Context) error {
	authUUID, _ := uuid.Parse(ctx.Request().Header.Get("UUID"))

	// check permissions
	if !service.DoesUserHavePermission("update_user", authUUID) {
		return ctx.JSON(http.StatusUnauthorized, echo.Map{
			"message": "unauthorised",
		})
	}

	updateInfo := dto.UpdateUserDTO{}
	ctx.Bind(&updateInfo)

	// check if no fields are provided
	if (updateInfo.Name == "" && updateInfo.Email == "" && updateInfo.Password == "" && updateInfo.RoleName == "" && updateInfo.RoleId == 0) || updateInfo.UserId.ID() == 0 {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "bad request",
		})
	}

	err := service.UpdateUser(updateInfo)

	// error checking
	if err != nil {

		if ent.IsNotFound(err) {
			return ctx.JSON(http.StatusNotFound, echo.Map{
				"message": "role not found",
			})
		}
		if ent.IsValidationError(err) {
			return ctx.JSON(http.StatusBadRequest, echo.Map{
				"message": "validation error",
			})
		}
		if ent.IsConstraintError(err) {
			return ctx.JSON(http.StatusBadRequest, echo.Map{
				"message": "constraint error",
			})
		}
		if err.Error() == "user deleted" {
			return ctx.JSON(http.StatusGone, echo.Map{
				"message": "user deleted",
			})
		}

		log.Errorf("uncaught error updating user: %v", err)

		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{"message": "ok"})
}

func setPfp(ctx echo.Context) error {
	// get params from multipart
	userId := ctx.FormValue("uuid")
	file, err := ctx.FormFile("file")

	// validate request
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": "bad request"})
	}
	userUUID, err := uuid.Parse(userId)
	if !service.DoesUserWithUUIDExist(userUUID) {
		return ctx.JSON(http.StatusNotFound, echo.Map{"message": "user not found"})
	}
	fmt.Println(userUUID, fmt.Sprintf("%s/%s", config.Config.CDN.Directory, userId))

	// open file
	src, err := file.Open()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"message": "internal server error"})
	}
	defer src.Close()

	// create file
	dst, err := os.Create(fmt.Sprintf("%s/%s", config.Config.CDN.Directory, userId))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"message": "internal server error"})
	}

	// write to file
	if _, err = io.Copy(dst, src); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"message": "internal server error"})
	}

	return ctx.JSON(http.StatusOK, echo.Map{"message": "ok"})
}

func handleDeleteUser(ctx echo.Context) error {
	authUUID, _ := uuid.Parse(ctx.Request().Header.Get("UUID"))

	// check permissions
	if !service.DoesUserHavePermission("delete_user", authUUID) {
		return ctx.JSON(http.StatusUnauthorized, echo.Map{
			"message": "unauthorised",
		})
	}

	deleteInfo := dto.DeleteUserDTO{}
	ctx.Bind(&deleteInfo)

	// check if uuid is provided
	if deleteInfo.UserId.ID() == 0 {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "bad request",
		})
	}

	err := service.DeleteUser(deleteInfo.UserId)

	// error checking
	if err != nil {
		if ent.IsNotFound(err) {
			return ctx.JSON(http.StatusNotFound, echo.Map{
				"message": "user not found",
			})
		}
		if err.Error() == "already deleted" {
			return ctx.JSON(http.StatusGone, echo.Map{
				"message": "user already deleted",
			})
		}

		log.Errorf("uncaught error deleting user: %v", err)

		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{"message": "ok"})
}
