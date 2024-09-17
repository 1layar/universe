package infra

import (
	"go.uber.org/fx"

	"github.com/1layar/universe/internal/email_service/internal/infra/db"
	"github.com/1layar/universe/internal/email_service/internal/infra/queue"
)

func Module() fx.Option {
	return fx.Module("infra", db.Module(), queue.Module())
}
