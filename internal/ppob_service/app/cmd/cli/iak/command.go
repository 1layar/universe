package iak

import (
	"github.com/1layar/universe/internal/ppob_service/app/appconfig"
	cliapp "github.com/1layar/universe/internal/ppob_service/app/cmd/cli"
	"github.com/1layar/universe/internal/ppob_service/service"
	"github.com/1layar/universe/pkg/logger"
	"github.com/urfave/cli/v2"
	"go.uber.org/fx"
)

type iakCommandDeps struct {
	fx.In

	IakService     *service.IakService
	ProductService *service.ProductService
}

func Command() *cli.Command {
	log := logger.GetLogger()
	var deps iakCommandDeps
	log.Info("prepare iak service...")
	cliapp.Start(fx.Populate(&deps))

	log.Info("init iak service...")

	iakService := deps.IakService
	productService := deps.ProductService

	return &cli.Command{
		Name:  "iak",
		Usage: "manage iak service",
		Subcommands: []*cli.Command{
			{
				Name:  "check-balance",
				Usage: "check iak balance",
				Action: func(c *cli.Context) error {
					log.Info("check iak balance...")
					balance, err := iakService.GetBalance()

					if err != nil {
						log.Error("failed to get iak balance")
						return cli.Exit(err, 1)
					}

					return cli.Exit(balance, 0)
				},
			},
			{
				Name:  "price-list",
				Usage: "get iak price list",
				Action: func(c *cli.Context) error {
					log.Info("get iak price list...")
					priceList, err := iakService.GetPriceList(
						appconfig.Pulsa,
						appconfig.PulsaAxis,
						appconfig.All,
					)

					if err != nil {
						log.Error("failed to get iak price list")
						return cli.Exit(err, 1)
					}

					return cli.Exit(priceList, 0)
				},
			},
			{
				Name:  "import-product",
				Usage: "sync iak product with own product list",
				Action: func(c *cli.Context) error {
					log.Info("import iak price list...")
					priceList, err := iakService.GetPriceList(
						appconfig.Pulsa,
						appconfig.PulsaAxis,
						appconfig.All,
					)

					if err != nil {
						log.Error("failed to get iak price list")
						return cli.Exit(err, 1)
					}

					err = productService.ImportIak(c.Context, priceList.Data.PriceList)

					if err != nil {
						log.Error("failed to import iak price list")
						return cli.Exit(err, 1)
					}

					return cli.Exit("Successed", 0)
				},
			},
		},
	}
}
