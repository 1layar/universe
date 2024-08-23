package handler

import (
	"github.com/1layar/universe/pkg/shared/command"
	"github.com/1layar/universe/pkg/shared/constant"
	"github.com/1layar/universe/pkg/shared/transport"
	"github.com/ThreeDotsLabs/watermill/components/requestreply"
)

func RegLogin(c AuthHandler) {
	backendConfig := transport.GetBackendConfig(constant.AUTH_LOGIN_CMD, c.Publisher, c.Subscriber)
	backend, err := requestreply.NewPubSubBackend(
		backendConfig,
		requestreply.BackendPubsubJSONMarshaler[command.LoginResult]{},
	)

	if err != nil {
		panic(err)
	}

	err = c.CommandProcessor.AddHandlers(
		requestreply.NewCommandHandlerWithResult(
			c.GetHandlerName("login"),
			backend,
			c.HandleLogin,
		),
	)

	if err != nil {
		panic(err)
	}
}
