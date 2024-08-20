package main

import (
	"github.com/1layar/merasa/backend/src/product_catalog_service/cmd/app"
	"github.com/rs/zerolog/log"
)

var (
	name    = "product_catalog_service"
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
