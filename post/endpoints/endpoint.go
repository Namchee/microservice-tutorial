package endpoints

import (
	"context"
	"reflect"

	"github.com/Namchee/microservice-tutorial/post/entity"
	"github.com/Namchee/microservice-tutorial/post/service"
	"github.com/go-kit/kit/endpoint"
)

type UserEndpoints struct {
	GetUsers    endpoint.Endpoint
	GetUserById endpoint.Endpoint
	CreateUser  endpoint.Endpoint
}

func NewUserEndpoint(svc service.PostService) *UserEndpoints {
	return &UserEndpoints{
		GetUsers:    makeGetPostsEndpoint(svc),
		GetUserById: makeGetPostByIdEndpoint(svc),
		CreateUser:  makeCreatePostEndpoint(svc),
	}
}

func makeGetPostsEndpoint(svc service.PostService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entity.Pagination)
		result, err := svc.GetPosts(ctx, req)

		if err != nil {
			return nil, err
		}

		return result, nil
	}
}

func makeGetPostByIdEndpoint(svc service.PostService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		reflection := reflect.ValueOf(request).Elem()
		id := reflection.FieldByName("id").Interface().(int)

		result, err := svc.GetPostById(ctx, id)

		if err != nil {
			return nil, err
		}

		return result, nil
	}
}

func makeCreatePostEndpoint(svc service.PostService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*entity.Post)
		result, err := svc.CreatePost(ctx, req)

		if err != nil {
			return nil, err
		}

		return result, nil
	}
}
