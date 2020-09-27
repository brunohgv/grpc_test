package main

import (
	"context"
	"log"
	"os"

	"github.com/brunohgv/grpc_test/protobuffer"
	"google.golang.org/grpc"
)

func main() {
	arguments := os.Args[1:]
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not stablish connection:\n %v", err)
	}
	defer connection.Close()

	client := protobuffer.NewHelloServiceClient(connection)
	res, err := hello(client, arguments[0])
	if err != nil {
		log.Fatalf("Error during procedure call:\n%v", err)
	}
	log.Println(res)
}

func hello(client protobuffer.HelloServiceClient, name string) (*protobuffer.HelloResponse, error) {
	request := &protobuffer.HelloRequest{
		Name: name,
	}
	res, err := client.Hello(context.Background(), request)
	return res, err
}
