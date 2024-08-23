package service

import (
	"context"

	"github.com/1layar/universe/pkg/email_service/internal/repo"
	"github.com/1layar/universe/pkg/email_service/model"
)

type emailEventService struct {
	repo *repo.EmailEventRepository
}

// define const interface here
var _ EmailEventService = (*emailEventService)(nil)

func NewEmailEventService(repo *repo.EmailEventRepository) *emailEventService {
	return &emailEventService{repo: repo}
}

func (s *emailEventService) CreateEmailEvent(ctx context.Context, event *model.EmailEvent) error {
	return s.repo.Create(ctx, event)
}

func (s *emailEventService) GetEmailEventByID(ctx context.Context, id int) (*model.EmailEvent, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *emailEventService) UpdateEmailEvent(ctx context.Context, event *model.EmailEvent) error {
	return s.repo.Update(ctx, event)
}

func (s *emailEventService) DeleteEmailEvent(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
