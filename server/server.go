package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"panel/config"
	"panel/service"
)

func ServerInit() {
	e := echo.New()
	e.Use(middleware.CORS())
	RouteInit(e, service.InitDB())
	e.Logger.Fatal(e.Start(config.Config.Server.URI()))
}
