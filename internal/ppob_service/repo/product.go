package repo

import (
	"github.com/1layar/universe/internal/ppob_service/model"
	"github.com/1layar/universe/pkg/shared/repository"
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
