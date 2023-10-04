package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/gabriel-dzul/grpc-go-course/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Printf("doMax function was invoked")

	stream, err := c.Max(context.Background())

	if err != nil {
		log.Fatalf("Error while calling Max: %v\n", err)
	}

	reqs := []*pb.MaxRequest{
		{Number: 2},
		{Number: -52},
		{Number: 8},
		{Number: 34},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Send request: %v\n", req)
			stream.Send(req)
			time.Sleep(2 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving: %v\n", res)
				break
			}

			log.Printf("Max at the moment: %d\n", res.Max)
		}

		close(waitc)
	}()

	<-waitc
}
