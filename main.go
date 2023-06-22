package main

import (
	"panel/config"
	"panel/server"
	"panel/service"
)

func main() {
	config.InitConfig()
	service.InitDB()
	server.ServerInit()
}
