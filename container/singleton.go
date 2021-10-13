package container

import "github.com/qbhy/goal/contracts"

var (
	instance contracts.Container
)

func Singleton() contracts.Container {
	if instance != nil {
		return instance
	}

	instance = New()

	return instance
}

func SetSingleton(container contracts.Container) {
	instance = container
}
