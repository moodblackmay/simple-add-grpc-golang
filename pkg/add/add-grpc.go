package add

import (
	"context"
	"simple-adder-grpc-golang/pkg/api"
)

type AddService struct {
}

func (*AddService) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
		x := req.X
		y := req.Y
		sum := x + y
		return &api.AddResponse{Sum: sum}, nil
}