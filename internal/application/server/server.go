package server

import (
	"context"
	"fmt"
	"grpchello/pb"
)

type Server struct {
	pb.UnimplementedHelloServer
}

func (s *Server) SayHello(ctx context.Context, hr *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: buildMessage(hr.GetName())}, nil
}

func buildMessage(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
