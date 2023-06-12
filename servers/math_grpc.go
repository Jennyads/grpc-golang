package servers

import (
	"context"
	"grpcmaratona/pb"
	"time"
)

//criar objeto (class) para que possa rodar o método/função  de somar

type Math struct {
}

// toda vez que alterar m, modificará math
func (m *Math) Sum(ctx context.Context, in *pb.NewSumRequest) (*pb.NewSumResponse, error) {
	res := in.Sum.GetA() + in.Sum.GetB()

	return &pb.NewSumResponse{
		Result: res,
	}, nil
}

func (m *Math) Fibonacci(in *pb.FibonacciRequest, stream pb.MathService_FibonacciServer) error {
	n := in.GetN()

	var i int32
	for i = 1; i <= n; i++ {
		res := &pb.FibonacciResponse{
			Result: FibonacciRecursion(i),
		}
		stream.Send(res)
		time.Sleep(2000 * time.Millisecond)
	}
	return nil

}

func FibonacciRecursion(n int32) int32 {
	if n <= 1 {
		return n
	}
	return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
}
