package mux

import (
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("mux", fx.Provide(Create))
}
