package providers

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/goal/app/services"
	"github.com/goal-web/micro"
	"github.com/goal-web/microdemo"
	micro2 "go-micro.dev/v4"
)

func Micro() contracts.ServiceProvider {
	return &micro.ServiceProvider{
		ServiceRegister: register,
	}
}

// register 返回错误将阻止 app 启动
func register(service micro2.Service) error {
	return microdemo.RegisterHelloServiceHandler(service.Server(), new(services.HelloService))
}
