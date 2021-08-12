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

	var req *pb.GetUserByIdRequest = &pb.GetUserByIdRequest{
		Id: 1,
	}

	user, err := client.GetUserById(context.Background(), req)

	if err != nil {
		log.Fatalln(err.Error())
		log.Fatalln("failed to call user service")
	}

	if user == nil {
		fmt.Println("Yay!!")
	}
}
