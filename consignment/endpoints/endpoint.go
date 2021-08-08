package endpoints

import (
	"context"

	"github.com/Namchee/microservice-tutorial/consignment/pb"
	"github.com/Namchee/microservice-tutorial/consignment/service"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateConsignment endpoint.Endpoint
	GetAll            endpoint.Endpoint
}

func MakeEndpoints(svc service.Service) *Endpoints {
	return &Endpoints{
		CreateConsignment: makeCreateConsignmentEndpoint(svc),
		GetAll:            makeGetAllEndpoint(svc),
	}
}

func makeCreateConsignmentEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, _ := request.(pb.Consignment)

		result, serviceError := svc.CreateConsignment(ctx, &req)

		if serviceError != nil {
			return nil, serviceError
		}

		return result, nil
	}
}

func makeGetAllEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, _ := request.(pb.GetRequest)

		result, serviceError := svc.GetAll(ctx, &req)

		if serviceError != nil {
			return nil, serviceError
		}

		return result, nil
	}
}
