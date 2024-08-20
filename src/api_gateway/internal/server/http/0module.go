package http

import (
	"go.uber.org/fx"

	"github.com/1layar/merasa/backend/src/api_gateway/internal/server/http/route"
)

func Module() fx.Option {
	return fx.Module("http", fx.Provide(Create), route.Module())
}
