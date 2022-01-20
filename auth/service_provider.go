package auth

import (
	"github.com/goal-web/contracts"
)

type ServiceProvider struct {
}

func (this ServiceProvider) Start() error {
	return nil
}

func (this ServiceProvider) Stop() {

}

func (this ServiceProvider) Register(container contracts.Application) {
	container.Singleton("auth", func(config contracts.Config) contracts.Auth {

		if authConfig, hasAuthConfig := config.Get("auth").(Config); hasAuthConfig {
			return &Auth{
				config:        config,
				guardDrivers:  authConfig.Guards,
				guards:        make(map[string]contracts.Guard),
				userDrivers:   authConfig.Users,
				userProviders: make(map[string]contracts.UserProvider),
			}
		}

		return &Auth{
			config:        config,
			guardDrivers:  make(map[string]contracts.GuardProvider),
			guards:        make(map[string]contracts.Guard),
			userDrivers:   make(map[string]contracts.UserProviderProvider),
			userProviders: make(map[string]contracts.UserProvider),
		}
	})
}
