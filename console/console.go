package console

import (
	"errors"
	"github.com/qbhy/goal/contracts"
)

var CommandDontExists = errors.New("命令不存在！")

type Console struct {
	commands map[string]contracts.Command
}

func (this *Console) Call(cmd string, arguments contracts.ConsoleArguments) interface{} {
	for signature, command := range this.commands {
		if cmd == signature {
			return command.Handle(arguments)
		}
	}
	return CommandDontExists
}

func (this *Console) Run(input contracts.ConsoleInput) interface{} {
	for signature, command := range this.commands {
		if input.GetSignature() == signature {
			return command.Handle(input.GetArguments())
		}
	}
	return CommandDontExists
}
