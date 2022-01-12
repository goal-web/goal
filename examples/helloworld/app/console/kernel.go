package console

import (
	"github.com/qbhy/goal/console"
	"github.com/qbhy/goal/console/commands"
	"github.com/qbhy/goal/contracts"
)

func NewKernel(app contracts.Application) contracts.Console {
	return &Kernel{console.NewKernel(app, []console.CommandProvider{
		commands.Runner,
	})}
}

type Kernel struct {
	*console.Kernel
}
