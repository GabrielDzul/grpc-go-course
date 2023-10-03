package main

import (
	"context"
	"io"
	"log"

	pb "github.com/gabriel-dzul/grpc-go-course/calculator/proto"
)

func doPrimes(c pb.CalculatorServiceClient) {
	log.Println("doPrimes was invoked")
	//Build request
	req := &pb.PrimesRequest{
		Number: 120,
	}
	//Make request to the primes endpoint
	stream, err := c.Primes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling Primes: %v\n", err)
	}

	for {
		//Receive response
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("Primes: %d\n", msg.Factor)
	}
}
