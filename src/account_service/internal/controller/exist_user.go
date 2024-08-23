package controller

import (
	"github.com/1layar/universe/src/shared/command"
	"github.com/1layar/universe/src/shared/constant"
	"github.com/1layar/universe/src/shared/transport"
	"github.com/ThreeDotsLabs/watermill/components/requestreply"
)

func RegExistUsername(c UserHandler) {
	backendConfig := transport.GetBackendConfig(constant.USER_EXIST_USERNAME_CMD, c.Publisher, c.Subscriber)
	backend, err := requestreply.NewPubSubBackend(
		backendConfig,
		requestreply.BackendPubsubJSONMarshaler[command.GetUsernameExistsResult]{},
	)

	if err != nil {
		panic(err)
	}

	err = c.CommandProcessor.AddHandlers(
		requestreply.NewCommandHandlerWithResult(
			c.GetHandlerName("exist_username"),
			backend,
			c.HandleUsernameExist,
		),
	)

	if err != nil {
		panic(err)
	}
}

func RegExistEmail(c UserHandler) {
	backendConfig := transport.GetBackendConfig(constant.USER_EXIST_EMAIL_CMD, c.Publisher, c.Subscriber)
	backend, err := requestreply.NewPubSubBackend(
		backendConfig,
		requestreply.BackendPubsubJSONMarshaler[command.GetEmailExistsResult]{},
	)

	if err != nil {
		panic(err)
	}

	err = c.CommandProcessor.AddHandlers(
		requestreply.NewCommandHandlerWithResult(
			c.GetHandlerName("exist_email"),
			backend,
			c.HandleEmailExist,
		),
	)

	if err != nil {
		panic(err)
	}
}
