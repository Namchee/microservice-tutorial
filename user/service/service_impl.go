package service

import (
	"context"
	"errors"

	"github.com/Namchee/microservice-tutorial/user/entity"
	"github.com/Namchee/microservice-tutorial/user/repository"
)

var (
	ErrLimitOutOfRange  = errors.New("`limit` must be a positive integer")
	ErrOffsetOutOfRange = errors.New("`offset` must be an integer greater than -1")
	ErrIDOutOfRange     = errors.New("`id` must be a positive integer")
)

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *userService {
	return &userService{
		repository: repository,
	}
}

func (svc *userService) GetUsers(ctx context.Context, req *entity.Pagination) ([]*entity.User, error) {
	if req.Limit < 1 {
		return nil, ErrLimitOutOfRange
	}

	if req.Offset < 0 {
		return nil, ErrOffsetOutOfRange
	}

	users, err := svc.repository.GetUsers(ctx, req)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (svc *userService) GetUserById(ctx context.Context, id int) (*entity.User, error) {
	if id < 1 {
		return nil, ErrIDOutOfRange
	}

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
	if id < 1 {
		return nil, ErrIDOutOfRange
	}

	user, err := svc.repository.DeleteUser(ctx, id)

	if err != nil {
		return nil, err
	}

	return user, nil
}
