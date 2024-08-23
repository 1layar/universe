package watermill

import (
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("watermill", fx.Provide(CreatePub), fx.Provide(CreateSub))
}
