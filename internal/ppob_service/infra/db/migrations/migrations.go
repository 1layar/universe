package migrations

import (
	"embed"
	"os"
	"path/filepath"

	"github.com/1layar/universe/pkg/logger"
	"github.com/uptrace/bun/migrate"
)

var (
	fullDirectory = filepath.Join(
		os.Getenv("PWD"),
		"infra", "db", "migrations",
		"sql",
	)
)

// A collection of migrations.
var Migrations = migrate.NewMigrations(migrate.WithMigrationsDirectory(fullDirectory))

//go:embed sql/*.sql
var sqlMigrations embed.FS

func init() {
	logger := logger.GetLogger()
	if err := Migrations.Discover(sqlMigrations); err != nil {
		logger.Panic(err)
	}
}
