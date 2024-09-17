package dto

import (
	"github.com/1layar/universe/internal/auth_service/model"
)

type (
	CreateSessionDto struct {
		UserID    int
		IP        string
		UserAgent string
		Kind      model.SessionKind
		Retry     int
	}
)
