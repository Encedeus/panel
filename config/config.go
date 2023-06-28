package config

import (
	"errors"
	"fmt"
	"github.com/hashicorp/hcl/v2/hclsimple"
	_ "github.com/joho/godotenv"
	"log"
	"os"
)

// const DefaultLocation = "/etc/encedeus"
const DefaultLocation = "./"

type Configuration struct {
	Server ServerConfiguration   `hcl:"server,block"`
	DB     DatabaseConfiguration `hcl:"database,block"`
	Auth   AuthConfiguration     `hcl:"auth,block"`
	CDN    CDNConfiguration      `hcl:"cdn,block"`
}

type ServerConfiguration struct {
	Host string `hcl:"host"`
	Port int    `hcl:"port"`
}

type DatabaseConfiguration struct {
	Host     string `hcl:"host"`
	Port     int    `hcl:"port"`
	User     string `hcl:"user"`
	DBName   string `hcl:"name"`
	Password string `hcl:"password"`
}

type AuthConfiguration struct {
	JWTSecretAccess  string `hcl:"jwt_secret_access"`
	JWTSecretRefresh string `hcl:"jwt_secret_refresh"`
}

type CDNConfiguration struct {
	Directory string `hcl:"dir"`
}

func (s *ServerConfiguration) URI() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

var Config Configuration

func InitConfig() {
	if _, err := os.Stat(fmt.Sprintf("%s/config.hcl", DefaultLocation)); errors.Is(err, os.ErrNotExist) {
		log.Fatal("Config file does not exist")
	}

	err := hclsimple.DecodeFile(fmt.Sprintf("%s/config.hcl", DefaultLocation), nil, &Config)
	if err != nil {
		log.Fatalf("Failed to load configuration file: %v", err)
	}
}
