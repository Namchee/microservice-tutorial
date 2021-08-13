module github.com/Namchee/microservice-tutorial/post

go 1.16

replace github.com/Namchee/microservice-tutorial/user => ../user

require (
	github.com/Namchee/microservice-tutorial/user v0.0.0-00010101000000-000000000000
	github.com/go-kit/kit v0.11.0
	github.com/lib/pq v1.10.2
	google.golang.org/grpc v1.39.1
	google.golang.org/protobuf v1.27.1
)
