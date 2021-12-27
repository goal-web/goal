package contracts

type EventName string

type Event interface {
	Name() EventName
}

type SyncEvent interface {
	Event
	Sync() bool
}

type EventListener interface {
	Handle(event Event)
}

type EventDispatcher interface {
	Register(name EventName, listener EventListener)
	Dispatch(event Event)
}
