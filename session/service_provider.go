package session

import "github.com/qbhy/goal/contracts"

type ServiceProvider struct {
	app contracts.Application
}

func (this *ServiceProvider) Register(application contracts.Application) {
	this.app = application
}

func (this *ServiceProvider) Start() error {
	this.app.Call(func(dispatcher contracts.EventDispatcher) {
	})
	return nil
}

func (this *ServiceProvider) Stop() {
}
