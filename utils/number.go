package utils

import "strconv"

// 把能转换成 int64 的值转换成 int64
func ConvertToInt64(i interface{}, defaultValue int64) int64 {
	switch value := i.(type) {
	case int64:
		return value
	case int:
		return int64(value)
	case int8:
		return int64(value)
	case int16:
		return int64(value)
	case int32:
		return int64(value)
	case float64:
		return int64(value)
	case float32:
		return int64(value)
	case string:
		i64, _ := strconv.ParseInt(value, 10, 64)
		return i64
	}

	return defaultValue
}

// 把能转换成 float64 的值转换成 float64
func ConvertToFloat64(f interface{}, defaultValue float64) float64 {
	switch value := f.(type) {
	case float64:
		return value
	case int64:
		return float64(value)
	case int:
		return float64(value)
	case int8:
		return float64(value)
	case int16:
		return float64(value)
	case int32:
		return float64(value)
	case float32:
		return float64(value)
	case string:
		f64, _ := strconv.ParseFloat(value, 64)
		return f64
	}

	return defaultValue
}
