package main

import (
	"context"
	"log"

	pb "github.com/gabriel-dzul/grpc-go-course/blog/proto"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Printf("---ReadBlog function was invoked---")

	req := &pb.BlogId{Id: id}

	res, err := c.ReadBlog(context.Background(), req)

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog was read %v\n", res)
	return res
}
