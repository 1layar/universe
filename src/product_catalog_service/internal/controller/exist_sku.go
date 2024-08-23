package controller

import (
	"github.com/1layar/universe/src/shared/command"
	"github.com/1layar/universe/src/shared/constant"
	"github.com/1layar/universe/src/shared/transport"
	"github.com/ThreeDotsLabs/watermill/components/requestreply"
)

func RegExistSKU(c ProductHandler) {
	backendConfig := transport.GetBackendConfig(constant.PRODUCT_EXIST_SKU_CMD, c.Publisher, c.Subscriber)
	backend, err := requestreply.NewPubSubBackend(
		backendConfig,
		requestreply.BackendPubsubJSONMarshaler[command.GetSKUExistsResult]{},
	)

	if err != nil {
		panic(err)
	}

	err = c.CommandProcessor.AddHandlers(
		requestreply.NewCommandHandlerWithResult(
			c.GetHandlerName("exist_sku"),
			backend,
			c.HandleSKUExist,
		),
	)

	if err != nil {
		panic(err)
	}
}
