package handler

import (
	"golang.org/x/net/context"
	"Core/grpc-gen/core"
	"math/rand"
	"time"
)
type Handler struct {
	core.CoreServer
}

func (*Handler) Calculate(ctx context.Context, in *core.CalculateRequest) (*core.CalculateResponse, error) {
	time.Sleep(time.Duration(in.X))
	randNum := rand.Intn(int(in.X))
	return &core.CalculateResponse{Result: int64(randNum)}, nil
}