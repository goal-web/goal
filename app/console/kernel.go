package console

import (
	"github.com/goal-web/application/commands"
	"github.com/goal-web/console"
	"github.com/goal-web/contracts"
	commands2 "github.com/goal-web/goal/app/console/commands"
)

func Service() contracts.ServiceProvider {
	return &console.ServiceProvider{ConsoleProvider: NewKernel}
}

func NewKernel(app contracts.Application) contracts.Console {
	return &Kernel{console.NewKernel(app, []contracts.CommandProvider{
		commands.Runner,
		commands2.NewHello,
	}), app}
}

type Kernel struct {
	*console.Kernel
	app contracts.Application
}

func (this *Kernel) Exists(schedule string) bool {
	return true
}

func (this *Kernel) Schedule(schedule contracts.Schedule) {
	//schedule.Call(func() {
	//	logs.Default().Info("周日每5秒钟打印 周日愉快")
	//}).EveryFiveSeconds().Sundays()
}
