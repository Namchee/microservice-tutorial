package endpoints

import (
	"context"

	"github.com/Namchee/microservice-tutorial/post/entity"
	"github.com/Namchee/microservice-tutorial/post/service"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

type PostEndpoints struct {
	GetPosts    endpoint.Endpoint
	GetPostById endpoint.Endpoint
	CreatePost  endpoint.Endpoint
}

func NewPostEndpoint(logger log.Logger, svc service.PostService) *PostEndpoints {
	return &PostEndpoints{
		GetPosts:    MakeGetPostsLoggingMiddleware(logger)(makeGetPostsEndpoint(svc)),
		GetPostById: MakeGetPostByIdLoggingMiddleware(logger)(makeGetPostByIdEndpoint(svc)),
		CreatePost:  MakeCreatePostLoggingMiddleware(logger)(makeCreatePostEndpoint(svc)),
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
