package console

import (
	"github.com/qbhy/goal/console"
	"github.com/qbhy/goal/console/commands"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/logs"
)

func NewKernel(app contracts.Application) contracts.Console {
	return &Kernel{console.NewKernel(app, []console.CommandProvider{
		commands.Runner,
	})}
}

type Kernel struct {
	*console.Kernel
}

func (this *Kernel) Schedule(schedule contracts.Schedule) {
	schedule.Call(func() {
		logs.Default().Info("每10秒钟打印 goal")
	}).EveryTenSeconds()

	schedule.Call(func() {
		logs.Default().Info("周日每5秒钟打印 周日愉快")
	}).EveryFiveSeconds().Sundays()
}
