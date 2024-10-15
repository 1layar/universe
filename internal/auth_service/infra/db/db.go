package db

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"

	"github.com/1layar/universe/internal/auth_service/app/appconfig"
)

func New(config *appconfig.Config) *bun.DB {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(config.DatabaseUrl)))
	db := bun.NewDB(sqldb, pgdialect.New())
	return db
}
