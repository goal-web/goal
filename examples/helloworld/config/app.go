package config

import "github.com/qbhy/goal/config"

var (
	configs = make(map[string]config.ConfigProvider)
)

func Configs() map[string]config.ConfigProvider {
	return configs
}
