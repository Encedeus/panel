package main

import (
	"github.com/Encedeus/panel/config"
	"github.com/Encedeus/panel/controllers"
	"github.com/Encedeus/panel/module"
	"github.com/labstack/gommon/log"
	"path/filepath"
)

func main() {
	config.InitConfig()
	db := config.InitDB()

	// create module store and load modules
	s := module.NewStore(filepath.Join(config.Config.StorageLocationPath, config.Config.Modules.ModulesDirectory))
	err := s.LoadAll()
	if err != nil {
		log.Fatalf("Load all error: %v", err)
	}

	controllers.StartDefaultServer(db, s)
}
