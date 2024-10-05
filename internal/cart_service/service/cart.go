package service

import (
	"context"

	"github.com/1layar/universe/internal/cart_service/model"
	"github.com/1layar/universe/internal/cart_service/repo"
	"github.com/1layar/universe/pkg/shared/utils"
)

type CartService struct {
	repo *repo.CartRepository
}

func NewCartService(repo *repo.CartRepository) *CartService {
	return &CartService{
		repo: repo,
	}
}

func (s *CartService) AddItem(ctx context.Context, sessionId string, productId int, quantity int, source string) (string, error) {
	// generate sessionId when empty
	if sessionId == "" {
		nwSessionId, err := utils.GenerateSessionID(15)

		if err != nil {
			return "", err
		}

		hasCart, _ := s.repo.GetCart(ctx, nwSessionId)

		if len(hasCart) > 0 {
			return s.AddItem(ctx, sessionId, productId, quantity, source)
		}

		sessionId = nwSessionId
	}

	err := s.repo.AddItem(ctx, sessionId, productId, quantity, source)

	return sessionId, err
}
func (s *CartService) RemoveItem(ctx context.Context, sessionId string, productId int, source string) error {
	return s.repo.RemoveItem(ctx, sessionId, productId, source)
}
func (s *CartService) UpdateQuantity(ctx context.Context, sessionId string, productId int, quantity int, source string) error {
	return s.repo.UpdateQuantity(ctx, sessionId, productId, quantity, source)
}
func (s *CartService) GetCart(ctx context.Context, sessionId string) ([]model.CartItem, error) {
	return s.repo.GetCart(ctx, sessionId)
}
func (s *CartService) EmptyCart(ctx context.Context, sessionId string) error {
	return s.repo.EmptyCart(ctx, sessionId)
}
