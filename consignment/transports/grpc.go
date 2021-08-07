package transport

import (
	"context"
	"log"
	"errors"

	endpoints "github.com/Namchee/microservice-tutorial/consignment/endpoints"
	"github.com/Namchee/microservice-tutorial/consignment/pb"
	gt "github.com/go-kit/kit/transport/grpc"
)

var (
	assertionError: error
)

type gRPCServer struct {
	createConsignment gt.Handler
	getAll            gt.Handler
}

func NewGRPCServer(endpoints endpoints.Endpoints, logger log.Logger) pb.ConsignmentServiceServer {
	return &gRPCServer{
		createConsignment: gt.NewServer(
			endpoints.CreateConsignment,

		)
	}
}

func decodeCreateConsignmentRequest(_ context.Context, request interface{}) (interface{}, error) {
	req, err := request.(*pb.Consignment)
	
	if err != nil {
		return nil, 
	}

	return 
}
