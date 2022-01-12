package console

import "github.com/qbhy/goal/contracts"

type CommandArgumentException struct {
	contracts.Exception
}

type CommandDontExistsException struct {
	contracts.Exception
}
