package events

import "github.com/qbhy/goal/contracts"

type ServiceProvider struct {
}

func (this ServiceProvider) OnStop() {

}

func (this ServiceProvider) OnStart() error {
	return nil
}


func (provider ServiceProvider) Register(container contracts.Application) {
	container.ProvideSingleton(func(handler contracts.ExceptionHandler) contracts.EventDispatcher {
		return NewDispatcher(handler)
	})
}
