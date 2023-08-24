package controllers

import (
    "fmt"
    "github.com/Encedeus/panel/config"
    "github.com/Encedeus/panel/dto"
    "github.com/Encedeus/panel/ent"
    "github.com/Encedeus/panel/hashing"
    "github.com/Encedeus/panel/middleware"
    "github.com/Encedeus/panel/services"
    "github.com/google/uuid"
    "github.com/labstack/echo/v4"
    "github.com/labstack/gommon/log"
    "io"
    "net/http"
    "os"
    "strings"
)

type UserController struct {
    Controller
}

func (uc UserController) registerRoutes(srv *Server) {
    userEndpoint := srv.Group("user")
    {
        userEndpoint.Static("/pfp", config.Config.CDN.Directory)

        userEndpoint.Use(middleware.AccessJWTAuth)

        userEndpoint.GET("/:id", func(c echo.Context) error {
            return handleFindUser(c, srv.DB)
        })
        userEndpoint.POST("", func(c echo.Context) error {
            return handleCreateUser(c, srv.DB)
        })
        userEndpoint.PUT("", func(c echo.Context) error {
            return handleSetPfp(c, srv.DB)
        })
        userEndpoint.PATCH("", func(c echo.Context) error {
            return handleUpdateUser(c, srv.DB)
        })
        userEndpoint.DELETE("/:id", func(c echo.Context) error {
            return handleDeleteUser(c, srv.DB)
        })
    }
}

func handleFindUser(c echo.Context, db *ent.Client) error {
    ctx := c.Request().Context()
    rawUserId := c.Param("id")

    userId, err := uuid.Parse(rawUserId)
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"message": "bad request"})
    }

    userData, err := services.GetUser(ctx, db, userId)

    if err != nil {
        if ent.IsNotFound(err) {
            return c.JSON(http.StatusNotFound, echo.Map{"message": "user not found"})
        }
        if err.Error() == "user deleted" {
            return c.JSON(http.StatusGone, echo.Map{"message": "user deleted"})
        }

        log.Errorf("error querying user: %v", err)

        return c.JSON(http.StatusInternalServerError, echo.Map{
            "message": "internal server error",
        })
    }

    return c.JSON(http.StatusOK, echo.Map{
        "id":        userData.ID,
        "name":      userData.Name,
        "email":     userData.Email,
        "createdAt": userData.CreatedAt,
        "updatedAt": userData.UpdatedAt,
        "roleId":    userData.RoleID,
    })
}

func handleCreateUser(c echo.Context, db *ent.Client) error {
    // get uuid from header provided by the middleware
    ctx := c.Request().Context()
    authUUID, _ := middleware.IDFromAccessContext(ctx)
    // check permissions
    if !services.DoesUserHavePermission(ctx, db, "create_user", authUUID) {
        return c.JSON(http.StatusUnauthorized, echo.Map{
            "message": "unauthorised",
        })
    }

    userInfo := dto.CreateUserDTO{}
    c.Bind(&userInfo)

    // check if all the fields are provided

    if strings.TrimSpace(userInfo.Name) == "" || strings.TrimSpace(userInfo.Password) == "" || strings.TrimSpace(userInfo.Email) == "" {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "message": "bad request",
        })
    }

    var (
        err    error
        userId *uuid.UUID
    )
    // check which method was used for role assignment
    if userInfo.RoleName != "" {
        userId, err = services.CreateUserRoleName(ctx, db, userInfo.Name, userInfo.Email, hashing.HashPassword(userInfo.Password), userInfo.RoleName)
    } else if userInfo.RoleId.String() == "" {
        userId, err = services.CreateUserRoleId(ctx, db, userInfo.Name, userInfo.Email, hashing.HashPassword(userInfo.Password), userInfo.RoleId)
    } else {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "message": "either role name or id must be specified",
        })
    }

    // error checking
    if err != nil {
        if ent.IsNotFound(err) {
            return c.JSON(http.StatusNotFound, echo.Map{
                "message": "role not found",
            })
        }

        if ent.IsConstraintError(err) {
            return c.JSON(http.StatusConflict, echo.Map{
                "message": "username taken",
            })
        }

        // log any uncaught errors
        log.Errorf("uncaught error querying role: %v", err)

        return c.JSON(http.StatusInternalServerError, echo.Map{
            "message": "internal server error",
        })
    }

    return c.JSON(http.StatusCreated, echo.Map{"uuid": userId.String()})
}

