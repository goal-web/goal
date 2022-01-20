package scheduling

import (
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/qbhy/goal/supports/utils"
)

func NewCallbackEvent(mutex *Mutex, callback interface{}, timezone string) contracts.CallbackEvent {
	return &CallbackEvent{
		Event:       NewEvent(mutex, callback, timezone),
		description: "",
	}
}

type CallbackEvent struct {
	*Event
	description string
}

func (this *CallbackEvent) Description(description string) contracts.CallbackEvent {
	this.description = description
	return this
}

func (this *CallbackEvent) MutexName() string {
	if this.mutexName == "" {
		return fmt.Sprintf("goal.schedule-%s", utils.Md5(this.expression+this.description))
	}
	return this.mutexName
}
