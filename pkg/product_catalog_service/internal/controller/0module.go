package controller

import (
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("controller",
		fx.Invoke(RegAddProduct),
		fx.Invoke(RegExistSKU),
		fx.Invoke(RegUpdateProduct),
		fx.Invoke(RegGetAllProduct),
	)
}
