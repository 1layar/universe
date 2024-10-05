package controller

import (
	"github.com/1layar/universe/pkg/shared/command"
	"github.com/1layar/universe/pkg/shared/constant"
	"github.com/1layar/universe/pkg/shared/transport"
	"github.com/ThreeDotsLabs/watermill/components/requestreply"
)

func ReqGetCart(c CartHandler) {
	backendConfig := transport.GetBackendConfig(constant.CART_GET_CART_CMD, c.Publisher, c.Subscriber)
	backend, err := requestreply.NewPubSubBackend(
		backendConfig,
		requestreply.BackendPubsubJSONMarshaler[command.CartGetCartCommandResult]{},
	)

	if err != nil {
		panic(err)
	}

	err = c.CommandProcessor.AddHandlers(
		requestreply.NewCommandHandlerWithResult(
			c.GetHandlerName(constant.CART_GET_CART_CMD),
			backend,
			c.GetCart,
		),
	)

	if err != nil {
		panic(err)
	}
}
