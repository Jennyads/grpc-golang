package main

import (
	"google.golang.org/grpc"
	"grpcmaratona/pb"
	"grpcmaratona/servers"
	"log"
	"net"
)

func main() {
	grpcServer := grpc.NewServer()
	pb.RegisterMathServiceServer(grpcServer, &servers.Math{})

	listener, err := net.Listen("tcp", ":50055")
	if err != nil {
		log.Fatalf("Cannot start server: %v", err)
	}

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatalf("Cannot start server: %v", err)
	}
}
