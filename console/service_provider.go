package console

import "github.com/qbhy/goal/contracts"

type CommandProvider func(application contracts.Application) contracts.Command

type ServiceProvider struct {
	Commands []CommandProvider
}

func (this *ServiceProvider) Register(application contracts.Application) {
	application.Singleton("console", func() contracts.Console {
		commands := make(map[string]contracts.Command)

		for _, commandProvider := range this.Commands {
			command := commandProvider(application)
			commands[command.GetSignature()] = command
		}

		return &Console{commands}
	})
}

func (this *ServiceProvider) Start() error {
	return nil
}

func (this *ServiceProvider) Stop() {

}
