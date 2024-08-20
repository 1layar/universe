package guard

import (
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("guard", fx.Provide(New))
}
