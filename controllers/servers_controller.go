package controllers

import (
	"errors"
	"fmt"
	"github.com/Encedeus/panel/ent"
	"github.com/Encedeus/panel/module"
	"github.com/Encedeus/panel/proto"
	protoapi "github.com/Encedeus/panel/proto/go"
	"github.com/Encedeus/panel/services"
	"github.com/docker/docker/api/types/container"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"io"
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
		serversEndpoint.GET("/:id/status", func(c echo.Context) error {
			return sc.handleGetServerStatus(c, srv.DB)
		})
		serversEndpoint.GET("/:id/console", func(c echo.Context) error {
			return sc.handleUpgradeConsole(c, srv.DB)
		})
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

func (ServersController) handleGetServerStatus(c echo.Context, db *ent.Client) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	serverId, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "bad request",
		})
	}

	state, err := services.InspectServerContainer(ctx, db, proto.UUIDToProtoUUID(serverId))
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

	return c.JSON(http.StatusOK, echo.Map{
		"status": state.State.Status,
	})
}

func (ServersController) handleUpgradeConsole(c echo.Context, db *ent.Client) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	serverId, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "bad request",
		})
	}

	srv, err := services.FindServerByID(ctx, db, serverId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	cli, err := services.CreateNodeDockerClient(ctx, db, proto.UUIDToProtoUUID(serverId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	resp, err := cli.ContainerAttach(ctx, srv.ContainerId, container.AttachOptions{
		Stdin:  true,
		Stream: true,
		Logs:   true,
		Stdout: true,
		Stderr: true,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}
	defer resp.Close()

	type ConsoleCommand struct {
		Cmd string `json:"cmd"`
	}

	upgrader := websocket.Upgrader{}
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	cl := make(chan struct{}, 1)

	go func(ws *websocket.Conn) {
		for {
			select {
			case <-cl:
				break
			default:
				fmt.Println("going")
				//logs := make([]byte, 0, resp.Reader.Size())
				//_, err := io.ReadAtLeast(resp.R, logs)
				//io.ReadFull(resp.Reader, logs)
				//logs, err := io.ReadAll(resp.Reader)
				for {
					r, err := resp.Reader.ReadString('\n')
					if err != nil {
						if errors.Is(err, io.EOF) {
							break
						}
						cl <- struct{}{}
						close(cl)
					}
					err = ws.WriteMessage(websocket.TextMessage, []byte(string(r)))
					if err != nil {
						c.Logger().Error(err)
						cl <- struct{}{}
						close(cl)
						break
					}
				}
				/*				if err != nil {
									_ = ws.WriteJSON(echo.Map{
										"message": "internal error",
									})
								}

								if len(logs) > 0 {
									err = ws.WriteJSON(echo.Map{
										"logs": base64.StdEncoding.EncodeToString(logs),
									})
									if err != nil {
										c.Logger().Error(err)
										cl <- struct{}{}
										close(cl)
										break
									}
								}*/
			}
		}
	}(ws)

	go func(ws *websocket.Conn) {
		for {
			select {
			case <-cl:
				break
			default:
				cmd := new(ConsoleCommand)
				err = ws.ReadJSON(cmd)
				if err != nil {
					cl <- struct{}{}
					close(cl)
					c.Logger().Error(err)
					break
				}

				_, err = resp.Conn.Write([]byte(fmt.Sprintf("%s\x0D", cmd.Cmd)))
				if err != nil {
					_ = ws.WriteJSON(echo.Map{
						"message": "internal error",
					})
				}
			}
		}
	}(ws)

	/*	for {
		_, _, err := ws.UnderlyingConn()
		if errors.Is(err, &websocket.CloseError{}) {
			cl <- struct{}{}
			close(cl)
			return c.NoContent(http.StatusNoContent)
		}
	}*/

	for {
		select {
		case <-cl:
			break
		default:
		}
	}
}
