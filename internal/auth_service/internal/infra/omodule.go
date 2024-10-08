package infra

import (
	"go.uber.org/fx"

	"github.com/1layar/universe/internal/auth_service/internal/infra/db"
	"github.com/1layar/universe/internal/auth_service/internal/infra/guard"
)

func Module() fx.Option {
	return fx.Module("infra", db.Module(), guard.Module())
}
