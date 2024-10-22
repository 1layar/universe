package validator

import (
	"github.com/1layar/universe/internal/api_gateway/validator/rule"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("validator", fx.Provide(CreateValidator), fx.Invoke(rule.RegisterUserExistRule))
}
