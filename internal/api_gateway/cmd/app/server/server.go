package server

import (
	"context"
	"net"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"

	"github.com/1layar/universe/pkg/api_gateway/docs"
	"github.com/1layar/universe/pkg/api_gateway/internal/app"
	"github.com/1layar/universe/pkg/api_gateway/internal/app/appconfig"
	"github.com/1layar/universe/pkg/api_gateway/internal/app/appcontext"
)

func Run() {
	app.New(appcontext.Declare(appcontext.EnvServer), fx.Invoke(run)).Run()
}

func run(lc fx.Lifecycle, app *fiber.App, conf *appconfig.Config) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", conf.ServiceListenAddress)
			if err != nil {
				log.Error().Err(err).Msg("failed to listen")
				return err
			}
			docs.SwaggerInfo.Host = ln.Addr().String()
			go func() {
				if err := app.Listener(ln); err != nil {
					log.Error().Err(err).Msg("server terminated unexpectedly")
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info().Msg("gracefully shutting down server")
			if err := app.Shutdown(); err != nil {
				log.Error().Err(err).Msg("error occurred while gracefully shutting down server")
				return err
			}
			log.Info().Msg("graceful server shut down completed")
			return nil
		},
	})
}
