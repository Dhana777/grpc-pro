package main

import (
	"context"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/Dhana777/grpc-pro/greet/proto"
)

var addr string = "0.0.0.0:50051"

func main() {

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect %v\n", err)
	}

	c := pb.NewGreetServiceClient(conn)

	// doGreet(c)
	//	doGreetMany(c)
	doLongGreet(c)

	defer conn.Close()

}

func doGreet(c pb.GreetServiceClient) {

	log.Println("DoGreet was invoked")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "dhana",
	})

	if err != nil {
		log.Fatalf("could not greet %v\n", err)
	}

	log.Printf("Greeting %s\n", res.Result)
}

func doGreetMany(c pb.GreetServiceClient) {
	log.Println("Do Greet many times was invoked")

	Stream, err := c.GreetManyTimes(context.Background(), &pb.GreetRequest{
		FirstName: "dhana",
	})

	if err != nil {
		log.Fatalf("could not greet %v\n", err)
	}
	for {
		result, err := Stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading the stream %v\n", err)
		}
		log.Printf("printing the result %+v", result.Result)
	}

}

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("Do Greet many times was invoked")

	Stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("could not greet %v\n", err)
	}
	names := []string{"dhana", "chitty"}
	for i := 0; i <= len(names)-1; i++ {
		Stream.Send(&pb.GreetRequest{FirstName: names[i]})
		time.Sleep(1 * time.Second)
	}

	res, err := Stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receving response %v", err)
	}
	log.Printf("result %v", res)
}
