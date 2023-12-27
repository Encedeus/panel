package controllers

type ModulesController struct {
    Controller
}

func (mc ModulesController) registerRoutes(srv *Server) {
    modulesEndpoint := srv.Group("modules")
    {
        _ = modulesEndpoint
    }
}
