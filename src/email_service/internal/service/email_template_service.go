package service

import (
	"context"

	"github.com/1layar/universe/src/email_service/model"
)

type EmailTemplateService interface {
	CreateEmailTemplate(ctx context.Context, template *model.EmailTemplate) error
	GetEmailTemplateByID(ctx context.Context, id int) (*model.EmailTemplate, error)
	UpdateEmailTemplate(ctx context.Context, template *model.EmailTemplate) error
	DeleteEmailTemplate(ctx context.Context, id int) error
}
