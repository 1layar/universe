package cli

import (
	"context"

	"go.uber.org/fx"

	"github.com/1layar/universe/pkg/email_service/internal/app"
	"github.com/1layar/universe/pkg/email_service/internal/app/appcontext"
)

func Start(module fx.Option) {
	app.New(appcontext.Declare(appcontext.EnvCLI), module).Start(context.Background())
}
