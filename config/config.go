package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port          string
	Env           string
	MaxFileSizeMb int
}

type ServiceConfig struct {
	Auth    string
	Backend string
	CheckIn string
	Store   string
}

type CorsConfig struct {
	AllowOrigins string
}

type Config struct {
	App  AppConfig
	Svc  ServiceConfig
	Cors CorsConfig
}

func LoadConfig() (*Config, error) {
	if os.Getenv("APP_ENV") == "" {
		err := godotenv.Load(".env")
		if err != nil {
			return nil, err
		}
	}

	maxFileSizeMb, err := strconv.ParseInt(os.Getenv("APP_MAX_FILE_SIZE_MB"), 10, 64)
	if err != nil {
		return nil, err
	}

	appConfig := AppConfig{
		Port:          os.Getenv("APP_PORT"),
		Env:           os.Getenv("APP_ENV"),
		MaxFileSizeMb: int(maxFileSizeMb),
	}

	serviceConfig := ServiceConfig{
		Auth:    os.Getenv("SERVICE_AUTH"),
		Backend: os.Getenv("SERVICE_BACKEND"),
		CheckIn: os.Getenv("SERVICE_CHECKIN"),
		Store:   os.Getenv("SERVICE_STORE"),
	}

	corsConfig := CorsConfig{
		AllowOrigins: os.Getenv("CORS_ORIGINS"),
	}

	return &Config{
		App:  appConfig,
		Svc:  serviceConfig,
		Cors: corsConfig,
	}, nil
}

func (ac *AppConfig) IsDevelopment() bool {
	return ac.Env == "development"
}
