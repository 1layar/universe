package controller

import (
	"context"
	"fmt"
	"time"

	"github.com/1layar/universe/src/shared/command"
	"github.com/1layar/universe/src/shared/constant"
	"github.com/1layar/universe/src/shared/dto"
	"github.com/1layar/universe/src/shared/transport"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type AuthController struct {
	fx.In
	CommandBus *cqrs.CommandBus
	Publisher  *amqp.Publisher
	Subscriber *amqp.Subscriber
	Validator  *dto.XValidator
	Route      fiber.Router `name:"api-v1"`
}

func RegisterAuth(c AuthController) {
	c.Route.Post("/auth/login", c.Login)
	c.Route.Post("/auth/register", c.Register)
}

// Login users
//
//	@Summary      Login accounts
//	@Description  get accounts
//	@Tags         auth
//	@Accept       json
//	@Produce      json
//	@Param request body dto.LoginDTO true "query params"
//	@Success      200  {object}  dto.GlobalHandlerResp[command.LoginResult]
//	@Router       /api/v1/auth/login [post]
func (c *AuthController) Login(ctx *fiber.Ctx) error {
	credential := &dto.LoginDTO{}
	requestBody := ctx.Body()
	if err := json.Unmarshal(requestBody, credential); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"source": "gateway_service",
			"msg":    err.Error(),
		})
	}

	if err := c.Validator.Validate(credential); len(err) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"source": "gateway_service",
			"rules":  err,
			"msg":    "invalid request",
		})
	}
	replyCh, cancel, err := transport.GetReqRep[command.LoginResult](
		constant.AUTH_LOGIN_CMD,
		c.Publisher,
		c.Subscriber,
		context.Background(),
		c.CommandBus,
		&command.LoginCommand{
			Email:    credential.Email,
			Password: credential.Password,
			Ip:       ctx.IP(),
			Ua:       ctx.Get("User-Agent", ""),
			DevId:    ctx.Get("X-Request-ID", ""),
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

// Register users
//
//	@Summary      Register accounts
//	@Description  get accounts
//	@Tags         auth
//	@Accept       json
//	@Produce      json
//	@Param request body dto.RegisterDTO true "query params"
//	@Success      200  {object}  dto.GlobalHandlerResp[command.RegisterResult]
//	@Router       /api/v1/auth/register [post]
func (c *AuthController) Register(ctx *fiber.Ctx) error {
	credential := &dto.RegisterDTO{}
	requestBody := ctx.Body()
	if err := json.Unmarshal(requestBody, credential); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"source": "gateway_service",
			"msg":    err.Error(),
		})
	}

	if err := c.Validator.Validate(credential); len(err) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"source": "gateway_service",
			"rules":  err,
			"msg":    "invalid request",
		})
	}

	replyCh, cancel, err := transport.GetReqRep[command.RegisterResult](
		constant.AUTH_REGISTER_CMD,
		c.Publisher,
		c.Subscriber,
		context.Background(),
		c.CommandBus,
		&command.RegisterCommand{
			Email:    credential.Email,
			Password: credential.Password,
			Username: credential.Username,
			Ip:       ctx.IP(),
			Ua:       ctx.Get("User-Agent", ""),
			DevId:    ctx.Get("X-Request-ID", ""),
		},
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

		return ctx.JSON(dto.NewApiResp(reply.HandlerResult))
	case <-time.After(time.Second * time.Duration(5)):
		return ctx.Status(fiber.StatusBadGateway).JSON(map[string]any{
			"source": "gateway_service",
			"msg":    "timeout",
		})
	}
}
