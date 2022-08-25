package protobuf

import (
	"fmt"
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) WhoIs(ctx context.Context, request *MessageResponseRequest) (*MessageResponseRequest, error) {
	//request.Message = "hello world"
	fmt.Printf("request %v \n", request)
	return request, nil
}

func (s *Server) mustEmbedUnimplementedMessageServer() {
	//TODO implement me
	panic("implement me")
}

func (s *Server) SayHello(ctx context.Context, in *MessageResponseRequest) (*MessageResponseRequest, error) {
	log.Printf("Receive message body from client: %s", in.Message)
	return &MessageResponseRequest{Message: "Hello world"}, nil
}
