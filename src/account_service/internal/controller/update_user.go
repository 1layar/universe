package controller

import (
	"github.com/1layar/merasa/backend/src/shared/command"
	"github.com/1layar/merasa/backend/src/shared/constant"
	"github.com/1layar/merasa/backend/src/shared/transport"
	"github.com/ThreeDotsLabs/watermill/components/requestreply"
)

func RegUpdateUser(c UserHandler) {
	backendConfig := transport.GetBackendConfig(constant.UPDATE_USER_CMD, c.Publisher, c.Subscriber)
	backend, err := requestreply.NewPubSubBackend(
		backendConfig,
		requestreply.BackendPubsubJSONMarshaler[command.UpdateUserResult]{},
	)

	if err != nil {
		panic(err)
	}

	err = c.CommandProcessor.AddHandlers(
		requestreply.NewCommandHandlerWithResult(
			c.GetHandlerName("update_user"),
			backend,
			c.HandleUpdateUser,
		),
	)

	if err != nil {
		panic(err)
	}
}
