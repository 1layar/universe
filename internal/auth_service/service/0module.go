package service

import (
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("service",
		fx.Provide(NewAuthService),
		fx.Provide(NewSessionService),
		fx.Provide(NewAuthorizeService),
	)
}
