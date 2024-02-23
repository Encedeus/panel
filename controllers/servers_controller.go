package controllers

import (
	"errors"
	"github.com/Encedeus/panel/ent"
	"github.com/Encedeus/panel/module"
	"github.com/Encedeus/panel/proto"
	protoapi "github.com/Encedeus/panel/proto/go"
	"github.com/Encedeus/panel/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ServersController struct {
	Controller
}

func (sc ServersController) registerRoutes(srv *Server) {
	serversEndpoint := srv.Group("servers")
	{
		serversEndpoint.POST("", func(c echo.Context) error {
			return sc.handleCreateServer(c, srv.DB, srv.ModuleStore)
		})
	}
}

func (ServersController) handleCreateServer(c echo.Context, db *ent.Client, store *module.Store) error {
	ctx := c.Request().Context()

	createReq := new(protoapi.ServersCreateRequest)
	err := proto.UnmarshalProtoBody(c, createReq)
	if err != nil {
		return err
	}

	resp, err := services.CreateServer(ctx, db, store, createReq)
	if err != nil {
		if errors.Is(err, services.ValidationError{}) {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return proto.MarshalControllerProtoResponseToJSON(&c, http.StatusCreated, resp)
}
