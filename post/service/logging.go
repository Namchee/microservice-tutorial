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

func (mw *loggingMiddleware) GetPosts(ctx context.Context, pagination *entity.Pagination) ([]*entity.Post, error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "GetPosts",
			"time", time.Since(begin),
		)
	}(time.Now())

	return mw.next.GetPosts(ctx, pagination)
}

func (mw *loggingMiddleware) GetPostById(ctx context.Context, id int) (*entity.Post, error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "GetPostById",
			"time", time.Since(begin),
		)
	}(time.Now())

	return mw.next.GetPostById(ctx, id)
}

func (mw *loggingMiddleware) CreatePost(ctx context.Context, post *entity.Post) (*entity.Post, error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "CreatePost",
			"time", time.Since(begin),
		)
	}(time.Now())

	return mw.next.CreatePost(ctx, post)
}

func (mw *loggingMiddleware) DeletePost(ctx context.Context, postId int) (*entity.Post, error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "DeletePost",
			"time", time.Since(begin),
		)
	}(time.Now())

	return mw.next.DeletePost(ctx, postId)
}

func (mw *loggingMiddleware) DeletePostByUser(ctx context.Context, userId int) ([]*entity.Post, error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "DeletePostByUser",
			"time", time.Since(begin),
		)
	}(time.Now())

	return mw.next.DeletePostByUser(ctx, userId)
}
