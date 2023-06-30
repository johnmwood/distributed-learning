package server

import (
	"context"

	pb "github.com/johnmwood/distributed-learning/protos"
)

type Server struct {
	Cache map[string]string

	pb.UnimplementedBoraServiceServer
}

func (s *Server) GetValue(ctx context.Context, req *pb.KeyRequest) (*pb.ValueResponse, error) {
	return &pb.ValueResponse{
		Value: s.Cache[req.GetKey()],
	}, nil
}
