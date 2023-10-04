package main

import (
	"io"
	"log"
	"slices"

	pb "github.com/gabriel-dzul/grpc-go-course/calculator/proto"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Printf("Max function was invoked")
	var numbers []int32

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading the stream %v\n", err)
		}

		numbers = append(numbers, req.Number)

		err = stream.Send(&pb.MaxResponse{
			Max: slices.Max(numbers),
		})

		if err != nil {
			log.Fatalf("Error while sending data to client %v\n", err)
		}
	}
}
