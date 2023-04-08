package commands

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/commands"
	"github.com/goal-web/supports/logs"
)

func NewHello(app contracts.Application) contracts.Command {
	return &Hello{
		Command: commands.Base("hello {say}", "打印 hello goal"),
	}
}

type Hello struct {
	commands.Command
}

func (hello Hello) Handle() any {
	logs.Default().Info("hello goal " + hello.GetString("say"))
	return nil
}