func handleUpdateUser(c echo.Context, db *ent.Client) error {
    ctx := c.Request().Context()
    authUUID, _ := middleware.IDFromAccessContext(ctx)

    // check permissions
    if !services.DoesUserHavePermission(ctx, db, "update_user", authUUID) {
        return c.JSON(http.StatusUnauthorized, echo.Map{
            "message": "unauthorised",
        })
    }

    updateInfo := dto.UpdateUserDTO{}
    _ = c.Bind(&updateInfo)

    // check if no fields are provided
    if (updateInfo.Name == "" && updateInfo.Email == "" && updateInfo.Password == "" && updateInfo.RoleName == "" && updateInfo.RoleId.String() == "") || updateInfo.UserId.ID() == 0 {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "message": "bad request",
        })
    }

    updateInfo.Password = hashing.HashPassword(updateInfo.Password)
    err := services.UpdateUser(ctx, db, updateInfo)

    // error checking
    if err != nil {

        if ent.IsNotFound(err) {
            if strings.Contains(err.Error(), "role") {
                return c.JSON(http.StatusNotFound, echo.Map{
                    "message": "role not found",
                })
            }
            return c.JSON(http.StatusNotFound, echo.Map{
                "message": "user not found",
            })
        }
        if ent.IsValidationError(err) {
            return c.JSON(http.StatusBadRequest, echo.Map{
                "message": "validation error",
            })
        }
        if ent.IsConstraintError(err) {
            return c.JSON(http.StatusBadRequest, echo.Map{
                "message": "constraint error",
            })
        }
        if err.Error() == "user deleted" {
            return c.JSON(http.StatusGone, echo.Map{
                "message": "user deleted",
            })
        }

        log.Errorf("uncaught error updating user: %v", err)

        return c.JSON(http.StatusInternalServerError, echo.Map{
            "message": "internal server error",
        })
    }

    return c.NoContent(http.StatusOK)
}

func handleSetPfp(c echo.Context, db *ent.Client) error {
    ctx := c.Request().Context()
    // get params from multipart
    userId := c.FormValue("uuid")
    file, err := c.FormFile("file")

    // validate request
    if err != nil || userId == "" {
        return c.JSON(http.StatusBadRequest, echo.Map{"message": "bad request"})
    }

    userUUID, err := uuid.Parse(userId)
    if !services.DoesUserWithUUIDExist(ctx, db, userUUID) {
        return c.JSON(http.StatusNotFound, echo.Map{"message": "user not found"})
    }

    // open file
    src, err := file.Open()
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"message": "invalid file format"})
    }
    defer src.Close()

    // create file
    dst, err := os.Create(fmt.Sprintf("%s/%s", config.Config.CDN.Directory, userId))
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"message": "internal server error"})
    }

    // write to file
    if _, err = io.Copy(dst, src); err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"message": "internal server error"})
    }

    return c.NoContent(http.StatusOK)
}

func handleDeleteUser(c echo.Context, db *ent.Client) error {
    ctx := c.Request().Context()
    authUUID, _ := middleware.IDFromRefreshContext(ctx)

    // check permissions
    if !services.DoesUserHavePermission(ctx, db, "delete_user", authUUID) {
        return c.JSON(http.StatusUnauthorized, echo.Map{
            "message": "unauthorised",
        })
    }

    rawUserId := c.Param("id")

    userId, err := uuid.Parse(rawUserId)
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"message": "bad request"})
    }

    err = services.DeleteUser(ctx, db, userId)

    // error checking
    if err != nil {
        if ent.IsNotFound(err) {
            return c.JSON(http.StatusNotFound, echo.Map{
                "message": "user not found",
            })
        }
        if err.Error() == "already deleted" {
            return c.JSON(http.StatusGone, echo.Map{
                "message": "user already deleted",
            })
        }

        log.Errorf("uncaught error deleting user: %v", err)

        return c.JSON(http.StatusInternalServerError, echo.Map{
            "message": "internal server error",
        })
    }

    return c.NoContent(http.StatusOK)
}
