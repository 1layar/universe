package controller

import (
	"context"

	"github.com/1layar/universe/internal/ppob_service/service"
	"github.com/1layar/universe/pkg/shared/command"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"go.uber.org/fx"
)

type TypeHandler struct {
	fx.In
	TypeService      *service.ProductTypeService
	CommandBus       *cqrs.CommandBus
	Publisher        *amqp.Publisher
	Subscriber       *amqp.Subscriber
	CommandProcessor *cqrs.CommandProcessor
}

func (h TypeHandler) GetHandlerName(name string) string {
	return "ppob.type." + name
}

func (h TypeHandler) HandleGetAllType(ctx context.Context, cmd *command.PpobGetTypeCommand) (command.PpobGetTypeResult, error) {

	types, total, err := h.TypeService.GetAllPaginate(ctx,
		cmd.Page, cmd.Limit,
		[]string{},
		map[string]any{},
	)

	if err != nil {
		return command.PpobGetTypeResult{}, err
	}

	typeResults := make([]command.PpobTypeResult, len(types))

	for i, type_ := range types {
		typeResults[i] = command.PpobTypeResult{
			ID:   type_.ID,
			Name: type_.Name,
		}
	}

	return command.PpobGetTypeResult{
		PaginateResponse: command.PaginateResponse[[]command.PpobTypeResult]{
			Data:  typeResults,
			Total: total,
		},
	}, nil
}
