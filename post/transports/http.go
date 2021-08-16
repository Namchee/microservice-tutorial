package transports

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/Namchee/microservice-tutorial/post/endpoints"
	"github.com/Namchee/microservice-tutorial/post/entity"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

type httpResponse struct {
	Data interface{} `json:"data"`
	Err  string      `json:"error"`
}

func NewHTTPRouter(endpoints *endpoints.PostEndpoints, logger log.Logger) *mux.Router {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeErrorResponse),
	}

	getPostsHandler := httptransport.NewServer(
		endpoints.GetPosts,
		decodeGetPostsHTTPRequest,
		encodeHTTPResponse,
		options...,
	)

	getPostByIdHandler := httptransport.NewServer(
		endpoints.GetPostById,
		decodeGetPostByIdHTTPRequest,
		encodeHTTPResponse,
		options...,
	)

	createPostHandler := httptransport.NewServer(
		endpoints.CreatePost,
		decodeCreatePostHTTPRequest,
		encodeHTTPResponse,
		options...,
	)

	deletePostHandler := httptransport.NewServer(
		endpoints.DeletePost,
		decodeDeletePostHTTPRequest,
		encodeHTTPResponse,
		options...,
	)

	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api").Subrouter()

	apiRouter.Methods("GET").Path("/posts").Handler(getPostsHandler)
	apiRouter.Methods("GET").Path("/post/{id:[0-9]+}").Handler(getPostByIdHandler)
	apiRouter.Methods("POST").Path("/post").Handler(createPostHandler)
	apiRouter.Methods("DELETE").Path("/post/{id:[0-9]+}").Handler(deletePostHandler)

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

func decodeGetPostsHTTPRequest(_ context.Context, request *http.Request) (interface{}, error) {
	limitRaw := request.URL.Query().Get("limit")
	offsetRaw := request.URL.Query().Get("offset")

	limit := 1
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

func decodeGetPostByIdHTTPRequest(_ context.Context, request *http.Request) (interface{}, error) {
	params := mux.Vars(request)
	id := params["id"]

	num, err := strconv.Atoi(id)

	if err != nil {
		return nil, err
	}

	return num, nil
}

func decodeCreatePostHTTPRequest(_ context.Context, request *http.Request) (interface{}, error) {
	var post *entity.Post

	if err := json.NewDecoder(request.Body).Decode(&post); err != nil {
		return nil, err
	}

	return post, nil
}

func decodeDeletePostHTTPRequest(_ context.Context, request *http.Request) (interface{}, error) {
	params := mux.Vars(request)
	id := params["id"]

	num, err := strconv.Atoi(id)

	if err != nil {
		return nil, err
	}

	return num, nil
}
