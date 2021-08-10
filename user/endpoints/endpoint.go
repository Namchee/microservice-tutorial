package endpoints

import (
	"context"

	"github.com/Namchee/microservice-tutorial/user/pb"
	"github.com/Namchee/microservice-tutorial/user/service"
	"github.com/go-kit/kit/endpoint"
)

type Response struct {
	Data  interface{}
	Error error
}

type UserEndpoints struct {
	GetUsers    endpoint.Endpoint
	GetUserById endpoint.Endpoint
	CreateUser  endpoint.Endpoint
}

func NewUserEndpoint(svc service.UserService) *UserEndpoints {
	return &UserEndpoints{
		GetUsers:    makeGetUsersEndpoint(svc),
		GetUserById: makeGetUserByIdEndpoint(svc),
		CreateUser:  makeCreateUserEndpoint(svc),
	}
}

func makeGetUsersEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.GetUsersRequest)
		result, err := svc.GetUsers(ctx, req)

		if err != nil {
			return &Response{
				Data:  nil,
				Error: err,
			}, nil
		}

		return &Response{
			Data:  result,
			Error: nil,
		}, nil
	}
}

func makeGetUserByIdEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.GetUserByIdRequest)
		result, err := svc.GetUserById(ctx, req)

		if err != nil {
			return &Response{
				Data:  nil,
				Error: err,
			}, nil
		}

		return &Response{
			Data:  result,
			Error: nil,
		}, nil
	}
}

func makeCreateUserEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.CreateUserRequest)
		result, err := svc.CreateUser(ctx, req)

		if err != nil {
			return &Response{
				Data:  nil,
				Error: err,
			}, nil
		}

		return &Response{
			Data:  result,
			Error: nil,
		}, nil
	}
}
