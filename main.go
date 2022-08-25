package main

import (
	"fmt"
	"github.com/devazizi/golang-grpc/protobuf"
	"google.golang.org/grpc"
	"log"
	"net"
)

func netGrpc() {

	fmt.Println("start project ... ")
	lis, err := net.Listen("tcp", "0.0.0.0:3000")
	if err != nil {
		log.Fatalf("fail to listen on 3000 %v \n", err)
	}

	s := protobuf.Server{}

	grpcServer := grpc.NewServer()

	protobuf.RegisterMessageServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("fail to run grpc serve %v \n", lis)
	}

}

func main() {

	//app := fiber.New(fiber.Config{
	//	AppName: "fiber classic",
	//})
	//
	//app.Get("/", controller.Home())

	netGrpc()

	//log.Fatal(app.Listen("0.0.0.0:5000"))
}
