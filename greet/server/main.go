package main

import (
	pb "/greet"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	pb.GreetServiceServer
}

var addr string = "0.0.0.0:50051"

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v\n", err)
	}
}
