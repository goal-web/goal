package application

import (
	"github.com/qbhy/goal/container"
	"github.com/qbhy/goal/contracts"
)

var (
	instance contracts.Application
)

func Singleton(environment string) contracts.Application {
	if instance != nil {
		return instance
	}

	instance = &application{
		environment: environment,
		Container:   container.New(),
		services:    make([]contracts.ServiceProvider, 0),
	}

	return instance
}

func SetSingleton(app contracts.Application) {
	instance = app
}
