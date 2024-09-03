package cli

import (
	"context"

	"go.uber.org/fx"

	"github.com/1layar/universe/internal/payment_service/app"
	"github.com/1layar/universe/internal/payment_service/app/appcontext"
	"github.com/1layar/universe/pkg/logger"
)

func Start(module fx.Option) {
	logger := logger.GetLogger()
	err := app.NewApp(appcontext.Declare(appcontext.EnvCLI), module).Start(context.Background())
	if err != nil {
		logger.Panic(err)
	}
}
