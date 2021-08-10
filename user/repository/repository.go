package repository

import (
	"context"

	pb "github.com/Namchee/microservice-tutorial/user/pb"
)

type UserRepository interface {
	GetUsers(context.Context, *pb.GetUsersRequest) ([]*pb.User, error)
	GetUserById(context.Context, int32) (*pb.User, error)
	CreateUser(context.Context, *pb.CreateUserRequest) (*pb.User, error)
}
