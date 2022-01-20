package config

import (
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
	"github.com/qbhy/goal/application"
	"github.com/qbhy/goal/config"
	"os"
)

var (
	configs = make(map[string]config.ConfigProvider)
)

func Configs() map[string]config.ConfigProvider {
	return configs
}

func init() {
	hostname, _ := os.Hostname()
	userHome, _ := os.UserHomeDir()
	configs["app"] = func(env contracts.Env) interface{} {
		return application.Config{
			ServerId: fmt.Sprintf("%s:%s.%s", hostname, userHome, utils.RandStr(6)),
			Name:     env.GetString("app.name"),
			Debug:    env.GetBool("app.debug"),
			Timezone: env.GetString("app.timezone"),
			Env:      env.GetString("app.env"),
			Locale:   env.GetString("app.locale"),
			Key:      env.GetString("app.key"),
		}
	}
}
