package controller

import (
	"github.com/1layar/merasa/backend/src/shared/command"
	"github.com/1layar/merasa/backend/src/shared/constant"
	"github.com/1layar/merasa/backend/src/shared/transport"
	"github.com/ThreeDotsLabs/watermill/components/requestreply"
)

func RegAddUser(c UserHandler) {
	backendConfig := transport.GetBackendConfig(constant.ADD_USER_CMD, c.Publisher, c.Subscriber)
	backend, err := requestreply.NewPubSubBackend(
		backendConfig,
		requestreply.BackendPubsubJSONMarshaler[command.AddUserResult]{},
	)

	if err != nil {
		panic(err)
	}

	err = c.CommandProcessor.AddHandlers(
		requestreply.NewCommandHandlerWithResult(
			c.GetHandlerName("add_user"),
			backend,
			c.HandleAddUser,
		),
	)

	if err != nil {
		panic(err)
	}
}
