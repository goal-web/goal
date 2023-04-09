package providers

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/goal/app/services"
	"github.com/goal-web/micro"
	"github.com/goal-web/microdemo"
	micro2 "go-micro.dev/v4"
)

type MicroServiceProvider struct {
	micro.ServiceProvider
}

func Micro() contracts.ServiceProvider {
	return &MicroServiceProvider{micro.ServiceProvider{
		ServiceRegister: register,
	}}
}

// register 返回错误将阻止 app 启动
func register(service micro2.Service) error {
	return microdemo.RegisterHelloServiceHandler(service.Server(), new(services.HelloService))
}

func (provider *MicroServiceProvider) Register(app contracts.Application) {
	provider.ServiceProvider.Register(app)

	app.Singleton("hello", func(service micro2.Service) microdemo.HelloService {
		return microdemo.NewHelloService("hello", service.Client())
	})
}
