package controllers

import (
	"github.com/Encedeus/panel/ent"
	"github.com/Encedeus/panel/middleware"
	"github.com/Encedeus/panel/proto"
	protoapi "github.com/Encedeus/panel/proto/go"
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

		roleEndpoint.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
			return middleware.AccessJWTAuth(srv.DB, next)
		})

		roleEndpoint.GET("", func(c echo.Context) error {
			return rc.handleFindAllRoles(c, srv.DB)
		})
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

	resp, err := services.FindRole(ctx, db, &protoapi.RoleFindOneRequest{
		Id: proto.UUIDToProtoUUID(id),
	})
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

	return proto.MarshalControllerProtoResponseToJSON(&c, http.StatusOK, resp)
}
func (RoleController) handleFindAllRoles(c echo.Context, db *ent.Client) error {
	ctx := c.Request().Context()

	resp, err := services.FindAllRoles(ctx, db)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "internal server error",
		})
	}

	return proto.MarshalControllerProtoResponseToJSON(&c, http.StatusOK, resp)
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

	createReq := new(protoapi.RoleCreateRequest)
	err := c.Bind(createReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "bad request",
		})
	}

	// check if name is specified
	if createReq.Name == "" || len(createReq.Permissions) == 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "bad request",
		})
	}

	resp, err := services.CreateRole(ctx, db, createReq)
	if err != nil {
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "role not found",
			})
		}
		if ent.IsValidationError(err) {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "validate error",
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

	return proto.MarshalControllerProtoResponseToJSON(&c, http.StatusOK, resp)
}

func (RoleController) handleUpdateRole(c echo.Context, db *ent.Client) error {
	ctx := c.Request().Context()
	userId, _ := middleware.IDFromAccessContext(ctx)

	if !services.DoesUserHavePermission(ctx, db, "update_role", userId) {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "unauthorised",
		})
	}

	updateReq := new(protoapi.RoleUpdateRequest)
	err := c.Bind(updateReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "bad request",
		})
	}

	// check if all fields are provided
	if (updateReq.Name == "" && len(updateReq.Permissions) == 0) || updateReq.Id.Value == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "bad request",
		})
	}

	resp, err := services.UpdateRole(ctx, db, updateReq)

	if err != nil {
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "role not found",
			})
		}
		if ent.IsValidationError(err) {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "validate error",
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

		log.Errorf("uncaught error updating role: %e", err)

		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "internal server error",
		})
	}

	return proto.MarshalControllerProtoResponseToJSON(&c, http.StatusOK, resp)
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

	_, err = services.DeleteRole(ctx, db, &protoapi.RoleDeleteRequest{
		Id: proto.UUIDToProtoUUID(id),
	})
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
