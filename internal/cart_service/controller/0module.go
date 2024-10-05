package controller

import (
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("controller",
		fx.Invoke(ReqAddItem),
		fx.Invoke(ReqEmpty),
		fx.Invoke(ReqRemoveItem),
		fx.Invoke(ReqGetCart),
	)
}
