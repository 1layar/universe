package server

import (
	"github.com/1layar/universe/internal/cart_service/server/watermill"
	"github.com/1layar/universe/internal/cart_service/server/watermill/command"
	"github.com/1layar/universe/internal/cart_service/server/watermill/route"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("server", watermill.Module(), route.Module(), command.Module())
}
