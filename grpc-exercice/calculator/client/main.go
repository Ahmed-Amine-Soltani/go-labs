package main

import (
	"log"

	pb "example.com/grpc-exercice/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "0.0.0.0:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect %w\n", err)
	}
	defer conn.Close()
	c := pb.NewCServiceClient(conn)
	doCal(c)
}
