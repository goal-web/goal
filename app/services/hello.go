package services

import (
	"context"
	"github.com/goal-web/microdemo"
)

type HelloService struct {
}

func (h *HelloService) SayHello(ctx context.Context, request *microdemo.HelloRequest, response *microdemo.HelloResponse) error {
	response.Message = "goal: hello " + request.Name
	return nil
}
