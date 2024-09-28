package main

import (
	pb "belajargolangpart2/session8introgrpc/proto/helloword/v1"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreaterServiceServer
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Fatal error to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreaterServiceServer(s, &server{})
	log.Println("server is running on port 50051")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Fatal error to server: %v", err)
	}
}

func (s *server) SayHello(ctx context.Context, in *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	return &pb.SayHelloResponse{
		Message: fmt.Sprintf("Hello world : %s", in.Name),
	}, nil
}
