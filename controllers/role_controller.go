package controllers

import (
    "github.com/Encedeus/panel/dto"
    "github.com/Encedeus/panel/ent"
    "github.com/Encedeus/panel/middleware"
    "github.com/Encedeus/panel/services"
    "github.com/google/uuid"
    "github.com/labstack/echo/v4"
    "github.com/labstack/gommon/log"
    "net/http"
    "strconv"
)

func init() {
    addController(func(server *echo.Echo, db *ent.Client) {
        roleEndpoint := server.Group("role")

        roleEndpoint.Use(middleware.AccessJWTAuth)

        roleEndpoint.GET("/:id", getRole)
        roleEndpoint.POST("", handleCreateRole)
        roleEndpoint.PATCH("", handleUpdateRole)
        roleEndpoint.DELETE("/:id", handleDeleteRole)
    })
}

func getRole(ctx echo.Context) error {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil || id <= 0 {
        return ctx.JSON(http.StatusBadRequest, echo.Map{
            "message": "bad request",
        })
    }

    roleData, err := services.GetRole(id)

    if err != nil {
        if ent.IsNotFound(err) {
            return ctx.JSON(http.StatusNotFound, echo.Map{
                "message": "role not found",
            })
        }

        if err.Error() == "role deleted" {
            return ctx.JSON(http.StatusGone, echo.Map{
                "message": "role deleted",
            })
        }

        log.Errorf("uncaught error querying role: %v", err)

        return ctx.JSON(http.StatusInternalServerError, echo.Map{
            "message": "internal server error",
        })
    }

    return ctx.JSON(http.StatusOK, echo.Map{
        "name":        roleData.Name,
        "permissions": roleData.Permissions,
        "createdAt":   roleData.CreatedAt,
        "updatedAt":   roleData.UpdatedAt,
    })
}

func handleCreateRole(ctx echo.Context) error {
    userId, _ := uuid.Parse(ctx.Request().Header.Get("UUID"))

    // permission check
    if !services.DoesUserHavePermission("create_role", userId) {
        return ctx.JSON(http.StatusUnauthorized, echo.Map{
            "message": "unauthorised",
        })
    }

    roleInfo := dto.CreateRoleDTO{}
    ctx.Bind(&roleInfo)

    // check if name is specified
    if roleInfo.Name == "" || len(roleInfo.Permissions) == 0 {
        return ctx.JSON(http.StatusBadRequest, echo.Map{
            "message": "bad request",
        })
    }

    roleId, err := services.CreateRole(roleInfo.Name, roleInfo.Permissions)

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

        log.Errorf("uncaught error creating role: %v", err)

        return ctx.JSON(http.StatusInternalServerError, echo.Map{
            "message": "internal server error",
        })
    }

    return ctx.JSON(http.StatusOK, echo.Map{"id": roleId})
}

func handleUpdateRole(ctx echo.Context) error {
    userId, _ := uuid.Parse(ctx.Request().Header.Get("UUID"))

    if !services.DoesUserHavePermission("update_role", userId) {
        return ctx.JSON(http.StatusUnauthorized, echo.Map{
            "message": "unauthorised",
        })
    }

    roleInfo := dto.UpdateRoleDTO{}
    ctx.Bind(&roleInfo)

    // check if all fields are provided
    if (roleInfo.Name == "" && len(roleInfo.Permissions) == 0) || roleInfo.Id == 0 {
        return ctx.JSON(http.StatusBadRequest, echo.Map{
            "message": "bad request",
        })
    }

    err := services.UpdateRole(roleInfo)

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
        if err.Error() == "role deleted" {
            return ctx.JSON(http.StatusGone, echo.Map{
                "message": "role deleted",
            })
        }

        log.Errorf("uncaught error updating role: %v", err)

        return ctx.JSON(http.StatusInternalServerError, echo.Map{
            "message": "internal server error",
        })
    }

    return ctx.NoContent(http.StatusOK)
}

func handleDeleteRole(ctx echo.Context) error {
    userId, _ := uuid.Parse(ctx.Request().Header.Get("UUID"))

    if !services.DoesUserHavePermission("delete_role", userId) {
        return ctx.JSON(http.StatusUnauthorized, echo.Map{
            "message": "unauthorised",
        })
    }

    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil || id <= 0 {
        return ctx.JSON(http.StatusBadRequest, echo.Map{
            "message": "bad request",
        })
    }

    err = services.DeleteRole(id)
    if err != nil {
        if ent.IsNotFound(err) {
            return ctx.JSON(http.StatusNotFound, echo.Map{
                "message": "role not found",
            })
        }

        if err.Error() == "already deleted" {
            return ctx.JSON(http.StatusGone, echo.Map{
                "message": "role already deleted",
            })
        }
        log.Errorf("uncaught error deleting role: %v", err)

        return ctx.JSON(http.StatusInternalServerError, echo.Map{
            "message": "internal server error",
        })
    }

    return ctx.NoContent(http.StatusOK)
}
