package service

import (
	"context"
	"testing"

	"github.com/Namchee/microservice-tutorial/post/entity"
)

type mockRepository struct{}

func (repo *mockRepository) GetPosts(ctx context.Context, pagination *entity.Pagination) ([]*entity.Post, error) {
	return []*entity.Post{}, nil
}

func (repo *mockRepository) GetPostById(ctx context.Context, _ int) (*entity.Post, error) {
	return &entity.Post{}, nil
}

func (repo *mockRepository) CreatePost(ctx context.Context, _ *entity.Post) (*entity.Post, error) {
	return &entity.Post{}, nil
}

func (repo *mockRepository) DeletePost(ctx context.Context, _ int) (*entity.Post, error) {
	return &entity.Post{}, nil
}

func (repo *mockRepository) DeletePostByUser(ctx context.Context, _ int) ([]*entity.Post, error) {
	return []*entity.Post{}, nil
}

var ctx context.Context
var service *postService

func init() {
	ctx = context.Background()
	service = NewPostService(&mockRepository{})
}

func TestGetPosts(t *testing.T) {
	pager := &entity.Pagination{
		Limit:  0,
		Offset: 0,
	}

	_, err := service.GetPosts(ctx, pager)

	if err != ErrLimitOutOfRange {
		t.Fatalf("TestGetPosts: expected ErrLimitOutOfRange")
	}

	pager.Limit = 1
	pager.Offset = -1

	_, err = service.GetPosts(ctx, pager)

	if err != ErrOffsetOutOfRange {
		t.Fatalf("TestGetPosts: expected ErrOffsetOutOfRange")
	}

	pager.Limit = 1
	pager.Offset = 0

	_, err = service.GetPosts(ctx, pager)

	if err != nil {
		t.Fatalf("TestGetPosts: expected nil error")
	}
}

func TestGetPostById(t *testing.T) {
	_, err := service.GetPostById(ctx, -1)

	if err != ErrIDOutOfRange {
		t.Fatalf("TestGetPostById: expected ErrIDOutOfRange")
	}

	_, err = service.GetPostById(ctx, 1)

	if err != nil {
		t.Fatalf("TestGetPostById: expected nil errro")
	}
}

func TestCreatePost(t *testing.T) {
	post := &entity.Post{
		Text: "WASM",
		User: -1,
	}

	_, err := service.CreatePost(ctx, post)

	if err != ErrIDOutOfRange {
		t.Fatalf("TestCreatePost: expected ErrIDOutOfRange")
	}

	post.User = 1

	_, err = service.CreatePost(ctx, post)

	if err != nil {
		t.Fatalf("TestCreatePost: expected nil error")
	}
}

func TestDeletePost(t *testing.T) {
	_, err := service.DeletePost(ctx, -1)

	if err != ErrIDOutOfRange {
		t.Fatalf("TestDeletePost: expected ErrIDOutOfRange")
	}

	_, err = service.DeletePost(ctx, 1)

	if err != nil {
		t.Fatalf("TestDeletePost: expected nil errro")
	}
}

func TestDeletePostByUser(t *testing.T) {
	_, err := service.DeletePostByUser(ctx, -1)

	if err != ErrIDOutOfRange {
		t.Fatalf("TestDeletePostByUser: expected ErrIDOutOfRange")
	}

	_, err = service.DeletePostByUser(ctx, 1)

	if err != nil {
		t.Fatalf("TestPostUser: expected nil errro")
	}
}
