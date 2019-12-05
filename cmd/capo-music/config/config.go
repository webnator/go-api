package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config is global object that holds all application level variables.
var Config appConfig

// LoadConfig loads config from files
func LoadConfig() error {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	Config = appConfig{
		ServerPort: 8081,
		DBConfig: dBConfig{
			Host:     getEnv("DB_HOST"),
			Port:     getEnv("DB_PORT"),
			DBName:   getEnv("DB_DATABASE"),
			User:     getEnv("DB_USER"),
			Password: getEnv("DB_PASSWORD"),
		},
		Collections: map[string]string{
			"songs":      "songs",
			"view_count": "view_count",
		},
	}
	return nil
}

type dBConfig struct {
	Host     string
	Port     string
	DBName   string
	User     string
	Password string
}
type appConfig struct {
	// the server port. Defaults to 8080
	ServerPort  int
	DBConfig    dBConfig
	Collections map[string]string
}

func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	panic(fmt.Errorf("Config var not provided: %s", key))
}
