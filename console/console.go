package console

import (
	"errors"
	"fmt"
	"github.com/modood/table"
	"github.com/qbhy/goal/contracts"
)

var CommandDontExists = errors.New("命令不存在！")

type Console struct {
	commands map[string]contracts.Command
}

type CommandItem struct {
	Command     string
	Description string
}

func (this Console) Help() {
	cmdTable := make([]CommandItem, 0)
	for _, command := range this.commands {
		cmdTable = append(cmdTable, CommandItem{
			Command:     command.GetName(),
			Description: command.GetDescription(),
		})
	}
	fmt.Println("支持的命令：")
	table.Output(cmdTable)
}

func (this *Console) Call(cmd string, arguments contracts.CommandArguments) interface{} {
	if cmd == "" {
		this.Help()
		return nil
	}
	for signature, command := range this.commands {
		if cmd == signature {
			if arguments.Exists("h") || arguments.Exists("help") {
				fmt.Println(command.GetDescription())
				fmt.Println(command.GetHelp())
				return nil
			}
			return command.Handle(arguments)
		}
	}
	return CommandDontExists
}

func (this *Console) Run(input contracts.ConsoleInput) interface{} {
	return this.Call(input.GetCommand(), input.GetArguments())
}
