package server

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"

	"github.com/1layar/universe/pkg/auth_service/internal/app"
	"github.com/1layar/universe/pkg/auth_service/internal/app/appconfig"
	"github.com/1layar/universe/pkg/auth_service/internal/app/appcontext"
)

func Run() {
	log.Info().Msg("starting server...")
	app.New(appcontext.Declare(appcontext.EnvServer), fx.Invoke(run)).Run()
}

func run(lc fx.Lifecycle, app *message.Router, conf *appconfig.Config) {
	log.Info().Msg("run server...")
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := app.Run(context.Background()); err != nil {
					log.Error().Err(err).Msg("server terminated unexpectedly")
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info().Msg("gracefully shutting down server")
			if err := app.Close(); err != nil {
				log.Error().Err(err).Msg("error occurred while gracefully shutting down server")
				return err
			}
			log.Info().Msg("graceful server shut down completed")
			return nil
		},
	})
}
