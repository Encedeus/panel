package main

import (
    "fmt"
    "github.com/Encedeus/panel/config"
    "github.com/Encedeus/panel/controllers"
    "github.com/Encedeus/panel/module"
)

func main() {
    config.InitConfig()
    db := config.InitDB()

    s := module.NewStore("/home/optimuseprime/Projects/Encedeus/test/panel/modules")
    // m, err := s.LoadOne("test_module_2.ema")
    err := s.LoadAll()
    // go s.LoadOne("./daemon_test.ema")
    fmt.Println(err)
    // fmt.Printf("%+v\n", m)

    /*    f, err := os.Open("/home/optimuseprime/Projects/Encedeus/test/panel/modules/cache/test_module/test_module.wasm")
          if err != nil {
              log.Errorf("%e", err)
          }
          err = module.ExecuteModuleWasmFile(f, 8082)
          if err != nil {
              log.Errorf("%e", err)
          }*/

    controllers.StartDefaultServer(db)
}
