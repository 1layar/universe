package infra

import (
	"go.uber.org/fx"

	"github.com/1layar/universe/internal/cart_service/infra/breaker"
	"github.com/1layar/universe/internal/cart_service/infra/cache"
)

func Module() fx.Option {
	return fx.Module("infra", breaker.Module(), cache.Module())
}
