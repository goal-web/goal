package events

import (
	"github.com/goal-web/contracts"
	"github.com/qbhy/goal/exceptions"
)

func NewDispatcher(handler contracts.ExceptionHandler) contracts.EventDispatcher {
	return &EventDispatcher{
		eventListenersMap: make(map[string][]contracts.EventListener, 0),
		exceptionHandler:  handler,
	}
}

type EventDispatcher struct {
	eventListenersMap map[string][]contracts.EventListener

	// 依赖异常处理器
	exceptionHandler contracts.ExceptionHandler
}

func (dispatcher EventDispatcher) Register(name string, listener contracts.EventListener) {
	dispatcher.eventListenersMap[name] = append(dispatcher.eventListenersMap[name], listener)
}

func (dispatcher EventDispatcher) Dispatch(event contracts.Event) {
	// 处理异常
	defer func() {
		if err := recover(); err != nil {
			go func() {
				dispatcher.exceptionHandler.Handle(EventException{
					exception: exceptions.ResolveException(err),
					event:     event,
				})
			}()
		}
	}()

	if _, isSync := event.(contracts.SyncEvent); isSync {
		// 同步执行事件
		for _, listener := range dispatcher.eventListenersMap[event.Event()] {
			listener.Handle(event)
		}
	} else {
		// 协程执行
		go func() {
			for _, listener := range dispatcher.eventListenersMap[event.Event()] {
				listener.Handle(event)
			}
		}()
	}
}
