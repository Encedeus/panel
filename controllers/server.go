package controllers

import (
    "github.com/Encedeus/panel/config"
    "github.com/Encedeus/panel/ent"
    "github.com/Encedeus/panel/middleware"
    "github.com/labstack/echo/v4"
)

func InitRouter(server *echo.Echo, db *ent.Client) {
    server.Use(middleware.JSONSyntaxMiddleware)

    AuthController{}.RegisterRoutes(server, db)

    roleEndpoint := server.Group("role")
    {
        roleEndpoint.Use(middleware.AccessJWTAuth)

        roleEndpoint.GET("/:id", getRole)
        roleEndpoint.POST("", handleCreateRole)
        roleEndpoint.PATCH("", handleUpdateRole)
        roleEndpoint.DELETE("/:id", handleDeleteRole)
    }

    userEndpoint := server.Group("user")
    {
        userEndpoint.Static("/pfp", config.Config.CDN.Directory)

        userEndpoint.Use(middleware.AccessJWTAuth)

        userEndpoint.GET("/:id", getUser)
        userEndpoint.POST("", handleCreateUser)
        userEndpoint.PUT("", setPfp)
        userEndpoint.PATCH("", handleUpdateUser)
        userEndpoint.DELETE("/:id", handleDeleteUser)
    }
    // controllers.InitControllers(server, db)
}
