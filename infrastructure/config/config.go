package config

import (
	"os"

	"github.com/spf13/viper"
)

type (
	AppConfig struct {
		Application Application
		MySQL       MySQL
	}

	Application struct {
		Name     string
		LevelLog string
		Server   Server
	}

	MySQL struct {
		Host string
	}

	Server struct {
		Port    string
		Timeout string
	}
)

func NewViperConfig() *viper.Viper {
	os.Setenv("ENVIRONMENT", "local")
	var env = os.Getenv("ENVIRONMENT")
	viper.SetConfigName("config." + env)
	viper.AddConfigPath("./infrastructure/config/env")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	viper.SetDefault("log_level", "debug")
	viper.SetDefault("log_format", "json")

	return viper.GetViper()
}

func LoadConfig() (*AppConfig, error) {
	var appConfig AppConfig

	viper := NewViperConfig()
	if err := viper.ReadInConfig(); err != nil {
		return &AppConfig{}, err
	}

	if err := viper.Unmarshal(&appConfig); err != nil {
		return &AppConfig{}, err
	}

	return &appConfig, nil
}