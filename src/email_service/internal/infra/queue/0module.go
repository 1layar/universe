package queue

import (
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module(
		"queue",
		fx.Provide(NewClient),
		fx.Provide(NewServer),
	)
}
