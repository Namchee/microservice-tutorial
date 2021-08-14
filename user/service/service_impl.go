package service

import (
	"context"

	"github.com/Namchee/microservice-tutorial/user/entity"
	"github.com/Namchee/microservice-tutorial/user/repository"
)

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserService {
	return &userService{
		repository: repository,
	}
}

func (svc *userService) GetUsers(ctx context.Context, req *entity.Pagination) ([]*entity.User, error) {
	users, err := svc.repository.GetUsers(ctx, req)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (svc *userService) GetUserById(ctx context.Context, id int) (*entity.User, error) {
	user, err := svc.repository.GetUserById(ctx, id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (svc *userService) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	user, err := svc.repository.CreateUser(ctx, user)

	if err != nil {
		return nil, err
	}

	return user, err
}

func (svc *userService) DeleteUser(ctx context.Context, id int) (*entity.User, error) {
	user, err := svc.repository.DeleteUser(ctx, id)

	if err != nil {
		return nil, err
	}

	return user, nil
}
