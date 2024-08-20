package cli

import (
	"context"

	"go.uber.org/fx"

	"github.com/1layar/merasa/backend/src/api_gateway/internal/app"
	"github.com/1layar/merasa/backend/src/api_gateway/internal/app/appcontext"
)

func Start(module fx.Option) {
	app.New(appcontext.Declare(appcontext.EnvCLI), module).Start(context.Background())
}
