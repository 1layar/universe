package controller

import (
	"github.com/1layar/universe/pkg/shared/command"
	"github.com/1layar/universe/pkg/shared/constant"
	"github.com/1layar/universe/pkg/shared/transport"
	"github.com/ThreeDotsLabs/watermill/components/requestreply"
)

func RegGetAllType(c TypeHandler) {
	backendConfig := transport.GetBackendConfig(constant.PPOB_GET_ALL_PRODUCT_TYPE_CMD, c.Publisher, c.Subscriber)
	backend, err := requestreply.NewPubSubBackend(
		backendConfig,
		requestreply.BackendPubsubJSONMarshaler[command.PpobGetTypeResult]{},
	)

	if err != nil {
		panic(err)
	}

	err = c.CommandProcessor.AddHandlers(
		requestreply.NewCommandHandlerWithResult(
			c.GetHandlerName("ppob.get_all_type"),
			backend,
			c.HandleGetAllType,
		),
	)

	if err != nil {
		panic(err)
	}
}
