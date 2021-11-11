package config

import (
	"github.com/qbhy/goal/contracts"
)

type ServiceProvider struct {
	Env   string
	Paths []string
	Sep   string
}

func (this ServiceProvider) OnStop() {

}

func (this ServiceProvider) OnStart() error {
	return nil
}

func (provider ServiceProvider) Register(application contracts.Application) {
	application.ProvideSingleton(func() contracts.Config {

		configInstance := New(provider.Env)

		configInstance.Load(EnvFieldsProvider{
			Paths: provider.Paths,
			Sep:   provider.Sep,
		})

		return configInstance
	})
}
