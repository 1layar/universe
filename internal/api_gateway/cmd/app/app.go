package app

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"

	"github.com/1layar/universe/internal/api_gateway/cmd/app/cli/db"
	"github.com/1layar/universe/internal/api_gateway/cmd/app/server"
)

func Run() {
	app := &cli.App{
		Name: "app",
		Commands: []*cli.Command{
			server.Command(),
			db.Command(),
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal().Err(err).Msg("failed to run app")
	}
}
