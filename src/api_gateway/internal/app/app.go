package app

import (
	"go.uber.org/fx"

	"github.com/1layar/merasa/backend/src/api_gateway/internal/app/appconfig"
	"github.com/1layar/merasa/backend/src/api_gateway/internal/app/appcontext"
	"github.com/1layar/merasa/backend/src/api_gateway/internal/controller"
	"github.com/1layar/merasa/backend/src/api_gateway/internal/server"
	"github.com/1layar/merasa/backend/src/api_gateway/internal/validator"
	"github.com/1layar/merasa/backend/src/api_gateway/internal/x/logger"
	"github.com/1layar/merasa/backend/src/api_gateway/internal/x/logger/fxlogger"
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
