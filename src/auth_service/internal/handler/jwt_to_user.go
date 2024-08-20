package handler

import (
	"github.com/1layar/merasa/backend/src/shared/command"
	"github.com/1layar/merasa/backend/src/shared/constant"
	"github.com/1layar/merasa/backend/src/shared/transport"
	"github.com/ThreeDotsLabs/watermill/components/requestreply"
)

func RegJwtToUser(c AuthHandler) {
	backendConfig := transport.GetBackendConfig(constant.AUTH_JWT_TO_USER_CMD, c.Publisher, c.Subscriber)
	backend, err := requestreply.NewPubSubBackend(
		backendConfig,
		requestreply.BackendPubsubJSONMarshaler[command.JwtToUserResponse]{},
	)

	if err != nil {
		panic(err)
	}

	err = c.CommandProcessor.AddHandlers(
		requestreply.NewCommandHandlerWithResult(
			c.GetHandlerName("jwtToUser"),
			backend,
			c.HandleJwtToUser,
		),
	)

	if err != nil {
		panic(err)
	}
}
