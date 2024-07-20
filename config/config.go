package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

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

type RegConfig struct {
	RpkmStart      time.Time
	FirstdateStart time.Time
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
	Reg    RegConfig
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

	parsedRpkmTime, err := time.Parse(time.RFC3339, os.Getenv("REG_RPKM_START"))
	if err != nil {
		return nil, err
	}
	parsedFirstdateTime, err := time.Parse(time.RFC3339, os.Getenv("REG_FIRSTDATE_START"))
	if err != nil {
		return nil, err
	}

	const gmtPlus7 = 7 * 60 * 60
	gmtPlus7Location := time.FixedZone("GMT+7", gmtPlus7)

	localRpkmTime := time.Date(
		parsedRpkmTime.Year(), parsedRpkmTime.Month(), parsedRpkmTime.Day(),
		parsedRpkmTime.Hour(), parsedRpkmTime.Minute(), parsedRpkmTime.Second(),
		parsedRpkmTime.Nanosecond(), gmtPlus7Location)
	fmt.Println("Local RPKM time (GMT+7):", localRpkmTime)

	localFirstdateTime := time.Date(
		parsedFirstdateTime.Year(), parsedFirstdateTime.Month(), parsedFirstdateTime.Day(),
		parsedFirstdateTime.Hour(), parsedFirstdateTime.Minute(), parsedFirstdateTime.Second(),
		parsedFirstdateTime.Nanosecond(), gmtPlus7Location)
	fmt.Println("Local Firstdate time (GMT+7):", localFirstdateTime)

	regConfig := RegConfig{
		RpkmStart:      localRpkmTime,
		FirstdateStart: localFirstdateTime,
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
		Reg:    regConfig,
		Svc:    serviceConfig,
		Cors:   corsConfig,
		Db:     DbConfig,
		Tracer: tracerConfig,
	}, nil
}

func (ac *AppConfig) IsDevelopment() bool {
	return ac.Env == "development"
}
