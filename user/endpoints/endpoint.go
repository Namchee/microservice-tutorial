package endpoints

import (
	"context"
	"reflect"

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
		GetUsers:    MakeGetUsersLoggingMiddleware(logger)(makeGetUsersEndpoint(svc)),
		GetUserById: MakeGetUserByIdLoggingMiddleware(logger)(makeGetUserByIdEndpoint(svc)),
		CreateUser:  MakeCreateUserLoggingMiddleware(logger)(makeCreateUserEndpoint(svc)),
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
		reflection := reflect.ValueOf(request).Elem()
		id := reflection.FieldByName("id").Interface().(int)

		result, err := svc.GetUserById(ctx, id)

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
