package controller

import (
	"context"
	"time"

	"github.com/1layar/universe/internal/api_gateway/app/appconfig"
	"github.com/1layar/universe/pkg/shared/command"
	"github.com/1layar/universe/pkg/shared/constant"
	"github.com/1layar/universe/pkg/shared/dto"
	"github.com/1layar/universe/pkg/shared/transport"
	"github.com/1layar/universe/pkg/shared/utils"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"go.uber.org/fx"
)

type PPOBController struct {
	fx.In
	CommandBus *cqrs.CommandBus
	Config     *appconfig.Config
	Publisher  *amqp.Publisher
	Subscriber *amqp.Subscriber
	Validator  *dto.XValidator
	Route      fiber.Router `name:"api-v1"`
}

func RegisterPPOB(c PPOBController) {
	ppob := c.Route.Group("/ppob")
	ppob.Get("/products", c.GetProducts)
	ppob.Get("/types", c.GetProductTypes)
	ppob.Get("/iak-sync", c.SyncIak)
}

// Product
//
//	@Summary      Get All Product
//	@Description  get all product with pagination
//	@Tags         ppob
//	@Accept       json
//	@Produce      json
//	@Param request query dto.GetAllProductPaginateReqDTO true "query params"
//	@Success      201  {object}  dto.GlobalHandlerResp[dto.GetAllProductPaginateResDTO]
//	@Router       /api/v1/ppob/products [get]
func (c *PPOBController) GetProducts(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// Product Type
//
//	@Summary      Get All Product
//	@Description  get all product with pagination
//	@Tags         ppob
//	@Accept       json
//	@Produce      json
//	@Param request query dto.PaginateReqDto true "query params"
//	@Success      201  {object}  dto.GlobalHandlerResp[dto.GetAllPpobTypeResDTO]
//	@Router       /api/v1/ppob/types [get]
func (c *PPOBController) GetProductTypes(ctx *fiber.Ctx) error {
	queryType := &dto.PaginateReqDto{}

	if err := ctx.QueryParser(queryType); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"source": "gateway_service",
			"msg":    err.Error(),
		})
	}

	if err := c.Validator.Validate(queryType); len(err) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"source": "gateway_service",
			"rules":  err,
			"msg":    "invalid request",
		})
	}

	getAllCmd := &command.PpobGetTypeCommand{}

	err := copier.Copy(getAllCmd, queryType)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]any{
			"source": "gateway_service",
			"msg":    err.Error(),
		})
	}

	replyCh, cancel, err := transport.GetReqRep[command.PpobGetTypeResult](
		constant.PPOB_GET_ALL_PRODUCT_TYPE_CMD,
		c.Publisher,
		c.Subscriber,
		context.Background(),
		c.CommandBus,
		getAllCmd,
	)
	if err != nil {
		return err
	}

	defer cancel()

	select {
	case reply := <-replyCh:
		if err := reply.Error; err != nil {
			return ctx.Status(fiber.StatusBadGateway).JSON(map[string]any{
				"source": "ppob_service",
				"msg":    err.Error(),
			})
		}

		result := &dto.GetAllPpobTypeResDTO{}

		err = copier.CopyWithOption(result, reply.HandlerResult, copier.Option{
			DeepCopy: true,
		})

		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]any{
				"source": "gateway_service",
				"msg":    err.Error(),
			})
		}

		result.Page = utils.CalculateTotalPages(int(result.Total), queryType.Limit)
		result.Limit = queryType.Limit

		return ctx.JSON(dto.NewApiResp(result))
	case <-time.After(time.Second * time.Duration(5)):
		return ctx.Status(fiber.StatusBadGateway).JSON(map[string]any{
			"source": "gateway_service",
			"msg":    "timeout",
		})
	}
}

// Sync IAK
//
//	@Summary      Get All Product
//	@Description  get all product with pagination
//	@Tags         ppob
//	@Accept       json
//	@Produce      json
//	@Param request query dto.GetAllProductPaginateReqDTO true "query params"
//	@Success      201  {object}  dto.GlobalHandlerResp[dto.GetAllProductPaginateResDTO]
//	@Router       /api/v1/ppob/sync-iak [get]
func (c *PPOBController) SyncIak(ctx *fiber.Ctx) error {
	panic("unimplemented")
}
