package server

import (
	"github.com/1layar/merasa/backend/src/auth_service/internal/server/watermill"
	"github.com/1layar/merasa/backend/src/auth_service/internal/server/watermill/command"
	"github.com/1layar/merasa/backend/src/auth_service/internal/server/watermill/route"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("server", watermill.Module(), route.Module(), command.Module())
}
