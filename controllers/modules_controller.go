package controllers

import (
	"context"
	"errors"
	"github.com/Encedeus/panel/api"
	"github.com/Encedeus/panel/module"
	"github.com/Encedeus/panel/proto"
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
		modulesEndpoint.GET("/:id", mc.handleFindOneModule)
		modulesEndpoint.POST("", mc.handleInstall)
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

	req := &protoapi.ModulesFindAllRequest{
		BackendOnly:  backendOnly,
		FrontendOnly: frontendOnly,
	}

	resp := services.FindAllModules(c.Request().Context(), mc.ModuleStore, req)

	return proto.MarshalControllerProtoResponseToJSON(&c, http.StatusOK, resp)
}

func (mc ModulesController) handleFindOneModule(c echo.Context) error {
	moduleId := c.Param("id")

	req := &protoapi.ModulesFindOneRequest{
		Id: &protoapi.UUID{
			Value: moduleId,
		},
	}

	resp, err := services.FindOneModule(context.Background(), mc.ModuleStore, req)
	if errors.Is(err, services.ErrModuleNotFound) {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": err.Error(),
		})
	}

	return proto.MarshalControllerProtoResponseToJSON(&c, http.StatusOK, resp)
}

// todo: unstupid after friday
func (mc ModulesController) handleInstall(c echo.Context) error {
	installReq := new(protoapi.ModuleInstallRequest)
	err := c.Bind(installReq)
	if err != nil || installReq.ModuleId == nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "bad request",
		})
	}

	fileName, uri, err := api.GetLatestReleaseDownloadURI(installReq.ModuleId.Value)

	if err != nil {
		return err
	}

	err = services.InstallModule(fileName, uri)
	if err != nil {
		return err
	}

	return c.NoContent(200)
}
