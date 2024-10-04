package app

import (
	"os"

	"github.com/1layar/universe/internal/cart_service/app/cmd/server"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

func Run(name string, version string) {
	app := &cli.App{
		Name:    name,
		Version: version,
		Commands: []*cli.Command{
			server.Command(),
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal().Err(err).Msg("failed to run app")
	}
}
