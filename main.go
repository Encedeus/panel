package main

import (
    "fmt"
    "github.com/Encedeus/panel/config"
    "github.com/Encedeus/panel/controllers"
    "github.com/Encedeus/panel/module"
    "time"
)

func main() {
    config.InitConfig()
    db := config.InitDB()

    s := module.NewStore("/home/optimuseprime/Projects/Encedeus/test/panel/modules")
    go s.InitRPCServer()
    // m, err := s.LoadOne("test_module_2.ema")
    // time.Sleep(5 * time.Second)
    err := s.LoadAll()
    // go s.LoadOne("./daemon_test.ema")
    fmt.Println(err)
    go func() {
        for {
            if s.Modules[0] != nil {
                fmt.Println(s.Modules[0].Store.HasRegisteredCommand("test_cmd"))
            }
            time.Sleep(1 * time.Second)
        }
    }()
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
