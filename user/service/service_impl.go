package service

import (
	"context"

	"github.com/Namchee/microservice-tutorial/user/entity"
	"github.com/Namchee/microservice-tutorial/user/repository"
	"github.com/go-kit/kit/log"
)

type userService struct {
	repository repository.UserRepository
	logger     log.Logger
}

func NewUserService(logger log.Logger, repository repository.UserRepository) *userService {
	return &userService{
		repository: repository,
		logger:     logger,
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
