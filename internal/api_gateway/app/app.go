package app

import (
	"go.uber.org/fx"

	"github.com/1layar/universe/internal/api_gateway/app/appconfig"
	"github.com/1layar/universe/internal/api_gateway/app/appcontext"
	"github.com/1layar/universe/internal/api_gateway/controller"
	"github.com/1layar/universe/internal/api_gateway/server"
	"github.com/1layar/universe/internal/api_gateway/validator"
	"github.com/1layar/universe/internal/api_gateway/x/logger"
	"github.com/1layar/universe/internal/api_gateway/x/logger/fxlogger"
)

func New(ctx appcontext.Ctx, additionalOpts ...fx.Option) *fx.App {
	conf, err := appconfig.Parse(ctx)
	if err != nil {
		panic(err)
	}

	// logger and configuration are the only two things that are not in the fx graph
	// because some other packages need them to be initialized before fx starts
	logger.Configure(conf)

	baseOpts := []fx.Option{
		fx.WithLogger(fxlogger.Logger),
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