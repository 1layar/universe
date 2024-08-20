package app

import (
	"github.com/1layar/merasa/backend/src/email_service/internal/app/appconfig"
	"github.com/1layar/merasa/backend/src/email_service/internal/app/appcontext"
	"github.com/1layar/merasa/backend/src/email_service/internal/handler"
	"github.com/1layar/merasa/backend/src/email_service/internal/infra"
	"github.com/1layar/merasa/backend/src/email_service/internal/processor"
	"github.com/1layar/merasa/backend/src/email_service/internal/repo"
	"github.com/1layar/merasa/backend/src/email_service/internal/server"
	"github.com/1layar/merasa/backend/src/email_service/internal/service"
	"github.com/1layar/merasa/backend/src/email_service/internal/x/logger"
	"github.com/1layar/merasa/backend/src/email_service/internal/x/logger/fxlogger"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
)

func New(ctx appcontext.Ctx, additionalOpts ...fx.Option) *fx.App {
	log.Info().Msg("starting app...")
	conf, err := appconfig.Parse(ctx)
	if err != nil {
		panic(err)
	}
	log.Info().Msg("parse config...")
	// logger and configuration are the only two things that are not in the fx graph
	// because some other packages need them to be initialized before fx starts
	logger.Configure(conf)

	log.Info().Msg("configure logger...")

	baseOpts := []fx.Option{
		fx.WithLogger(fxlogger.Logger),
		fx.Supply(conf),
		infra.Module(),
		repo.Module(),
		handler.Module(),
		processor.Module(),
		service.Module(),
		server.Module(),
	}

	log.Info().Msg("create fx app...")

	defer func() {
		if err := recover(); err != nil {
			log.Err(err.(error)).Msg("app terminated unexpectedly")
		}
	}()

	app := fx.New(append(baseOpts, additionalOpts...)...)

	if err := app.Err(); err != nil {
		panic(err)
	}

	return app
}
