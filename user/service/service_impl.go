package service

import (
	"context"

	"github.com/Namchee/microservice-tutorial/user/pb"
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

func (svc *userService) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	users, err := svc.repository.GetUsers(ctx, req)

	if err != nil {
		return nil, err
	}

	return &pb.GetUsersResponse{
		Users: users,
	}, nil
}

func (svc *userService) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.User, error) {
	user, err := svc.repository.GetUserById(ctx, req.GetId())

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (svc *userService) CreateUser(ctx context.Context, data *pb.CreateUserRequest) (*pb.User, error) {
	user, err := svc.repository.CreateUser(ctx, data)

	if err != nil {
		return nil, err
	}

	return user, err
}
