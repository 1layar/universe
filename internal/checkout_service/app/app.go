package app

import (
	"github.com/1layar/universe/internal/checkout_service/app/appconfig"
	"github.com/1layar/universe/internal/checkout_service/app/appcontext"
	"github.com/1layar/universe/internal/checkout_service/controller"
	"github.com/1layar/universe/internal/checkout_service/infra"
	"github.com/1layar/universe/internal/checkout_service/repo"
	"github.com/1layar/universe/internal/checkout_service/server"
	"github.com/1layar/universe/internal/checkout_service/service"
	"github.com/1layar/universe/internal/checkout_service/store"
	"github.com/1layar/universe/pkg/logger"
	"go.uber.org/fx"
)

func NewApp(ctx appcontext.Ctx, additionalOpts ...fx.Option) *fx.App {
	logger := logger.GetLogger()
	logger.Debug("starting app...")
	conf, err := appconfig.Parse(ctx)
	if err != nil {
		panic(err)
	}
	logger.Debug("parse config...")
	// logger and configuration are the only two things that are not in the fx graph
	// because some other packages need them to be initialized before fx starts

	logger.Debug("configure logger...")
	baseOpts := []fx.Option{
		fx.WithLogger(logger.GetFxLogger),
		fx.Supply(conf),
		infra.Module(),
		repo.Module(),
		service.Module(),
		store.Module(),
	}

	if ctx.Env == appcontext.EnvServer {
		baseOpts = append(baseOpts, controller.Module(), server.Module())
	}

	logger.Debug("create fx app...")

	defer func() {
		if err := recover(); err != nil {
			logger.Error(err)
		}
	}()

	app := fx.New(append(baseOpts, additionalOpts...)...)

	if err := app.Err(); err != nil {
		panic(err)
	}

	return app
}
