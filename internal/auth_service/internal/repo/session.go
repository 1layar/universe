package repo

import (
	"context"

	"github.com/1layar/universe/internal/auth_service/model"
	"github.com/1layar/universe/internal/auth_service/model/dto"
	"github.com/uptrace/bun"
)

type SessionRepository struct {
	db *bun.DB
}

func NewSession(db *bun.DB) *SessionRepository {
	return &SessionRepository{db: db}
}

func (r *SessionRepository) CreateSession(ctx context.Context, data dto.CreateSessionDto) (*model.Session, error) {
	session := &model.Session{
		UserID:    data.UserID,
		IP:        data.IP,
		UserAgent: data.UserAgent,
		Kind:      data.Kind,
		Retry:     data.Retry,
	}

	if _, err := r.db.NewInsert().Model(session).Exec(ctx); err != nil {
		return nil, err
	}

	return session, nil
}
