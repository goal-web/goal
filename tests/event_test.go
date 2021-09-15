package tests

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"goal/contracts"
	"goal/events"
	"goal/logs"
	"testing"
)

const Demo contracts.EventName = "demo"

type DemoEvent struct {
}

func (d DemoEvent) Name() contracts.EventName {
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
	events.SetEventListeners(map[contracts.EventName][]contracts.EventListener{
		Demo: {
			DemoListener{},
			DemoPanicListener{},
		},
	})
	events.Dispatch(DemoEvent{})

	assert.Nil(t, recover())
}
