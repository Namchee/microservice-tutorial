package endpoints

import (
	"context"
	"strconv"
	"time"

	"github.com/Namchee/microservice-tutorial/user/entity"
	"github.com/Namchee/microservice-tutorial/user/service"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

type UserEndpoints struct {
	GetUsers    endpoint.Endpoint
	GetUserById endpoint.Endpoint
	CreateUser  endpoint.Endpoint
	DeleteUser  endpoint.Endpoint
}

func NewUserEndpoint(logger log.Logger, svc service.UserService, mq service.PublisherService) *UserEndpoints {
	return &UserEndpoints{
		GetUsers:    makeGetUsersEndpoint(svc),
		GetUserById: makeGetUserByIdEndpoint(svc),
		CreateUser:  makeCreateUserEndpoint(svc),
		DeleteUser:  makeDeleteUserEndpoint(svc, mq),
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
		result, err := svc.GetUserById(ctx, request.(int))

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

func makeDeleteUserEndpoint(svc service.UserService, mq service.PublisherService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		result, err := svc.DeleteUser(ctx, request.(int))

		if err != nil {
			return nil, err
		}

		msg := &entity.Message{
			Name:      "delete",
			Content:   strconv.Itoa(result.Id),
			Timestamp: time.UTC.String(),
		}

		err = mq.Publish(ctx, msg)

		if err != nil {
			return nil, err
		}

		return result, nil
	}
}
