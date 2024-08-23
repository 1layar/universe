package service

import (
	"context"

	"github.com/1layar/universe/src/auth_service/internal/repo"
	"github.com/1layar/universe/src/auth_service/model"
	"github.com/1layar/universe/src/auth_service/model/dto"
)

type SessionService struct {
	sessionRepo *repo.SessionRepository
}

func NewSessionService(sessionRepo *repo.SessionRepository) *SessionService {
	return &SessionService{
		sessionRepo: sessionRepo,
	}
}

func (s *SessionService) CreateSession(ctx context.Context, data dto.CreateSessionDto) (*model.Session, error) {
	return s.sessionRepo.CreateSession(ctx, data)
}
