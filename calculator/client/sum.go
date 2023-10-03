package main

import (
	"context"
	"log"

	pb "github.com/gabriel-dzul/grpc-go-course/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		Num1: 10,
		Num2: 3,
	})

	if err != nil {
		log.Fatalf("Could not greet %v\n", err)
	}

	log.Printf("Greeting: %d\n", res.Result)
}
