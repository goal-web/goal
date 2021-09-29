package config

import (
	"fmt"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/utils"
	"os"
	"strings"
)

func New(env string) contracts.Config {
	return &config{
		env:     env,
		fields:  make(contracts.Fields),
		configs: make(map[string]contracts.Config, 0),
	}
}

func WithFields(fields contracts.Fields) contracts.Config {
	return &config{
		env:     "",
		fields:  fields,
		configs: make(map[string]contracts.Config, 0),
	}
}

type config struct {
	env     string
	fields  contracts.Fields
	configs map[string]contracts.Config
}

func (this *config) Load(provider contracts.FieldsProvider) {
	utils.MergeFields(this.fields, provider.Get())
}

func (this *config) Merge(key string, config contracts.Config) {
	this.configs[key] = config
}

func (this *config) Set(key string, value interface{}) {
	this.fields[this.getKey(key)] = value
}

func (this *config) Get(key string, defaultValue ...interface{}) interface{} {

	// 环境变量优先级最高
	if envValue := os.Getenv(key); envValue != "" {
		return envValue
	}

	// 指定 env 配置次之
	if this.env != "" && !strings.Contains(key, ":") {
		if value := this.Get(fmt.Sprintf("%s:%s", this.env, key)); value != nil {
			return value
		}
	}

	if field, existsField := this.fields[key]; existsField {
		return field
	}

	keys := strings.Split(key, ".")

	if len(keys) > 1 {
		if subConfig, existsSubConfig := this.configs[keys[0]]; existsSubConfig {
			return subConfig.Get(strings.Join(keys[1:], "."), defaultValue...)
		}
	}

	if len(defaultValue) > 0 {
		return defaultValue[0]
	}

	return nil
}

func (this *config) getKey(key string) string {
	if this.env != "" {
		return utils.IfString(strings.Contains(key, ":"), key, fmt.Sprintf("%s:%s", this.env, key))
	}
	return key
}

func (this *config) GetConfig(key string) contracts.Config {
	return this.configs[key]
}

func (this *config) GetFields(key string) contracts.Fields {
	if field, isTypeRight := this.Get(key).(contracts.Fields); isTypeRight {
		return field
	}

	return nil
}

func (this *config) GetString(key string) string {
	if field, isTypeRight := this.Get(key).(string); isTypeRight {
		return field
	}

	return ""
}

func (this *config) GetInt(key string) int64 {
	if field := this.Get(key); field != nil {
		value := utils.ConvertToInt64(field, 0)
		if value != 0 { // 缓存转换结果
			this.Set(key, value)
		}
		return value
	}

	return 0
}

func (this *config) GetFloat(key string) float64 {
	if field := this.Get(key); field != nil {
		value := utils.ConvertToFloat64(field, 0)
		if value != 0 { // 缓存转换结果
			this.Set(key, value)
		}
		return value
	}

	return 0
}

func (this *config) GetBool(key string) bool {
	if field := this.Get(key); field != nil {
		switch value := field.(type) {
		case bool:
			return value
		case string:
			if value == "false" || value == "(false)" || value == "0" {
				this.Set(key, false) // 缓存结果
				return false
			}
			if value == "true" || value == "(true)" || value == "1" {
				this.Set(key, true) // 缓存结果
				return true
			}
			return len(value) > 0
		case int, float64, int64, int16, int8, float32, int32:
			return utils.ConvertToInt64(value, 0) >= 1
		}
	}

	return false
}
