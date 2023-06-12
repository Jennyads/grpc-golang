package main

import (
	"context"
	"google.golang.org/grpc"
	"grpcmaratona/pb"
	"io"
	"log"
)

func main() {
	connection, err := grpc.Dial("localhost:50055", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC server: %v", &err)
	}
	client := pb.NewMathServiceClient(connection)

	defer connection.Close()

	go Sum(10, 55, err, client)
	FibonacciGrpc(err, client)

}
func FibonacciGrpc(err error, client pb.MathServiceClient) {
	request := &pb.FibonacciRequest{
		N: 10,
	}
	responseStream, err := client.Fibonacci(context.Background(), request)
	if err != nil {
		log.Fatalf("Error during the request: %v", err)
	}

	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error receiving streaming: %v", err)
		}

		log.Printf("Fibonacci: %v", stream.GetResult())
	}
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
