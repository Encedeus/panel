storage_location = "../test/panel"

server {
  host = "0.0.0.0"
  port = 8080
}

database {
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  name     = "PanelDB"
  password = "root"
}

auth {
  jwt_secret_access  = "^¿Œõ%D:Ll$žö]Œ°áe7Óç]«"
  jwt_secret_refresh = "±ùk5fÈŽ?Åuunûs<­'™Ž’Hf"
}

cdn {
  dir = "./pfp"
}

modules {
  modules_dir                 = "modules"
  compiler_optimization_level = 2
}

skyhook {
  storage_location = "/etc/skyhook"
  default_port = 8000
  min_free_ram = 0
  min_free_logical_cores = 0
  min_free_disk = 0
}