package transports

import (
	"context"

	"github.com/Namchee/microservice-tutorial/post/endpoints"
	"github.com/Namchee/microservice-tutorial/post/entity"
	"github.com/Namchee/microservice-tutorial/post/pb"
	gt "github.com/go-kit/kit/transport/grpc"
)

type gRPCServer struct {
	getPosts    gt.Handler
	getPostById gt.Handler
	createPost  gt.Handler
}

func NewGRPCServer(endpoints *endpoints.PostEndpoints) pb.PostServiceServer {
	return &gRPCServer{
		getPosts: gt.NewServer(
			endpoints.GetPosts,
			decodeGetPostsRequest,
			encodeGetUsersResponse,
		),
		getPostById: gt.NewServer(
			endpoints.GetPostById,
			decodeGetPostByIdRequest,
			encodeGetPostByIdResponse,
		),
		createPost: gt.NewServer(
			endpoints.CreatePost,
			decodeCreatePostRequest,
			encodeCreatePostResponse,
		),
	}
}

func (s *gRPCServer) GetPosts(ctx context.Context, req *pb.GetPostsRequest) (*pb.GetPostsResponse, error) {
	_, resp, err := s.getPosts.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.GetPostsResponse), nil
}

func (s *gRPCServer) GetPostById(ctx context.Context, req *pb.GetPostByIdRequest) (*pb.Post, error) {
	_, resp, err := s.getPostById.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.Post), nil
}

func (s *gRPCServer) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.Post, error) {
	_, resp, err := s.createPost.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.Post), nil
}

func decodeGetPostsRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetPostsRequest)

	return &entity.Pagination{
		Limit:  int(req.Limit),
		Offset: int(req.Offset),
	}, nil
}

func encodeGetUsersResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.([]*entity.Post)

	var posts []*pb.Post

	for _, val := range res {
		pbPost := &pb.Post{
			Id:   int32(val.Id),
			Text: val.Text,
			User: int32(val.User),
		}

		posts = append(posts, pbPost)
	}

	return posts, nil
}

func decodeGetPostByIdRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetPostByIdRequest)
	return req.Id, nil
}

func encodeGetPostByIdResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(*entity.Post)

	return &pb.Post{
		Id:   int32(res.Id),
		Text: res.Text,
		User: int32(res.Id),
	}, nil
}

func decodeCreatePostRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreatePostRequest)

	return &entity.Post{
		Text: req.Text,
		User: int(req.User),
	}, nil
}

func encodeCreatePostResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(*entity.Post)

	return &pb.Post{
		Id:   int32(res.Id),
		Text: res.Text,
		User: int32(res.User),
	}, nil
}
