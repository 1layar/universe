package store

import (
	"context"

	"github.com/1layar/universe/internal/checkout_service/model"
)

type IStore interface {
	Initialize(ctx context.Context) error
	AddItem(ctx context.Context, sessionId string, productId int, quantity int, source string) error
	RemoveItem(ctx context.Context, sessionId string, productId int, source string) error
	UpdateQuantity(ctx context.Context, sessionId string, productId int, quantity int, source string) error
	GetCart(ctx context.Context, sessionId string) ([]model.CartItem, error)
	EmptyCart(ctx context.Context, sessionId string) error
}
