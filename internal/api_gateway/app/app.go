package app

import (
	"go.uber.org/fx"

	"github.com/1layar/universe/internal/api_gateway/app/appconfig"
	"github.com/1layar/universe/internal/api_gateway/app/appcontext"
	"github.com/1layar/universe/internal/api_gateway/controller"
	"github.com/1layar/universe/internal/api_gateway/server"
	"github.com/1layar/universe/internal/api_gateway/validator"
	"github.com/1layar/universe/pkg/logger"
)

func New(ctx appcontext.Ctx, additionalOpts ...fx.Option) *fx.App {
	logger := logger.GetLogger()
	logger.Debug("starting app...")
	conf, err := appconfig.Parse(ctx)
	if err != nil {
		panic(err)
	}

	baseOpts := []fx.Option{
		fx.WithLogger(logger.GetFxLogger),
		fx.Supply(conf),
		controller.Module(),
		validator.Module(),
		// middleware.Module(),
		// infra.Module(),
		// repo.Module(),
		// service.Module(),
		server.Module(),
	}

	return fx.New(append(baseOpts, additionalOpts...)...)
}
