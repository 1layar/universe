package middleware

import (
	"context"
	"time"

	"github.com/1layar/universe/src/shared/command"
	"github.com/1layar/universe/src/shared/constant"
	"github.com/1layar/universe/src/shared/transport"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type AuthorizeMiddleware struct {
	fx.In
	CommandBus *cqrs.CommandBus
	Publisher  *amqp.Publisher
	Subscriber *amqp.Subscriber
}

func (m *AuthorizeMiddleware) New(group string, paramName string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Locals("user").(command.JwtToUserResponse)

		if token.User.Id == 0 {
			return c.Status(fiber.StatusUnauthorized).JSON(map[string]any{
				"source": "gateway_service",
				"msg":    "unauthorized",
			})
		}

		id, err := c.ParamsInt(paramName)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(map[string]any{
				"source": "gateway_service",
				"msg":    err.Error(),
			})
		}

		var permission string

		switch c.Method() {
		case fiber.MethodGet:
			permission = "reader"
		case fiber.MethodPost, fiber.MethodPut, fiber.MethodDelete:
			permission = "writer"
		default:
			permission = "reader"
		}

		replyCh, cancel, err := transport.GetReqRep[command.CheckAccessResponse](
			constant.AUTH_CHECK_ACCESS_CMD,
			m.Publisher,
			m.Subscriber,
			context.Background(),
			m.CommandBus,
			&command.CheckAccessCommand{
				UserId:     token.User.Id,
				Permission: permission,
				Subject:    group,
				SubjectId:  id,
			},
		)
		if err != nil {
			return err
		}

		defer cancel()

		select {
		case reply := <-replyCh:
			if err := reply.Error; err != nil {
				return c.Status(fiber.StatusBadGateway).JSON(map[string]any{
					"source": "account_service",
					"msg":    err.Error(),
				})
			}
			// Store the user ID from the token in the context for further processing
			if !reply.HandlerResult.Allow {
				return c.Status(fiber.StatusForbidden).JSON(map[string]any{
					"source": "gateway_service",
					"msg":    "forbidden",
				})
			}

			// Pass control to the next handler
			return c.Next()
		case <-time.After(time.Second * time.Duration(15)):
			return c.Status(fiber.StatusBadGateway).JSON(map[string]any{
				"source": "gateway_service",
				"msg":    "timeout",
			})
		}

	}
}
