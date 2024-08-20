package service

import (
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module(
		"service",
		fx.Provide(NewEmailEventService),
		fx.Provide(NewEmailMessageService),
		fx.Provide(NewEmailTemplateService),
		fx.Provide(NewAccountService),
		fx.Provide(NewEmailManager),
		fx.Provide(NewEmailAgent),
	)
}
