package config

import (
	"github.com/goal-web/application"
	"github.com/goal-web/contracts"
)

var configs = make(map[string]contracts.ConfigProvider)

func GetConfigProviders() map[string]contracts.ConfigProvider {
	return configs
}

func init() {
	configs["app"] = func(env contracts.Env) any {
		return application.Config{
			Name:     env.GetString("app.name"),
			Debug:    env.GetBool("app.debug"),
			Timezone: env.GetString("app.timezone"),
			Env:      env.GetString("app.env"),
			Locale:   env.GetString("app.locale"),
			Key:      env.GetString("app.key"),
		}
	}
}
