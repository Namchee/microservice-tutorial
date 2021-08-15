package transports

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/Namchee/microservice-tutorial/user/endpoints"
	"github.com/Namchee/microservice-tutorial/user/entity"
	httptransport "github.com/go-kit/kit/transport/http"
)

type httpResponse struct {
	data interface{} `json:"data"`
	err  string      `json:"error"`
}

type httpServer struct {
	getUsers    httptransport.Server
	getUserById httptransport.Server
	createUser  httptransport.Server
	deleteUser  httptransport.Server
}

func NewHTTPServer(endpoints endpoints.UserEndpoints) httpServer {
	return &httpServer{
		getUsers: httptransport.NewServer(
			endpoints.GetUsers,
			decodeGetUsersHTTPRequest,
			encodeGetUsersHTTPResponse,
		),
		getUserById: httptransport.NewServer(
			endpoints.GetUserById,
			decodeGetUserByIdRequest,
			encodeGetUserByIdResponse,
		),
		createUser: httptransport.NewServer(
			endpoints.CreateUser,
			decodeCreateUserRequest,
			encodeCreateUserResponse,
		),
		deleteUser: httptransport.NewServer(
			endpoints.DeleteUser,
			decodeDeleteUserRequest,
			encodeDeleteUserResponse,
		),
	}
}

func encodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("error with nil err")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500) // just give it 500 for now
	json.NewEncoder(w).Encode(httpResponse{
		data: nil,
		err:  err.Error(),
	})
}

func encodeHTTPResponse(_ context.Context, w http.ResponseWriter, res interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(httpResponse{
		data: res,
	})
}

func decodeGetUsersHTTPRequest(_ context.Context, request *http.Request) (interface{}, error) {
	limitRaw := request.URL.Query().Get("limit")
	offsetRaw := request.URL.Query().Get("offset")

	limit, err := strconv.Atoi(limitRaw)

	if err != nil {
		return nil, err
	}

	offset, err := strconv.Atoi(offsetRaw)

	if err != nil {
		return nil, err
	}

	return &entity.Pagination{
		Limit:  limit,
		Offset: offset,
	}, nil
}

func decodeGetUserByIdHTTPRequest(_ context.Context, request *http.Request) (interface{}, error) {
	req := mux.Vars(request)
	return req.Id, nil
}

/*

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
	return req.Id, nil
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
*/
