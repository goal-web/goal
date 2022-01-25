package tests

import (
	"errors"
	"github.com/goal-web/contracts"
	"github.com/goal-web/events"
	"github.com/goal-web/supports/logs"
	"github.com/goal-web/goal/exceptions"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const Demo string = "demo"

type DemoEvent struct {
	IsSync bool
}

func (d DemoEvent) Sync() bool {
	return d.IsSync
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
	dispatcher.Dispatch(DemoEvent{true})

	time.Sleep(time.Second)
	assert.Nil(t, recover())
}
