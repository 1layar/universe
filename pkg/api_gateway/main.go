package main

import (
	"github.com/1layar/universe/pkg/api_gateway/cmd/app"
	"github.com/rs/zerolog/log"
)

// @title Merasa API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
// @securityDefinitions.apiKey JWT
// @in header
// @name Authorization
// @description "Type 'Bearer TOKEN' to correctly set the API Key"
func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Error().Msgf("%+v", err)
		}
	}()
	app.Run()
}
