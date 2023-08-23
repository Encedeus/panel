package main

import (
    "github.com/Encedeus/panel/config"
    "github.com/Encedeus/panel/controllers"
)

func main() {
    config.InitConfig()
    db := config.InitDB()
    // go module.Init()
    controllers.StartDefaultServer(db)
}
