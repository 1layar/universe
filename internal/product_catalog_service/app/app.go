package app

import (
	"fmt"

	"github.com/1layar/universe/internal/product_catalog_service/app/appconfig"
	"github.com/1layar/universe/internal/product_catalog_service/app/appcontext"
	"github.com/1layar/universe/internal/product_catalog_service/controller"
	"github.com/1layar/universe/internal/product_catalog_service/infra"
	"github.com/1layar/universe/internal/product_catalog_service/repo"
	"github.com/1layar/universe/internal/product_catalog_service/server"
	"github.com/1layar/universe/internal/product_catalog_service/service"
	"github.com/1layar/universe/pkg/logger"
	"go.uber.org/fx"
)

func New(ctx appcontext.Ctx, additionalOpts ...fx.Option) *fx.App {
	fmt.Println("starting app...")
	conf, err := appconfig.Parse(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("parse config...")
	// logger and configuration are the only two things that are not in the fx graph
	// because some other packages need them to be initialized before fx starts

	fmt.Println("configure logger...")
	logger := logger.GetLogger()
	baseOpts := []fx.Option{
		fx.WithLogger(logger.GetFxLogger()),
		fx.Supply(conf),
		controller.Module(),
		infra.Module(),
		repo.Module(),
		service.Module(),
		server.Module(),
	}

	fmt.Println("create fx app...")

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	app := fx.New(append(baseOpts, additionalOpts...)...)

	if err := app.Err(); err != nil {
		panic(err)
	}

	return app
}
