package config

import (
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Configuration struct {
	Server ServerConfiguration
	Db     DatabaseConfiguration
}

type ServerConfiguration struct {
	Host string
	Port int
	URI  string
}
type DatabaseConfiguration struct {
	Host     string
	Port     int
	User     string
	DbName   string
	Password string
}

var Config Configuration

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Config = Configuration{
		ServerConfiguration{
			os.Getenv("HOST"),
			getIntEnv("PORT"),
			fmt.Sprintf("%s:%d", os.Getenv("HOST"), getIntEnv("PORT")),
		},
		DatabaseConfiguration{
			os.Getenv("DB_HOST"),
			getIntEnv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PASS"),
		},
	}
}

func getIntEnv(key string) int {
	value, err := strconv.Atoi(os.Getenv(key))

	if err != nil {
		log.Fatal("Failed to convert a environment variable to integer")
	}

	return value
}
