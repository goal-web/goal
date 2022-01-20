package console

import "github.com/goal-web/contracts"

type CommandArgumentException struct {
	contracts.Exception
}

type CommandDontExistsException struct {
	contracts.Exception
}

type ScheduleEventException struct {
	contracts.Exception
}
