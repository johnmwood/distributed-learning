package main

import (
	"context"
	"log"

	pb "github.com/johnmwood/distributed-learning/protos"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc server failed with error: %v", err)
	}
	defer conn.Close()

	client := pb.NewBoraServiceClient(conn)

	ctx := context.Background()
	request := &pb.KeyRequest{
		Key: "one",
	}

	response, err := client.GetValue(ctx, request)
	if err != nil {
		log.Printf("grpc server return err: %v", err)
	}

	log.Printf("received value from server: %s", response.Value)
}
