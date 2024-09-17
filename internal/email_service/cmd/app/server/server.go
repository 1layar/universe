package server

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"

	"github.com/1layar/universe/internal/email_service/internal/app"
	"github.com/1layar/universe/internal/email_service/internal/app/appconfig"
	"github.com/1layar/universe/internal/email_service/internal/app/appcontext"
	"github.com/hibiken/asynq"
)

func Run() {
	log.Info().Msg("starting server...")
	app.New(appcontext.Declare(appcontext.EnvServer), fx.Invoke(run)).Run()
}

func run(
	lc fx.Lifecycle,
	app *message.Router,
	conf *appconfig.Config,
	asynqServer *asynq.Server,
	mux *asynq.ServeMux,
) {
	log.Info().Msg("run server...")
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Info().Msg("starting watermill server...")
				if err := app.Run(context.Background()); err != nil {
					log.Error().Err(err).Msg("server terminated unexpectedly")
				}
			}()

			go func() {
				log.Info().Msg("starting asynq server...")
				if err := asynqServer.Run(mux); err != nil {
					log.Error().Err(err).Msg("Could not run Asynq server")
				}
				log.Info().Msg("asynq server stopped")

			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info().Msg("gracefully shutting down server")
			if err := app.Close(); err != nil {
				log.Error().Err(err).Msg("error occurred while gracefully shutting down server")
				return err
			}

			asynqServer.Shutdown()

			log.Info().Msg("graceful server shut down completed")
			return nil
		},
	})
}
