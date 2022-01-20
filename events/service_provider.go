package events

import "github.com/goal-web/contracts"

type ServiceProvider struct {
}

func (this ServiceProvider) Stop() {

}

func (this ServiceProvider) Start() error {
	return nil
}

func (provider ServiceProvider) Register(container contracts.Application) {
	container.Singleton("events", func(handler contracts.ExceptionHandler) contracts.EventDispatcher {
		return NewDispatcher(handler)
	})
}
