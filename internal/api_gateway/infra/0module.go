package infra

import (
	"go.uber.org/fx"

	"github.com/1layar/universe/internal/api_gateway/infra/db"
)

func Module() fx.Option {
	return fx.Module("infra", db.Module())
}