package providers

import (
	"github.com/goal-web/contracts"
)

type Console struct {
	Commands []contracts.CommandProvider
	Schedule func(schedule contracts.Schedule)
}

func NewConsoleService(commands []contracts.CommandProvider, schedule func(schedule2 contracts.Schedule)) contracts.ServiceProvider {
	return Console{
		Commands: commands,
		Schedule: schedule,
	}
}

func (c Console) Register(application contracts.Application) {
	application.Call(func(console contracts.Console, schedule contracts.Schedule) {
		for _, provider := range c.Commands {
			console.RegisterCommand(provider)
		}

		c.Schedule(schedule)
	})
}

func (c Console) Start() error {
	return nil
}

func (c Console) Stop() {
}
