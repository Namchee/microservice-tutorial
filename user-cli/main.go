package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Namchee/microservice-tutorial/user/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalln("failed to connect to gRPC server")
	}

	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	var req *pb.CreateUserRequest = &pb.CreateUserRequest{
		Username: "namchee",
		Name:     "Budi",
		Bio:      "Hello World",
	}

	var id int
	inserted, err := client.CreateUser(context.Background(), req)

	if err != nil {
		log.Fatalln(err.Error())
		log.Fatalln("failed to call user service")
	}

	if inserted != nil {
		id = int(inserted.Id)
	}

	var res *pb.GetUserByIdRequest = &pb.GetUserByIdRequest{
		Id: int32(id),
	}

	user, err := client.GetUserById(context.Background(), res)

	if err != nil || user.Data == nil {
		panic("user should exist")
	}

	fmt.Println(user.Data.Username)
}
