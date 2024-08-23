package repository

import (
	"context"
	"fmt"

	"github.com/uptrace/bun"
)

type ICrudRepository[T any] interface {
	Create(ctx context.Context, model *T, options ...Option) error
	GetByID(ctx context.Context, id int) (*T, error)
	GetByField(ctx context.Context, field string, value any, opt ...map[string]string) (*T, error)
	GetAll(ctx context.Context, opt ...map[string]string) ([]*T, error)
	Update(ctx context.Context, model *T) error
	GetAllPaginate(ctx context.Context, options ...Option) ([]*T, int, error)
	Delete(ctx context.Context, id int) error
}

type CrudRepository[T any] struct {
	db *bun.DB
}

func NewCrudRepository[T any](db *bun.DB) CrudRepository[T] {
	return CrudRepository[T]{db: db}
}

func (r *CrudRepository[T]) Create(ctx context.Context, model *T, options ...Option) error {
	var err error
	param := parseOptions(options...)
	qb := param.InsertQuery(r.db, model)

	param.InjectReturning(qb)

	_, err = qb.Exec(ctx)
	if err != nil {
		return err
	}

	return err
}

func (r *CrudRepository[T]) GetByID(ctx context.Context, id int) (*T, error) {
	product := new(T)
	err := r.db.NewSelect().Model(product).Where("id = ?", id).Scan(ctx)
	return product, err
}

func (r *CrudRepository[T]) GetByField(ctx context.Context, field string, value any, opt ...map[string]string) (*T, error) {
	product := new(T)
	model := r.db.NewSelect().Model(product)

	model.Where(fmt.Sprintf("%s = ?", field), value)

	for _, v := range opt {
		for k, v := range v {
			model.Where(fmt.Sprintf("%s <> ?", k), v)
		}
	}

	if err := model.Limit(1).Scan(ctx); err != nil {
		return nil, err
	}

	return product, nil
}

func (r *CrudRepository[T]) GetAll(ctx context.Context, opt ...map[string]string) ([]*T, error) {
	products := []*T{}
	model := r.db.NewSelect().Model(&products)

	for _, v := range opt {
		for k, v := range v {
			model.Where(fmt.Sprintf("%s = ?", k), v)
		}
	}

	if err := model.Scan(ctx); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *CrudRepository[T]) GetAllPaginate(ctx context.Context, options ...Option) ([]*T, int, error) {
	products := []*T{}
	param := parseOptions(options...)
	model := r.db.NewSelect().Model(&products)

	param.InjectFilter(model)
	param.InjectSort(model)
	param.InjectInclude(model)

	total, err := model.ScanAndCount(ctx)

	if err != nil {
		return nil, 0, err
	}

	param.InjectPagination(model)

	err = model.Scan(ctx)

	if err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

func (r *CrudRepository[T]) Delete(ctx context.Context, id int) error {
	_, err := r.db.NewDelete().Model((*T)(nil)).Where("id = ?", id).Exec(ctx)
	return err
}

func (r *CrudRepository[T]) Update(ctx context.Context, product *T) error {
	_, err := r.db.NewUpdate().Model(product).WherePK().Exec(ctx)
	return err
}
