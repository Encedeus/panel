package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"github.com/Encedeus/panel/controller"
	"github.com/Encedeus/panel/ent"
	"github.com/Encedeus/panel/middleware"
)

func RouteInit(server *echo.Echo, db *ent.Client) {
	server.Use(middleware.JSONSyntaxMiddleware)

	controller.InitControllers(server, db)

	server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}
