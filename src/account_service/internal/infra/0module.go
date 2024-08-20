package infra

import (
	"go.uber.org/fx"

	"github.com/1layar/merasa/backend/src/account_service/internal/infra/db"
)

func Module() fx.Option {
	return fx.Module("infra", db.Module())
}
