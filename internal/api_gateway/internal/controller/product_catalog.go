package controller

import (
	"context"
	"fmt"
	"time"

	"github.com/1layar/universe/pkg/api_gateway/internal/app/appconfig"
	"github.com/1layar/universe/pkg/api_gateway/internal/middleware"
	"github.com/1layar/universe/pkg/shared/command"
	"github.com/1layar/universe/pkg/shared/constant"
	"github.com/1layar/universe/pkg/shared/dto"
	"github.com/1layar/universe/pkg/shared/transport"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"go.uber.org/fx"
)

type ProductCatalogController struct {
	fx.In
	CommandBus *cqrs.CommandBus
	Config     *appconfig.Config
	Publisher  *amqp.Publisher
	Subscriber *amqp.Subscriber
	Validator  *dto.XValidator
	Route      fiber.Router `name:"api-v1"`
}

func RegisterProductCatalog(c ProductCatalogController, authM middleware.AuthenticMiddleware, autho middleware.AuthorizeMiddleware) {
	c.Route.Use("/products", authM.New())
	c.Route.Post("/products", c.AddProduct)
	c.Route.Get("/products", c.GetProducts)
	c.Route.Use("/products/:id", autho.New("product", "id"))
	c.Route.Put("/products/:id", c.UpdateProduct)
}

// Product
//
//	@Summary      Get All Product
//	@Description  get all product with pagination
//	@Tags         products
//	@Accept       json
//	@Produce      json
//	@Param request query dto.GetAllProductPaginateReqDTO true "query params"
//	@Success      201  {object}  dto.GlobalHandlerResp[dto.GetAllProductPaginateResDTO]
//	@Router       /api/v1/products [get]
//	@Security JWT
func (c *ProductCatalogController) GetProducts(ctx *fiber.Ctx) error {
	queryProduct := &dto.GetAllProductPaginateReqDTO{}
	if err := ctx.QueryParser(queryProduct); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"source": "gateway_service",
			"msg":    err.Error(),
		})
	}

	if err := c.Validator.Validate(queryProduct); len(err) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"source": "gateway_service",
			"rules":  err,
			"msg":    "invalid request",
		})
	}

	getAllCmd := &command.GetAllProductCommand{}

	err := copier.Copy(getAllCmd, queryProduct)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]any{
			"source": "gateway_service",
			"msg":    err.Error(),
		})
	}

	replyCh, cancel, err := transport.GetReqRep[command.GetAllProductResult](
		constant.GET_ALL_PRODUCT_CMD,
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
				"source": "account_service",
				"msg":    err.Error(),
			})
		}

		result := &dto.GetAllProductPaginateResDTO{}

		copier.CopyWithOption(result, reply.HandlerResult, copier.Option{
			DeepCopy: true,
		})

		result.Page = queryProduct.Page
		result.Limit = queryProduct.Limit

		return ctx.JSON(dto.NewApiResp(result))
	case <-time.After(time.Second * time.Duration(5)):
		return ctx.Status(fiber.StatusBadGateway).JSON(map[string]any{
			"source": "gateway_service",
			"msg":    "timeout",
		})
	}
}

// Product
//
//	@Summary      Create Product
//	@Description  add new product
//	@Tags         products
//	@Accept       json
//	@Produce      json
//	@Param request body dto.AddProductDTO true "query params"
//	@Success      201  {object}  dto.GlobalHandlerResp[command.AddProductResult]
//	@Router       /api/v1/products [post]
//	@Security JWT
func (c *ProductCatalogController) AddProduct(ctx *fiber.Ctx) error {
	token := ctx.Locals("user").(command.JwtToUserResponse)
	addProduct := &dto.AddProductDTO{}
	requestBody := ctx.Body()
	if err := json.Unmarshal(requestBody, addProduct); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"source": "gateway_service",
			"msg":    err.Error(),
		})
	}

	if err := c.Validator.Validate(addProduct); len(err) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"source": "gateway_service",
			"rules":  err,
			"msg":    "invalid request",
		})
	}

	replyCh, cancel, err := transport.GetReqRep[command.AddProductResult](
		constant.ADD_PRODUCT_CMD,
		c.Publisher,
		c.Subscriber,
		context.Background(),
		c.CommandBus,
		&command.AddProductCommand{
			Name:        addProduct.Name,
			SKU:         addProduct.SKU,
			Description: addProduct.Description,
			PictureURL:  addProduct.PictureURL,
			Quantity:    addProduct.Quantity,
			Price:       addProduct.Price,
			Categories:  addProduct.CategoryIDs,
			UserID:      token.User.Id,
		},
	)
	if err != nil {
		return err
	}

	defer cancel()

	select {
	case reply := <-replyCh:
		fmt.Println(reply)
		if err := reply.Error; err != nil {
			return ctx.Status(fiber.StatusBadGateway).JSON(map[string]any{
				"source": "account_service",
				"msg":    err.Error(),
			})
		}
		return ctx.JSON(dto.NewApiResp(reply.HandlerResult))
	case <-time.After(time.Second * time.Duration(5)):
		return ctx.Status(fiber.StatusBadGateway).JSON(map[string]any{
			"source": "gateway_service",
			"msg":    "timeout",
		})
	}
}

// Product
//
//	@Summary      Update Product
//	@Description  update product
//	@Tags         products
//	@Accept       json
//	@Produce      json
//	@Param id path int true "product id"
//	@Param request body dto.UpdateProductDTO true "query params"
//	@Success      201  {object}  dto.GlobalHandlerResp[command.UpdateProductResult]
//	@Router       /api/v1/products/{id} [put]
//	@Security JWT
func (c *ProductCatalogController) UpdateProduct(ctx *fiber.Ctx) error {
	token := ctx.Locals("user").(command.JwtToUserResponse)
	updateProduct := &dto.UpdateProductDTO{}
	requestBody := ctx.Body()
	if err := json.Unmarshal(requestBody, updateProduct); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"source": "gateway_service",
			"msg":    err.Error(),
		})
	}

	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"source": "gateway_service",
			"msg":    err.Error(),
		})
	}
	updateProduct.ID = id

	if err := c.Validator.Validate(updateProduct); len(err) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"source": "gateway_service",
			"rules":  err,
			"msg":    "invalid request",
		})
	}

	replyCh, cancel, err := transport.GetReqRep[command.UpdateProductCommand](
		constant.UPDATE_PRODUCT_CMD,
		c.Publisher,
		c.Subscriber,
		context.Background(),
		c.CommandBus,
		&command.UpdateProductCommand{
			ID:          updateProduct.ID,
			Name:        updateProduct.Name,
			SKU:         updateProduct.SKU,
			Description: updateProduct.Description,
			PictureURL:  updateProduct.PictureURL,
			Quantity:    updateProduct.Quantity,
			Price:       updateProduct.Price,
			Categories:  updateProduct.CategoryIDs,
			UserID:      token.User.Id,
		},
	)
	if err != nil {
		return err
	}

	defer cancel()

	select {
	case reply := <-replyCh:
		fmt.Println(reply)
		if err := reply.Error; err != nil {
			return ctx.Status(fiber.StatusBadGateway).JSON(map[string]any{
				"source": "account_service",
				"msg":    err.Error(),
			})
		}
		return ctx.JSON(dto.NewApiResp(reply.HandlerResult))
	case <-time.After(time.Second * time.Duration(5)):
		return ctx.Status(fiber.StatusBadGateway).JSON(map[string]any{
			"source": "gateway_service",
			"msg":    "timeout",
		})
	}
}
