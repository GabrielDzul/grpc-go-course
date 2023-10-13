package main

import (
	"context"
	"log"

	pb "github.com/gabriel-dzul/grpc-go-course/blog/proto"
)

func DeleteBlog(c pb.BlogServiceClient, id string) {
	log.Printf("---DeleteBlog function was invoked---")

	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Println("Blog was deleted")
}
