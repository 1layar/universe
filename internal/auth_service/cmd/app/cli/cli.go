package cli

import (
	"context"

	"go.uber.org/fx"

	"github.com/1layar/universe/internal/auth_service/app"
	"github.com/1layar/universe/internal/auth_service/app/appcontext"
)

func Start(module fx.Option) {
	app.New(appcontext.Declare(appcontext.EnvCLI), module).Start(context.Background())
}
