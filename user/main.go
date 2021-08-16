package main

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"
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
	"github.com/nsqio/go-nsq"
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
		level.Error(logger).Log("err", "failed to connect to database")
		os.Exit(1)
	}

	producer, err := nsq.NewProducer(fmt.Sprintf("%s:%s", os.Getenv("NSQ_HOST"), os.Getenv("NSQ_PORT")), nsq.NewConfig())

	if err != nil {
		level.Error(logger).Log("err", "failed to connect to message queue")
		os.Exit(1)
	}

	repository := repository.NewPgUserRepository(db)
	userService := service.NewUserService(repository)
	userService = service.NewLoggingMiddleware(logger)(userService)
	publisher := service.NewNSQPublisher(producer)
	userEndpoint := endpoints.NewUserEndpoint(logger, userService, publisher)

	errs := make(chan error)
	c := make(chan os.Signal)
	go func() {
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	grpcServer := transports.NewGRPCServer(userEndpoint)
	grpcListener, err := net.Listen("tcp", ":50051")
	if err != nil {
		level.Error(logger).Log("failed to open grpc server")
		os.Exit(1)
	}
	httpRouter := transports.NewHTTPRouter(userEndpoint, logger)
	httpServer := http.Server{
		Addr:    ":8080",
		Handler: httpRouter,
	}

	httpServer.ListenAndServe()

	go func() {
		baseServer := grpc.NewServer()
		pb.RegisterUserServiceServer(baseServer, grpcServer)
		level.Info(logger).Log("msg", "Server started successfully 🚀")
		baseServer.Serve(grpcListener)
	}()

	level.Error(logger).Log("exit", <-errs)
}
