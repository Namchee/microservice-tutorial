package service

import (
	"context"
	"errors"

	"github.com/Namchee/microservice-tutorial/post/entity"
	"github.com/Namchee/microservice-tutorial/post/repository"
)

var (
	ErrOutOfRangeLimit  = errors.New("`limit` must be a positive integer")
	ErrOutOfRangeOffset = errors.New("`offset` must be an integer bigger than -1")
	ErrOutOfRangeId     = errors.New("`id` must be a positive integer")
)

type postService struct {
	repository repository.PostRepository
}

func NewPostService(repository repository.PostRepository) *postService {
	return &postService{
		repository: repository,
	}
}

func (svc *postService) GetPosts(ctx context.Context, pagination *entity.Pagination) ([]*entity.Post, error) {
	if pagination.Limit < 1 {
		return nil, ErrOutOfRangeLimit
	}

	if pagination.Offset < 0 {
		return nil, ErrOutOfRangeOffset
	}

	posts, err := svc.repository.GetPosts(ctx, pagination)

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (svc *postService) GetPostById(ctx context.Context, id int) (*entity.Post, error) {
	if id < 1 {
		return nil, ErrOutOfRangeId
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

func (svc *postService) DeletePost(ctx context.Context, postId int) (*entity.Post, error) {
	if postId < 0 {
		return nil, ErrOutOfRangeId
	}

	post, err := svc.repository.DeletePost(ctx, postId)

	if err != nil {
		return nil, err
	}

	return post, nil
}

func (svc *postService) DeletePostByUser(ctx context.Context, userId int) ([]*entity.Post, error) {
	posts, err := svc.repository.DeletePostByUser(ctx, userId)

	if err != nil {
		return nil, err
	}

	return posts, nil
}
