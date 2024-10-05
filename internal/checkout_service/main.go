package main

import (
	app "github.com/1layar/universe/internal/checkout_service/app/cmd"
	"github.com/rs/zerolog/log"
)

var (
	name    = "cart_service"
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