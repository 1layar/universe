package service

import (
	"context"

	"github.com/1layar/universe/internal/email_service/dto"
	"github.com/1layar/universe/internal/email_service/model"
	"github.com/hibiken/asynq"
)

type IEmailManager interface {
	ProcessEmailDeliveryTask(ctx context.Context, t *asynq.Task) error
	Compose(ctx context.Context, compose dto.Compose) (*model.EmailEvent, error)
}
