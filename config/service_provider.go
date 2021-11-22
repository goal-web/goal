package config

import (
	"github.com/qbhy/goal/contracts"
)

type ServiceProvider struct {
	Env   string
	Paths []string
	Sep   string
}

func (this ServiceProvider) Stop() {

}

func (this ServiceProvider) Start() error {
	return nil
}

func (provider ServiceProvider) Register(application contracts.Application) {
	application.Singleton("env", func() contracts.Env {
		return NewEnv(provider.Paths, provider.Sep)
	})

	application.Singleton("config", func(env contracts.Env) contracts.Config {

		configInstance := New(provider.Env)

		configInstance.Load(env)

		return configInstance
	})
}
