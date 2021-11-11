package application

import (
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/parallel"
)

type application struct {
	contracts.Container
	services []contracts.ServiceProvider
}

func (this *application) Start() (errors []error) {
	queue := parallel.NewParallel(len(this.services))

	for _, service := range this.services {
		(func(service contracts.ServiceProvider) {
			queue.Add(func() interface{} {
				return service.OnStart()
			})
		})(service)
	}

	for _, result := range queue.Wait() {
		if err, isErr := result.(error); isErr {
			errors = append(errors, err)
		}
	}

	return
}

func (this *application) OnStop() {
	for _, service := range this.services {
		service.OnStop()
	}
}

func (this *application) RegisterServices(services ...contracts.ServiceProvider) {
	this.services = append(this.services, services...)

	for _, service := range services {
		service.Register(this)
	}
}
