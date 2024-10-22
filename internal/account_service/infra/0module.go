package infra

import (
	"go.uber.org/fx"

	"github.com/1layar/universe/internal/account_service/infra/db"
)

func Module() fx.Option {
	return fx.Module("infra", db.Module())
}
