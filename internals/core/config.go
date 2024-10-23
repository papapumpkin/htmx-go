package config

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type Settings struct {
	DBName            string `envconfig:"DB_NAME"`
	DBHost            string `envconfig:"DB_HOST"`
	DBPort            string `envconfig:"DB_PORT"`
	DBUser            string `envconfig:"DB_USER"`
	DBPassword        string `envconfig:"DB_PASSWORD"`
	AppTimeoutSeconds int    `envconfig:"APP_TIMEOUT_SECONDS"`
}

var (
	settingsInstance *Settings
	settingsOnce     sync.Once
)

func LoadSettings() (*Settings, error) {
	var err error

	settingsOnce.Do(func() {
		var settings Settings
		err = envconfig.Process("", &settings)
		if err != nil {
			return
		}

		if settings.DBName == "" {
			settings.DBName = "database"
		}
		if settings.DBHost == "" {
			settings.DBHost = "localhost"
		}
		if settings.DBPort == "" {
			settings.DBPort = "5432"
		}
		if settings.DBUser == "" {
			settings.DBUser = "user"
		}
		if settings.DBPassword == "" {
			settings.DBPassword = "password"
		}
		if settings.AppTimeoutSeconds == 0 {
			settings.AppTimeoutSeconds = 10
		}
		settingsInstance = &settings
	})
	return settingsInstance, err
}
