package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/Namchee/microservice-tutorial/post/endpoints"
	"github.com/Namchee/microservice-tutorial/post/entity"
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
	messageChannel = "chan"
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

	lookupAddres := fmt.Sprintf("%s:%s", os.Getenv("LOOKUP_HOST"), os.Getenv("LOOKUP_PORT"))
	consumer, _ := nsq.NewConsumer(messageTopic, messageChannel, &nsq.Config{})
	err = consumer.ConnectToNSQD(lookupAddres)

	if err != nil {
		level.Error(logger).Log("err", "failed to connect to message queue")
		os.Exit(1)
	}

	repository := repository.NewPgPostRepository(db)
	postService := service.NewPostService(repository)
	postService = service.NewPostServiceProxy(userClient)(postService)
	postService = service.NewPostLoggingMiddleware(logger)(postService)
	postEndpoints := endpoints.NewPostEndpoint(postService)
	grpcServer := transports.NewGRPCServer(postEndpoints)

	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		var msg *entity.Message

		err := json.Unmarshal(message.Body, &msg)

		if err != nil {
			return err
		}

		user, err := strconv.ParseInt(msg.Content, 10, 32)

		_, err = postEndpoints.DeletePostByUser(context.Background(), user)

		if err != nil {
			return err
		}

		return nil
	}))

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
