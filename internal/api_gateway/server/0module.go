package server

import (
	"go.uber.org/fx"

	"github.com/1layar/universe/internal/api_gateway/server/http"
	"github.com/1layar/universe/internal/api_gateway/server/watermill"
	"github.com/1layar/universe/internal/api_gateway/server/watermill/command"
)

func Module() fx.Option {
	return fx.Module("server", http.Module(), watermill.Module(), command.Module())
}
