package controllers

import (
	"errors"
	"github.com/Encedeus/panel/ent"
	"github.com/Encedeus/panel/proto"
	protoapi "github.com/Encedeus/panel/proto/go"
	"github.com/Encedeus/panel/services"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/encoding/protojson"
	"io"
	"net/http"
)

type NodesController struct {
	Controller
}

func (nc NodesController) registerRoutes(srv *Server) {
	nodesEndpoint := srv.Group("nodes")
	{
		nodesEndpoint.POST("", func(c echo.Context) error {
			return nc.handleCreateNode(c, srv.DB)
		})
		nodesEndpoint.GET("", func(c echo.Context) error {
			return nc.handleFindAllNodes(c, srv.DB)
		})
		nodesEndpoint.DELETE("", func(c echo.Context) error {
			return nc.handleDeleteNode(c, srv.DB)
		})
		nodesEndpoint.GET("/:id", func(c echo.Context) error {
			return nc.handleFindOneNode(c, srv.DB)
		})
	}
}

func (NodesController) handleCreateNode(c echo.Context, db *ent.Client) error {
	ctx := c.Request().Context()

	b, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "bad request",
		})
	}
	createReq := new(protoapi.NodesCreateRequest)
	err = protojson.Unmarshal(b, createReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "bad request",
		})
	}
	//err = c.Bind(createReq)
	/*	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "bad request",
		})
	}*/

	resp, err := services.CreateNode(ctx, db, createReq)
	if err != nil {
		if errors.Is(err, services.ValidationError{}) || errors.Is(err, services.ErrFailedConnectingToSkyhook) || errors.Is(err, services.ErrFailedGettingHardwareInfo) {
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

func (NodesController) handleFindAllNodes(c echo.Context, db *ent.Client) error {
	ctx := c.Request().Context()

	resp, err := services.FindAllNodes(ctx, db, &protoapi.NodesFindAllRequest{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return proto.MarshalControllerProtoResponseToJSON(&c, http.StatusOK, resp)
}

func (NodesController) handleDeleteNode(c echo.Context, db *ent.Client) error {
	ctx := c.Request().Context()

	deleteReq := new(protoapi.NodesDeleteRequest)
	err := proto.UnmarshalProtoBody(c, deleteReq)
	if err != nil {
		return err
	}
	/*	b, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "bad request",
			})
		}
		err = protojson.Unmarshal(b, deleteReq)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "bad request",
			})
		}*/

	resp, err := services.DeleteNode(ctx, db, deleteReq)
	if err != nil {
		if errors.Is(err, services.ErrNodeNotFound) {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return proto.MarshalControllerProtoResponseToJSON(&c, http.StatusNoContent, resp)
}

func (NodesController) handleFindOneNode(c echo.Context, db *ent.Client) error {
	ctx := c.Request().Context()

	paramId := c.Param("id")
	id, err := uuid.Parse(paramId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "invalid uuid",
		})
	}

	findReq := &protoapi.NodesFindOneRequest{
		Id: proto.UUIDToProtoUUID(id),
	}

	resp, err := services.FindNodeByID(ctx, db, findReq)
	if err != nil {
		if errors.Is(err, services.ErrNodeNotFound) {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return proto.MarshalControllerProtoResponseToJSON(&c, http.StatusOK, resp)
}
