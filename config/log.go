package config

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/log"
	"github.com/goal-web/supports/utils"
)

func init() {
	configs["log"] = func(env contracts.Env) any {
		return log.Config{
			Default: utils.StringOr(env.GetString("log.default"), "stack"),
			Channels: map[string]contracts.Fields{
				"console": {
					log.KeyDriver: "console",
					log.KeyLevel:  utils.StringOr(env.GetString("log.level"), "debug"),
					log.KeyFormat: "console",
				},
				"file": {
					log.KeyDriver: "file",
					log.KeyLevel:  "info",
					log.KeyFormat: "json",
					log.KeyPath:   utils.StringOr(env.GetString("log.path"), "storage/logs/app.log"), // 按天切分为 app-YYYYMMDD.log
				},
				"stack": {
					log.KeyDriver: "stack",
					log.KeyLevel:  "debug",
					// "channels" 是 stack 驱动的保留配置键（log.KeyChannels），声明被组合扇出的子通道
					log.KeyChannels: []string{"console", "file"},
				},
			},
		}
	}
}
