package commands

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/logs"
	"github.com/qbhy/goal/console/commands"
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
