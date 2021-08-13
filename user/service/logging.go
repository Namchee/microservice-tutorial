package service

import (
	"context"
	"time"

	"github.com/Namchee/microservice-tutorial/user/entity"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type LoggingMiddleware func(UserService) UserService

type loggingMiddleware struct {
	logger log.Logger
	next   UserService
}

func NewLoggingMiddleware(logger log.Logger) LoggingMiddleware {
	return func(next UserService) UserService {
		return &loggingMiddleware{logger, next}
	}
}

func (mw *loggingMiddleware) GetUsers(ctx context.Context, pagination *entity.Pagination) ([]*entity.User, error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "GetUsers",
			"time", time.Since(begin),
		)
	}(time.Now())
	return mw.next.GetUsers(ctx, pagination)
}

func (mw *loggingMiddleware) GetUserById(ctx context.Context, id int) (*entity.User, error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "GetUsers",
			"time", time.Since(begin),
		)
	}(time.Now())
	return mw.next.GetUserById(ctx, id)
}

func (mw *loggingMiddleware) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "GetUsers",
			"time", time.Since(begin),
		)
	}(time.Now())
	return mw.next.CreateUser(ctx, user)
}
