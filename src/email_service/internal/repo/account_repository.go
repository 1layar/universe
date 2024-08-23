package repo

import (
	"context"
	"fmt"

	"github.com/1layar/universe/src/email_service/model"

	"github.com/uptrace/bun"
)

type AccountRepository struct {
	db *bun.DB
}

func NewAccountRepository(db *bun.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (r *AccountRepository) Create(ctx context.Context, account *model.Account) error {
	_, err := r.db.NewInsert().Model(account).Exec(ctx)
	return err
}

func (r *AccountRepository) GetByID(ctx context.Context, id int) (*model.Account, error) {
	account := new(model.Account)
	err := r.db.NewSelect().Model(account).Where("id = ?", id).Scan(ctx)
	return account, err
}

func (r *AccountRepository) GetUserByField(ctx context.Context, field string, value any, opt ...map[string]string) (*model.Account, error) {
	user := &model.Account{}
	model := r.db.NewSelect().Model(user)

	model.Where(fmt.Sprintf("%s = ?", field), value)

	for _, v := range opt {
		for k, v := range v {
			model.Where(fmt.Sprintf("%s <> ?", k), v)
		}
	}

	if err := model.Scan(ctx); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *AccountRepository) Update(ctx context.Context, account *model.Account) error {
	_, err := r.db.NewUpdate().Model(account).Where("id = ?", account.ID).Exec(ctx)
	return err
}

func (r *AccountRepository) Delete(ctx context.Context, id int) error {
	_, err := r.db.NewDelete().Model(&model.Account{}).Where("id = ?", id).Exec(ctx)
	return err
}
