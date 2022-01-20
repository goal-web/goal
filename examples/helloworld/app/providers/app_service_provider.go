package providers

import (
	"github.com/golang-module/carbon/v2"
	"github.com/qbhy/goal/application"
	"github.com/goal-web/contracts"
)

type AppServiceProvider struct {
}

func (this AppServiceProvider) Register(app contracts.Application) {
	app.Call(func(config contracts.Config) {
		appConfig := config.Get("app").(application.Config)
		carbon.SetLocale(appConfig.Locale)
		carbon.SetTimezone(appConfig.Timezone)
	})
}

func (this AppServiceProvider) Start() error {
	return nil
}

func (this AppServiceProvider) Stop() {
}
