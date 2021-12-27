package tests

import (
	"errors"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/events"
	"github.com/qbhy/goal/exceptions"
	"github.com/qbhy/goal/logs"
	"github.com/stretchr/testify/assert"
	"testing"
)

const Demo string = "demo"

type DemoEvent struct {
}

func (d DemoEvent) Sync() bool {
	return true
}

func (d DemoEvent) Event() string {
	return Demo
}

type DemoPanicListener struct {
}

func (d DemoPanicListener) Handle(event contracts.Event) {
	panic(errors.New("报错啦"))
}

type DemoListener struct {
}

func (d DemoListener) Handle(event contracts.Event) {
	logs.Default().Info("正常处理事件")
}

func TestEvent(t *testing.T) {
	dispatcher := events.NewDispatcher(exceptions.DefaultExceptionHandler{})
	dispatcher.Register(Demo, DemoPanicListener{})
	dispatcher.Register(Demo, DemoListener{})
	dispatcher.Dispatch(DemoEvent{})

	assert.Nil(t, recover())
}
