package controllers

import (
    "github.com/Encedeus/panel/ent"
    "github.com/Encedeus/panel/proto"
    protoapi "github.com/Encedeus/panel/proto/go"
    "github.com/Encedeus/panel/services"
    "github.com/google/uuid"
    "github.com/labstack/echo/v4"
    "google.golang.org/protobuf/encoding/protojson"
    "net/http"
)

type APIKeyController struct {
    Controller
}

func (akc APIKeyController) registerRoutes(srv *Server) {
    keyEndpoint := srv.Group("key/account")
    {
        /*        keyEndpoint.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
                  return middleware.AccessJWTAuth(srv.DB, next)
              })*/

        keyEndpoint.POST("", func(c echo.Context) error {
            return akc.handleCreateAccountAPIKey(c, srv.DB)
        })
        keyEndpoint.DELETE("/:id", func(c echo.Context) error {
            return akc.handleDeleteAccountAPIKey(c, srv.DB)
        })
        /*            accountEndpoint.GET("/:id", func(c echo.Context) error {
                      return akc.handleFindAccountAPIKeyByID(c, srv.DB)
                  })*/
        keyEndpoint.GET("/:userId", func(c echo.Context) error {
            return akc.handleFindAccountAPIKeysByUserId(c, srv.DB)
        })
    }
}

func (APIKeyController) handleCreateAccountAPIKey(c echo.Context, db *ent.Client) (err error) {
    ctx := c.Request().Context()

    createReq := new(protoapi.AccountAPIKeyCreateRequest)

    body := make([]byte, c.Request().ContentLength)
    _, err = c.Request().Body.Read(body)
    err = protojson.Unmarshal(body, createReq)
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "message": err.Error(),
        })
    }

    resp, err := services.CreateAccountAPIKey(ctx, db, createReq)
    if err != nil {
        if services.IsValidationError(err) {
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

func (APIKeyController) handleDeleteAccountAPIKey(c echo.Context, db *ent.Client) (err error) {
    ctx := c.Request().Context()

    id, err := uuid.Parse(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "message": err.Error(),
        })
    }

    _, err = services.DeleteAccountAPIKey(ctx, db, &protoapi.AccountAPIKeyDeleteRequest{
        Id: proto.UUIDToProtoUUID(id),
    })
    if err != nil {
        if ent.IsNotFound(err) {
            return c.JSON(http.StatusNotFound, echo.Map{
                "message": "api key not found",
            })
        }

        if err.Error() == "already deleted" {
            return c.JSON(http.StatusGone, echo.Map{
                "message": "api key already deleted",
            })
        }

        return c.JSON(http.StatusInternalServerError, echo.Map{
            "message": err.Error(),
        })
    }

    return c.NoContent(http.StatusOK)
}

func (APIKeyController) handleFindAccountAPIKeysByUserId(c echo.Context, db *ent.Client) (err error) {
    ctx := c.Request().Context()

    userId, err := uuid.Parse(c.Param("userId"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "message": "invalid UUID",
        })
    }

    resp, err := services.FindAccountAPIKeysByUserID(ctx, db, &protoapi.AccountAPIkeyFindManyByUserRequest{
        UserId: proto.UUIDToProtoUUID(userId),
    })
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

        return c.JSON(http.StatusInternalServerError, echo.Map{
            "message": "internal server error",
        })
    }

    return proto.MarshalControllerProtoResponseToJSON(&c, http.StatusOK, resp)
}

func (APIKeyController) handleFindAccountAPIKeyByID(c echo.Context, db *ent.Client) (err error) {
    ctx := c.Request().Context()

    id, err := uuid.Parse(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "message": "invalid UUID",
        })
    }

    resp, err := services.FindAccountAPIKeyByID(ctx, db, &protoapi.AccountAPIKeyFindOneRequest{
        Id: proto.UUIDToProtoUUID(id),
    })
    if err != nil {
        return c.JSON(http.StatusNotFound, echo.Map{
            "message": "api key not found",
        })
    }

    return proto.MarshalControllerProtoResponseToJSON(&c, http.StatusOK, resp)
}
