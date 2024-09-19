package repo

import (
	"github.com/1layar/universe/internal/ppob_service/model"
	"github.com/1layar/universe/pkg/shared/repository"
	"github.com/uptrace/bun"
)

type CategoryRepository struct {
	repository.CrudRepository[model.ProductCategory]
}

func NewCategoryRepository(db *bun.DB) *CategoryRepository {
	return &CategoryRepository{
		CrudRepository: repository.NewCrudRepository[model.ProductCategory](db),
	}
}
