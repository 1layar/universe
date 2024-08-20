package repo

import (
	"github.com/1layar/merasa/backend/src/product_catalog_service/model"
	"github.com/1layar/merasa/backend/src/shared/repository"
	"github.com/uptrace/bun"
)

type ProductRepository struct {
	repository.CrudRepository[model.Product]
}

func NewProductRepository(db *bun.DB) *ProductRepository {
	return &ProductRepository{
		CrudRepository: repository.NewCrudRepository[model.Product](db),
	}
}
