package db

import (
	"strings"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
	"github.com/urfave/cli/v2"
	"go.uber.org/fx"

	cliapp "github.com/1layar/universe/internal/product_catalog_service/cmd/app/cli"
	"github.com/1layar/universe/internal/product_catalog_service/infra/db/fixture"
	"github.com/1layar/universe/internal/product_catalog_service/infra/db/migrations"
	"github.com/rs/zerolog/log"
)

type dbCommandDeps struct {
	fx.In

	DB *bun.DB
}

func Command() *cli.Command {
	migrator := func() *migrations.Migrator {
		var deps dbCommandDeps
		cliapp.Start(fx.Populate(&deps))
		println("prepare migrations...")
		m := migrate.NewMigrator(deps.DB, migrations.Migrations)
		println("init migrations...")
		return migrations.NewMigrator(m)
	}

	return &cli.Command{
		Name:  "db",
		Usage: "manage database migrations",
		Subcommands: []*cli.Command{
			{
				Name:  "init",
				Usage: "create migration tables",
				Action: func(c *cli.Context) error {
					return migrator().Init(c.Context)
				},
			},
			{
				Name:  "seed",
				Usage: "seed database",
				Action: func(c *cli.Context) error {
					var deps dbCommandDeps
					cliapp.Start(fx.Populate(&deps))
					println("prepare fixtures...")
					return fixture.Load(c.Context, deps.DB)
				},
			},
			{
				Name:  "migrate",
				Usage: "migrate database",
				Action: func(c *cli.Context) error {
					log.Info().Msg("migrate database...")
					instance := migrator()
					err := instance.Init(c.Context)

					if err != nil {
						return err
					}

					err = instance.Migrate(c.Context)

					if err != nil {
						if err := instance.Rollback(c.Context); err != nil {
							return err
						}
						return err
					}
					log.Info().Msg("migrate database done")
					return nil
				},
			},
			{
				Name:  "rollback",
				Usage: "rollback the last migration group",
				Action: func(c *cli.Context) error {
					return migrator().Rollback(c.Context)
				},
			},
			{
				Name:  "lock",
				Usage: "lock migrations",
				Action: func(c *cli.Context) error {
					return migrator().Lock(c.Context)
				},
			},
			{
				Name:  "unlock",
				Usage: "unlock migrations",
				Action: func(c *cli.Context) error {
					return migrator().Unlock(c.Context)
				},
			},
			{
				Name:  "create_go",
				Usage: "create Go migration",
				Action: func(c *cli.Context) error {
					name := strings.Join(c.Args().Slice(), "_")
					return migrator().CreateGoMigration(c.Context, name)
				},
			},
			{
				Name:  "create_sql",
				Usage: "create up and down SQL migrations",
				Action: func(c *cli.Context) error {
					name := strings.Join(c.Args().Slice(), "_")
					return migrator().CreateSQLMigrations(c.Context, name)
				},
			},
			{
				Name:  "status",
				Usage: "print migrations status",
				Action: func(c *cli.Context) error {
					return migrator().Status(c.Context)
				},
			},
			{
				Name:  "mark_applied",
				Usage: "mark migrations as applied without actually running them",
				Action: func(c *cli.Context) error {
					return migrator().MarkApplied(c.Context)
				},
			},
		},
	}
}
