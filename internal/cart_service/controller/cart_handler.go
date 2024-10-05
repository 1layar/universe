package controller

import (
	"context"

	"github.com/1layar/universe/internal/cart_service/service"
	"github.com/1layar/universe/pkg/shared/command"
	"github.com/1layar/universe/pkg/shared/dto"
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

func (h CartHandler) AddItem(ctx context.Context, cmd *command.CartAddItemCommand) (command.CartAddItemCommandResult, error) {
	session, err := h.CartService.AddItem(ctx, cmd.SessionID, cmd.ProductID, cmd.ProductID, cmd.Source)

	if err != nil {
		return command.CartAddItemCommandResult{
			Success: false,
		}, err
	}

	return command.CartAddItemCommandResult{
		SessionID: session,
		Success:   true,
	}, nil
}

func (h CartHandler) Empty(ctx context.Context, cmd *command.CartEmptyCommand) (command.CartEmptyCommandResult, error) {
	err := h.CartService.EmptyCart(ctx, cmd.SessionId)

	if err != nil {
		return command.CartEmptyCommandResult{
			Success: false,
		}, err
	}

	return command.CartEmptyCommandResult{
		Success: true,
	}, nil
}

func (h CartHandler) RemoveItem(ctx context.Context, cmd *command.CartRemoveItemCommand) (command.CartRemoveItemCommandResult, error) {
	err := h.CartService.RemoveItem(ctx, cmd.SessionID, cmd.ProductID, cmd.Source)

	if err != nil {
		return command.CartRemoveItemCommandResult{
			Success: false,
		}, err
	}

	return command.CartRemoveItemCommandResult{
		Success: true,
	}, nil
}

func (h CartHandler) GetCart(ctx context.Context, cmd *command.CartGetCartCommand) (command.CartGetCartCommandResult, error) {
	data, err := h.CartService.GetCart(ctx, cmd.SessionId)
	resp := command.CartGetCartCommandResult{}

	if err != nil {
		return resp, err
	}

	cartItem := make([]dto.CartItem, len(data))

	for i := 0; i < len(data); i++ {
		cartItem[i] = dto.CartItem{
			ProductID: data[i].ProductID,
			Quantity:  data[i].Quantity,
		}
	}

	resp.Cart = &dto.CartResp{
		SessionId: cmd.SessionId,
		Cart: dto.Cart{
			Items: cartItem,
		},
	}

	return resp, err
}
