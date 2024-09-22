package cache

import (
	"github.com/1layar/universe/internal/cart_service/app/appconfig"
	"github.com/redis/go-redis/v9"
)

func New(config *appconfig.Config) *redis.Client {
	rdb := redis.NewClient(
		&redis.Options{
			Addr:     config.RedisAddr,
			Password: config.RedisPassword,
			DB:       0,
		},
	)

	return rdb
}
