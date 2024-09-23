package store

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/1layar/universe/internal/cart_service/model"
	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	store *redis.Client
}

// EmptyCart implements IStore.
func (r *RedisStore) EmptyCart(ctx context.Context, sessionId string) error {
	key := fmt.Sprintf("cart:%s", sessionId)

	return r.store.Del(ctx, key).Err()
}

func NewRedisStore(store *redis.Client) IStore {
	return &RedisStore{
		store: store,
	}
}

// AddItem implements IStore.
func (r *RedisStore) AddItem(ctx context.Context, sessionId string, productId int, quantity int, source string) error {
	key := fmt.Sprintf("cart:%s", sessionId)
	// get cart from redis
	err := r.store.HIncrBy(ctx, key, fmt.Sprintf("product:%s:%d", source, productId), int64(quantity)).Err()

	if err != nil {
		return err
	}

	return nil
}

// GetCart implements IStore.
func (r *RedisStore) GetCart(ctx context.Context, sessionId string) ([]model.CartItem, error) {
	cartItems := []model.CartItem{}
	key := fmt.Sprintf("cart:%s", sessionId)
	items, err := r.store.HGetAll(ctx, key).Result()
	if err != nil {
		return cartItems, err
	}

	for productIDStr, quantityStr := range items {
		// get product id from split string by :
		// product:{source}:{product_id}
		items := strings.Split(productIDStr, ":")

		if len(items) != 3 {
			return cartItems, fmt.Errorf("invalid product id: %s", productIDStr)
		}

		productIDStr = items[2]

		quantity, _ := strconv.Atoi(quantityStr)
		productID, _ := strconv.Atoi(productIDStr)
		cartItems = append(cartItems, model.CartItem{
			ProductID: productID,
			Quantity:  quantity,
		})
	}
	return cartItems, nil

}

// Initialize implements IStore.
func (r *RedisStore) Initialize(ctx context.Context) error {
	// ping redis server
	_, err := r.store.Ping(ctx).Result()

	return err
}

// RemoveItem implements IStore.
func (r *RedisStore) RemoveItem(ctx context.Context, sessionId string, productId int, source string) error {
	key := fmt.Sprintf("cart:%s", sessionId)

	err := r.store.HDel(ctx, key, fmt.Sprintf("product:%s:%d", source, productId)).Err()

	if err != nil {
		return err
	}

	return nil
}

// UpdateQuantity implements IStore.
func (r *RedisStore) UpdateQuantity(ctx context.Context, sessionId string, productId int, quantity int, source string) error {

	key := fmt.Sprintf("cart:%s", sessionId)

	err := r.store.HSet(ctx, key, fmt.Sprintf("product:%s:%d", source, productId), quantity).Err()

	if err != nil {
		return err
	}

	return nil
}
