package breaker

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module("breaker", fx.Provide(New))
}
