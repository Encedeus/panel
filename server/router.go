package server

import (
	"github.com/Encedeus/panel/controllers"
	"github.com/Encedeus/panel/ent"
	"github.com/Encedeus/panel/middleware"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RouteInit(server *echo.Echo, db *ent.Client) {
    server.Use(middleware.JSONSyntaxMiddleware)

    controllers.InitControllers(server, db)

    server.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })
}
