package controllers

import (
	"github.com/Encedeus/panel/config"
	"github.com/Encedeus/panel/ent"
	encMiddleware "github.com/Encedeus/panel/middleware"
	"github.com/Encedeus/panel/module"
	"github.com/labstack/echo/v4"
)

type Controller interface {
	registerRoutes(*Server)
}

func registerControllerRoutes(srv *Server, cs ...Controller) {
	for _, c := range cs {
		c.registerRoutes(srv)
	}
}

type Server struct {
	*echo.Echo
	DB          *ent.Client
	ModuleStore *module.Store
}

func NewEmptyServer(db *ent.Client, store *module.Store) *Server {
	srv := &Server{
		Echo:        echo.New(),
		DB:          db,
		ModuleStore: store,
	}

	return srv
}

func WrapServerWithDefaults(srv *Server, _ *ent.Client) {
	srv.Use(encMiddleware.JSONSyntaxMiddleware)
	/*srv.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods: []string{"GET", "POST", "DELETE", "PUT", "PATCH", "HEAD"},
		AllowHeaders: []string{"Accept", "Content-Type", "Authorization"},
		//AllowOrigins:     []string{"http://localhost:5173"},
		AllowCredentials: true,
	}))*/
	srv.Use(encMiddleware.CORSMiddleware)

	InitRouter(srv)
}

func NewDefaultServer(db *ent.Client, store *module.Store) *Server {
	srv := NewEmptyServer(db, store)
	WrapServerWithDefaults(srv, db)

	return srv
}

func InitRouter(srv *Server) {
	registerControllerRoutes(srv,
		AuthController{},
		RoleController{},
		UserController{},
		APIKeyController{},
		ModulesController{
			ModuleStore: srv.ModuleStore,
		},
		NodesController{},
		ServersController{},
	)
}

func StartServer(srv *Server) {
	srv.Logger.Fatal(srv.Start(config.Config.Server.URI()))
}

func StartDefaultServer(db *ent.Client, store *module.Store) {
	StartServer(NewDefaultServer(db, store))
}
