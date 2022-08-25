package main

import (
	"context"
	"github.com/devazizi/golang-grpc/protobuf"
	"google.golang.org/grpc"
	"log"
	"os"
)

func main() {

	if len(os.Args) == 1 {
		panic("fail")
	}

	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:3000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("not connect: %s", err)
	}

	defer conn.Close()

	c := protobuf.NewMessageClient(conn)

	response, err := c.WhoIs(context.Background(), &protobuf.MessageResponseRequest{
		Message: os.Args[1],
	})

	if err != nil {
		log.Fatalf("Error message: %s", err)
	}
	log.Printf("Response from server: %s", response.Message)
}
