package service

import (
	"context"
	"database/sql"

	"github.com/1layar/universe/internal/ppob_service/app/appconfig"
	"github.com/1layar/universe/internal/ppob_service/model"
	"github.com/1layar/universe/internal/ppob_service/repo"
	"github.com/1layar/universe/pkg/shared/repository"
	"github.com/1layar/universe/pkg/shared/service"
	"github.com/uptrace/bun"
)

type ProductService struct {
	service.CrudBunService[model.Product]
	categoryS *CategoryService
	typeS     *ProductTypeService
	db        *bun.DB
}

func NewProductService(
	repo *repo.ProductRepository,
	catgoryS *CategoryService,
	typeS *ProductTypeService,
	db *bun.DB,
) *ProductService {
	return &ProductService{
		CrudBunService: service.NewCrudBunService(repo),
		db:             db,
		categoryS:      catgoryS,
		typeS:          typeS,
	}
}

func (s *ProductService) HasSku(ctx context.Context, sku string, opt ...map[string]string) (bool, error) {
	model, err := s.Repo.GetByField(ctx, "sku", sku, opt...)

	if err != nil && err != sql.ErrNoRows {
		return false, err
	}

	return model != nil, nil
}

func (s *ProductService) Create(ctx context.Context, product *model.Product) (int, error) {
	// star tx

	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return 0, err
	}

	err = s.Repo.Create(ctx, product, repository.WithTx(&tx))

	if err != nil {
		if err := tx.Rollback(); err != nil {
			return 0, err
		}
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return product.ID, err
}

func (s *ProductService) ImportIak(ctx context.Context, product []IakProduct) error {
	for _, p := range product {
		_, err := s.Repo.GetByField(ctx, "product_code", p.ProductCode)
		if err != nil && err != sql.ErrNoRows {
			return err
		} else if err == nil {
			continue
		}

		prodType, err := s.typeS.GetByField(ctx, "type_name", p.ProductType)
		if err != nil && err != sql.ErrNoRows {
			return err
		} else if prodType == nil {
			prodType = &model.ProductType{
				Name: p.ProductType,
			}

			err = s.typeS.Create(ctx, prodType)

			if err != nil {
				return err
			}
		}

		category, err := s.categoryS.GetByField(ctx, "category_name", p.ProductCategory)
		if err != nil && err != sql.ErrNoRows {
			return err
		} else if category == nil {
			category = &model.ProductCategory{
				Name: p.ProductCategory,
			}

			err = s.categoryS.Create(ctx, category)

			if err != nil {
				return err
			}
		}

		_, err = s.Create(ctx, &model.Product{
			Code:        p.ProductCode,
			Kind:        model.KindPrepaid,
			Description: p.ProductDescription,
			Details:     p.ProductDetails,
			Nominal:     p.ProductNominal,
			Price:       p.ProductPrice,
			IconURL:     p.IconURL,
			Status:      IakStatusToProductStatus(p.Status),
			CategoryId:  category.ID,
			TypeId:      prodType.ID,
		})

		if err != nil {
			return err
		}
	}
	return nil
}

func IakStatusToProductStatus(status appconfig.IakProductStatus) model.ProductStatus {
	switch status {
	case appconfig.Active:
		return model.StatusActive
	case appconfig.Inactive:
		return model.StatusInactive
	default:
		return model.StatusInactive
	}
}
