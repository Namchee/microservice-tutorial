package service

import (
	"context"
	"testing"

	"github.com/Namchee/microservice-tutorial/user/entity"
)

type mockRepository struct{}

func (repo *mockRepository) GetUsers(ctx context.Context, pagination *entity.Pagination) ([]*entity.User, error) {
	return []*entity.User{}, nil
}

func (repo *mockRepository) GetUserById(ctx context.Context, _ int) (*entity.User, error) {
	return &entity.User{}, nil
}

func (repo *mockRepository) CreateUser(ctx context.Context, _ *entity.User) (*entity.User, error) {
	return &entity.User{}, nil
}

func (repo *mockRepository) DeleteUser(ctx context.Context, _ int) (*entity.User, error) {
	return &entity.User{}, nil
}

var ctx context.Context
var service *userService

func init() {
	ctx = context.Background()
	service = NewUserService(&mockRepository{})
}

func TestGetUsers(t *testing.T) {
	pager := &entity.Pagination{
		Limit:  0,
		Offset: 0,
	}

	_, err := service.GetUsers(ctx, pager)

	if err != ErrLimitOutOfRange {
		t.Fatalf("TestGetUsers: expected ErrLimitOutOfRange")
	}

	pager.Limit = 1
	pager.Offset = -1

	_, err = service.GetUsers(ctx, pager)

	if err != ErrOffsetOutOfRange {
		t.Fatalf("TestGetUsers: expected ErrOffsetOutOfRange")
	}

	pager.Limit = 1
	pager.Offset = 0

	_, err = service.GetUsers(ctx, pager)

	if err != nil {
		t.Fatalf("TestGetUsers: expected nil error")
	}
}

func TestGetUserById(t *testing.T) {
	_, err := service.GetUserById(ctx, -1)

	if err != ErrIDOutOfRange {
		t.Fatalf("TestGetUserById: expected ErrIDOutOfRange")
	}

	_, err = service.GetUserById(ctx, 1)

	if err != nil {
		t.Fatalf("TestGetUserById: expected nil errro")
	}
}

func TestCreateUser(t *testing.T) {
	user := &entity.User{
		Username: "foo",
		Name:     "bar",
		Bio:      "baz",
	}

	_, err := service.CreateUser(ctx, user)

	if err != nil {
		t.Fatalf("TestCreateUser: expected nil error")
	}
}

func TestDeleteUser(t *testing.T) {
	_, err := service.DeleteUser(ctx, -1)

	if err != ErrIDOutOfRange {
		t.Fatalf("TestDeleteUser: expected ErrIDOutOfRange")
	}

	_, err = service.DeleteUser(ctx, 1)

	if err != nil {
		t.Fatalf("TestDeleteUser: expected nil errro")
	}
}
