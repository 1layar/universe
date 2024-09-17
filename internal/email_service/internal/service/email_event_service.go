package service

import (
	"context"

	"github.com/1layar/universe/internal/email_service/model"
)

type EmailEventService interface {
	CreateEmailEvent(ctx context.Context, event *model.EmailEvent) error
	GetEmailEventByID(ctx context.Context, id int) (*model.EmailEvent, error)
	UpdateEmailEvent(ctx context.Context, event *model.EmailEvent) error
	DeleteEmailEvent(ctx context.Context, id int) error
}
