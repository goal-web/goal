package routes

import (
	"github.com/qbhy/goal/contracts"
)

type ServiceProvider struct {
}

func (provider ServiceProvider) Register(container contracts.Container) {
	routerInstance := New(container)
	container.ProvideSingleton(func() contracts.Router {
		return routerInstance
	}, "router")
}
