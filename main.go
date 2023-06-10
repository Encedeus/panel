package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

var e = echo.New()

func main() {
	routeInit(e)
	e.Logger.Fatal(e.Start(":42069"))
}

func routeInit(ech *echo.Echo) {
	ech.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}
