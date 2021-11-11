package auth

import (
	"github.com/qbhy/goal/contracts"
)

type ServiceProvider struct {
}

func (this ServiceProvider) OnStop() {

}

func (this ServiceProvider) Register(container contracts.Application) {
	container.ProvideSingleton(func(config contracts.Config) contracts.Auth {
		return &Auth{
			config:        config,
			guardDrivers:  make(map[string]contracts.GuardProvider),
			guards:        make(map[string]contracts.Guard),
			userDrivers:   make(map[string]contracts.UserProviderProvider),
			userProviders: make(map[string]contracts.UserProvider),
		}
	})
}
