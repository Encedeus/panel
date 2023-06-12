package main

import (
	"Panel/config"
	"Panel/server"
	"Panel/service"
)

func main() {
	config.InitConfig()
	server.ServerInit()
	service.DBInit()
}
