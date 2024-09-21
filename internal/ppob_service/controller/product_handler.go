package controller

import (
	"context"

	"github.com/1layar/universe/internal/ppob_service/model"
	"github.com/1layar/universe/internal/ppob_service/service"
	"github.com/1layar/universe/pkg/shared/command"
	"github.com/1layar/universe/pkg/shared/constant"
	"github.com/1layar/universe/pkg/shared/event"
	"github.com/1layar/universe/pkg/shared/transport"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/components/requestreply"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/goccy/go-json"
	"go.uber.org/fx"
)

type ProductHandler struct {
	fx.In
	ProductService   *service.ProductService
	CommandBus       *cqrs.CommandBus
	Publisher        *amqp.Publisher
	Subscriber       *amqp.Subscriber
	CommandProcessor *cqrs.CommandProcessor
}

type RouterItem[T any, R any] struct {
	Cmd     string
	Handler func(ctx context.Context, cmd *T) (R, error)
	Config  func(cmd string) requestreply.PubSubBackendConfig
}

func (i RouterItem[T, R]) Build() cqrs.CommandHandler {
	backendConfig := i.Config(i.Cmd)
	backend, err := requestreply.NewPubSubBackend(
		backendConfig,
		requestreply.BackendPubsubJSONMarshaler[R]{},
	)

	if err != nil {
		panic(err)
	}

	return requestreply.NewCommandHandlerWithResult(
		i.Cmd,
		backend,
		i.Handler,
	)
}

func (h ProductHandler) GetHandlerName(name string) string {
	return "product_catalog_service.product." + name
}

func (h ProductHandler) RegHandler(handlers ...RouterItem[any, any]) {

	config := func(cmd string) requestreply.PubSubBackendConfig {
		return transport.GetBackendConfig(cmd, h.Publisher, h.Subscriber)
	}

	for _, handler := range handlers {
		handler.Config = config

		err := h.CommandProcessor.AddHandlers(
			handler.Build(),
		)

		if err != nil {
			panic(err)
		}
	}
}

func (h ProductHandler) HandleGetAllProduct(ctx context.Context, cmd *command.GetAllProductCommand) (command.GetAllProductResult, error) {
	products, total, err := h.ProductService.GetAllPaginate(ctx,
		cmd.Page, cmd.Limit,
		[]string{"Categories"},
		map[string]any{
			"name": cmd.Name,
			"sku":  cmd.SKU,
		})

	if err != nil {
		return command.GetAllProductResult{}, err
	}

	productResults := make([]command.ProductResult, len(products))

	for i, product := range products {
		productResults[i] = command.ProductResult{
			ID:          product.ID,
			SKU:         product.Code,
			Description: product.Description,
		}
	}

	return command.GetAllProductResult{
		PaginateResponse: command.PaginateResponse[[]command.ProductResult]{
			Data:  productResults,
			Total: total,
		},
	}, nil
}

func (h ProductHandler) HandleAddProduct(ctx context.Context, cmd *command.AddProductCommand) (command.AddProductResult, error) {
	product := &model.Product{
		Code:        cmd.SKU,
		Description: cmd.Description,
	}

	productID, err := h.ProductService.Create(ctx, product)
	if err != nil {
		return command.AddProductResult{}, err
	}

	sendData := event.ProductAddEvent{
		ID:     productID,
		UserID: cmd.UserID,
	}

	payload, err := json.Marshal(sendData)
	if err != nil {
		return command.AddProductResult{}, err
	}

	err = h.Publisher.Publish(constant.ADD_PRODUCT_EVENT, &message.Message{
		Payload: payload,
	})

	if err != nil {
		return command.AddProductResult{}, err
	}

	return command.AddProductResult{
		ID: productID,
	}, nil
}

func (h ProductHandler) HandleAddUpdate(ctx context.Context, cmd *command.UpdateProductCommand) (command.UpdateProductResult, error) {
	product := &model.Product{
		Code:        cmd.SKU,
		Description: cmd.Description,
	}

	err := h.ProductService.Update(ctx, product)
	if err != nil {
		return command.UpdateProductResult{}, err
	}

	return command.UpdateProductResult{
		ID: cmd.ID,
	}, nil
}

func (h ProductHandler) HandleSKUExist(ctx context.Context, cmd *command.GetSKUExistsCommand) (command.GetSKUExistsResult, error) {
	var exists bool
	exists, err := h.ProductService.HasSku(ctx, cmd.SKU, cmd.Field)

	if err != nil {
		return command.GetSKUExistsResult{}, err
	}

	return command.GetSKUExistsResult{
		Exists: exists,
	}, nil
}
