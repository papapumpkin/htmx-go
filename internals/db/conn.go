package db

import (
	"fmt"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	config "htmx-go/internals/core"
)

var (
	conn     *gorm.DB
	connOnce sync.Once
)

func Connect() *gorm.DB {
	connOnce.Do(func() {
		settings, err := config.LoadSettings()
		if err != nil {
			panic("Failed to load settings.")
		}
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable port=%s", settings.DBHost, settings.DBUser, settings.DBPassword, settings.DBName, settings.DBPort)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("Failed to connect to DB.")
		}
		conn = db
	})
	return conn
}
