package routes

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func RouteInit(ech *echo.Echo) {
	// this will later point to controller functions

	ech.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}
