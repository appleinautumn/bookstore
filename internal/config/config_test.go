package config

import (
	"os"
	"testing"
)

func TestNew(t *testing.T) {

	t.Run("APP_NAME not set", func(t *testing.T) {
		_, err := New()
		if err == nil {
			t.Errorf(`Value: %v`, err)
		}
	})

	t.Run("APP_ENV not set", func(t *testing.T) {
		os.Setenv("APP_NAME", "super-app")

		_, err := New()
		if err == nil {
			t.Errorf(`Value: %v`, err)
		}
	})

	t.Run("APP_VERSION not set", func(t *testing.T) {
		os.Setenv("APP_NAME", "super-app")
		os.Setenv("APP_ENV", "super-env")

		_, err := New()
		if err == nil {
			t.Errorf(`Value: %v`, err)
		}
	})

	t.Run("DB_URL not set", func(t *testing.T) {
		os.Setenv("APP_NAME", "super-app")
		os.Setenv("APP_ENV", "super-env")
		os.Setenv("APP_VERSION", "12.34.5")

		_, err := New()
		if err == nil {
			t.Errorf(`Value: %v`, err)
		}
	})

	appName := "super-app"
	appEnv := "super-env"
	appPort := "9000"
	appVersion := "12.34.5"
	dbUrl := "http://rds.com/aaa"

	os.Setenv("APP_NAME", appName)
	os.Setenv("APP_ENV", appEnv)
	os.Setenv("APP_PORT", appPort)
	os.Setenv("APP_VERSION", appVersion)
	os.Setenv("DB_URL", dbUrl)

	cfg, _ := New()

	tests := []struct {
		name string
		env  string
		want string
	}{
		{name: "APP_NAME", env: cfg.AppName, want: appName},
		{name: "APP_ENV", env: cfg.AppEnv, want: appEnv},
		{name: "APP_VERSION", env: cfg.AppVersion, want: appVersion},
		{name: "DB_URL", env: cfg.DbUrl, want: dbUrl},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.env != tt.want {
				t.Errorf(`Value: %s. want: %s`, tt.env, tt.want)
			}
		})
	}
}

func TestGetEnv(t *testing.T) {

	a := "aaa"
	defaultValue := "default-value"

	os.Setenv("ENV_A", a)

	tests := []struct {
		name string
		env  string
		def  string
		want string
	}{
		{name: "ENV_A", env: "ENV_A", def: defaultValue, want: a},
		{name: "DONTEXIST", env: "DONTEXIST", def: defaultValue, want: defaultValue},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			res := getEnv(tt.env, tt.def)

			if res != tt.want {
				t.Errorf(`Value: %s. want: %s`, res, tt.want)
			}
		})
	}
}
