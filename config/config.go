package config

import (
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port string
	Env  string
}

type Config struct {
	App AppConfig
}

func LoadConfig() (*Config, error) {
	if os.Getenv("APP_ENV") == "" {
		err := godotenv.Load(".env")
		if err != nil {
			return nil, err
		}
	}

	appConfig := AppConfig{
		Port: os.Getenv("APP_PORT"),
		Env:  os.Getenv("APP_ENV"),
	}

	return &Config{
		App: appConfig,
	}, nil
}

func (ac *AppConfig) IsDevelopment() bool {
	return ac.Env == "development"
}
