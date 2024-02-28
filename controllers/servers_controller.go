package controllers

import (
	"errors"
	"github.com/Encedeus/panel/ent"
	"github.com/Encedeus/panel/module"
	"github.com/Encedeus/panel/proto"
	protoapi "github.com/Encedeus/panel/proto/go"
	"github.com/Encedeus/panel/services"
	"github.com/google/uuid"
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
		serversEndpoint.GET("", func(c echo.Context) error {
			return sc.handleFindAllServers(c, srv.DB, srv.ModuleStore)
		})
		serversEndpoint.GET("/:id", func(c echo.Context) error {
			return sc.handleFindOneServer(c, srv.DB, srv.ModuleStore)
		})
		/*		serversEndpoint.GET("/:id/status", func(c echo.Context) error {
				return sc.handleGetServerStatus(c, srv.DB, srv.ModuleStore)
			})*/
		serversEndpoint.DELETE("/:id", func(c echo.Context) error {
			return sc.handleDeleteServer(c, srv.DB)
		})
		serversEndpoint.POST("/:id/start", func(c echo.Context) error {
			return sc.handleStartServer(c, srv.DB, srv.ModuleStore)
		})
		serversEndpoint.POST("/:id/restart", func(c echo.Context) error {
			return sc.handleRestartServer(c, srv.DB, srv.ModuleStore)
		})
		serversEndpoint.POST("/:id/stop", func(c echo.Context) error {
			return sc.handleStopServer(c, srv.DB, srv.ModuleStore)
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

func (sc ServersController) handleFindAllServers(c echo.Context, db *ent.Client, store *module.Store) error {
	ctx := c.Request().Context()

	srvs, err := services.FindAllServers(ctx, db, &protoapi.ServersFindAllRequest{})
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

	protoSrvs := make([]*protoapi.Server, len(srvs))
	for idx, srv := range srvs {
		protoSrvs[idx] = proto.EntServerToProtoServer(*srv, store)
	}

	resp := &protoapi.ServersFindAllResponse{
		Servers: protoSrvs,
	}

	return proto.MarshalControllerProtoResponseToJSON(&c, http.StatusOK, resp)
}

func (sc ServersController) handleFindOneServer(c echo.Context, db *ent.Client, store *module.Store) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	serverId, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "bad request",
		})
	}

	req := &protoapi.ServersFindOneRequest{
		Id: proto.UUIDToProtoUUID(serverId),
	}

	srv, err := services.FindOneServer(ctx, db, req)
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

	resp := &protoapi.ServersFindAllResponse{
		Servers: []*protoapi.Server{
			proto.EntServerToProtoServer(*srv, store),
		},
	}

	return proto.MarshalControllerProtoResponseToJSON(&c, http.StatusOK, resp)
}

func (ServersController) handleDeleteServer(c echo.Context, db *ent.Client) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	serverId, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "bad request",
		})
	}

	req := &protoapi.ServersDeleteRequest{
		Id: proto.UUIDToProtoUUID(serverId),
	}

	err = services.DeleteOneServer(ctx, db, req)
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

	return c.NoContent(http.StatusNoContent)
}

func (ServersController) handleStartServer(c echo.Context, db *ent.Client, store *module.Store) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	serverId, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "bad request",
		})
	}

	err = services.StartServer(ctx, store, db, serverId)
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

	return c.NoContent(http.StatusNoContent)
}

func (ServersController) handleRestartServer(c echo.Context, db *ent.Client, store *module.Store) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	serverId, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "bad request",
		})
	}

	err = services.RestartServer(ctx, store, db, serverId)
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

	return c.NoContent(http.StatusNoContent)
}

func (ServersController) handleStopServer(c echo.Context, db *ent.Client, store *module.Store) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	serverId, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "bad request",
		})
	}

	err = services.StopServer(ctx, store, db, serverId)
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

	return c.NoContent(http.StatusNoContent)
}

/*func (ServersController) handleGetServerStatus(c echo.Context, db *ent.Client) error {

}*/
