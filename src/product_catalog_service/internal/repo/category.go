package repo

import (
	"github.com/1layar/universe/src/product_catalog_service/model"
	"github.com/1layar/universe/src/shared/repository"
	"github.com/uptrace/bun"
)

type CategoryRepository struct {
	repository.CrudRepository[model.Category]
}

func NewCategoryRepository(db *bun.DB) *CategoryRepository {
	return &CategoryRepository{
		CrudRepository: repository.NewCrudRepository[model.Category](db),
	}
}
