package server

import (
	"panel/config"
	"panel/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "PATCH", "HEAD"},
		AllowHeaders:     []string{"Accept", "Content-Type", "Authorization"},
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowCredentials: true,
	}))
	RouteInit(e, service.InitDB())

	e.Logger.Fatal(e.Start(config.Config.Server.URI()))
}
