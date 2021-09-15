package events

import (
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/exceptions"
)

var (
	listenerPool = make(map[contracts.EventName][]contracts.EventListener)
)

// SetEventListeners 批量设置监听器
func SetEventListeners(listeners map[contracts.EventName][]contracts.EventListener) {
	listenerPool = listeners
}

// RegisterListener 注册事件监听器
func RegisterListener(event contracts.EventName, listener contracts.EventListener) {
	listenerPool[event] = append(listenerPool[event], listener)
}

// Dispatch 触发事件
func Dispatch(event contracts.Event) {
	// 加个协程
	defer func() {
		if err := recover(); err != nil {
			exceptions.Handle(EventException{
				exception: exceptions.ResolveException(err),
				event:     event,
			})
		}
	}()
	for _, listener := range listenerPool[event.Name()] {
		listener.Handle(event)
	}
}
