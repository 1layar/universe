package server

import (
	"go.uber.org/fx"

	"github.com/1layar/merasa/backend/src/api_gateway/internal/server/http"
	"github.com/1layar/merasa/backend/src/api_gateway/internal/server/watermill"
	"github.com/1layar/merasa/backend/src/api_gateway/internal/server/watermill/command"
)

func Module() fx.Option {
	return fx.Module("server", http.Module(), watermill.Module(), command.Module())
}
