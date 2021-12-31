package config

import (
	"github.com/qbhy/goal/contracts"
)

type ConfigProvider func(env contracts.Env) interface{}

type ServiceProvider struct {
	app             contracts.Application
	Env             string
	Paths           []string
	Sep             string
	ConfigProviders map[string]ConfigProvider
}

func (this *ServiceProvider) Stop() {

}

func (this *ServiceProvider) Start() error {
	envInstance := this.app.Get("env").(contracts.Env)
	configInstance := this.app.Get("config").(contracts.Config)
	for key, provider := range this.ConfigProviders {
		configInstance.Set(key, provider(envInstance))
	}
	return nil
}

func (this *ServiceProvider) Register(application contracts.Application) {
	this.app = application
	application.Singleton("env", func() contracts.Env {
		return NewEnv(this.Paths, this.Sep)
	})

	application.Singleton("config", func(env contracts.Env) contracts.Config {

		configInstance := New(this.Env)

		configInstance.Load(env)

		return configInstance
	})
}
