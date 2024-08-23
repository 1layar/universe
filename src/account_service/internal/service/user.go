package service

import (
	"context"

	"github.com/1layar/universe/src/account_service/internal/repo"
	"github.com/1layar/universe/src/account_service/model"
	"github.com/1layar/universe/src/account_service/model/dto"
	"github.com/1layar/universe/src/shared/utils"
)

type User struct {
	userRepo *repo.User
}

func NewUser(userRepo *repo.User) *User {
	return &User{
		userRepo: userRepo,
	}
}

func (s *User) AddUser(ctx context.Context, data dto.CreateUserDto) (*model.User, error) {
	// hash password
	password, err := utils.HashPassword(data.Password)

	if err != nil {
		return nil, err
	}

	data.Password = password

	return s.userRepo.AddUser(ctx, data)
}

func (s *User) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	return s.userRepo.GetUserByField(ctx, "id", id)
}

func (s *User) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	return s.userRepo.GetUserByField(ctx, "email", email)
}

func (s *User) HasUsername(ctx context.Context, username string, opt ...map[string]string) bool {
	_, err := s.userRepo.GetUserByField(ctx, "username", username, opt...)

	return err == nil
}

func (s *User) HasEmail(ctx context.Context, email string, opt ...map[string]string) bool {
	_, err := s.userRepo.GetUserByField(ctx, "email", email, opt...)

	return err == nil
}

func (s *User) SearchUser(ctx context.Context, search dto.SearchUserDto) ([]model.User, error) {

	if search.Username == nil {
		return s.userRepo.FindAll(ctx)
	}
	user, err := s.userRepo.FindUserByField(ctx, "username", *search.Username)

	return user, err
}

func (s *User) Update(ctx context.Context, data dto.UpdateUser) (*model.User, error) {

	if data.Password != "" {
		password, err := utils.HashPassword(data.Password)
		if err != nil {
			return nil, err
		}

		data.Password = password
	}

	result, err := s.userRepo.Update(ctx, data)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *User) DeleteByID(ctx context.Context, id int) (*model.User, error) {
	return s.userRepo.DeleteByID(ctx, id)
}
