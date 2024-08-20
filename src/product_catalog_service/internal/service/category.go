package service

import (
	"github.com/1layar/merasa/backend/src/product_catalog_service/internal/repo"
	"github.com/1layar/merasa/backend/src/product_catalog_service/model"
	"github.com/1layar/merasa/backend/src/shared/service"
)

type CategoryService struct {
	service.CrudBunService[model.Category]
}

func NewCategoryService(categoryRepo *repo.CategoryRepository) *CategoryService {
	return &CategoryService{
		CrudBunService: service.NewCrudBunService(categoryRepo),
	}
}
