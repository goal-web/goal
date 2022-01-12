package commands

import (
	"github.com/qbhy/goal/console/commands"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/logs"
)

func NewHello(app contracts.Application) contracts.Command {
	return &Hello{
		Base: commands.BaseCommand("hello {say}", "打印 hello goal"),
	}
}

type Hello struct {
	commands.Base
}

func (this Hello) Handle() interface{} {
	logs.Default().Info("hello goal " + this.GetString("say"))
	return nil
}
