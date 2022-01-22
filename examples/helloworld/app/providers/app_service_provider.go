package providers

import (
	"github.com/goal-web/application"
	"github.com/goal-web/contracts"
	"github.com/golang-module/carbon/v2"
	"github.com/qbhy/goal/examples/helloworld/app/listeners"
)

type AppServiceProvider struct {
}

func (this AppServiceProvider) Register(app contracts.Application) {
	app.Call(func(config contracts.Config, dispatcher contracts.EventDispatcher) {
		appConfig := config.Get("app").(application.Config)
		carbon.SetLocale(appConfig.Locale)
		carbon.SetTimezone(appConfig.Timezone)

		dispatcher.Register("QUERY_EXECUTED", listeners.DebugQuery{})
	})
}

func (this AppServiceProvider) Start() error {

	return nil
}

func (this AppServiceProvider) Stop() {
}
