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

	if len(users.Data) > 0 {
		log.Fatalln("wrong user database state")
	}

	user, err := userClient.CreateUser(context.Background(), &pbu.CreateUserRequest{
		Username: "namchee",
		Name:     "Namchee",
		Bio:      "Hello World",
	})

	if err != nil {
		log.Fatalln("failed to insert user")
	}

	id := user.Id

	_, err = postClient.CreatePost(context.Background(), &pbp.CreatePostRequest{
		Text: "This is a test post",
		User: id,
	})

	if err != nil {
		log.Fatalln("failed to insert new post")
	}

	deleted, err := userClient.DeleteUser(context.Background(), &pbu.DeleteUserRequest{
		Id: id,
	})

	if err != nil {
		log.Fatalln("failed to delete user")
	}

	fmt.Println(deleted.User.Id)
}
