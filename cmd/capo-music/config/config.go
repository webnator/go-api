package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

// Config is global object that holds all application level variables.
var Config appConfig

type appConfig struct {
	// the shared DB ORM object
	DB *gorm.DB
	// the error thrown be GORM when using DB ORM object
	DBErr error
	// the server port. Defaults to 8080
	ServerPort int `mapstructure:"server_port"`
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