package main

import (
	"panel/config"
	"panel/module"
	"panel/server"
	"panel/service"
)

func main() {
	config.InitConfig()
	service.InitDB()
    go module.Init()
    server.Init()
}
