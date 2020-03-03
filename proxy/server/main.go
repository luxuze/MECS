package main

import (
	"context"
	"log"
	"net"

	pb "mecs/pkg/code"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedMescServer
}

func (s *server) Register(ctx context.Context, in *pb.Code) (*pb.Response, error) {
	log.Printf("Register request: %v", in.GetCode())
	return &pb.Response{Result: "Registed success: " + in.GetCode()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMescServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
