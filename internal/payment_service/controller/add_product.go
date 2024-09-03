package controller

import (
	"github.com/1layar/universe/pkg/shared/command"
	"github.com/1layar/universe/pkg/shared/constant"
	"github.com/1layar/universe/pkg/shared/transport"
	"github.com/ThreeDotsLabs/watermill/components/requestreply"
)

func RegAddProduct(c ProductHandler) {
	backendConfig := transport.GetBackendConfig(constant.GET_ALL_PRODUCT_CMD, c.Publisher, c.Subscriber)
	backend, err := requestreply.NewPubSubBackend(
		backendConfig,
		requestreply.BackendPubsubJSONMarshaler[command.AddProductResult]{},
	)

	if err != nil {
		panic(err)
	}

	err = c.CommandProcessor.AddHandlers(
		requestreply.NewCommandHandlerWithResult(
			c.GetHandlerName("add_product"),
			backend,
			c.HandleAddProduct,
		),
	)

	if err != nil {
		panic(err)
	}
}
