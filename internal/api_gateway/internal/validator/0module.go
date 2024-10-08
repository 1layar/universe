package validator

import (
	"github.com/1layar/universe/pkg/api_gateway/internal/validator/rule"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("validator", fx.Provide(CreateValidator), fx.Invoke(rule.RegisterUserExistRule))
}
