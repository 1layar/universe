package cache

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module("redis", fx.Provide(New))
}
