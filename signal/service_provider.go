package signal

import (
	"github.com/goal-web/contracts"
	"os"
	"os/signal"
	"syscall"
)

type ServiceProvider struct {
	signals       []os.Signal
	signalChannel chan os.Signal
	app           contracts.Application
}

func (this *ServiceProvider) Register(application contracts.Application) {
	this.app = application
}

func (this *ServiceProvider) Start() (err error) {
	this.signalChannel = make(chan os.Signal)
	signal.Notify(this.signalChannel, this.signals...)
	for sign := range this.signalChannel {
		this.app.Call(func(dispatcher contracts.EventDispatcher) {
			dispatcher.Dispatch(&Received{sign})
		})

		switch sign {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			this.app.Stop()
		}
	}

	return err
}

func (this *ServiceProvider) Stop() {
	close(this.signalChannel)
}
