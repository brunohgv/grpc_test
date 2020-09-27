package main

import (
	"context"
	"log"

	"github.com/brunohgv/grpc_test/protobuffer"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not stablish connection:\n %v", err)
	}
	defer connection.Close()

	client := protobuffer.NewHelloServiceClient(connection)

	request := &protobuffer.HelloRequest{
		Name: "Bruno",
	}

	res, err := client.Hello(context.Background(), request)
	if err != nil {
		log.Fatalf("Error during execution:\n %v", err)
	}
	log.Println(res)
}
