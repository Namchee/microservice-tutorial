package service

import (
	"context"

	"github.com/Namchee/microservice-tutorial/user/entity"
	"github.com/Namchee/microservice-tutorial/user/repository"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
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
	level.Info(svc.logger).Log("UserService: executing get users query")
	users, err := svc.repository.GetUsers(ctx, req)
	level.Info(svc.logger).Log("UserService: get users query execution finished")

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (svc *userService) GetUserById(ctx context.Context, id int) (*entity.User, error) {
	level.Info(svc.logger).Log("UserService: executing get user by id query")
	user, err := svc.repository.GetUserById(ctx, id)
	level.Info(svc.logger).Log("UserService: get user by id query execution finished")

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (svc *userService) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	level.Info(svc.logger).Log("UserService: executing create user query")
	user, err := svc.repository.CreateUser(ctx, user)
	level.Info(svc.logger).Log("UserService: create user query execution finished")

	if err != nil {
		return nil, err
	}

	return user, err
}
