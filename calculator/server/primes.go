package main

import (
	"log"

	pb "github.com/gabriel-dzul/grpc-go-course/calculator/proto"
)

func (s *Server) Primes(in *pb.PrimesRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("Primes function was invoked with: %v\n", in)
	k := uint32(2)
	n := in.Number

	for n > 1 {
		if n%k == 0 {
			stream.Send(&pb.PrimesResponse{
				Factor: k,
			})
			n = n / k
		} else {
			k = k + 1
		}
	}

	return nil
}
