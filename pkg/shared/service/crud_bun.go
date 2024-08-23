package service

import (
	"context"
	"fmt"

	"github.com/1layar/universe/pkg/shared/repository"
)

type CrudBunService[T any] struct {
	Repo repository.ICrudRepository[T]
}

func NewCrudBunService[T any](repo repository.ICrudRepository[T]) CrudBunService[T] {
	return CrudBunService[T]{Repo: repo}
}

func (s *CrudBunService[T]) Create(ctx context.Context, model *T) error {
	return s.Repo.Create(ctx, model)
}

func (s *CrudBunService[T]) GetByID(ctx context.Context, id int) (*T, error) {
	return s.Repo.GetByID(ctx, id)
}

func (s *CrudBunService[T]) GetByField(ctx context.Context, field string, value any, opt ...map[string]string) (*T, error) {
	return s.Repo.GetByField(ctx, field, value, opt...)
}

func (s *CrudBunService[T]) GetAll(ctx context.Context, opt ...map[string]string) ([]*T, error) {
	return s.Repo.GetAll(ctx, opt...)
}

func (s *CrudBunService[T]) GetAllPaginate(ctx context.Context, page int, limit int, include []string, opt ...map[string]any) ([]*T, int, error) {
	filterItemOption := []repository.Option{}

	for _, v := range opt {
		for k, v := range v {
			if v == "" {
				continue
			}

			strFal := fmt.Sprintf("%v", v)
			filterItemOption = append(filterItemOption, repository.WithEqFilter(k, strFal))
		}
	}

	for _, v := range include {
		filterItemOption = append(filterItemOption, repository.WithInclude(v))
	}

	filterItemOption = append(filterItemOption, repository.WithPaginate(page, limit))

	return s.Repo.GetAllPaginate(ctx, filterItemOption...)
}

func (s *CrudBunService[T]) Update(ctx context.Context, model *T) error {
	return s.Repo.Update(ctx, model)
}

func (s *CrudBunService[T]) Delete(ctx context.Context, id int) error {
	return s.Repo.Delete(ctx, id)
}
