package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port        string
	Env         string
	ServiceName string
}

type ImageConfig struct {
	MaxFileSizeMb int
	CropWidth     int
	CropHeight    int
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

type DbConfig struct {
	Url string
}

type TracerConfig struct {
	Endpoint string
}

type Config struct {
	App    AppConfig
	Img    ImageConfig
	Svc    ServiceConfig
	Cors   CorsConfig
	Db     DbConfig
	Tracer TracerConfig
}

func LoadConfig() (*Config, error) {
	if os.Getenv("APP_ENV") == "" {
		err := godotenv.Load(".env")
		if err != nil {
			return nil, err
		}
	}

	appConfig := AppConfig{
		Port:        os.Getenv("APP_PORT"),
		Env:         os.Getenv("APP_ENV"),
		ServiceName: os.Getenv("APP_SERVICE_NAME"),
	}

	maxFileSizeMb, err := strconv.ParseInt(os.Getenv("IMG_MAX_FILE_SIZE_MB"), 10, 64)
	if err != nil {
		return nil, err
	}
	cropWidth, err := strconv.ParseInt(os.Getenv("IMG_CROP_WIDTH"), 10, 64)
	if err != nil {
		return nil, err
	}
	cropHeight, err := strconv.ParseInt(os.Getenv("IMG_CROP_HEIGHT"), 10, 64)
	if err != nil {
		return nil, err
	}
	imageConfig := ImageConfig{
		MaxFileSizeMb: int(maxFileSizeMb),
		CropWidth:     int(cropWidth),
		CropHeight:    int(cropHeight),
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

	DbConfig := DbConfig{
		Url: os.Getenv("DB_URL"),
	}

	tracerConfig := TracerConfig{
		Endpoint: os.Getenv("TRACER_ENDPOINT"),
	}

	return &Config{
		App:    appConfig,
		Img:    imageConfig,
		Svc:    serviceConfig,
		Cors:   corsConfig,
		Db:     DbConfig,
		Tracer: tracerConfig,
	}, nil
}

func (ac *AppConfig) IsDevelopment() bool {
	return ac.Env == "development"
}
