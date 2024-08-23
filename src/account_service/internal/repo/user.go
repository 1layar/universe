package repo

import (
	"context"
	"fmt"

	"github.com/1layar/universe/src/account_service/model"
	"github.com/1layar/universe/src/account_service/model/dto"
	"github.com/uptrace/bun"
)

type User struct {
	db *bun.DB
}

func NewUser(db *bun.DB) *User {
	return &User{db: db}
}

func (r *User) AddUser(ctx context.Context, data dto.CreateUserDto) (*model.User, error) {
	user := &model.User{
		Email:    data.Email,
		Username: data.Username,
		Password: data.Password,
	}

	if _, err := r.db.NewInsert().Model(user).Exec(ctx); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *User) GetUserByField(ctx context.Context, field string, value any, opt ...map[string]string) (*model.User, error) {
	user := &model.User{}
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

func (r *User) FindAll(ctx context.Context) ([]model.User, error) {
	var users []model.User

	model := r.db.NewSelect().Model(&users)

	if err := model.Scan(ctx); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *User) FindUserByField(ctx context.Context, field string, value any) ([]model.User, error) {
	var users []model.User

	model := r.db.NewSelect().Model(&users)
	kv := fmt.Sprintf("%%%s%%", value)
	model.Where(field+" like ?", kv)

	if err := model.Scan(ctx); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *User) DeleteByID(ctx context.Context, id int) (*model.User, error) {
	user, err := r.GetUserByField(ctx, "id", id)

	if err != nil {
		return nil, err
	}

	result, err := r.db.NewDelete().Model(user).Where("id = ?", user.ID).Exec(ctx)
	if err != nil {
		return nil, err
	}

	aff, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if aff < 0 {
		return nil, fmt.Errorf("nothing deleted")
	}

	return user, nil
}

func (r *User) Update(ctx context.Context, data dto.UpdateUser) (*model.User, error) {
	user := &model.User{
		ID:       data.Id,
		Username: data.Username,
		Email:    data.Email,
		Password: data.Password,
		Role:     data.Role,
	}

	res, err := r.db.NewUpdate().
		Model(user).
		OmitZero().
		WherePK().
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	aff, err := res.RowsAffected()

	if err != nil {
		return nil, err
	}

	if aff < 0 {
		return nil, fmt.Errorf("nothing updated")
	}

	return user, nil
}
