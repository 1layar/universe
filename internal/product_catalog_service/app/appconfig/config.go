package appconfig

import (
	"github.com/1layar/universe/internal/product_catalog_service/app/appcontext"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
)

func Parse(ctx appcontext.Ctx) (*Config, error) {
	var conf ConfigSpec
	if err := envconfig.Process("app", &conf); err != nil {
		return nil, err
	}

	return &Config{
		ConfigSpec: conf,
		AppContext: ctx,
	}, nil
}
