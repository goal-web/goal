package supports

import (
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/utils"
)

type BaseFields struct { // 具体方法
	contracts.FieldsProvider // 抽象方法，继承 interface
}

func (this *BaseFields) StringOption(key string, defaultValue string) string {
	if value, exists := this.Fields()[key]; exists {
		return utils.ConvertToString(value, defaultValue)
	}
	return defaultValue
}

func (this *BaseFields) Int64Option(key string, defaultValue int64) int64 {
	if value, exists := this.Fields()[key]; exists {
		return utils.ConvertToInt64(value, defaultValue)
	}
	return defaultValue
}

func (this *BaseFields) IntOption(key string, defaultValue int) int {
	if value, exists := this.Fields()[key]; exists {
		return utils.ConvertToInt(value, defaultValue)
	}
	return defaultValue
}

func (this *BaseFields) Float64Option(key string, defaultValue float64) float64 {
	if value, exists := this.Fields()[key]; exists {
		return utils.ConvertToFloat64(value, defaultValue)
	}
	return defaultValue
}

func (this *BaseFields) FloatOption(key string, defaultValue float32) float32 {
	if value, exists := this.Fields()[key]; exists {
		return utils.ConvertToFloat(value, defaultValue)
	}
	return defaultValue
}

func (this *BaseFields) BoolOption(key string, defaultValue bool) bool {
	if value, exists := this.Fields()[key]; exists {
		return utils.ConvertToBool(value, defaultValue)
	}
	return defaultValue
}

func (this *BaseFields) FieldsOption(key string, defaultValue contracts.Fields) contracts.Fields {
	if value, exists := this.Fields()[key]; exists {
		if fields, err := utils.ConvertToFields(value); err == nil {
			return fields
		}
	}
	return defaultValue
}

func (this *BaseFields) GetString(key string) string {
	return utils.GetStringField(this.Fields(), key)
}

func (this *BaseFields) GetInt64(key string) int64 {
	return utils.GetInt64Field(this.Fields(), key)
}

func (this *BaseFields) GetInt(key string) int {
	return utils.GetIntField(this.Fields(), key)
}

func (this *BaseFields) GetFloat64(key string) float64 {
	return utils.GetFloat64Field(this.Fields(), key)
}

func (this *BaseFields) GetFloat(key string) float32 {
	return utils.GetFloatField(this.Fields(), key)
}

func (this *BaseFields) GetBool(key string) bool {
	return utils.GetBoolField(this.Fields(), key)
}

func (this *BaseFields) GetFields(key string) contracts.Fields {
	return utils.GetSubField(this.Fields(), key)
}
