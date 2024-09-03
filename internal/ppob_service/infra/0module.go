package infra

import (
	"go.uber.org/fx"

	"github.com/1layar/universe/internal/ppob_service/infra/db"
)

func Module() fx.Option {
	return fx.Module("infra", db.Module())
}
