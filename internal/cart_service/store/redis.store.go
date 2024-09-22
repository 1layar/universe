package store

import (
	"context"
	"fmt"

	"github.com/1layar/universe/internal/cart_service/model"
	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	store *redis.Client
}

func NewRedisStore(store *redis.Client) IStore {
	return &RedisStore{
		store: store,
	}
}

// AddItem implements IStore.
func (r *RedisStore) AddItem(sessionId string, productId int, quantity int, source string) error {
	// get cart from redis
	cmd := r.store.HSet(context.Background(), fmt.Sprintf("cart:%s", sessionId), sessionId, fmt.Sprintf("product:%d", productId))
	_, err := cmd.Result()
	return err
}

// GetCart implements IStore.
func (r *RedisStore) GetCart(sessionId string) (model.Cart, error) {
	// get cart from redis
	cmd := r.store.HGetAll(context.Background(), fmt.Sprintf("cart:%s", sessionId))
	return model.Cart{}, cmd.Err()
}

// Initialize implements IStore.
func (r *RedisStore) Initialize() error {
	panic("unimplemented")
}

// RemoveItem implements IStore.
func (r *RedisStore) RemoveItem(sessionId string, productId int, source string) error {
	panic("unimplemented")
}

// UpdateQuantity implements IStore.
func (r *RedisStore) UpdateQuantity(sessionId string, productId int, quantity int, source string) error {
	panic("unimplemented")
}
