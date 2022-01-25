package config

import (
	"fmt"
	"github.com/goal-web/application"
	"github.com/goal-web/config"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
	"os"
)

var (
	configs = make(map[string]config.Provider)
)

func Configs() map[string]config.Provider {
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
