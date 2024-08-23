package repo

import (
	"github.com/1layar/universe/src/product_catalog_service/model"
	"github.com/1layar/universe/src/shared/repository"
	"github.com/uptrace/bun"
)

type ProductCategoryRepository struct {
	repository.CrudRepository[model.ProductCategoryRelation]
}

func NewProductCategoryRepository(db *bun.DB) *ProductCategoryRepository {
	return &ProductCategoryRepository{
		CrudRepository: repository.NewCrudRepository[model.ProductCategoryRelation](db),
	}
}
