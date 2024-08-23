package migrations

import (
	"embed"
	"os"
	"path/filepath"

	"github.com/uptrace/bun/migrate"
)

var (
	fullDirectory = filepath.Join(
		os.Getenv("PWD"),
		"internal", "infra", "db", "migrations",
		"sql",
	)
)

// A collection of migrations.
var Migrations = migrate.NewMigrations(migrate.WithMigrationsDirectory(fullDirectory))

//go:embed sql/*.sql
var sqlMigrations embed.FS

func init() {
	if err := Migrations.Discover(sqlMigrations); err != nil {
		panic(err)
	}
}
