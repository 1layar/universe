package store

import (
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("store",
		fx.Provide(NewRedisStore),
	)
}
