package service

import (
	"context"
	"errors"

	"github.com/Namchee/microservice-tutorial/post/entity"
	"github.com/Namchee/microservice-tutorial/post/repository"
)

var (
	errOutOfRangeLimit  = errors.New("limit is out of range")
	errOutOfRangeOffset = errors.New("offset is out of range")
	errOutOfRangeId     = errors.New("id is out of range")
)

type postService struct {
	repository repository.PostRepository
}

func NewPostService(repository repository.PostRepository) PostService {
	return &postService{
		repository: repository,
	}
}

func (svc *postService) GetPosts(ctx context.Context, pagination *entity.Pagination) ([]*entity.Post, error) {
	if pagination.Limit < 1 {
		return nil, errOutOfRangeLimit
	}

	if pagination.Offset < 0 {
		return nil, errOutOfRangeOffset
	}

	posts, err := svc.repository.GetPosts(ctx, pagination)

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (svc *postService) GetPostById(ctx context.Context, id int) (*entity.Post, error) {
	if id < 1 {
		return nil, errOutOfRangeId
	}

	post, err := svc.repository.GetPostById(ctx, id)

	if err != nil {
		return nil, err
	}

	return post, nil
}

func (svc *postService) CreatePost(ctx context.Context, data *entity.Post) (*entity.Post, error) {
	post, err := svc.repository.CreatePost(ctx, data)

	if err != nil {
		return nil, err
	}

	return post, nil
}
