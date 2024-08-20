package service

import (
	"context"

	"github.com/1layar/merasa/backend/src/email_service/internal/repo"
	"github.com/1layar/merasa/backend/src/email_service/model"
)

type emailMessageService struct {
	repo *repo.EmailMessageRepository
}

// define const interface here
var _ EmailMessageService = (*emailMessageService)(nil)

func NewEmailMessageService(repo *repo.EmailMessageRepository) *emailMessageService {
	return &emailMessageService{repo: repo}
}

func (s *emailMessageService) CreateEmailMessage(ctx context.Context, message *model.EmailMessage) error {
	return s.repo.Create(ctx, message)
}

func (s *emailMessageService) GetEmailMessageByID(ctx context.Context, id int) (*model.EmailMessage, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *emailMessageService) UpdateEmailMessage(ctx context.Context, message *model.EmailMessage) error {
	return s.repo.Update(ctx, message)
}

func (s *emailMessageService) DeleteEmailMessage(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
