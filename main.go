package main

import (
	"Panel/routes"
	"github.com/labstack/echo/v4"
)

var e = echo.New()

func main() {
	routes.RouteInit(e)
	e.Logger.Fatal(e.Start(":42069"))
}
