package service

import (
	"context"

	"github.com/1layar/merasa/backend/src/email_service/model"
)

type AccountService interface {
	CreateAccount(ctx context.Context, account *model.Account) error
	GetAccountByID(ctx context.Context, id int) (*model.Account, error)
	GetAccountByCode(ctx context.Context, code string) (*model.Account, error)
	UpdateAccount(ctx context.Context, account *model.Account) error
	DeleteAccount(ctx context.Context, id int) error
}
