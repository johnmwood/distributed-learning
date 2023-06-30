package main

import (
	"context"
	"log"
	"net"
	"os"
	"testing"

	srv "github.com/johnmwood/distributed-learning/internal/api"
	pb "github.com/johnmwood/distributed-learning/protos"

	"google.golang.org/grpc"
)

const (
	testAddress = "localhost:50051"
)

func TestMain(m *testing.M) {
	// Start the gRPC server in a separate goroutine
	go startServer()

	// Run the tests
	exitCode := m.Run()

	// Perform any necessary cleanup or teardown here

	// Exit the test
	os.Exit(exitCode)
}

func startServer() {
	// Create a listener
	listener, err := net.Listen("tcp", testAddress)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create the gRPC server
	server := grpc.NewServer()
	boraServer := &srv.Server{
		Cache: map[string]string{
			"something": "something",
			"please":    "work",
			"one":       "two",
		},
	}
	pb.RegisterBoraServiceServer(server, boraServer)

	// Start serving gRPC requests
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func TestIntegration(t *testing.T) {
	// Create a gRPC client connection
	conn, err := grpc.Dial(testAddress, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client
	client := pb.NewBoraServiceClient(conn)

	// Prepare the request
	request := &pb.KeyRequest{
		Key: "something",
	}

	// Send the gRPC request to the server
	response, err := client.GetValue(context.Background(), request)
	if err != nil {
		t.Fatalf("Failed to get value: %v", err)
	}

	// Verify the response
	expectedValue := "something"
	if response.Value != expectedValue {
		t.Errorf("Unexpected value received. Expected: %s, Got: %s", expectedValue, response.Value)
	}
}
