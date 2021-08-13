package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Namchee/microservice-tutorial/post/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50052", grpc.WithInsecure())

	if err != nil {
		log.Fatalln("failed to connect to gRPC server")
	}

	defer conn.Close()

	client := pb.NewPostServiceClient(conn)

	var req *pb.CreatePostRequest = &pb.CreatePostRequest{
		Text: "Hello World",
		User: 1,
	}

	var id int
	inserted, err := client.CreatePost(context.Background(), req)

	if err != nil {
		log.Fatalln(err.Error())
		log.Fatalln("failed to call post service")
	}

	if inserted != nil {
		id = int(inserted.Id)
	}

	var res *pb.GetPostByIdRequest = &pb.GetPostByIdRequest{
		Id: int32(id),
	}

	post, err := client.GetPostById(context.Background(), res)

	if err != nil || post.Data == nil {
		panic("post should exist")
	}

	fmt.Println(post.Data.Text)
}
