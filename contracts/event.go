package contracts

type EventName string

type Event interface {
	Name() EventName
}

type EventListener interface {
	Handle(event Event)
}
