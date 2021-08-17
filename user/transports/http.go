package transports

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/Namchee/microservice-tutorial/user/endpoints"
	"github.com/Namchee/microservice-tutorial/user/entity"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type httpResponse struct {
	Data interface{} `json:"data"`
	Err  string      `json:"error"`
}

func NewHTTPRouter(endpoints *endpoints.UserEndpoints, logger log.Logger) *mux.Router {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeErrorResponse),
	}

	getUsersHandler := httptransport.NewServer(
		endpoints.GetUsers,
		decodeGetUsersHTTPRequest,
		encodeHTTPResponse,
		options...,
	)

	getUserByIdHandler := httptransport.NewServer(
		endpoints.GetUserById,
		decodeGetUserByIdHTTPRequest,
		encodeHTTPResponse,
		options...,
	)

	createUserHandler := httptransport.NewServer(
		endpoints.CreateUser,
		decodeCreateUserHTTPRequest,
		encodeHTTPResponse,
		options...,
	)

	deleteUserHandler := httptransport.NewServer(
		endpoints.DeleteUser,
		decodeDeleteUserHTTPRequest,
		encodeHTTPResponse,
		options...,
	)

	router := mux.NewRouter()
	router.Handle("/metrics", promhttp.Handler())

	apiRouter := router.PathPrefix("/api").Subrouter()

	apiRouter.Methods("GET").Path("/users").Handler(getUsersHandler)
	apiRouter.Methods("GET").Path("/user/{id:[0-9]+}").Handler(getUserByIdHandler)
	apiRouter.Methods("POST").Path("/user").Handler(createUserHandler)
	apiRouter.Methods("DELETE").Path("/user/{id:[0-9]+}").Handler(deleteUserHandler)

	return router
}

func encodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("error with nil err")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500) // just give it 500 for now
	json.NewEncoder(w).Encode(httpResponse{
		Data: nil,
		Err:  err.Error(),
	})
}

func encodeHTTPResponse(_ context.Context, w http.ResponseWriter, res interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(httpResponse{
		Data: res,
	})
}

func decodeGetUsersHTTPRequest(_ context.Context, request *http.Request) (interface{}, error) {
	limitRaw := request.URL.Query().Get("limit")
	offsetRaw := request.URL.Query().Get("offset")

	limit := 0
	offset := 0

	if len(limitRaw) > 0 {
		lim, err := strconv.Atoi(limitRaw)

		if err != nil {
			return nil, err
		}

		limit = lim
	}

	if len(offsetRaw) > 0 {
		off, err := strconv.Atoi(offsetRaw)

		if err != nil {
			return nil, err
		}

		offset = off
	}

	return &entity.Pagination{
		Limit:  limit,
		Offset: offset,
	}, nil
}

func decodeGetUserByIdHTTPRequest(_ context.Context, request *http.Request) (interface{}, error) {
	params := mux.Vars(request)
	id := params["id"]

	num, err := strconv.Atoi(id)

	if err != nil {
		return nil, err
	}

	return num, nil
}

func decodeCreateUserHTTPRequest(_ context.Context, request *http.Request) (interface{}, error) {
	var user *entity.User

	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		return nil, err
	}

	return user, nil
}

func decodeDeleteUserHTTPRequest(_ context.Context, request *http.Request) (interface{}, error) {
	params := mux.Vars(request)
	id := params["id"]

	num, err := strconv.Atoi(id)

	if err != nil {
		return nil, err
	}

	return num, nil
}
