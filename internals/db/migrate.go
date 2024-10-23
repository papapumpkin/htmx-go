package db

import (
	"io"
	"log/slog"
	"os"

	"ariga.io/atlas-provider-gorm/gormschema"
	"gorm.io/gorm"

	"htmx-go/internals/models"
)

func Migrate(db *gorm.DB, models ...interface{}) error {
	err := db.AutoMigrate(models...)
	slog.Info("Migration complete...")
	if err != nil {
		return err
	}
	return nil
}

func Migrate2() {
	stmts, err := gormschema.New("users").Load(&models.User{})
	if err != nil {
		slog.Error("Failed to load gorm schema", "err", err.Error())
	}
	_, err = io.WriteString(os.Stdout, stmts)
	if err != nil {
		slog.Error("Failed to write gorm schema", "err", err.Error())
		os.Exit(1)
	}
}
