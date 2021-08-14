package endpoints

import (
	"context"

	"github.com/Namchee/microservice-tutorial/post/entity"
	"github.com/Namchee/microservice-tutorial/post/service"
	"github.com/go-kit/kit/endpoint"
)

type PostEndpoints struct {
	GetPosts         endpoint.Endpoint
	GetPostById      endpoint.Endpoint
	CreatePost       endpoint.Endpoint
	DeletePost       endpoint.Endpoint
	DeletePostByUser endpoint.Endpoint
}

func NewPostEndpoint(svc service.PostService) *PostEndpoints {
	return &PostEndpoints{
		GetPosts:         makeGetPostsEndpoint(svc),
		GetPostById:      makeGetPostByIdEndpoint(svc),
		CreatePost:       makeCreatePostEndpoint(svc),
		DeletePost:       makeDeletePostEndpoint(svc),
		DeletePostByUser: makeDeletePostByUserEndpoint(svc),
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
		id := int(request.(int32))

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

func makeDeletePostEndpoint(svc service.PostService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := int(request.(int32))
		result, err := svc.DeletePost(ctx, req)

		if err != nil {
			return nil, err
		}

		return result, nil
	}
}

func makeDeletePostByUserEndpoint(svc service.PostService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := int(request.(int32))
		result, err := svc.DeletePostByUser(ctx, req)

		if err != nil {
			return nil, err
		}

		return result, nil
	}
}
