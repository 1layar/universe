package service

import (
	"github.com/1layar/universe/internal/ppob_service/model"
	"github.com/1layar/universe/internal/ppob_service/repo"
	"github.com/1layar/universe/pkg/shared/service"
)

type ProductTypeService struct {
	service.CrudBunService[model.ProductType]
}

func NewProductTypeService(categoryRepo *repo.ProductTypeRepository) *ProductTypeService {
	return &ProductTypeService{
		CrudBunService: service.NewCrudBunService(categoryRepo),
	}
}
