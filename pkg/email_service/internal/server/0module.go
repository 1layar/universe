package server

import (
	"github.com/1layar/universe/pkg/email_service/internal/server/mux"
	"github.com/1layar/universe/pkg/email_service/internal/server/watermill"
	"github.com/1layar/universe/pkg/email_service/internal/server/watermill/command"
	"github.com/1layar/universe/pkg/email_service/internal/server/watermill/route"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("server", watermill.Module(), route.Module(), command.Module(), mux.Module())
}
