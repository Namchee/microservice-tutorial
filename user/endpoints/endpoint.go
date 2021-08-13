package endpoints

import (
	"context"

	"github.com/Namchee/microservice-tutorial/user/entity"
	"github.com/Namchee/microservice-tutorial/user/service"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

type UserEndpoints struct {
	GetUsers    endpoint.Endpoint
	GetUserById endpoint.Endpoint
	CreateUser  endpoint.Endpoint
}

func NewUserEndpoint(logger log.Logger, svc service.UserService) *UserEndpoints {
	return &UserEndpoints{
		GetUsers:    makeGetUsersEndpoint(svc),
		GetUserById: makeGetUserByIdEndpoint(svc),
		CreateUser:  makeCreateUserEndpoint(svc),
	}
}

func makeGetUsersEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entity.Pagination)
		result, err := svc.GetUsers(ctx, req)

		if err != nil {
			return nil, err
		}

		return result, nil
	}
}

func makeGetUserByIdEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := svc.GetUserById(ctx, int(request.(int32)))

		if err != nil {
			return nil, err
		}

		return result, nil
	}
}

func makeCreateUserEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entity.User)
		result, err := svc.CreateUser(ctx, req)

		if err != nil {
			return nil, err
		}

		return result, nil
	}
}
