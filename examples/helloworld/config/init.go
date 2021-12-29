package config

import "github.com/qbhy/goal/contracts"

type ConfigProvider func(env contracts.Env) interface{}

var (
	configs = make(map[string]ConfigProvider)
)

func RegisterConfigs(config contracts.Config, env contracts.Env) {
	for key, conf := range configs {
		config.Set(key, conf(env))
	}
}
