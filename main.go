package main

import (
    "github.com/Encedeus/panel/config"
    // "panel/module"
    "github.com/Encedeus/panel/server"
    "github.com/Encedeus/panel/services"
)

func main() {
    config.InitConfig()
    services.InitDB()
    // go module.Init()
    server.Init()
}
