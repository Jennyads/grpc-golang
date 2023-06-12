package servers

import (
	"context"
	"grpcmaratona/pb"
)

//criar objeto (class) para que possa rodar o método/função  de somar

type Math struct {
}

// toda vez que alterar m, modificará math
func (c *Math) Sum(ctx context.Context, in *pb.NewSumRequest) (*pb.NewSumResponse, error) {
	res := in.Sum.GetA() + in.Sum.GetB()

	return &pb.NewSumResponse{
		Result: res,
	}, nil
}
