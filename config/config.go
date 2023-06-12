package config

import (
	"fmt"
	"github.com/hashicorp/hcl/v2/hclsimple"
	_ "github.com/joho/godotenv"
	"log"
	"os"
)

const DefaultLocation = "/etc/encedeus"

type Configuration struct {
	Server ServerConfiguration   `hcl:"server"`
	DB     DatabaseConfiguration `hcl:"database"`
}

type ServerConfiguration struct {
	Host string `hcl:"host"`
	Port int    `hcl:"port"`
}

func (s *ServerConfiguration) URI() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

type DatabaseConfiguration struct {
	Host     string `hcl:"host"`
	Port     int    `hcl:"port"`
	User     string `hcl:"user"`
	DBName   string `hcl:"dbname"`
	Password string `hcl:"password"`
}

var Config Configuration

func InitConfig() {
	_, err := os.Create(fmt.Sprintf("%s/config.hcl", DefaultLocation))
	if err != nil && !os.IsExist(err) {
		log.Fatalf("Failed creating configuration file: %v", err)
	}

	err = hclsimple.DecodeFile(fmt.Sprintf("%s/config.hcl", DefaultLocation), nil, &Config)
	if err != nil {
		log.Fatalf("Failed to load configuration file: %v", err)
	}
}
