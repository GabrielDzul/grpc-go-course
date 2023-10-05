package main

import (
	"log"
	"time"

	pb "github.com/gabriel-dzul/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "localhost:50051"

func main() {
	tls := true // change if necessary
	opts := []grpc.DialOption{}

	if tls {
		cerFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(cerFile, "")

		if err != nil {
			log.Fatalf("Error while loading the cert: %s\n", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	conn, err := grpc.Dial(addr, opts...)

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)
	//doGreet(c)
	//doGreetManyTimes(c)
	//doLongGreet(c)
	//doGreetEveryone(c)
	doGreetWithDeadline(c, 5*time.Second)

}
