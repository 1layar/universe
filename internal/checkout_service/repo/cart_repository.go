package repo

import (
	"context"

	"github.com/1layar/universe/internal/checkout_service/model"
	"github.com/1layar/universe/internal/checkout_service/store"
)

type CartRepository struct {
	cartStore *store.RedisStore
}

func NewCartRepository(cartStore *store.RedisStore) *CartRepository {
	return &CartRepository{
		cartStore: cartStore,
	}
}

func (r *CartRepository) AddItem(ctx context.Context, sessionId string, productId int, quantity int, source string) error {
	return r.cartStore.AddItem(ctx, sessionId, productId, quantity, source)
}
func (r *CartRepository) RemoveItem(ctx context.Context, sessionId string, productId int, source string) error {
	return r.cartStore.RemoveItem(ctx, sessionId, productId, source)
}
func (r *CartRepository) UpdateQuantity(ctx context.Context, sessionId string, productId int, quantity int, source string) error {
	return r.cartStore.UpdateQuantity(ctx, sessionId, productId, quantity, source)
}
func (r *CartRepository) GetCart(ctx context.Context, sessionId string) ([]model.CartItem, error) {
	return r.cartStore.GetCart(ctx, sessionId)
}
func (r *CartRepository) EmptyCart(ctx context.Context, sessionId string) error {
	return r.cartStore.EmptyCart(ctx, sessionId)
}
