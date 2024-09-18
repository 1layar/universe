package service

import (
	"context"
	"database/sql"

	"github.com/1layar/universe/internal/ppob_service/repo"
	"github.com/1layar/universe/internal/product_catalog_service/model"
	"github.com/1layar/universe/pkg/shared/repository"
	"github.com/1layar/universe/pkg/shared/service"
	"github.com/uptrace/bun"
)

type ProductService struct {
	service.CrudBunService[model.Product]
	productCategory *repo.ProductCategoryRepository
	db              *bun.DB
}

func NewProductService(repo *repo.ProductRepository, pr *repo.ProductCategoryRepository, db *bun.DB) *ProductService {
	return &ProductService{
		CrudBunService:  service.NewCrudBunService(repo),
		productCategory: pr,
		db:              db,
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
	product.Categories = nil

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

		_, err = s.Create(ctx, &model.Product{
			Code:        p.ProductCode,
			Name:        p.ProductDetails,
			Description: p.ProductDescription,
			Price:       p.ProductPrice,
		})

		if err != nil {
			return err
		}
	}

	return nil
}
