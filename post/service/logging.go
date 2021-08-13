package service

import (
	"context"
	"time"

	"github.com/Namchee/microservice-tutorial/post/entity"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type LoggingMiddleware func(PostService) PostService

type loggingMiddleware struct {
	logger log.Logger
	next   PostService
}

func NewPostLoggingMiddleware(logger log.Logger) LoggingMiddleware {
	return func(next PostService) PostService {
		return &loggingMiddleware{logger, next}
	}
}

func (mw *loggingMiddleware) GetPosts(ctx context.Context, pagination *entity.Pagination) (posts []*entity.Post, err error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "GetPosts",
			"time", time.Since(begin),
		)
	}(time.Now())

	return mw.next.GetPosts(ctx, pagination)
}

func (mw *loggingMiddleware) GetPostById(ctx context.Context, id int) (post *entity.Post, err error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "GetPostById",
			"time", time.Since(begin),
		)
	}(time.Now())

	return mw.next.GetPostById(ctx, id)
}

func (mw *loggingMiddleware) CreatePost(ctx context.Context, post *entity.Post) (entity *entity.Post, err error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "CreatePost",
			"time", time.Since(begin),
		)
	}(time.Now())

	return mw.next.CreatePost(ctx, post)
}
