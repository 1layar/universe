package handler

import (
	"github.com/1layar/merasa/backend/src/shared/command"
	"github.com/1layar/merasa/backend/src/shared/constant"
	"github.com/1layar/merasa/backend/src/shared/transport"
	"github.com/ThreeDotsLabs/watermill/components/requestreply"
)

func RegCheckAccess(c AuthHandler) {
	backendConfig := transport.GetBackendConfig(constant.AUTH_CHECK_ACCESS_CMD, c.Publisher, c.Subscriber)
	backend, err := requestreply.NewPubSubBackend(
		backendConfig,
		requestreply.BackendPubsubJSONMarshaler[command.CheckAccessResponse]{},
	)

	if err != nil {
		panic(err)
	}

	err = c.CommandProcessor.AddHandlers(
		requestreply.NewCommandHandlerWithResult(
			c.GetHandlerName("checkAccess"),
			backend,
			c.HandleCheckAccess,
		),
	)

	if err != nil {
		panic(err)
	}
}
