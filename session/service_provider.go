package session

import "github.com/goal-web/contracts"

type ServiceProvider struct {
	app contracts.Application
}

func (this *ServiceProvider) Register(application contracts.Application) {
	this.app = application

	application.Bind("session", func(config contracts.Config, request contracts.HttpRequest) contracts.Session {
		if session, isSession := request.Get("session").(contracts.Session); isSession {
			return session
		}
		session := New(config.GetString("session.name"), config.GetString("session.id"), config, request)

		request.Set("session", session)
		return session
	})
}

func (this *ServiceProvider) Start() error {
	this.app.Call(func(dispatcher contracts.EventDispatcher) {
		dispatcher.Register("REQUEST_AFTER", &RequestAfterListener{})
	})
	return nil
}

func (this *ServiceProvider) Stop() {
}
