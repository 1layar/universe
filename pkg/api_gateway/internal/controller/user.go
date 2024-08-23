package controller

import (
	"context"
	"fmt"
	"time"

	"github.com/goccy/go-json"

	"github.com/1layar/universe/pkg/api_gateway/internal/middleware"
	"github.com/1layar/universe/pkg/shared/command"
	"github.com/1layar/universe/pkg/shared/constant"
	"github.com/1layar/universe/pkg/shared/dto"
	"github.com/1layar/universe/pkg/shared/transport"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type UserController struct {
	fx.In
	CommandBus *cqrs.CommandBus
	Publisher  *amqp.Publisher
	Subscriber *amqp.Subscriber
	Validator  *dto.XValidator
	Route      fiber.Router `name:"api-v1"`
}

func RegisterUser(c UserController, authM middleware.AuthenticMiddleware) {
	c.Route.Use("/users", authM.New())
	c.Route.Get("/users", c.GetUsers)
	c.Route.Get("/users/:id", c.GetUser)
	c.Route.Post("/users", c.AddUser)
	c.Route.Put("/users/:id", c.UpdateUser)
	c.Route.Delete("/users/:id", c.DeleteUser)
}

// Accounts delete a accounts
//
//	@Summary      Delete accounts
//	@Description  delete accounts by id
//	@Tags         accounts
//	@Accept       json
//	@Produce      json
//	@Param        id   path      int  true  "Account ID"
//	@Success      200  {object}  dto.GlobalHandlerResp[command.DeleteUserResult]
//	@Router       /api/v1/users/:id [delete]
//	@Security JWT
func (c *UserController) DeleteUser(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"source": "gateway_service",
			"msg":    err.Error(),
		})
	}

	deleteUserDto := &dto.FindUser{
		ID: id,
	}

	replyCh, cancel, err := transport.GetReqRep[command.DeleteUserResult](
		constant.DELETE_USER_CMD,
		c.Publisher,
		c.Subscriber,
		context.Background(),
		c.CommandBus,
		&command.DeleteUserCommand{
			Id: deleteUserDto.ID,
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

// Accounts update a accounts
//
//	@Summary      Get accounts
//	@Description  get accounts
//	@Tags         accounts
//	@Accept       json
//	@Produce      json
//	@Param        id   path      int  true  "Account ID"
//	@Param request body dto.UpdateUser true "query params"
//	@Success      200  {object}  dto.GlobalHandlerResp[command.UpdateUserResult]
//	@Router       /api/v1/users/:id [put]
//	@Security JWT
func (c *UserController) UpdateUser(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"source": "gateway_service",
			"msg":    err.Error(),
		})
	}

	updateUserDto := &dto.UpdateUser{
		Id: id,
	}
	requestBody := ctx.Body()
	if err := json.Unmarshal(requestBody, updateUserDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"source": "gateway_service",
			"msg":    err.Error(),
		})
	}
	if err := c.Validator.Validate(updateUserDto); len(err) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"source": "gateway_service",
			"rules":  err,
			"msg":    "invalid request",
		})
	}
	replyCh, cancel, err := transport.GetReqRep[command.UpdateUserResult](
		constant.UPDATE_USER_CMD,
		c.Publisher,
		c.Subscriber,
		context.Background(),
		c.CommandBus,
		&command.UpdateUserCommand{
			ID:       updateUserDto.Id,
			Name:     updateUserDto.Username,
			Email:    updateUserDto.Email,
			Password: updateUserDto.Password,
			Role:     updateUserDto.Role,
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

// Accounts get a accounts
//
//	@Summary      Get accounts
//	@Description  get accounts
//	@Tags         accounts
//	@Accept       json
//	@Produce      json
//	@Param        id   path      int  true  "Account ID"
//	@Success      200  {object}  dto.GlobalHandlerResp[command.GetUserResult]
//	@Router       /api/v1/users/:id [get]
//	@Security JWT
func (c *UserController) GetUser(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"source": "gateway_service",
			"msg":    err.Error(),
		})
	}

	searchUserDto := &dto.FindUser{
		ID: id,
	}

	replyCh, cancel, err := transport.GetReqRep[command.GetUserResult](
		constant.GET_USER_CMD,
		c.Publisher,
		c.Subscriber,
		context.Background(),
		c.CommandBus,
		&command.GetUserCommand{
			ID: searchUserDto.ID,
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

// Accounts lists all existing accounts
//
//	@Summary      List accounts
//	@Description  get accounts
//	@Tags         accounts
//	@Accept       json
//	@Produce      json
//	@Param        username    query     string  false  "user search by username"
//	@Success      200  {object}  dto.GlobalHandlerResp[command.SearchUserResult]
//	@Router       /api/v1/users [get]
//	@Security JWT
func (c *UserController) GetUsers(ctx *fiber.Ctx) error {
	searchUserDto := &dto.SearchUser{}
	if err := ctx.QueryParser(searchUserDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"source": "gateway_service",
			"msg":    err.Error(),
		})
	}

	replyCh, cancel, err := transport.GetReqRep[command.SearchUserResult](
		constant.SEARCH_USER_CMD,
		c.Publisher,
		c.Subscriber,
		context.Background(),
		c.CommandBus,
		&command.SearchUserCommand{
			Username: searchUserDto.Username,
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

// Accounts add accounts
//
//	@Summary      Add accounts
//	@Description  add new accounts
//	@Tags         accounts
//	@Accept       json
//	@Produce      json
//	@Param request body dto.AddUser true "query params"
//	@Success      200  {object}  dto.GlobalHandlerResp[command.AddUserResult]
//	@Security JWT
//	@Router       /api/v1/users [post]
func (c *UserController) AddUser(ctx *fiber.Ctx) error {
	addUserDto := &dto.AddUser{}
	requestBody := ctx.Body()
	if err := json.Unmarshal(requestBody, addUserDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"source": "gateway_service",
			"msg":    err.Error(),
		})
	}

	if err := c.Validator.Validate(addUserDto); len(err) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"source": "gateway_service",
			"rules":  err,
			"msg":    "invalid request",
		})
	}

	replyCh, cancel, err := transport.GetReqRep[command.AddUserResult](
		constant.ADD_USER_CMD,
		c.Publisher,
		c.Subscriber,
		context.Background(),
		c.CommandBus,
		&command.AddUserCommand{
			Username: addUserDto.Username,
			Email:    addUserDto.Email,
			Password: addUserDto.Password,
			Role:     addUserDto.Role,
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
