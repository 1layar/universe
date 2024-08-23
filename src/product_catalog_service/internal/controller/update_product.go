package controller

import (
	"github.com/1layar/universe/src/shared/command"
	"github.com/1layar/universe/src/shared/constant"
	"github.com/1layar/universe/src/shared/transport"
	"github.com/ThreeDotsLabs/watermill/components/requestreply"
)

func RegUpdateProduct(c ProductHandler) {
	backendConfig := transport.GetBackendConfig(constant.UPDATE_PRODUCT_CMD, c.Publisher, c.Subscriber)
	backend, err := requestreply.NewPubSubBackend(
		backendConfig,
		requestreply.BackendPubsubJSONMarshaler[command.UpdateProductResult]{},
	)

	if err != nil {
		panic(err)
	}

	err = c.CommandProcessor.AddHandlers(
		requestreply.NewCommandHandlerWithResult(
			c.GetHandlerName("update_product"),
			backend,
			c.HandleAddUpdate,
		),
	)

	if err != nil {
		panic(err)
	}
}
