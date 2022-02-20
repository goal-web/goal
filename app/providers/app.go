package providers

import (
	"github.com/goal-web/application"
	"github.com/goal-web/contracts"
	"github.com/goal-web/goal/app/listeners"
	"github.com/golang-module/carbon/v2"
)

type App struct {
}

func (this App) Register(app contracts.Application) {
	app.Call(func(config contracts.Config, dispatcher contracts.EventDispatcher) {
		appConfig := config.Get("app").(application.Config)
		carbon.SetLocale(appConfig.Locale)
		carbon.SetTimezone(appConfig.Timezone)

		dispatcher.Register("QUERY_EXECUTED", listeners.DebugQuery{})
	})
}

func (this App) Start() error {

	return nil
}

func (this App) Stop() {
}
