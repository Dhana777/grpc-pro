package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/Dhana777/grpc-pro/greet/proto"
)

var addr string = "0.0.0.0:50001"

func main() {

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect %v\n", err)
	}

	c := pb.NewGreetServiceClient(conn)
	DoGreet(c)

	defer conn.Close()

}
