package arguments

import (
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/supports"
	"github.com/qbhy/goal/utils"
	"strings"
)

type Arguments struct {
	supports.BaseFields
	args    []string
	options contracts.Fields
}

func (this *Arguments) GetArg(index int) string {
	if index >= len(this.args) {
		return ""
	}
	return this.args[index]
}

func (this *Arguments) GetArgs() []string {
	return this.args
}

func NewArguments(args []string, options contracts.Fields) contracts.CommandArguments {
	arguments := &Arguments{
		args:       args,
		BaseFields: supports.BaseFields{},
		options:    options,
	}

	arguments.BaseFields.FieldsProvider = arguments
	return arguments
}

func (this *Arguments) StringArrayOption(key string, defaultValue []string) []string {
	if value := this.GetString(key); value != "" {
		return strings.Split(value, ",")
	}
	return defaultValue
}

func (this *Arguments) Int64ArrayOption(key string, defaultValue []int64) []int64 {
	if value := this.GetString(key); value != "" {
		values := make([]int64, 0)
		for _, value = range strings.Split(value, ",") {
			values = append(values, utils.ConvertToInt64(value, 0))
		}
		return values
	}
	return defaultValue
}

func (this *Arguments) IntArrayOption(key string, defaultValue []int) []int {
	if value := this.GetString(key); value != "" {
		values := make([]int, 0)
		for _, value = range strings.Split(value, ",") {
			values = append(values, utils.ConvertToInt(value, 0))
		}
		return values
	}
	return defaultValue
}

func (this *Arguments) Float64ArrayOption(key string, defaultValue []float64) []float64 {
	if value := this.GetString(key); value != "" {
		values := make([]float64, 0)
		for _, value = range strings.Split(value, ",") {
			values = append(values, utils.ConvertToFloat64(value, 0))
		}
		return values
	}
	return defaultValue
}

func (this *Arguments) FloatArrayOption(key string, defaultValue []float32) []float32 {
	if value := this.GetString(key); value != "" {
		values := make([]float32, 0)
		for _, value = range strings.Split(value, ",") {
			values = append(values, utils.ConvertToFloat(value, 0))
		}
		return values
	}
	return defaultValue
}

func (this *Arguments) Fields() contracts.Fields {
	return this.options
}

func (this *Arguments) Exists(key string) bool {
	_, exists := this.options[key]
	return exists
}
