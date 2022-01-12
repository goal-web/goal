package console

import (
	"github.com/qbhy/goal/console/inputs"
	"github.com/qbhy/goal/contracts"
)

type ConsoleProvider func(application contracts.Application) contracts.Console

type ServiceProvider struct {
	ConsoleProvider ConsoleProvider
}

func (this *ServiceProvider) Register(application contracts.Application) {
	application.Singleton("console", func() contracts.Console {
		return this.ConsoleProvider(application)
	})
	application.Singleton("console.inputs", func() contracts.ConsoleInput {
		return inputs.NewOSArgsInput()
	})
}

func (this *ServiceProvider) Start() error {
	return nil
}

func (this *ServiceProvider) Stop() {

}
