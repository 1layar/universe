package controller

import (
	"github.com/1layar/universe/internal/cart_service/service"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"go.uber.org/fx"
)

type CartHandler struct {
	fx.In
	CartService      *service.CartService
	CommandBus       *cqrs.CommandBus
	Publisher        *amqp.Publisher
	Subscriber       *amqp.Subscriber
	CommandProcessor *cqrs.CommandProcessor
}

func (h CartHandler) GetHandlerName(name string) string {
	return "cart." + name
}
