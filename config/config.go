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
	RpkmConfirmStart        time.Time
	RpkmDayOneStart         time.Time
	RpkmDayTwoStart         time.Time
	FreshyNightConfirmStart time.Time
	FreshyNightConfirmEnd   time.Time
	FreshyNightStart        time.Time
	RpkmStart               time.Time
	RpkmEnd                 time.Time
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

	parsedRpkmConfirmStart, err := parseLocalTime("REG_RPKM_CONFIRM_START")
	if err != nil {
		return nil, err
	}
	fmt.Printf("Parsed RPKM confirm start time: %v\n", parsedRpkmConfirmStart)

	parsedRpkmDayOneStart, err := parseLocalTime("REG_RPKM_DAY_ONE_START")
	if err != nil {
		return nil, err
	}
	fmt.Printf("Parsed RPKM day one start time: %v\n", parsedRpkmDayOneStart)

	parsedRpkmDayTwoStart, err := parseLocalTime("REG_RPKM_DAY_TWO_START")
	if err != nil {
		return nil, err
	}
	fmt.Printf("Parsed RPKM day two start time: %v\n", parsedRpkmDayTwoStart)

	parsedFreshyNightConfirmStart, err := parseLocalTime("REG_FRESHY_NIGHT_CONFIRM_START")
	if err != nil {
		return nil, err
	}
	fmt.Printf("Parsed Freshy Night confirm start time: %v\n", parsedFreshyNightConfirmStart)

	parsedFreshyNightConfirmEnd, err := parseLocalTime("REG_FRESHY_NIGHT_CONFIRM_END")
	if err != nil {
		return nil, err
	}
	fmt.Printf("Parsed Freshy Night confirm end time: %v\n", parsedFreshyNightConfirmEnd)

	parsedFreshyNightStart, err := parseLocalTime("REG_FRESHY_NIGHT_START")
	if err != nil {
		return nil, err
	}
	fmt.Printf("Parsed Freshy Night start time: %v\n", parsedFreshyNightStart)

	parsedRpkmStartTime, err := parseLocalTime("REG_RPKM_START")
	if err != nil {
		return nil, err
	}
	fmt.Printf("Parsed RPKM start time: %v\n", parsedRpkmStartTime)

	parsedRpkmEndTime, err := parseLocalTime("REG_RPKM_END")
	if err != nil {
		return nil, err
	}
	fmt.Printf("Parsed RPKM end time: %v\n", parsedRpkmEndTime)

	regConfig := RegConfig{
		RpkmConfirmStart:        parsedRpkmConfirmStart,
		RpkmDayOneStart:         parsedRpkmDayOneStart,
		RpkmDayTwoStart:         parsedRpkmDayTwoStart,
		FreshyNightConfirmStart: parsedFreshyNightConfirmStart,
		FreshyNightConfirmEnd:   parsedFreshyNightConfirmEnd,
		FreshyNightStart:        parsedFreshyNightStart,
		RpkmStart:               parsedRpkmStartTime,
		RpkmEnd:                 parsedRpkmEndTime,
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

func parseLocalTime(envName string) (time.Time, error) {
	parsedTime, err := time.Parse(time.RFC3339, os.Getenv(envName))
	if err != nil {
		return time.Time{}, err
	}

	const gmtPlus7 = 7 * 60 * 60
	gmtPlus7Location := time.FixedZone("GMT+7", gmtPlus7)

	return time.Date(
		parsedTime.Year(), parsedTime.Month(), parsedTime.Day(),
		parsedTime.Hour(), parsedTime.Minute(), parsedTime.Second(),
		parsedTime.Nanosecond(), gmtPlus7Location), nil
}
