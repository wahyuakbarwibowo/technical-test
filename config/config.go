package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DatabaseURL string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error loading config file", err)
	}

	cfg := &Config{
		DatabaseURL: viper.GetString("DATABASE_URL"),
	}

	return cfg, nil

}
