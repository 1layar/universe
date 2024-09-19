package service

import (
	"github.com/1layar/universe/internal/ppob_service/model"
	"github.com/1layar/universe/internal/ppob_service/repo"
	"github.com/1layar/universe/pkg/shared/service"
)

type CategoryService struct {
	service.CrudBunService[model.ProductCategory]
}

func NewCategoryService(categoryRepo *repo.CategoryRepository) *CategoryService {
	return &CategoryService{
		CrudBunService: service.NewCrudBunService(categoryRepo),
	}
}
