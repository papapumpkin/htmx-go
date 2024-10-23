package main

import (
	"ariga.io/atlas-provider-gorm/gormschema"

	"io"
	"log/slog"
	"os"

	"htmx-go/internals/db/models"
)

func main() {
	stmts, err := gormschema.New("postgres").Load(&models.User{})
	if err != nil {
		slog.Error("Failed to load gorm schema", "err", err.Error())
		os.Exit(1)
	}
	_, err = io.WriteString(os.Stdout, stmts)
	if err != nil {
		slog.Error("Failed to write gorm schema", "err" err.Error())
		os.Exit(1)
	}
}
