package main

import (
	"context"
	"log"
	"time"

	pb "github.com/gabriel-dzul/grpc-go-course/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{
		FirstName: "Gabo",
	}

	res, err := c.GreetWithDeadline(ctx, req)

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Printf("Deadline Exceeded!")
				return
			} else {
				log.Fatalf("Unexpected gRPC error!: %v\n", e)
			}
		} else {
			log.Printf("A non gRPC error: %v\n", err)
		}
	}

	log.Printf("Received: %v\n", res.Result)
}
