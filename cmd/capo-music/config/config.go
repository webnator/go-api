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
		ServerPort: getEnvDefault("PORT", "8081"),
		DBConfig: dBConfig{
			URI:    getEnv("MONGODB_URI"),
			DBName: getEnv("DB_DATABASE"),
		},
		Auth: authConfig{
			User:     getEnv("USER_NAME"),
			Password: getEnv("USER_PW"),
		},
		Collections: map[string]string{
			"songs":      "songs",
			"view_count": "view_count",
		},
	}
	return nil
}

type dBConfig struct {
	URI    string
	DBName string
}
type appConfig struct {
	// the server port. Defaults to 8080
	ServerPort  string
	DBConfig    dBConfig
	Auth        authConfig
	Collections map[string]string
}

type authConfig struct {
	User     string
	Password string
}

func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	panic(fmt.Errorf("Config var not provided: %s", key))
}

func getEnvDefault(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
