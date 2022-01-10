package input

import (
	"github.com/qbhy/goal/contracts"
	"os"
)

type ArgsInput struct {
	StringArrayInput
}

func NewOSArgsInput() contracts.ConsoleInput {
	return &ArgsInput{StringArrayInput{ArgsArray: os.Args[1:]}}
}
