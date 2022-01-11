package console

import (
	"github.com/qbhy/goal/console/inputs"
	"github.com/qbhy/goal/contracts"
)

type CommandProvider func(application contracts.Application) contracts.Command

type ServiceProvider struct {
	Commands []CommandProvider
}

func (this *ServiceProvider) Register(application contracts.Application) {
	application.Singleton("console", func() contracts.Console {
		commands := make(map[string]contracts.Command)

		for _, commandProvider := range this.Commands {
			command := commandProvider(application)
			commands[command.GetName()] = command
		}

		return &Console{commands}
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
