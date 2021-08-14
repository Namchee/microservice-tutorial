package repository

import (
	"context"

	"github.com/Namchee/microservice-tutorial/user/entity"
)

type UserRepository interface {
	GetUsers(context.Context, *entity.Pagination) ([]*entity.User, error)
	GetUserById(context.Context, int) (*entity.User, error)
	CreateUser(context.Context, *entity.User) (*entity.User, error)
	DeleteUser(context.Context, int) (*entity.User, error)
}
