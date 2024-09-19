package repo

import (
	"github.com/1layar/universe/internal/ppob_service/model"
	"github.com/1layar/universe/pkg/shared/repository"
	"github.com/uptrace/bun"
)

type ProductTypeRepository struct {
	repository.CrudRepository[model.ProductType]
}

func NewProductTypeRepository(db *bun.DB) *ProductTypeRepository {
	return &ProductTypeRepository{
		CrudRepository: repository.NewCrudRepository[model.ProductType](db),
	}
}
