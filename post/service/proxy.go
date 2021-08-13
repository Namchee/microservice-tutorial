package service

import (
	"context"
	"errors"

	"github.com/Namchee/microservice-tutorial/post/entity"
	pb "github.com/Namchee/microservice-tutorial/user/pb"
)

var (
	errUserNotFound = errors.New("user not found")
)

type ProxyMiddleware func(PostService) PostService

type proxyMiddleware struct {
	userClient pb.UserServiceClient
	next       PostService
}

func NewPostServiceProxy(client pb.UserServiceClient) ProxyMiddleware {
	return func(next PostService) PostService {
		return &proxyMiddleware{userClient: client, next: next}
	}
}

func (mw *proxyMiddleware) GetPosts(ctx context.Context, pagination *entity.Pagination) ([]*entity.Post, error) {
	return mw.next.GetPosts(ctx, pagination)
}

func (mw *proxyMiddleware) GetPostById(ctx context.Context, id int) (*entity.Post, error) {
	return mw.next.GetPostById(ctx, id)
}

func (mw *proxyMiddleware) CreatePost(ctx context.Context, post *entity.Post) (*entity.Post, error) {
	req := &pb.GetUserByIdRequest{
		Id: int32(post.User),
	}

	user, err := mw.userClient.GetUserById(ctx, req)

	if err != nil {
		return nil, err
	}

	if user.Data == nil {
		return nil, errUserNotFound
	}

	return mw.next.CreatePost(ctx, post)
}
