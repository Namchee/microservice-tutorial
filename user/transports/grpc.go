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
	deleteUser  gt.Handler
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
		deleteUser: gt.NewServer(
			endpoints.DeleteUser,
			decodeDeleteUserRequest,
			encodeDeleteUserResponse,
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

func (s *gRPCServer) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {
	_, resp, err := s.getUserById.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.GetUserByIdResponse), nil
}

func (s *gRPCServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	_, resp, err := s.createUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.User), nil
}

func (s *gRPCServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	_, resp, err := s.deleteUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.DeleteUserResponse), nil
}

func decodeGetUsersRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetUsersRequest)

	limit := 0
	offset := 0

	if req.Limit != nil {
		limit = int(*req.Limit)
	}

	if req.Offset != nil {
		offset = int(*req.Offset)
	}

	return &entity.Pagination{
		Limit:  limit,
		Offset: offset,
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

	return &pb.GetUsersResponse{
		Data: users,
	}, nil
}

func decodeGetUserByIdRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetUserByIdRequest)
	return int(req.Id), nil
}

func encodeGetUserByIdResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(*entity.User)
	var data *pb.User

	if res != nil {
		data = &pb.User{
			Id:       int32(res.Id),
			Username: res.Username,
			Name:     res.Name,
			Bio:      res.Bio,
		}
	}

	return &pb.GetUserByIdResponse{
		Data: data,
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

func decodeDeleteUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.DeleteUserRequest)
	return int(req.Id), nil
}

func encodeDeleteUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(*entity.User)
	var user *pb.User

	if res != nil {
		user = &pb.User{
			Id:       int32(res.Id),
			Username: res.Username,
			Name:     res.Name,
			Bio:      res.Bio,
		}
	}

	return &pb.DeleteUserResponse{
		User: user,
	}, nil
}
