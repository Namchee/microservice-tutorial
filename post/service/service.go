package service

import (
	"context"

	"github.com/Namchee/microservice-tutorial/post/entity"
)

type PostService interface {
	GetPosts(context.Context, *entity.Pagination) ([]*entity.Post, error)
	GetPostById(context.Context, int) (*entity.Post, error)
	CreatePost(context.Context, *entity.Post) (*entity.Post, error)
}
