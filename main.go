package main

import (
	"github.com/Encedeus/panel/config"
	"github.com/Encedeus/panel/controllers"
	"github.com/Encedeus/panel/module"
	"github.com/labstack/gommon/log"
)

func main() {
	config.InitConfig()
	db := config.InitDB()

	s := module.NewStore("/mnt/c/Projects/Encedeus/test/panel/modules")
	// go s.InitRPCServer()
	err := s.LoadAll()
	if err != nil {
		log.Fatalf("Load all error: %e", err)
	}

	controllers.StartDefaultServer(db, s)
}
