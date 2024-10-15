package app

import (
	"fmt"

	"github.com/1layar/universe/internal/account_service/app/appconfig"
	"github.com/1layar/universe/internal/account_service/app/appcontext"
	"github.com/1layar/universe/internal/account_service/controller"
	"github.com/1layar/universe/internal/account_service/infra"
	"github.com/1layar/universe/internal/account_service/repo"
	"github.com/1layar/universe/internal/account_service/server"
	"github.com/1layar/universe/internal/account_service/service"
	"github.com/1layar/universe/internal/account_service/x/logger"
	"github.com/1layar/universe/internal/account_service/x/logger/fxlogger"
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
