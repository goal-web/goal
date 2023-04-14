package controllers

import (
	"context"
	"github.com/goal-web/microdemo"
)

func RpcService(helloService microdemo.HelloService) any {
	res, err := helloService.SayHello(context.Background(), &microdemo.HelloRequest{Name: "测试"})
	if err != nil {
		return err.Error()
	}
	return res.Message
}
