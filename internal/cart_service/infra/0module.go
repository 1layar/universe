package infra

import (
	"go.uber.org/fx"

	"github.com/1layar/universe/internal/cart_service/infra/breaker"
	"github.com/1layar/universe/internal/cart_service/infra/cache"
	"github.com/1layar/universe/internal/cart_service/infra/db"
)

func Module() fx.Option {
	return fx.Module("infra", db.Module(), breaker.Module(), cache.Module())
}
