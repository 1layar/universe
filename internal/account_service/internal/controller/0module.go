package controller

import (
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("controller",
		fx.Invoke(RegAddUser),
		fx.Invoke(RegExistUsername),
		fx.Invoke(RegExistEmail),
		fx.Invoke(RegUpdateUser),
		fx.Invoke(RegDeleteUser),
		fx.Invoke(RegSearchUser),
		fx.Invoke(RegGetUser),
	)
}
