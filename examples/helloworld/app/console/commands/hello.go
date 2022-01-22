package commands

import (
	"github.com/goal-web/console/commands"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/logs"
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
