package server

import (
	"Panel/config"
	"github.com/labstack/echo/v4"
)

func ServerInit() {
	e := echo.New()
	RouteInit(e)
	e.Logger.Fatal(e.Start(config.Config.Server.URI()))
}
