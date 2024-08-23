package handler

import (
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("controller",
		fx.Invoke(RegLogin),
		fx.Invoke(RegJwtToUser),
		fx.Invoke(RegRegister),
		fx.Invoke(RegAsignAccess),
		fx.Invoke(RegCheckAccess),
	)
}
