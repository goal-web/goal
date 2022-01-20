package supports

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
	"strings"
)

type InstanceGetter func(key string) interface{}

type BaseFields struct { // 具体方法
	contracts.FieldsProvider // 抽象方法，继承 interface

	Getter InstanceGetter // 如果有设置 getter ，优先使用 getter
}

func (this *BaseFields) get(key string) interface{} {
	if this.Getter != nil {
		if value := this.Getter(key); value != nil && value != "" {
			return value
		}
	}
	return this.Fields()[key]
}

func (this *BaseFields) StringOption(key string, defaultValue string) string {
	if value := this.get(key); value != nil && value != "" {
		return utils.ConvertToString(value, defaultValue)
	}
	return defaultValue
}

func (this *BaseFields) Int64Option(key string, defaultValue int64) int64 {
	if value := this.get(key); value != nil && value != "" {
		return utils.ConvertToInt64(value, defaultValue)
	}
	return defaultValue
}

func (this *BaseFields) IntOption(key string, defaultValue int) int {
	if value := this.get(key); value != nil && value != "" {
		return utils.ConvertToInt(value, defaultValue)
	}
	return defaultValue
}

func (this *BaseFields) Float64Option(key string, defaultValue float64) float64 {
	if value := this.get(key); value != nil && value != "" {
		return utils.ConvertToFloat64(value, defaultValue)
	}
	return defaultValue
}

func (this *BaseFields) FloatOption(key string, defaultValue float32) float32 {
	if value := this.get(key); value != nil && value != "" {
		return utils.ConvertToFloat(value, defaultValue)
	}
	return defaultValue
}

func (this *BaseFields) BoolOption(key string, defaultValue bool) bool {
	if value := this.get(key); value != nil && value != "" {
		return utils.ConvertToBool(value, defaultValue)
	}
	return defaultValue
}

func (this *BaseFields) FieldsOption(key string, defaultValue contracts.Fields) contracts.Fields {
	if value := this.get(key); value != nil && value != "" {
		if fields, err := utils.ConvertToFields(value); err == nil {
			return fields
		}
	}
	fields := contracts.Fields{}
	for fieldKey, value := range this.Fields() {
		if strings.HasPrefix(fieldKey, key+".") {
			fields[strings.ReplaceAll(fieldKey, key+".", "")] = value
		}
	}
	if len(fields) > 0 {
		return fields
	}
	return defaultValue
}

func (this *BaseFields) GetString(key string) string {
	return this.StringOption(key, "")
}

func (this *BaseFields) GetInt64(key string) int64 {
	return this.Int64Option(key, 0)
}

func (this *BaseFields) GetInt(key string) int {
	return this.IntOption(key, 0)
}

func (this *BaseFields) GetFloat64(key string) float64 {
	return this.Float64Option(key, 0)
}

func (this *BaseFields) GetFloat(key string) float32 {
	return this.FloatOption(key, 0)
}

func (this *BaseFields) GetBool(key string) bool {
	return this.BoolOption(key, false)
}

func (this *BaseFields) GetFields(key string) contracts.Fields {
	return this.FieldsOption(key, contracts.Fields{})
}
