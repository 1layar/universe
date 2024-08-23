package controller

import (
	"github.com/1layar/universe/src/shared/command"
	"github.com/1layar/universe/src/shared/constant"
	"github.com/1layar/universe/src/shared/transport"
	"github.com/ThreeDotsLabs/watermill/components/requestreply"
)

func RegGetAllProduct(c ProductHandler) {
	backendConfig := transport.GetBackendConfig(constant.GET_ALL_PRODUCT_CMD, c.Publisher, c.Subscriber)
	backend, err := requestreply.NewPubSubBackend(
		backendConfig,
		requestreply.BackendPubsubJSONMarshaler[command.GetAllProductResult]{},
	)

	if err != nil {
		panic(err)
	}

	err = c.CommandProcessor.AddHandlers(
		requestreply.NewCommandHandlerWithResult(
			c.GetHandlerName("get_all_product"),
			backend,
			c.HandleGetAllProduct,
		),
	)

	if err != nil {
		panic(err)
	}
}
