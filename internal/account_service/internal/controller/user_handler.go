package controller

import (
	"context"

	"github.com/1layar/universe/internal/account_service/internal/service"
	"github.com/1layar/universe/internal/account_service/model/dto"
	"github.com/1layar/universe/pkg/shared/command"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"go.uber.org/fx"
)

type UserHandler struct {
	fx.In
	UserService      *service.User
	CommandBus       *cqrs.CommandBus
	Publisher        *amqp.Publisher
	Subscriber       *amqp.Subscriber
	CommandProcessor *cqrs.CommandProcessor
}

func (h UserHandler) GetHandlerName(name string) string {
	return "account_service.user." + name
}

func (h UserHandler) HandleAddUser(ctx context.Context, cmd *command.AddUserCommand) (command.AddUserResult, error) {
	user, err := h.UserService.AddUser(ctx, dto.CreateUserDto{
		Email:    cmd.Email,
		Username: cmd.Username,
		Password: cmd.Password,
		Role:     cmd.Role,
	})
	if err != nil {
		return command.AddUserResult{}, err
	}

	return command.AddUserResult{
		Id:    user.ID,
		Name:  user.Username,
		Email: user.Email,
	}, nil
}

func (h UserHandler) HandleUsernameExist(ctx context.Context, cmd *command.GetUsernameExistsCommand) (command.GetUsernameExistsResult, error) {
	var exists bool
	if len(cmd.Field) > 0 {
		exists = h.UserService.HasUsername(ctx, cmd.Username, cmd.Field)
	} else {
		exists = h.UserService.HasUsername(ctx, cmd.Username)
	}

	return command.GetUsernameExistsResult{
		Exists: exists,
	}, nil
}

func (h UserHandler) HandleEmailExist(ctx context.Context, cmd *command.GetEmailExistsCommand) (command.GetEmailExistsResult, error) {
	var exists bool

	if len(cmd.Field) > 0 {
		exists = h.UserService.HasEmail(ctx, cmd.Email, cmd.Field)
	} else {
		exists = h.UserService.HasEmail(ctx, cmd.Email)
	}

	return command.GetEmailExistsResult{
		Exists: exists,
	}, nil
}

func (h UserHandler) HandleSearchUser(ctx context.Context, cmd *command.SearchUserCommand) (command.SearchUserResult, error) {
	result, err := h.UserService.SearchUser(ctx, dto.SearchUserDto{
		Username: cmd.Username,
	})

	if err != nil {
		return command.SearchUserResult{Users: make([]command.AddUserResult, 0)}, err
	}

	users := make([]command.AddUserResult, len(result))

	for i, v := range result {
		users[i] = command.AddUserResult{
			Id:    v.ID,
			Name:  v.Username,
			Email: v.Email,
		}
	}

	return command.SearchUserResult{
		Users: users,
	}, nil
}

func (h UserHandler) HandleUpdateUser(ctx context.Context, cmd *command.UpdateUserCommand) (command.UpdateUserResult, error) {
	result, err := h.UserService.Update(ctx, dto.UpdateUser{
		Id:       cmd.ID,
		Username: cmd.Name,
		Email:    cmd.Email,
		Password: cmd.Password,
		Role:     cmd.Role,
	})

	if err != nil {
		return command.UpdateUserResult{}, err
	}

	return command.UpdateUserResult{
		Id:    result.ID,
		Name:  result.Username,
		Email: result.Email,
	}, nil
}

func (h UserHandler) HandleGetUser(ctx context.Context, cmd *command.GetUserCommand) (command.GetUserResult, error) {
	var resp command.GetUserResult

	if cmd.ID != 0 {
		result, err := h.UserService.GetUserByID(ctx, cmd.ID)

		if err != nil {
			return command.GetUserResult{}, err
		}

		resp = command.GetUserResult{
			Id:       result.ID,
			Name:     result.Username,
			Email:    result.Email,
			Password: result.Password,
		}
	}

	if cmd.Email != "" {
		result, err := h.UserService.GetUserByEmail(ctx, cmd.Email)

		if err != nil {
			return command.GetUserResult{}, err
		}

		resp = command.GetUserResult{
			Id:       result.ID,
			Name:     result.Username,
			Email:    result.Email,
			Password: result.Password,
		}
	}

	return resp, nil
}

func (h UserHandler) HandleDeleteUser(ctx context.Context, cmd *command.DeleteUserCommand) (command.DeleteUserResult, error) {
	result, err := h.UserService.DeleteByID(ctx, cmd.Id)

	if err != nil {
		return command.DeleteUserResult{}, err
	}

	return command.DeleteUserResult{
		Id:        result.ID,
		HasDelete: true,
	}, nil
}
