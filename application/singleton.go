package application

import (
	"github.com/goal-web/container"
	"github.com/goal-web/contracts"
)

var (
	instance contracts.Application
)

func Singleton() contracts.Application {
	if instance != nil {
		return instance
	}

	instance = &application{
		Container: container.New(),
		services:  make([]contracts.ServiceProvider, 0),
	}

	return instance
}

func SetSingleton(app contracts.Application) {
	instance = app
}

func Get(key string, args ...interface{}) interface{} {
	return Singleton().Get(key, args...)
}
