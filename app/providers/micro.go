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
	withServer bool
}

func NewMicro(withServer bool) contracts.ServiceProvider {
	return &MicroServiceProvider{micro.ServiceProvider{
		ServiceRegister: register,
	}, withServer}
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

func (provider *MicroServiceProvider) Start() error {
	if provider.withServer {
		return provider.ServiceProvider.Start()
	}
	return nil
}
func (provider *MicroServiceProvider) Stop() {
	if provider.withServer {
		provider.ServiceProvider.Stop()
	}
}
