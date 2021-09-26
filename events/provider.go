package events

import "github.com/qbhy/goal/contracts"

type ServiceProvider struct {
}

func (s ServiceProvider) Register(container contracts.Container) {
	container.ProvideSingleton(func(handler contracts.ExceptionHandler) contracts.EventDispatcher {
		return NewDispatcher(handler)
	})
}
