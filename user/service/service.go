package service

import (
	"context"

	"github.com/Namchee/microservice-tutorial/user/entity"
)

type UserService interface {
	GetUsers(context.Context, *entity.Pagination) ([]*entity.User, error)
	GetUserById(context.Context, int) (*entity.User, error)
	CreateUser(context.Context, *entity.User) (*entity.User, error)
	DeleteUser(context.Context, int) (*entity.User, error)
}

type PublisherService interface {
	Publish(context.Context, *entity.Message) error
}

type ConsumerService interface {
	Consume(context.Context, *entity.Message) error
}
