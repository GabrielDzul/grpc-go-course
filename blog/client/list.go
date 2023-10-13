package main

import (
	"context"
	"io"
	"log"

	pb "github.com/gabriel-dzul/grpc-go-course/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlog(c pb.BlogServiceClient) {
	log.Printf("---ListBlog function was invoked---")

	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Unexpected error ocurred: %v\n", err)
		}

		log.Print(res)
	}
}
