package client

import (
	"context"
	"grpchello/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CallHelloGRPC(address string, name string) (*pb.HelloResponse, error) {
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	cl := pb.NewHelloClient(conn)

	req := &pb.HelloRequest{
		Name: name,
	}

	res, err := cl.SayHello(context.Background(), req)

	if err != nil {
		return nil, err
	}

	return res, nil
}
