package exceptions

import (
	"github.com/goal-web/contracts"
)

type ServiceProvider struct {
	DontReportExceptions []contracts.Exception
}

func (provider ServiceProvider) Start() error {
	return nil
}

func (provider ServiceProvider) Stop() {
}

func (provider ServiceProvider) Register(container contracts.Application) {

	container.Singleton("exception.handler", func() contracts.ExceptionHandler {
		return NewDefaultHandler(provider.DontReportExceptions)
	})
}
