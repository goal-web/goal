package config

import (
	"github.com/qbhy/goal/application"
	"github.com/qbhy/goal/config"
	"github.com/qbhy/goal/contracts"
)

var (
	configs = make(map[string]config.ConfigProvider)
)

func Configs() map[string]config.ConfigProvider {
	return configs
}

func init() {
	configs["app"] = func(env contracts.Env) interface{} {
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
