package processor

import (
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("handler", fx.Invoke(RegHandleOutBoxProcess))
}
