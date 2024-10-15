package fixture

import (
	"context"
	"embed"
	"text/template"
	"time"

	"github.com/1layar/universe/internal/email_service/model"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dbfixture"
)

//go:embed data/*
var sqlMigrations embed.FS

func Load(ctx context.Context, db *bun.DB) error {
	funcMap := template.FuncMap{
		"now": func() string {
			return time.Now().Format(time.RFC3339Nano)
		},
		"loadFile": func(path string) string {
			b, err := sqlMigrations.ReadFile(path)
			if err != nil {
				panic(err)
			}
			return string(b)
		},
	}

	// load model
	db.RegisterModel(
		(*model.Account)(nil),
		(*model.EmailEvent)(nil),
		(*model.EmailMessage)(nil),
		(*model.EmailTemplate)(nil),
	)

	f := dbfixture.New(db, dbfixture.WithTruncateTables(), dbfixture.WithTemplateFuncs(funcMap))

	err := f.Load(ctx, sqlMigrations, "data/fixture.yaml")

	return err
}
