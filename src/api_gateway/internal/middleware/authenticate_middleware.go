package middleware

import (
	"context"
	"strings"
	"time"

	"github.com/1layar/universe/src/shared/command"
	"github.com/1layar/universe/src/shared/constant"
	"github.com/1layar/universe/src/shared/transport"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type AuthenticMiddleware struct {
	fx.In
	CommandBus *cqrs.CommandBus
	Publisher  *amqp.Publisher
	Subscriber *amqp.Subscriber
}

func (m *AuthenticMiddleware) New() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// get bearer token from header
		// Get the Authorization header
		authHeader := c.Get("Authorization")

		// Check if the Authorization header is missing
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Authorization header is missing",
			})
		}

		// Check if the Authorization header is in the correct format (Bearer <token>)
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid or missing Bearer token",
			})
		}

		// Extract the token from the header
		tokenString := parts[1]

		replyCh, cancel, err := transport.GetReqRep[command.JwtToUserResponse](
			constant.AUTH_JWT_TO_USER_CMD,
			m.Publisher,
			m.Subscriber,
			context.Background(),
			m.CommandBus,
			&command.JwtToUserCommand{
				Jwt: tokenString,
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
			c.Locals("user", reply.HandlerResult)

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
