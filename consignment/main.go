package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/Namchee/microservice-tutorial/consignment/endpoints"
	"github.com/Namchee/microservice-tutorial/consignment/pb"
	"github.com/Namchee/microservice-tutorial/consignment/repository"
	"github.com/Namchee/microservice-tutorial/consignment/service"
	"github.com/Namchee/microservice-tutorial/consignment/transports"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"google.golang.org/grpc"
)

var (
	port = ":50051"
)

func main() {
	var logger log.Logger
	logger = log.NewJSONLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	repository := repository.NewRepository()
	consignmentService := service.NewService(repository, logger)
	consignmentEndpoints := endpoints.MakeEndpoints(consignmentService)
	grpcServer := transports.NewGRPCServer(*consignmentEndpoints, logger)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	grpcListener, err := net.Listen("tcp", port)
	if err != nil {
		logger.Log("during", "Listen", "error", err)
		os.Exit(1)
	}

	go func() {
		baseServer := grpc.NewServer()
		pb.RegisterConsignmentServiceServer(baseServer, grpcServer)
		level.Info(logger).Log("msg", "Server started successfully 🚀")
		baseServer.Serve(grpcListener)
	}()

	level.Error(logger).Log("exit", <-errs)
}
