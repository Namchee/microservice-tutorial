package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

func MakeGetPostsLoggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			level.Info(logger).Log("msg", "GetPosts: calling endpoint")
			defer level.Info(logger).Log("msg", "GetPosts: called endpoint")
			return next(ctx, request)
		}
	}
}

func MakeGetPostByIdLoggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			level.Info(logger).Log("msg", "GetPostById: calling endpoint")
			defer level.Info(logger).Log("msg", "GetPostById: called endpoint")
			return next(ctx, request)
		}
	}
}

func MakeCreatePostLoggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			level.Info(logger).Log("msg", "CreatePost: calling endpoint")
			defer level.Info(logger).Log("msg", "CreatePost: called endpoint")
			return next(ctx, request)
		}
	}
}
