package service

import (
	"context"

	"github.com/1layar/universe/pkg/email_service/internal/repo"
	"github.com/1layar/universe/pkg/email_service/model"
)

type emailTemplateService struct {
	repo *repo.EmailTemplateRepository
}

func NewEmailTemplateService(repo *repo.EmailTemplateRepository) *emailTemplateService {
	return &emailTemplateService{repo: repo}
}

func (s *emailTemplateService) CreateEmailTemplate(ctx context.Context, template *model.EmailTemplate) error {
	return s.repo.Create(ctx, template)
}

func (s *emailTemplateService) GetEmailTemplateByID(ctx context.Context, id int) (*model.EmailTemplate, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *emailTemplateService) GetEmailTemplateByCode(ctx context.Context, code string) (*model.EmailTemplate, error) {
	return s.repo.GetByCode(ctx, code)
}

func (s *emailTemplateService) UpdateEmailTemplate(ctx context.Context, template *model.EmailTemplate) error {
	return s.repo.Update(ctx, template)
}

func (s *emailTemplateService) DeleteEmailTemplate(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
