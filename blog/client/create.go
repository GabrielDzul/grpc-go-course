package main

import (
	"context"
	"log"

	pb "github.com/gabriel-dzul/grpc-go-course/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Printf("---CreateBlog function was invoked---")

	blog := &pb.Blog{
		AuthorId: "Gabriel_id",
		Title:    "My first blog",
		Content:  "Content of the first blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog was created %s\n", res.Id)

	return res.Id
}
