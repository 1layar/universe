package command

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module("command", fx.Provide(CreateCommandBus))
}
