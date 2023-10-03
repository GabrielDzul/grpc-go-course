package main

import (
	"context"
	"log"
	"time"

	pb "github.com/gabriel-dzul/grpc-go-course/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Printf("doLongGreet function was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Gabriel "},
		{FirstName: "Flor"},
		{FirstName: "Pippin"},
		{FirstName: "Cow"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling LongGreet %v\n", err)
	}

	// Send a request every two seconds
	for _, req := range reqs {
		log.Printf("Sending req: %v", req)
		stream.Send(req)
		time.Sleep(2 * time.Second)
	}

	// After finishing sendin the request, get ready to receive a response
	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}
