package controller

import (
	"github.com/1layar/universe/src/product_catalog_service/internal/service"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"go.uber.org/fx"
)

type CategoryHandler struct {
	fx.In
	ProductService   *service.ProductService
	CommandBus       *cqrs.CommandBus
	Publisher        *amqp.Publisher
	Subscriber       *amqp.Subscriber
	CommandProcessor *cqrs.CommandProcessor
}

func (h CategoryHandler) GetHandlerName(name string) string {
	return "product_catalog_service.product." + name
}
