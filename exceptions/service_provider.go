package exceptions

import (
	"github.com/qbhy/goal/contracts"
)

type ServiceProvider struct {
	DontReportExceptions []contracts.Exception
}

func (provider ServiceProvider) Register(container contracts.Container) {

	container.ProvideSingleton(func() contracts.ExceptionHandler {
		return NewDefaultHandler(provider.DontReportExceptions)
	})
}
