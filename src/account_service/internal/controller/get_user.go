package controller

import (
	"github.com/1layar/merasa/backend/src/shared/command"
	"github.com/1layar/merasa/backend/src/shared/constant"
	"github.com/1layar/merasa/backend/src/shared/transport"
	"github.com/ThreeDotsLabs/watermill/components/requestreply"
)

func RegGetUser(c UserHandler) {
	backendConfig := transport.GetBackendConfig(constant.GET_USER_CMD, c.Publisher, c.Subscriber)
	backend, err := requestreply.NewPubSubBackend(
		backendConfig,
		requestreply.BackendPubsubJSONMarshaler[command.GetUserResult]{},
	)

	if err != nil {
		panic(err)
	}

	err = c.CommandProcessor.AddHandlers(
		requestreply.NewCommandHandlerWithResult(
			c.GetHandlerName("get_user"),
			backend,
			c.HandleGetUser,
		),
	)

	if err != nil {
		panic(err)
	}
}
