package console

import (
	"errors"
	"fmt"
	"github.com/modood/table"
	"github.com/qbhy/goal/contracts"
)

type CommandProvider func(application contracts.Application) contracts.Command

var CommandDontExists = errors.New("命令不存在！")

const logoText = "  ▄████  ▒█████   ▄▄▄       ██▓    \n ██▒ ▀█▒▒██▒  ██▒▒████▄    ▓██▒    \n▒██░▄▄▄░▒██░  ██▒▒██  ▀█▄  ▒██░    \n░▓█  ██▓▒██   ██░░██▄▄▄▄██ ▒██░    \n░▒▓███▀▒░ ████▓▒░ ▓█   ▓██▒░██████▒\n ░▒   ▒ ░ ▒░▒░▒░  ▒▒   ▓▒█░░ ▒░▓  ░\n  ░   ░   ░ ▒ ▒░   ▒   ▒▒ ░░ ░ ▒  ░\n░ ░   ░ ░ ░ ░ ▒    ░   ▒     ░ ░   \n      ░     ░ ░        ░  ░    ░  ░\n                                   "

type Kernel struct {
	commands map[string]contracts.Command
}

func NewKernel(app contracts.Application, commandProviders []CommandProvider) *Kernel {
	commands := make(map[string]contracts.Command)

	for _, commandProvider := range commandProviders {
		command := commandProvider(app)
		commands[command.GetName()] = command
	}

	return &Kernel{commands}
}

type CommandItem struct {
	Command     string
	Signature   string
	Description string
}

func (this Kernel) Help() {
	cmdTable := make([]CommandItem, 0)
	for _, command := range this.commands {
		cmdTable = append(cmdTable, CommandItem{
			Command:     command.GetName(),
			Signature:   command.GetSignature(),
			Description: command.GetDescription(),
		})
	}
	fmt.Println(logoText)
	table.Output(cmdTable)
}

func (this *Kernel) Call(cmd string, arguments contracts.CommandArguments) interface{} {
	if cmd == "" {
		this.Help()
		return nil
	}
	for signature, command := range this.commands {
		if cmd == signature {
			if arguments.Exists("h") || arguments.Exists("help") {
				fmt.Println(logoText)
				fmt.Printf(" %s 命令：%s\n", command.GetName(), command.GetDescription())
				fmt.Println(command.GetHelp())
				return nil
			}
			if err := command.InjectArguments(arguments); err != nil {
				fmt.Println(err.Error())
				fmt.Println(command.GetHelp())
				return nil
			}
			return command.Handle()
		}
	}
	return CommandDontExists
}

func (this *Kernel) Run(input contracts.ConsoleInput) interface{} {
	return this.Call(input.GetCommand(), input.GetArguments())
}
