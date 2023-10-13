package main

import (
	"context"
	"log"

	pb "github.com/gabriel-dzul/grpc-go-course/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Printf("---Updatelog function was invoked---")

	newBlog := &pb.Blog{
		Id:       id,
		Title:    "A cool title",
		AuthorId: "Not Gabo Id",
		Content:  "A cool content for this blog",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog was updated")
}
