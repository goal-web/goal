package contracts

type Event interface {
	Event() string
}

type SyncEvent interface {
	Event
	Sync() bool
}

type EventListener interface {
	Handle(event Event)
}

type EventDispatcher interface {
	Register(name string, listener EventListener)
	Dispatch(event Event)
}
