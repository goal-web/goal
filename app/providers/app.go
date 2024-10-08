package providers

import (
	"github.com/goal-web/application"
	"github.com/goal-web/contracts"
	"github.com/goal-web/goal/app/models"
	"github.com/goal-web/migration/migrate"
	"github.com/golang-module/carbon/v2"
)

type appServiceProvider struct {
	serviceProviders []contracts.ServiceProvider
}

func NewApp() contracts.ServiceProvider {
	return &appServiceProvider{
		serviceProviders: []contracts.ServiceProvider{},
	}
}

func (app appServiceProvider) Register(instance contracts.Application) {
	instance.RegisterServices(app.serviceProviders...)

	instance.Call(func(config contracts.Config, dispatcher contracts.EventDispatcher, factory contracts.DBFactory) {
		appConfig := config.Get("app").(application.Config)
		carbon.SetLocale(appConfig.Locale)
		carbon.SetTimezone(appConfig.Timezone)
		instance.Instance("app.env", appConfig.Env)

		migrate.Auto(
			factory,
			models.UserMigrator(),
		)
	})
}

func (app appServiceProvider) Start() error {
	return nil
}

func (app appServiceProvider) Stop() {
}
