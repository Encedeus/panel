package main

import (
	"github.com/Encedeus/panel/config"
	// "panel/module"
	"github.com/Encedeus/panel/server"
	"github.com/Encedeus/panel/service"
)

func main() {
	config.InitConfig()
	service.InitDB()
    // go module.Init()
    server.Init()
}
