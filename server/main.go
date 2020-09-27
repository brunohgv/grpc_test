package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/brunohgv/grpc_test/protobuffer"
)

type server struct {
}

func (*server) Hello(ctx context.Context, request *protobuffer.HelloRequest) (*protobuffer.HelloResponse, error) {
	result := "Hello " + request.GetName()
	response := &protobuffer.HelloResponse{
		Message: result,
	}
	return response, nil
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen:\n %v", err)
	}

	grpcServer := grpc.NewServer()
	protobuffer.RegisterHelloServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(listener); err == nil {
		log.Fatalf("Failed to serve application:\n %v", err)
	}
}
