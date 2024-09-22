package store

import "github.com/1layar/universe/internal/cart_service/model"

type IStore interface {
	Initialize() error
	AddItem(sessionId string, productId int, quantity int, source string) error
	RemoveItem(sessionId string, productId int, source string) error
	UpdateQuantity(sessionId string, productId int, quantity int, source string) error
	GetCart(sessionId string) (model.Cart, error)
}
