package scheduling

import (
	"github.com/qbhy/goal/console/inputs"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/logs"
)

type Schedule struct {
	timezone string
	mutex    *Mutex
	app      contracts.Application

	events []contracts.ScheduleEvent
}

func (this *Schedule) Call(callback interface{}, args ...interface{}) contracts.CallbackEvent {
	event := NewCallbackEvent(this.mutex, func() []interface{} {
		return this.app.Call(callback, args...)
	}, this.timezone)
	this.events = append(this.events, event)
	return event
}

func (this *Schedule) Command(command contracts.Command, args ...string) contracts.CommandEvent {
	args = append([]string{command.GetName()}, args...)
	input := inputs.StringArray(args)
	err := command.InjectArguments(input.GetArguments())
	if err != nil {
		logs.WithError(err).Debug("command 参数错误")
		panic(err)
	}
	event := NewCommandEvent(command.GetName(), this.mutex, func(console contracts.Console) []interface{} {
		return []interface{}{command.Handle()}
	}, this.timezone)
	this.events = append(this.events, event)
	return event
}

func (this *Schedule) Exec(command string, args ...string) contracts.CommandEvent {
	//TODO implement me
	panic("implement me")
}
