package providers

import (
	"github.com/goal-web/contracts"
	events2 "github.com/goal-web/database/events"
	"github.com/goal-web/goal/app/listeners"
)

type EventsServiceProvider struct {
	listeners map[contracts.Event][]contracts.EventListener
}

func NewEvents() contracts.ServiceProvider {
	return &EventsServiceProvider{
		listeners: map[contracts.Event][]contracts.EventListener{
			&events2.QueryExecuted{}: {listeners.DebugQuery{}},
		},
	}
}

func (provider EventsServiceProvider) Stop() {

}

func (provider EventsServiceProvider) Start() error {
	return nil
}

func (provider EventsServiceProvider) Register(container contracts.Application) {
	container.Call(func(dispatcher contracts.EventDispatcher) {
		for event, items := range provider.listeners {
			for _, listener := range items {
				dispatcher.Register(event.Event(), listener)
			}
		}
	})
}
