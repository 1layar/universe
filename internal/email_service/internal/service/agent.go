package service

import (
	"context"

	"github.com/1layar/universe/internal/email_service/model"
)

type Agent interface {
	SendEmail(ctx context.Context, email model.EmailMessage) error
}
