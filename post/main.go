package main

import (
	"database/sql"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/Namchee/microservice-tutorial/post/endpoints"
	"github.com/Namchee/microservice-tutorial/post/pb"
	"github.com/Namchee/microservice-tutorial/post/repository"
	"github.com/Namchee/microservice-tutorial/post/service"
	"github.com/Namchee/microservice-tutorial/post/transports"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"google.golang.org/grpc"

	_ "github.com/lib/pq"
)

func main() {
	var logger log.Logger
	logger = log.NewJSONLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		level.Error(logger).Log("msg", "failed to connect to db")
		os.Exit(1)
	}

	repository := repository.NewPgPostRepository(db)
	postService := service.NewPostService(logger, repository)
	postEndpoints := endpoints.NewPostEndpoint(logger, postService)
	grpcServer := transports.NewGRPCServer(postEndpoints)

	errs := make(chan error)
	c := make(chan os.Signal)
	go func() {
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	grpcListener, err := net.Listen("tcp", ":50051")
	if err != nil {
		level.Error(logger).Log("msg", "failed to start grpc server")
		os.Exit(1)
	}

	go func() {
		baseServer := grpc.NewServer()
		pb.RegisterPostServiceServer(baseServer, grpcServer)
		level.Info(logger).Log("msg", "server started successfully 🚀")
		baseServer.Serve(grpcListener)
	}()

	level.Error(logger).Log("exit", <-errs)
}
