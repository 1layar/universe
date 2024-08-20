package repo

import (
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("repo",
		fx.Provide(NewEmailEventRepository),
		fx.Provide(NewEmailMessageRepository),
		fx.Provide(NewEmailTemplateRepository),
		fx.Provide(NewAccountRepository),
	)
}
