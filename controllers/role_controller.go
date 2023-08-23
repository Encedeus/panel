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
)

type RoleController struct {
    Controller
}

func (rc RoleController) registerRoutes(srv *Server) {
    roleEndpoint := srv.Group("role")
    {
        roleEndpoint.Use(middleware.AccessJWTAuth)

        roleEndpoint.GET("/:id", func(c echo.Context) error {
            return rc.handleFindRole(c, srv.DB)
        })
        roleEndpoint.POST("", func(c echo.Context) error {
            return rc.handleCreateRole(c, srv.DB)
        })
        roleEndpoint.PATCH("", func(c echo.Context) error {
            return rc.handleUpdateRole(c, srv.DB)
        })
        roleEndpoint.DELETE("/:id", func(c echo.Context) error {
            return rc.handleDeleteRole(c, srv.DB)
        })
    }
}

func (RoleController) handleFindRole(c echo.Context, db *ent.Client) error {
    ctx := c.Request().Context()

    id, err := uuid.Parse(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "message": "bad request",
        })
    }

    roleData, err := services.FindRole(ctx, db, id)

    if err != nil {
        if ent.IsNotFound(err) {
            return c.JSON(http.StatusNotFound, echo.Map{
                "message": "role not found",
            })
        }

        if err.Error() == "role deleted" {
            return c.JSON(http.StatusGone, echo.Map{
                "message": "role deleted",
            })
        }

        log.Errorf("uncaught error querying role: %v", err)

        return c.JSON(http.StatusInternalServerError, echo.Map{
            "message": "internal server error",
        })
    }

    return c.JSON(http.StatusOK, echo.Map{
        "name":        roleData.Name,
        "permissions": roleData.Permissions,
        "createdAt":   roleData.CreatedAt,
        "updatedAt":   roleData.UpdatedAt,
    })
}

func (RoleController) handleCreateRole(c echo.Context, db *ent.Client) error {
    ctx := c.Request().Context()
    userId, _ := middleware.IDFromAccessContext(ctx)

    // permission check
    if !services.DoesUserHavePermission(ctx, db, "create_role", userId) {
        return c.JSON(http.StatusUnauthorized, echo.Map{
            "message": "unauthorised",
        })
    }

    roleInfo := dto.CreateRoleDTO{}
    c.Bind(&roleInfo)

    // check if name is specified
    if roleInfo.Name == "" || len(roleInfo.Permissions) == 0 {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "message": "bad request",
        })
    }

    roleId, err := services.CreateRole(ctx, db, roleInfo)

    if err != nil {
        if ent.IsNotFound(err) {
            return c.JSON(http.StatusNotFound, echo.Map{
                "message": "role not found",
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

        log.Errorf("uncaught error creating role: %v", err)

        return c.JSON(http.StatusInternalServerError, echo.Map{
            "message": "internal server error",
        })
    }

    return c.JSON(http.StatusOK, echo.Map{"id": roleId})
}

func (RoleController) handleUpdateRole(c echo.Context, db *ent.Client) error {
    ctx := c.Request().Context()
    userId, _ := middleware.IDFromAccessContext(ctx)

    if !services.DoesUserHavePermission(ctx, db, "update_role", userId) {
        return c.JSON(http.StatusUnauthorized, echo.Map{
            "message": "unauthorised",
        })
    }

    roleInfo := dto.UpdateRoleDTO{}
    _ = c.Bind(&roleInfo)

    // check if all fields are provided
    if (roleInfo.Name == "" && len(roleInfo.Permissions) == 0) || roleInfo.ID.String() == "" {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "message": "bad request",
        })
    }

    err := services.UpdateRole(ctx, db, roleInfo)

    if err != nil {
        if ent.IsNotFound(err) {
            return c.JSON(http.StatusNotFound, echo.Map{
                "message": "role not found",
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
        if err.Error() == "role deleted" {
            return c.JSON(http.StatusGone, echo.Map{
                "message": "role deleted",
            })
        }

        log.Errorf("uncaught error updating role: %v", err)

        return c.JSON(http.StatusInternalServerError, echo.Map{
            "message": "internal server error",
        })
    }

    return c.NoContent(http.StatusOK)
}

func (RoleController) handleDeleteRole(c echo.Context, db *ent.Client) error {
    ctx := c.Request().Context()
    userId, _ := middleware.IDFromAccessContext(ctx)

    if !services.DoesUserHavePermission(ctx, db, "delete_role", userId) {
        return c.JSON(http.StatusUnauthorized, echo.Map{
            "message": "unauthorised",
        })
    }

    id, err := uuid.Parse(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "message": "bad request",
        })
    }

    err = services.DeleteRole(ctx, db, id)
    if err != nil {
        if ent.IsNotFound(err) {
            return c.JSON(http.StatusNotFound, echo.Map{
                "message": "role not found",
            })
        }

        if err.Error() == "already deleted" {
            return c.JSON(http.StatusGone, echo.Map{
                "message": "role already deleted",
            })
        }
        log.Errorf("uncaught error deleting role: %v", err)

        return c.JSON(http.StatusInternalServerError, echo.Map{
            "message": "internal server error",
        })
    }

    return c.NoContent(http.StatusOK)
}
