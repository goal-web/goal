package config

import (
	"github.com/qbhy/goal/contracts"
)

type ServiceProvider struct {
	Env   string
	Paths []string
	Sep   string
}

func (provider ServiceProvider) Register(container contracts.Container) {
	container.ProvideSingleton(func() contracts.Config {

		configInstance := New(provider.Env)

		configInstance.Load(EnvFieldsProvider{
			Paths: provider.Paths,
			Sep:   provider.Sep,
		})

		return configInstance
	})
}
