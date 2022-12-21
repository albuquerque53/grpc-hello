package handler

import (
	"fmt"
	"grpchello/internal/application/server"
	"grpchello/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

func RunServer() {
	fmt.Println("Running gRPC server")

	ls, err := net.Listen("tcp", "localhost:9000")

	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &server.Server{})

	if err := s.Serve(ls); err != nil {
		log.Fatalf("failed do serve: %v", err)
	}
}
