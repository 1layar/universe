package handler

import (
	"context"

	"github.com/1layar/universe/src/auth_service/internal/service"
	"github.com/1layar/universe/src/shared/command"
	"github.com/1layar/universe/src/shared/dto"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"go.uber.org/fx"
)

type AuthHandler struct {
	fx.In
	AuthenticService *service.AuthenticService
	AuthorizeService *service.AuthorizeService
	CommandBus       *cqrs.CommandBus
	Publisher        *amqp.Publisher
	Subscriber       *amqp.Subscriber
	CommandProcessor *cqrs.CommandProcessor
}

func (h AuthHandler) GetHandlerName(name string) string {
	return "auth_service." + name
}

func (h AuthHandler) HandleLogin(ctx context.Context, cmd *command.LoginCommand) (command.LoginResult, error) {
	user, err := h.AuthenticService.Login(ctx, *cmd)
	if err != nil {
		return command.LoginResult{}, err
	}

	return *user, nil
}

func (h AuthHandler) HandleRegister(ctx context.Context, cmd *command.RegisterCommand) (command.RegisterResult, error) {
	user, err := h.AuthenticService.Register(ctx, *cmd)
	if err != nil {
		return command.RegisterResult{}, err
	}

	return *user, nil
}

func (h AuthHandler) HandleJwtToUser(ctx context.Context, cmd *command.JwtToUserCommand) (command.JwtToUserResponse, error) {
	user, err := h.AuthenticService.JwtToUser(ctx, *cmd)
	if err != nil {
		return command.JwtToUserResponse{}, err
	}

	return *user, nil
}

func (h AuthHandler) HandleCheckAccess(ctx context.Context, cmd *command.CheckAccessCommand) (command.CheckAccessResponse, error) {
	allow, err := h.AuthorizeService.CheckPermission(dto.PermissionCheckDTO{
		UserId:     cmd.UserId,
		Role:       "user",
		Permission: cmd.Permission,
		Subject:    cmd.Subject,
		SubjectID:  cmd.SubjectId,
	})
	if err != nil {
		return command.CheckAccessResponse{}, err
	}

	return command.CheckAccessResponse{
		Allow: allow,
	}, nil
}
