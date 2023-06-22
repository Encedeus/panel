package server

import (
	"github.com/labstack/echo/v4"
	"panel/config"
	"panel/service"
)

func ServerInit() {
	e := echo.New()
	RouteInit(e, service.InitDB())
	e.Logger.Fatal(e.Start(config.Config.Server.URI()))
}
