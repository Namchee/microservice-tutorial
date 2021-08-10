package endpoints

import (
	"context"
	"reflect"

	"github.com/Namchee/microservice-tutorial/post/entity"
	"github.com/Namchee/microservice-tutorial/post/service"
	"github.com/go-kit/kit/endpoint"
)

type PostEndpoints struct {
	GetPosts    endpoint.Endpoint
	GetPostById endpoint.Endpoint
	CreatePost  endpoint.Endpoint
}

func NewUserEndpoint(svc service.PostService) *PostEndpoints {
	return &PostEndpoints{
		GetPosts:    makeGetPostsEndpoint(svc),
		GetPostById: makeGetPostByIdEndpoint(svc),
		CreatePost:  makeCreatePostEndpoint(svc),
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
