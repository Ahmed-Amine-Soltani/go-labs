package main

import (
	"context"
	"log"

	pb "example.com/grpc-exercice/calculator/proto"
)

func (*Server) Cal(ctx context.Context, in *pb.CRequest) (*pb.CResponse, error) {
	log.Printf("Cal was invoked with %v\n", in)

	return &pb.CResponse{Result: in.A + in.B}, nil
}
