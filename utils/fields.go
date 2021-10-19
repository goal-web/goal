package utils

import "github.com/qbhy/goal/contracts"

func MergeFields(fields contracts.Fields, finalFields contracts.Fields) {
	for key, value := range finalFields {
		fields[key] = value
	}
}

func GetStringField(fields contracts.Fields, key string, defaultValues ...string) string {
	if value, existsString := fields[key]; existsString {
		if str, isString := value.(string); isString {
			return str
		}
		return StringOr(defaultValues...)
	} else {
		return StringOr(defaultValues...)
	}
}

func GetInt64Field(fields contracts.Fields, key string, defaultValues ...int64) int64 {
	var defaultValue int64 = 0
	if len(defaultValues) > 0 {
		defaultValue = defaultValues[0]
	}
	if value, existsValue := fields[key]; existsValue {
		if intValue, isInt := value.(int64); isInt {
			return intValue
		}
		return ConvertToInt64(value, defaultValue)
	} else {
		return defaultValue
	}
}

func GetBoolField(fields contracts.Fields, key string, defaultValues ...bool) bool {
	var defaultValue = false
	if len(defaultValues) > 0 {
		defaultValue = defaultValues[0]
	}
	if fieldValue, existsValue := fields[key]; existsValue {
		return ConvertToBool(fieldValue, defaultValue)
	}
	return defaultValue
}
