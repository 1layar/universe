package app

import (
	"fmt"

	"github.com/1layar/universe/src/account_service/internal/app/appconfig"
	"github.com/1layar/universe/src/account_service/internal/app/appcontext"
	"github.com/1layar/universe/src/account_service/internal/controller"
	"github.com/1layar/universe/src/account_service/internal/infra"
	"github.com/1layar/universe/src/account_service/internal/repo"
	"github.com/1layar/universe/src/account_service/internal/server"
	"github.com/1layar/universe/src/account_service/internal/service"
	"github.com/1layar/universe/src/account_service/internal/x/logger"
	"github.com/1layar/universe/src/account_service/internal/x/logger/fxlogger"
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
	logger.Configure(conf)

	fmt.Println("configure logger...")

	baseOpts := []fx.Option{
		fx.WithLogger(fxlogger.Logger),
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

	return fx.New(append(baseOpts, additionalOpts...)...)
}
