package server

import (
	"context"

	"github.com/johnmwood/distributed-learning/internal/api/stores/docustore"
	"github.com/johnmwood/distributed-learning/internal/api/stores/keyvaluestore"
	pb "github.com/johnmwood/distributed-learning/protos"
)

type Cache interface {
	Create(data any) error
	Read(key string) (any, error)
	Update(key string, data any) error
	Delete(key string) error
}

type Server struct {
	cache    keyvaluestore.KeyValueStore
	docCache docustore.DocuStore

	pb.UnimplementedBoraServiceServer
}

func (s *Server) GetValue(ctx context.Context, req *pb.KeyRequest) (*pb.ValueResponse, error) {
	value, err := s.cache.Read(req.GetKey())
	if err != nil {
		return &pb.ValueResponse{}, err
	}
	return &pb.ValueResponse{
		Value: value,
	}, nil
}

func (s *Server) GetDocument(ctx context.Context, req *pb.DocumentRequest) (*pb.DocumentResponse, error) {
	doc, err := s.docCache.Read(req.GetKey())
	if err != nil {
		return &pb.DocumentResponse{}, nil
	}
	return &pb.DocumentResponse{
		Data: doc.Data,
	}, nil
}
