package main

import (
	"github.com/1layar/universe/pkg/auth_service/cmd/app"

	"github.com/rs/zerolog/log"
)

var (
	name    = "auth_service"
	version = "1.0.0"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Error().Msgf("%+v", err)
		}
	}()

	app.Run(name, version)
}
