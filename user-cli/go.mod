module github.com/Namchee/microservice-tutorial/user-cli

go 1.16

replace github.com/Namchee/microservice-tutorial/user => ../user

require (
	github.com/Namchee/microservice-tutorial/user v0.0.0-00010101000000-000000000000
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20210811021853-ddbe55d93216 // indirect
	google.golang.org/grpc v1.40.0
)
