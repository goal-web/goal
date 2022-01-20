package commands

import (
	"github.com/goal-web/supports/utils"
	"github.com/modood/table"
	"regexp"
	"strings"
)

type ArgType int

const (
	RequiredArg ArgType = iota + 1 // 必要参数
	OptionalArg                    // 可选参数
	Option                         // 选项
)

type Arg struct {
	Name        string
	Type        ArgType
	Default     interface{}
	Description string
}

type Args []Arg

func (args Args) Help() string {
	if len(args) > 0 {
		return table.Table(args)
	}
	return "该命令无参数"
}

func NewArg(name string, argType ArgType, defaultValue interface{}) Arg {
	if names := strings.Split(name, ":"); len(names) > 1 { // 有定义描述
		return Arg{
			Name:        names[0],
			Type:        argType,
			Default:     defaultValue,
			Description: names[1],
		}
	} else {
		return Arg{
			Name:        name,
			Type:        argType,
			Default:     defaultValue,
			Description: "",
		}
	}
}

func ParseSignature(signature string) (string, Args) {
	cmd := strings.Split(signature, " ")[0]
	reg, _ := regexp.Compile(" {([^{}]*)}")
	args := make(Args, 0)

	for _, arg := range reg.FindAllString(signature, -1) {
		arg = utils.SubString(arg, 2, -1)
		if argArr := strings.Split(arg, "="); len(argArr) > 1 { // {name=goal} / {--name=goal}
			if strings.HasPrefix(argArr[0], "--") {
				args = append(args, NewArg(utils.SubString(argArr[0], 2, 0), Option, argArr[1]))
			} else {
				args = append(args, NewArg(argArr[0], OptionalArg, argArr[1]))
			}
		} else if strings.HasSuffix(arg, "?") { // {name?}
			arg = utils.SubString(arg, 0, -1)
			args = append(args, NewArg(arg, OptionalArg, nil))
		} else if strings.HasPrefix(arg, "--") { // {--name}
			arg = utils.SubString(arg, 2, 0)
			args = append(args, NewArg(arg, Option, nil))
		} else { // {name}
			args = append(args, NewArg(arg, RequiredArg, nil))
		}
	}
	return cmd, args
}
