package main

import (
	"log"
	"net"

	srv "github.com/johnmwood/distributed-learning/internal/api"
	pb "github.com/johnmwood/distributed-learning/protos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	address = "localhost:50051"
)

func main() {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	boraServer := srv.Server{
		Cache: map[string]string{
			"something": "something",
			"please":    "work",
			"one":       "two",
		},
	}

	pb.RegisterBoraServiceServer(server, &boraServer)
	reflection.Register(server)

	log.Printf("Starting gRPC server at: %s", address)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
