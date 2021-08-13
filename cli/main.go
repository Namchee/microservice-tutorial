package main

import (
	"context"
	"fmt"
	"log"

	pbp "github.com/Namchee/microservice-tutorial/post/pb"
	pbu "github.com/Namchee/microservice-tutorial/user/pb"
	"google.golang.org/grpc"
)

func main() {
	userConn, userErr := grpc.Dial(":50051", grpc.WithInsecure())
	postConn, postErr := grpc.Dial(":50052", grpc.WithInsecure())

	if userErr != nil {
		log.Fatalln("failed to connect to user gRPC server")
	}

	if postErr != nil {
		log.Fatalln("failed to connect to post gRPC server")
	}

	defer userConn.Close()
	defer postConn.Close()

	userClient := pbu.NewUserServiceClient(userConn)
	postClient := pbp.NewPostServiceClient(postConn)

	users, err := userClient.GetUsers(context.Background(), &pbu.GetUsersRequest{})

	if err != nil {
		log.Fatalln("failed to call user service")
	}

	var req *pbp.CreatePostRequest = &pbp.CreatePostRequest{
		Text: "Hello World",
		User: users.Data[len(users.Data)-1].Id,
	}

	inserted, err := postClient.CreatePost(context.Background(), req)

	if err != nil {
		log.Fatalln(err.Error())
		log.Fatalln("failed to call post service")
	}

	fmt.Println(inserted.Id)
}
