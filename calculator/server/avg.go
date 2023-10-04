package main

import (
	"io"
	"log"

	pb "github.com/gabriel-dzul/grpc-go-course/calculator/proto"
)

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Printf("AVG function was invoked")

	var sum int32 = 0
	items := 0

	for {
		req, err := stream.Recv()

		// Check if the client has finished the data sending
		if err == io.EOF {
			return stream.SendAndClose(&pb.AverageResponse{
				Average: float64(sum) / float64(items),
			})
		}

		if err != nil {
			log.Fatalf("Error while reading the stream %v\n", err)
		}

		log.Printf("Receiving: %s\n", req)
		sum += req.Number
		items++
	}
}
