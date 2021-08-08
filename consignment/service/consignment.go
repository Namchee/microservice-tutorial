package service

import (
	"context"

	"github.com/Namchee/microservice-tutorial/consignment/pb"
	"github.com/Namchee/microservice-tutorial/consignment/repository"
	"github.com/go-kit/kit/log"
)

type service struct {
	repo   repository.Repository
	logger log.Logger
}

type Service interface {
	CreateConsignment(context.Context, *pb.Consignment) (*pb.Response, error)
	GetAll(context.Context, *pb.GetRequest) (*pb.Response, error)
}

func NewService(repo repository.Repository, logger log.Logger) *service {
	return &service{
		repo:   repo,
		logger: logger,
	}
}

func (svc *service) CreateConsignment(ctx context.Context, consignment *pb.Consignment) (*pb.Response, error) {
	consignment, err := svc.repo.CreateConsignment(consignment)

	if err != nil {
		return nil, err
	}

	var res *pb.Response
	res.Created = true
	res.Consignment = consignment

	return res, nil
}

func (svc *service) GetAll(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	consignments, err := svc.repo.GetAll(req)

	if err != nil {
		return nil, err
	}

	var res *pb.Response
	res.Consignments = consignments
	return res, nil
}
