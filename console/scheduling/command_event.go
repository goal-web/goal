package scheduling

import (
	"fmt"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/utils"
)

func NewCommandEvent(command string, mutex *Mutex, callback interface{}, timezone string) contracts.CommandEvent {
	return &CommandEvent{
		Event:   NewEvent(mutex, callback, timezone),
		command: command,
	}
}

type CommandEvent struct {
	*Event
	command string
}

func (this *CommandEvent) Command(command string) contracts.CommandEvent {
	this.command = command
	return this
}

func (this *CommandEvent) MutexName() string {
	if this.mutexName == "" {
		return fmt.Sprintf("goal/schedule-%s", utils.Md5(this.expression+this.command))
	}
	return this.mutexName
}
