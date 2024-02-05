package main

import (
	"context"
	"log"

	pb "github.com/Dhana777/grpc-pro/greet/proto"
)

func DoGreet(c pb.GreetServiceClient) {

	log.Println("DoGreet was invoked")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "dhana",
	})

	if err != nil {
		log.Fatalf("could not greet %v\n", err)
	}

	log.Printf("Greeting %s\n", res.Result)
}
