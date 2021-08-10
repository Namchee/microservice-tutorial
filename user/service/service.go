package service

import (
	"context"

	"github.com/Namchee/microservice-tutorial/user/pb"
)

type UserService interface {
	GetUsers(context.Context, *pb.GetUsersRequest) (*pb.GetUsersResponse, error)
	GetUserById(context.Context, *pb.GetUserByIdRequest) (*pb.User, error)
	CreateUser(context.Context, *pb.CreateUserRequest) (*pb.User, error)
}
