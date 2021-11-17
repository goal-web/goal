package application

import (
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/utils"
	"github.com/qbhy/parallel"
	"reflect"
)

type application struct {
	contracts.Container
	services []contracts.ServiceProvider
}

func (this *application) Start() map[string]error {
	errors := make(map[string]error)
	queue := parallel.NewParallel(len(this.services))

	for _, service := range this.services {
		(func(service contracts.ServiceProvider) {
			queue.Add(func() interface{} {
				return service.Start()
			})
		})(service)
	}

	results := queue.Wait()
	for serviceIndex, result := range results {
		if err, isErr := result.(error); isErr {
			errors[utils.GetTypeKey(reflect.TypeOf(this.services[serviceIndex]))] = err
		}
	}

	return errors
}

func (this *application) Stop() {
	// 倒序执行各服务的关闭
	for serviceIndex := len(this.services) - 1; serviceIndex > -1; serviceIndex-- {
		this.services[serviceIndex].Stop()
	}
}

func (this *application) RegisterServices(services ...contracts.ServiceProvider) {
	this.services = append(this.services, services...)

	for _, service := range services {
		service.Register(this)
	}
}
