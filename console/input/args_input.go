package input

import (
	"github.com/qbhy/goal/contracts"
	"os"
)

type ArgsInput struct {
}

func (this *ArgsInput) GetSignature() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}
	return ""
}

func (this *ArgsInput) GetArguments() contracts.ConsoleArguments {
	// todo: parse arguments
	return nil
}
