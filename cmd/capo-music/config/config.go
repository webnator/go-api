package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config is global object that holds all application level variables.
var Config appConfig

type DBConfig struct {
	DBName string `mapstructure:"database"`
	Host   string `mapstructure:"host"`
	Port   string `mapstructure:"port"`
}
type appConfig struct {
	// the server port. Defaults to 8080
	ServerPort int      `mapstructure:"server_port"`
	DBConfig   DBConfig `mapstructure:"mongo"`
}

// LoadConfig loads config from files
func LoadConfig() error {
	v := viper.New()
	v.SetConfigName(".env")
	v.AddConfigPath(".")
	v.SetConfigType("json")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read the configuration file: %s", err)
	}
	return v.Unmarshal(&Config)
}
