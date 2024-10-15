package db

import (
	"database/sql"

	"github.com/1layar/universe/internal/product_catalog_service/app/appconfig"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"

	"github.com/1layar/universe/internal/product_catalog_service/model"
)

func New(config *appconfig.Config) *bun.DB {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(config.DatabaseUrl)))
	db := bun.NewDB(sqldb, pgdialect.New())

	db.RegisterModel(
		(*model.ProductCategoryRelation)(nil),
	)

	return db
}
