package controller

import (
	"github.com/1layar/universe/internal/ppob_service/service"
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
	return "ppob.category." + name
}
