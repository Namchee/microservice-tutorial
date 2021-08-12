package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

func MakeGetUsersLoggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			level.Info(logger).Log("msg", "GetUsers: calling endpoint")
			defer level.Info(logger).Log("msg", "GetUsers: called endpoint")
			return next(ctx, request)
		}
	}
}

func MakeGetUserByIdLoggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			level.Info(logger).Log("msg", "GetUserById: calling endpoint")
			defer level.Info(logger).Log("msg", "GetUserById: called endpoint")
			return next(ctx, request)
		}
	}
}

func MakeCreateUserLoggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			level.Info(logger).Log("msg", "CreateUser: calling endpoint")
			defer level.Info(logger).Log("msg", "CreateUser: called endpoint")
			return next(ctx, request)
		}
	}
}
