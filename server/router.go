package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func RouteInit(server *echo.Echo) {
	// this will later point to controller functions
	server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}
