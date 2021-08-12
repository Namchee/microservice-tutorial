package transports

import (
	"context"

	"github.com/Namchee/microservice-tutorial/user/endpoints"
	"github.com/Namchee/microservice-tutorial/user/entity"
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
	return resp.(*pb.GetUsersResponse), nil
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

	return &entity.Pagination{
		Limit:  int(*req.Limit),
		Offset: int(*req.Offset),
	}, nil
}

func encodeGetUsersResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.([]*entity.User)

	var users []*pb.User

	for _, val := range res {
		pbUser := &pb.User{
			Id:       int32(val.Id),
			Username: val.Username,
			Name:     val.Name,
			Bio:      val.Bio,
		}

		users = append(users, pbUser)
	}

	return users, nil
}

func decodeGetUserByIdRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetUserByIdRequest)
	return req.Id, nil
}

func encodeGetUserByIdResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(*entity.User)

	return &pb.User{
		Id:       int32(res.Id),
		Username: res.Username,
		Name:     res.Name,
		Bio:      res.Bio,
	}, nil
}

func decodeCreateUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateUserRequest)

	return &entity.User{
		Username: req.Username,
		Name:     req.Name,
		Bio:      req.Bio,
	}, nil
}

func encodeCreateUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(*entity.User)

	return &pb.User{
		Id:       int32(res.Id),
		Username: res.Username,
		Name:     res.Name,
		Bio:      res.Bio,
	}, nil
}
