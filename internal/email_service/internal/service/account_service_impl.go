package service

import (
	"context"

	"github.com/1layar/universe/internal/email_service/internal/repo"
	"github.com/1layar/universe/internal/email_service/model"
)

type accountService struct {
	repo *repo.AccountRepository
}

// define const interface here
var _ AccountService = (*accountService)(nil)

func NewAccountService(repo *repo.AccountRepository) *accountService {
	return &accountService{repo: repo}
}

func (s *accountService) CreateAccount(ctx context.Context, account *model.Account) error {
	return s.repo.Create(ctx, account)
}

func (s *accountService) GetAccountByID(ctx context.Context, id int) (*model.Account, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *accountService) GetAccountByCode(ctx context.Context, code string) (*model.Account, error) {
	return s.repo.GetUserByField(ctx, "code", code)
}

func (s *accountService) UpdateAccount(ctx context.Context, account *model.Account) error {
	return s.repo.Update(ctx, account)
}

func (s *accountService) DeleteAccount(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
