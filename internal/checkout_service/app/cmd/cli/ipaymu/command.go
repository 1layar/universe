package ipaymu

import (
	cliapp "github.com/1layar/universe/internal/checkout_service/app/cmd/cli"
	"github.com/1layar/universe/internal/checkout_service/service"
	"github.com/1layar/universe/pkg/logger"
	"github.com/urfave/cli/v2"
	"go.uber.org/fx"
)

type ipaymuCommandDeps struct {
	fx.In

	IpaymuService *service.IpaymuService
}

func Command() *cli.Command {
	log := logger.GetLogger()
	var deps ipaymuCommandDeps
	log.Info("prepare ipaymu service...")
	cliapp.Start(fx.Populate(&deps))
	log.Info("init ipaymu service...")

	ipaymuService := deps.IpaymuService
	return &cli.Command{
		Name:  "ipaymu",
		Usage: "manage ipaymu service",
		Subcommands: []*cli.Command{
			{
				Name:  "check-balance",
				Usage: "check ipaymu balance",
				Action: func(c *cli.Context) error {
					log.Info("check ipaymu balance...")
					balance, err := ipaymuService.CheckBalance()
					if err != nil {
						log.Error("failed to get ipaymu balance")
						return cli.Exit(err, 1)
					}

					return cli.Exit(balance, 0)
				},
			},
			{
				Name:  "payment-channels",
				Usage: "get ipaymu payment channels",
				Action: func(c *cli.Context) error {
					log.Info("get ipaymu payment channels...")
					paymentChannels, err := ipaymuService.IpaymuPaymentMethod()
					if err != nil {
						log.Error("failed to get ipaymu payment channels")
						return cli.Exit(err, 1)
					}

					return cli.Exit(paymentChannels, 0)
				},
			},
		},
	}
}
