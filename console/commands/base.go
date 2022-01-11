package commands

import (
	"errors"
	"fmt"
	"github.com/qbhy/goal/contracts"
)

type Base struct {
	contracts.CommandArguments
	Signature   string
	Description string
	Name        string
	Help        string
	args        []Arg
}

func BaseCommand(signature, description string) Base {
	name, args := ParseSignature(signature)
	return Base{
		Signature:   signature,
		Description: description,
		Name:        name,
		Help:        args.Help(),
		args:        args,
	}
}

func (this *Base) InjectArguments(arguments contracts.CommandArguments) error {
	argIndex := 0
	for _, arg := range this.args {
		switch arg.Type {
		case RequiredArg:
			argValue := arguments.GetArg(argIndex)
			if argValue == "" {
				if this.Exists(arg.Name) {
					arguments.SetOption(arg.Name, arguments.Fields()[arg.Name])
				} else {
					return errors.New(fmt.Sprintf("缺少必要参数：%s - %s", arg.Name, arg.Description))
				}
			} else {
				arguments.SetOption(arg.Name, argValue)
			}
			argIndex++
		case OptionalArg:
			argValue := arguments.GetArg(argIndex)
			if argValue == "" {
				arguments.SetOption(arg.Name, arg.Default)
			} else {
				arguments.SetOption(arg.Name, argValue)
			}
			argIndex++
		case Option:
			if !arguments.Exists(arg.Name) && arg.Default != nil {
				arguments.SetOption(arg.Name, arg.Default)
			}
		}
	}

	this.CommandArguments = arguments
	return nil
}

func (this *Base) GetSignature() string {
	return this.Signature
}
func (this *Base) GetDescription() string {
	return this.Description
}
func (this *Base) GetName() string {
	return this.Name
}
func (this *Base) GetHelp() string {
	return this.Help
}
