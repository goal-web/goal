package commands

import (
	"github.com/modood/table"
	"regexp"
	"strings"
)

type Arg struct {
	Name        string
	Require     bool
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

func NewArg(name string, require bool, defaultValue interface{}) Arg {
	if names := strings.Split(name, ":"); len(names) > 1 { // 有定义描述
		return Arg{
			Name:        names[0],
			Require:     require,
			Default:     defaultValue,
			Description: names[1],
		}
	} else {
		return Arg{
			Name:        name,
			Require:     require,
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
		r := []rune(arg)
		arg = string(r[2 : len(r)-1])
		if argArr := strings.Split(arg, "="); len(argArr) > 1 {
			args = append(args, NewArg(argArr[0], false, argArr[1]))
		} else if strings.HasSuffix(arg, "?") {
			arg = string([]rune(arg)[:len(arg)-1])
			args = append(args, NewArg(arg, false, nil))
		} else if strings.HasPrefix(arg, "--") {
			arg = string([]rune(arg)[2:])
			args = append(args, NewArg(arg, false, nil))
		} else {
			args = append(args, NewArg(arg, true, nil))
		}
	}
	return cmd, args
}
