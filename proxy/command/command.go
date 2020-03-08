package command

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "mecs/pkg/code"
	"net"
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

func WaitingUserCommand() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Printf("用户指令处理下发服务启动成功，监听端口%v", port)
	}
	s := grpc.NewServer()
	pb.RegisterMescServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
