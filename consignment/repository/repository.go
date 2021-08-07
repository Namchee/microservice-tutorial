package repository

import (
	"sync"

	"github.com/Namchee/microservice-tutorial/consignment/pb"
)

type Repository interface {
	CreateConsignment(*pb.Consignment) (*pb.Consignment, error)
	GetAll(*pb.GetRequest) ([]*pb.Consignment, error)
}

type dummyRepository struct {
	mu           sync.RWMutex
	consignments []*pb.Consignment
}

func (repo *dummyRepository) CreateConsignment(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.mu.Lock()
	updatedConsignments := append(repo.consignments, consignment)
	repo.consignments = updatedConsignments
	repo.mu.Unlock()
	return consignment, nil
}

func (repo *dummyRepository) GetAll(req *pb.GetRequest) ([]*pb.Consignment, error) {
	return repo.consignments, nil
}

func NewRepository() *dummyRepository {
	return &dummyRepository{}
}
