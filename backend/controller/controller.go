package controller

import (
	"github.com/labstack/echo/v4"
	"panel/ent"
)

type IController interface {
	Init(*echo.Echo, *ent.Client)
}

var controllers []IController

type templateController struct {
	init func(server *echo.Echo, db *ent.Client)
}

func (tc templateController) Init(server *echo.Echo, db *ent.Client) {
	tc.init(server, db)
}

func addController(init func(server *echo.Echo, db *ent.Client)) {
	temp := templateController{}
	temp.init = init
	controllers = append(controllers, temp)
}

func InitControllers(server *echo.Echo, db *ent.Client) {
	for _, controller := range controllers {
		controller.Init(server, db)
	}
}
