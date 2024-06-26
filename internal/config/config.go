package config

import (
	"errors"
	"os"
)

type Config struct {
	AppEnv     string
	AppName    string
	AppPort    string
	AppVersion string
	DbUrl      string
	LogLevel   string
}

var configImpl *Config

func New() (*Config, error) {

	appEnv := getEnv("APP_ENV", "")
	if appEnv == "" {
		return nil, errors.New("APP_ENV env is required")
	}

	appName := getEnv("APP_NAME", "")
	if appName == "" {
		return nil, errors.New("APP_NAME env is required")
	}

	appPort := getEnv("APP_PORT", "")
	if appPort == "" {
		return nil, errors.New("APP_PORT env is required")
	}

	appVersion := getEnv("APP_VERSION", "")
	if appVersion == "" {
		return nil, errors.New("APP_VERSION env is required")
	}

	dbUrl := getEnv("DB_URL", "")
	if dbUrl == "" {
		return nil, errors.New("DB_URL env is required")
	}

	logLevel := getEnv("LOG_LEVEL", "info")

	configImpl = &Config{
		AppEnv:     appEnv,
		AppName:    appName,
		AppPort:    appPort,
		AppVersion: appVersion,
		DbUrl:      dbUrl,
		LogLevel:   logLevel,
	}
	return configImpl, nil
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
