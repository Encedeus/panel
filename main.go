package main

import (
	"Panel/config"
	"Panel/server"
	"Panel/service"
	"github.com/labstack/echo/v4"
)

var e = echo.New()

func main() {
	config.InitConfig()
	server.RouteInit(e)
	service.DBInit()

	e.Logger.Fatal(config.Config.Server.URI)
}
