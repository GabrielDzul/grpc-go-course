package main

import (
	"context"
	"log"
	"time"

	pb "github.com/gabriel-dzul/grpc-go-course/calculator/proto"
)

func doAvg(c pb.CalculatorServiceClient) {
	log.Printf("doAvg function was invoked")

	reqs := []*pb.AverageRequest{
		{Number: 1},
		{Number: 43},
		{Number: 456},
		{Number: 4},
	}

	stream, err := c.Avg(context.Background())

	if err != nil {
		log.Fatalf("Error while calling LongGreet %v\n", err)
	}

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

	log.Printf("AVG: %f\n", res.Average)
}
