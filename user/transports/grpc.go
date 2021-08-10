package transports

import (
	"context"

	"github.com/Namchee/microservice-tutorial/user/endpoints"
	"github.com/Namchee/microservice-tutorial/user/pb"
	gt "github.com/go-kit/kit/transport/grpc"
)

type gRPCServer struct {
	getUsers    gt.Handler
	getUserById gt.Handler
	createUser  gt.Handler
}

func NewGRPCServer(endpoints *endpoints.UserEndpoints) pb.UserServiceServer {
	return &gRPCServer{
		getUsers: gt.NewServer(
			endpoints.GetUsers,
			decodeGetUsersRequest,
			encodeGetUsersResponse,
		),
		getUserById: gt.NewServer(
			endpoints.GetUserById,
			decodeGetUserByIdRequest,
			encodeGetUserByIdResponse,
		),
		createUser: gt.NewServer(
			endpoints.CreateUser,
			decodeCreateUserRequest,
			encodeCreateUserResponse,
		),
	}
}

func (s *gRPCServer) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	_, resp, err := s.getUsers.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.GetUsersResponse{Users: resp.([]*pb.User)}, nil
}

func (s *gRPCServer) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.User, error) {
	_, resp, err := s.getUserById.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.User), nil
}

func (s *gRPCServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	_, resp, err := s.createUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.User), nil
}

func decodeGetUsersRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetUsersRequest)
	return req, nil
}

func encodeGetUsersResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.Response)

	if res.Error != nil {
		return nil, res.Error
	}

	return res.Data.([]*pb.User), nil
}

func decodeGetUserByIdRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetUserByIdRequest)
	return req, nil
}

func encodeGetUserByIdResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.Response)

	if res.Error != nil {
		return nil, res.Error
	}

	return res.Data.(*pb.User), nil
}

func decodeCreateUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateUserRequest)
	return req, nil
}

func encodeCreateUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.Response)

	if res.Error != nil {
		return nil, res.Error
	}

	return res.Data.(*pb.User), nil
}
