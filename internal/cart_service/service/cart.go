package service

import (
	"context"

	"github.com/1layar/universe/internal/cart_service/model"
	"github.com/1layar/universe/internal/cart_service/repo"
)

type CartService struct {
	repo *repo.CartRepository
}

func NewCartService(repo *repo.CartRepository) *CartService {
	return &CartService{
		repo: repo,
	}
}

func (r *CartService) AddItem(ctx context.Context, sessionId string, productId int, quantity int, source string) error {
	return r.repo.AddItem(ctx, sessionId, productId, quantity, source)
}
func (r *CartService) RemoveItem(ctx context.Context, sessionId string, productId int, source string) error {
	return r.repo.RemoveItem(ctx, sessionId, productId, source)
}
func (r *CartService) UpdateQuantity(ctx context.Context, sessionId string, productId int, quantity int, source string) error {
	return r.repo.UpdateQuantity(ctx, sessionId, productId, quantity, source)
}
func (r *CartService) GetCart(ctx context.Context, sessionId string) ([]model.CartItem, error) {
	return r.repo.GetCart(ctx, sessionId)
}
func (r *CartService) EmptyCart(ctx context.Context, sessionId string) error {
	return r.repo.EmptyCart(ctx, sessionId)
}
