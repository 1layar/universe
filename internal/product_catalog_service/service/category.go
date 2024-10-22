package service

import (
	"github.com/1layar/universe/internal/product_catalog_service/model"
	"github.com/1layar/universe/internal/product_catalog_service/repo"
	"github.com/1layar/universe/pkg/shared/service"
)

type CategoryService struct {
	service.CrudBunService[model.Category]
}

func NewCategoryService(categoryRepo *repo.CategoryRepository) *CategoryService {
	return &CategoryService{
		CrudBunService: service.NewCrudBunService(categoryRepo),
	}
}
