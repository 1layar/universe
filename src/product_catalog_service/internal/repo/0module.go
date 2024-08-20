package repo

import (
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("repo", fx.Provide(NewProductRepository, NewCategoryRepository, NewProductCategoryRepository))
}
