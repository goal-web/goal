package inputs

import (
	"github.com/qbhy/goal/contracts"
	"os"
)

type ArgsInput struct {
	StringArrayInput
}

func NewOSArgsInput() contracts.ConsoleInput {
	return &ArgsInput{StringArray(os.Args[1:])}
}
