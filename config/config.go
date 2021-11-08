package config

import (
	"fmt"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/utils"
	"os"
	"strings"
)

func New(e string) contracts.Config {
	return &config{
		env:       e,
		envValues: make(map[string]env),
		fields:    make(contracts.Fields),
		configs:   make(map[string]contracts.Config, 0),
	}
}

func WithFields(fields contracts.Fields) contracts.Config {
	return &config{
		env:       "",
		envValues: make(map[string]env),
		fields:    fields,
		configs:   make(map[string]contracts.Config, 0),
	}
}

type config struct {
	env       string
	fields    contracts.Fields
	configs   map[string]contracts.Config
	envValues map[string]env
}
type env struct {
	value string
}

func (this *config) GetEnv(key string) string {
	if v, existsEnv := this.envValues[key]; existsEnv {
		return v.value
	} else if value := os.Getenv(key); value != "" {
		this.envValues[key] = env{value}
		return value
	}

	return ""
}

func (this *config) Fields() contracts.Fields {
	return this.fields
}

func (this *config) Load(provider contracts.FieldsProvider) {
	utils.MergeFields(this.fields, provider.Fields())
}

func (this *config) Merge(key string, config contracts.Config) {
	this.configs[key] = config
}

func (this *config) Set(key string, value interface{}) {
	this.fields[this.getKey(key)] = value
}

func (this *config) Get(key string, defaultValue ...interface{}) interface{} {

	// 环境变量优先级最高
	if envValue := this.GetEnv(key); envValue != "" {
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

	// 尝试获取 fields
	fields := contracts.Fields{}
	prefix := key + "."
	for fieldKey, fieldValue := range this.fields {
		if strings.HasPrefix(fieldKey, prefix) {
			fields[strings.Replace(fieldKey, prefix, "", 1)] = fieldValue
		}
	}
	if len(fields) > 0 {
		return fields
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

func (this *config) GetInt(key string) int {
	if field := this.Get(key); field != nil {
		value := utils.ConvertToInt(field, 0)
		if value != 0 { // 缓存转换结果
			this.Set(key, value)
		}
		return value
	}

	return 0
}
func (this *config) GetInt64(key string) int64 {
	if field := this.Get(key); field != nil {
		value := utils.ConvertToInt64(field, 0)
		if value != 0 { // 缓存转换结果
			this.Set(key, value)
		}
		return value
	}

	return 0
}

func (this *config) Unset(key string) {
	if this.env != "" && !strings.Contains(key, ":") {
		this.Unset(fmt.Sprintf("%s:%s", this.env, key))
	}
	delete(this.envValues, key)
	delete(this.fields, key)
	delete(this.configs, key)
}

func (this *config) GetFloat(key string) float32 {
	if field := this.Get(key); field != nil {
		value := utils.ConvertToFloat(field, 0)
		if value != 0 { // 缓存转换结果
			this.Set(key, value)
		}
		return float32(value)
	}

	return 0
}
func (this *config) GetFloat64(key string) float64 {
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
		result := utils.ConvertToBool(field, false)
		this.Set(key, result)
		return result
	}

	return false
}
