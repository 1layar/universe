package server

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/message"
	"go.uber.org/fx"

	"github.com/1layar/universe/internal/cart_service/app"
	"github.com/1layar/universe/internal/cart_service/app/appconfig"
	"github.com/1layar/universe/internal/cart_service/app/appcontext"
	"github.com/1layar/universe/pkg/logger"
)

func Run() {
	app.NewApp(appcontext.Declare(appcontext.EnvServer), fx.Invoke(run)).Run()
}

func run(lc fx.Lifecycle, app *message.Router, conf *appconfig.Config) {
	logger := logger.GetLogger()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := app.Run(context.Background()); err != nil {
					logger.Error("server terminated unexpectedly")
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("gracefully shutting down server")
			if err := app.Close(); err != nil {
				logger.Error("error occurred while gracefully shutting down server")
				return err
			}
			logger.Info("graceful server shut down completed")
			return nil
		},
	})
}
