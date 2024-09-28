package main

import (
	pb "belajargolangpart2/session8introgrpc/proto/helloword/v1"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Did not connect : %v", err)
	}

	defer conn.Close()

	greeterClient := pb.NewGreaterServiceClient(conn)

	res, err := greeterClient.SayHello(context.Background(), &pb.SayHelloRequest{
		Name: "Anto",
	})

	if err != nil {
		log.Fatalf("Could not greeting: %v", err)
	}

	log.Printf("greeting: %s", res.Message)
}
