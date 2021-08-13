package service

import (
	"context"
	"errors"

	"github.com/Namchee/microservice-tutorial/post/entity"
	"github.com/Namchee/microservice-tutorial/post/repository"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

var (
	errOutOfRangeLimit  = errors.New("limit is out of range")
	errOutOfRangeOffset = errors.New("offset is out of range")
	errOutOfRangeId     = errors.New("id is out of range")
)

type postService struct {
	repository repository.PostRepository
	logger     log.Logger
}

func NewPostService(logger log.Logger, repository repository.PostRepository) PostService {
	return &postService{
		repository: repository,
		logger:     logger,
	}
}

func (svc *postService) GetPosts(ctx context.Context, pagination *entity.Pagination) ([]*entity.Post, error) {
	if pagination.Limit < 1 {
		return nil, errOutOfRangeLimit
	}

	if pagination.Offset < 0 {
		return nil, errOutOfRangeOffset
	}

	level.Info(svc.logger).Log("msg", "PostService: executing GetPosts query")
	posts, err := svc.repository.GetPosts(ctx, pagination)
	level.Info(svc.logger).Log("msg", "PostService: executed GetPosts query")

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (svc *postService) GetPostById(ctx context.Context, id int) (*entity.Post, error) {
	if id < 1 {
		return nil, errOutOfRangeId
	}

	level.Info(svc.logger).Log("msg", "PostService: executing GetPostById query")
	post, err := svc.repository.GetPostById(ctx, id)
	level.Info(svc.logger).Log("msg", "PostService: executed GetPostById query")

	if err != nil {
		return nil, err
	}

	return post, nil
}

func (svc *postService) CreatePost(ctx context.Context, data *entity.Post) (*entity.Post, error) {
	level.Info(svc.logger).Log("msg", "PostService: executing CreatePost query")
	post, err := svc.repository.CreatePost(ctx, data)
	level.Info(svc.logger).Log("msg", "PostService: executed CreatePost query")

	if err != nil {
		return nil, err
	}

	return post, nil
}
