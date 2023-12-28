package controllers

import (
    "github.com/Encedeus/panel/module"
    protoapi "github.com/Encedeus/panel/proto/go"
    "github.com/Encedeus/panel/services"
    "github.com/labstack/echo/v4"
    "net/http"
    "strconv"
)

type ModulesController struct {
    Controller
    ModuleStore *module.Store
}

func NewModulesController(ms *module.Store) ModulesController {
    return ModulesController{
        ModuleStore: ms,
    }
}

func (mc ModulesController) registerRoutes(srv *Server) {
    modulesEndpoint := srv.Group("modules")
    {
        modulesEndpoint.GET("", mc.handleFindAllModules)
    }
}

func (mc ModulesController) handleFindAllModules(c echo.Context) error {
    frontendOnlyParam := c.QueryParam("frontendOnly")
    backendOnlyParam := c.QueryParam("backendOnly")
    frontendOnly, err := strconv.ParseBool(frontendOnlyParam)
    backendOnly, err := strconv.ParseBool(backendOnlyParam)
    if err != nil && (len(backendOnlyParam) != 0 || len(frontendOnlyParam) != 0) {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "message": "invalid query parameters",
        })
    }

    req := &protoapi.FindAllModulesRequest{
        BackendOnly:  backendOnly,
        FrontendOnly: frontendOnly,
    }

    resp := services.FindAllModules(c.Request().Context(), mc.ModuleStore, req)

    return c.JSON(http.StatusOK, resp)
}
