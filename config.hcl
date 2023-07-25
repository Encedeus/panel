server {
  host = "localhost"
  port = 8080
}

database {
  host = "localhost"
  port = 5432
  user = "postgres"
  name = "PanelDB"
  password = "root"
}

auth {
  jwt_secret_access = "i95FOB61kCoJjSt2SBSifhtwMHQ7Nasi"
  jwt_secret_refresh = "SjSt2fhtwi7BiFOS95MHQiasB61kCoJN"
}

cdn {
  dir = "./pfp"
}