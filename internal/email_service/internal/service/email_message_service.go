package service

import (
	"context"

	"github.com/1layar/universe/pkg/email_service/model"
)

type EmailMessageService interface {
	CreateEmailMessage(ctx context.Context, message *model.EmailMessage) error
	GetEmailMessageByID(ctx context.Context, id int) (*model.EmailMessage, error)
	UpdateEmailMessage(ctx context.Context, message *model.EmailMessage) error
	DeleteEmailMessage(ctx context.Context, id int) error
}
