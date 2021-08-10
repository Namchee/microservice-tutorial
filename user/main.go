package main

import (
	"database/sql"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/Namchee/microservice-tutorial/user/endpoints"
	"github.com/Namchee/microservice-tutorial/user/pb"
	"github.com/Namchee/microservice-tutorial/user/repository"
	"github.com/Namchee/microservice-tutorial/user/service"
	"github.com/Namchee/microservice-tutorial/user/transports"
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

	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	fmt.Println(connStr)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		logger.Log("during", "db", "err", err)
		os.Exit(1)
	}

	repository := repository.NewPgUserRepository(db)
	userService := service.NewUserService(logger, repository)
	userEndpoint := endpoints.NewUserEndpoint(userService)
	grpcServer := transports.NewGRPCServer(userEndpoint)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	grpcListener, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.Log("during", "Listen", "err", err)
		os.Exit(1)
	}

	go func() {
		baseServer := grpc.NewServer()
		pb.RegisterUserServiceServer(baseServer, grpcServer)
		level.Info(logger).Log("msg", "Server started successfully 🚀")
		baseServer.Serve(grpcListener)
	}()

	level.Error(logger).Log("exit", <-errs)
}
