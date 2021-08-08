package transport

import (
	"context"
	"errors"
	"log"

	endpoints "github.com/Namchee/microservice-tutorial/consignment/endpoints"
	"github.com/Namchee/microservice-tutorial/consignment/pb"
	gt "github.com/go-kit/kit/transport/grpc"
)

var (
	errAssertion = errors.New("failed to assert input type")
)

type gRPCServer struct {
	createConsignment gt.Handler
	getAll            gt.Handler
}

func NewGRPCServer(endpoints endpoints.Endpoints, logger log.Logger) pb.ConsignmentServiceServer {
	return &gRPCServer{
		createConsignment: gt.NewServer(
			endpoints.CreateConsignment,
			decodeCreateConsignmentRequest,
			encodeCreateConsignmentResponse,
		),
		getAll: gt.NewServer(
			endpoints.GetAll,
			decodeGetAllRequest,
			encodeGetAllResponse,
		),
	}
}

func (svr *gRPCServer) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
	_, res, err := svr.createConsignment.ServeGRPC(ctx, req)

	if err != nil {
		return nil, err
	}

	return res.(*pb.Response), nil
}

func (svr *gRPCServer) GetAll(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	_, res, err := svr.getAll.ServeGRPC(ctx, req)

	if err != nil {
		return nil, err
	}

	return res.(*pb.Response), nil
}

func decodeCreateConsignmentRequest(_ context.Context, request interface{}) (interface{}, error) {
	req, err := request.(*pb.Consignment)

	if err {
		return nil, errAssertion
	}

	return req, nil
}

func encodeCreateConsignmentResponse(_ context.Context, response interface{}) (interface{}, error) {
	res, err := response.(*pb.Response)

	if err {
		return nil, errAssertion
	}

	return res, nil
}

func decodeGetAllRequest(_ context.Context, request interface{}) (interface{}, error) {
	req, err := request.(*pb.GetRequest)

	if err {
		return nil, errAssertion
	}

	return req, nil
}

func encodeGetAllResponse(_ context.Context, response interface{}) (interface{}, error) {
	res, err := response.(*pb.Response)

	if err {
		return nil, errAssertion
	}

	return res, nil
}
