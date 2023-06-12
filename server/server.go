package server

import (
	"github.com/labstack/echo/v4"
	"panel/config"
)

func ServerInit() {
	e := echo.New()
	RouteInit(e)
	e.Logger.Fatal(e.Start(config.Config.Server.URI()))
}
