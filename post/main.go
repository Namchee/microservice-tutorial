package main

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Namchee/microservice-tutorial/post/endpoints"
	"github.com/Namchee/microservice-tutorial/post/mq"
	"github.com/Namchee/microservice-tutorial/post/pb"
	"github.com/Namchee/microservice-tutorial/post/repository"
	"github.com/Namchee/microservice-tutorial/post/service"
	"github.com/Namchee/microservice-tutorial/post/transports"
	upb "github.com/Namchee/microservice-tutorial/user/pb"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	nsq "github.com/nsqio/go-nsq"
	"google.golang.org/grpc"

	_ "github.com/lib/pq"
)

const (
	messageTopic   = "post"
	messageChannel = "channel"
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

	grpcStr := fmt.Sprintf("%s:%s", os.Getenv("USER_HOST"), os.Getenv("USER_PORT"))

	userConn, err := grpc.Dial(grpcStr, grpc.WithInsecure())

	if err != nil {
		level.Error(logger).Log("err", "failed to connect to user service")
		os.Exit(1)
	}
	userClient := upb.NewUserServiceClient(userConn)

	lookupAddress := fmt.Sprintf("%s:%s", os.Getenv("LOOKUP_HOST"), os.Getenv("LOOKUP_PORT"))
	consumer, _ := nsq.NewConsumer(messageTopic, messageChannel, nsq.NewConfig())

	if err != nil {
		level.Error(logger).Log("err", "failed to connect to message queue")
		os.Exit(1)
	}

	repository := repository.NewPgPostRepository(db)

	var postService service.PostService
	postService = service.NewPostService(repository)
	postService = service.NewPostServiceProxy(userClient)(postService)
	postService = service.NewPostLoggingMiddleware(logger)(postService)

	postEndpoints := endpoints.NewPostEndpoint(postService)

	errs := make(chan error)
	c := make(chan os.Signal)

	grpcServer := transports.NewGRPCServer(postEndpoints)

	consumer.AddHandler(&mq.MessageHandler{
		Endpoints: postEndpoints,
	})
	err = consumer.ConnectToNSQLookupd(lookupAddress)
	if err != nil {
		level.Error(logger).Log("err", "failed to connect to lookup service")
		os.Exit(1)
	}

	grpcListener, err := net.Listen("tcp", ":50051")
	if err != nil {
		level.Error(logger).Log("msg", "failed to start grpc server")
		os.Exit(1)
	}

	httpListener := transports.NewHTTPRouter(postEndpoints, logger)
	httpServer := http.Server{
		Addr:    ":8080",
		Handler: httpListener,
	}

	go func() {
		baseServer := grpc.NewServer()
		pb.RegisterPostServiceServer(baseServer, grpcServer)
		level.Info(logger).Log("msg", "server started successfully 🚀")
		baseServer.Serve(grpcListener)
	}()

	go func() {
		err = httpServer.ListenAndServe()

		if err != nil {
			level.Error(logger).Log("err", "failed to start http server")
			os.Exit(1)
		}
	}()

	go func() {
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	level.Error(logger).Log("exit", <-errs)
}
