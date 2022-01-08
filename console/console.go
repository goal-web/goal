package console

import "github.com/qbhy/goal/contracts"

type Console struct {
	app      contracts.Application
	commands map[string]contracts.Command
}

func (this *Console) Call(command string, args ...interface{}) error {
	//TODO implement me
	panic("implement me")
}
