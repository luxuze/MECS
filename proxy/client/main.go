package main

import (
	"context"
	"log"
	"time"

	pb "mecs/pkg/code"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "2020 code 123"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMescClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Register(ctx, &pb.Code{Code: defaultName})
	if err != nil {
		log.Fatalf("could not regist: %v", err)
	}
	log.Printf("Result: %s", r.GetResult())
}
