package input

import (
	"github.com/qbhy/goal/console/arguments"
	"github.com/qbhy/goal/contracts"
	"strings"
)

type StringArrayInput struct {
	ArgsArray []string
}

func (this *StringArrayInput) GetCommand() string {
	if len(this.ArgsArray) > 0 {
		return this.ArgsArray[0]
	}
	return ""
}

func (this *StringArrayInput) GetArguments() contracts.CommandArguments {
	if len(this.ArgsArray) > 0 {
		args := make([]string, 0)
		options := contracts.Fields{}

		for _, arg := range this.ArgsArray[1:] {
			if strings.HasPrefix(arg, "--") {
				if argArr := strings.Split(strings.ReplaceAll(arg, "--", ""), "="); len(argArr) > 1 {
					options[argArr[0]] = argArr[1]
				} else {
					options[argArr[0]] = true
				}
			} else if strings.HasPrefix(arg, "-") {
				if argArr := strings.Split(strings.ReplaceAll(arg, "-", ""), "="); len(argArr) > 1 {
					options[argArr[0]] = argArr[1]
				} else {
					options[argArr[0]] = true
				}
			} else {
				args = append(args, arg)
			}
		}

		return arguments.NewArguments(args, options)
	}
	return arguments.NewArguments([]string{}, nil)
}
