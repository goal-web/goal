package session

import "github.com/qbhy/goal/contracts"

type ServiceProvider struct {
}

func (this *ServiceProvider) Register(application contracts.Application) {
	panic("implement me")
}

func (this *ServiceProvider) Start() error {
	//TODO implement me
	panic("implement me")
}

func (this *ServiceProvider) Stop() {
	//TODO implement me
	panic("implement me")
}
