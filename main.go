package main

import (
	"Panel/server"
	"github.com/labstack/echo/v4"
)

var e = echo.New()

func main() {
	server.RouteInit(e)
	e.Logger.Fatal(e.Start(":42069"))
}
