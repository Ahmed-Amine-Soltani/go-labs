package main

import (
	"context"
	"log"

	pb "example.com/grpc-exercice/calculator/proto"
)

func doCal(c pb.CServiceClient) {
	log.Println("doCal was invoked")
	r, err := c.Cal(context.Background(), &pb.CRequest{A: 5, B: 3})
	if err != nil {
		log.Fatalf("Could not cal: %v\n", err)
	}
	log.Printf("Resultat: %d\n", r.Result)
}
