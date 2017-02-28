package main

import (
    "log"
    "net"
    "github.com/hellogrpc/messages"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

const port = ":5000"
type server struct{}

func(s *server) SayHello(ctx context.Context ,request *messages.HelloRequest)(response *messages.HelloResponse, err error) {
 	return &messages.HelloResponse{Message:"Hello " + request.Name}, nil
} 

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()

	messages.RegisterHelloServiceServer(s, &server{})

	s.Serve(lis)
}
