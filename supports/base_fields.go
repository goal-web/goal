package supports

import (
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/utils"
)

type BaseFields struct { // 具体方法
	contracts.FieldsProvider // 抽象方法，继承 interface
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
