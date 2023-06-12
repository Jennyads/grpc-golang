package main

import (
	"context"
	"google.golang.org/grpc"
	"grpcmaratona/pb"
	"log"
)

func main() {
	connection, err := grpc.Dial("localhost:50055", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC server: %v", &err)
	}
	client := pb.NewMathServiceClient(connection)

	defer connection.Close()

	Sum(10, 55, err, client)
}

func Sum(a float32, b float32, err error, client pb.MathServiceClient) {

	request := &pb.NewSumRequest{
		Sum: &pb.Sum{
			A: a,
			B: b,
		},
	}

	res, err := client.Sum(context.Background(), request)
	if err != nil {
		log.Fatalf("Error during the request: %v", err)
	}

	log.Println(res)
}
