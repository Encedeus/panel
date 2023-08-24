package controllers

import (
    "github.com/Encedeus/panel/dto"
    "github.com/Encedeus/panel/ent"
    "github.com/Encedeus/panel/services"
    "github.com/google/uuid"
    "github.com/labstack/echo/v4"
    "net/http"
)

type APIKeyController struct {
    Controller
}

func (akc APIKeyController) registerRoutes(srv *Server) {
    keyEndpoint := srv.Group("key/account")
    {
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

    apiKeyData := new(dto.AccountAPIKeyDTO)
    err = c.Bind(apiKeyData)
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "message": err.Error(),
        })
    }

    apiKey, err := services.CreateAccountAPIKey(ctx, db, *apiKeyData)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{
            "message": err.Error(),
        })
    }

    return c.JSON(http.StatusCreated, echo.Map{
        "data": apiKey,
    })
}

func (APIKeyController) handleDeleteAccountAPIKey(c echo.Context, db *ent.Client) (err error) {
    ctx := c.Request().Context()

    id, err := uuid.Parse(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "message": err.Error(),
        })
    }

    err = services.DeleteAccountAPIKey(ctx, db, id)
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

    apiKeys, err := services.FindAccountAPIKeysByUserID(ctx, db, userId)
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

    return c.JSON(http.StatusOK, echo.Map{
        "keys": apiKeys,
    })
}

func (APIKeyController) handleFindAccountAPIKeyByID(c echo.Context, db *ent.Client) (err error) {
    ctx := c.Request().Context()

    id, err := uuid.Parse(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "message": "invalid UUID",
        })
    }

    apiKey, err := services.FindAccountAPIKeyByID(ctx, db, id)
    if err != nil {
        return c.JSON(http.StatusNotFound, echo.Map{
            "message": "api key not found",
        })
    }

    return c.JSON(http.StatusOK, echo.Map{
        "data": []ent.ApiKey{*apiKey},
    })
}
