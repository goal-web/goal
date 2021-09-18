package utils

import (
	"github.com/qbhy/goal/contracts"
	"reflect"
	"strings"
)

// IsSameStruct 判断是否同一个结构体
func IsSameStruct(v1, v2 interface{}) bool {
	var (
		f1 reflect.Type
		f2 reflect.Type
		ok bool
	)

	if f1, ok = v1.(reflect.Type); !ok {
		f1 = reflect.TypeOf(v1)
	}

	if f2, ok = v2.(reflect.Type); !ok {
		f2 = reflect.TypeOf(v2)
	}

	return f1.PkgPath() == f2.PkgPath() && f1.Name() == f2.Name()
}

// ConvertToTypes 把变量转换成反射类型
func ConvertToTypes(args ...interface{}) []reflect.Type {
	types := make([]reflect.Type, 0)
	for _, arg := range args {
		types = append(types, reflect.TypeOf(arg))
	}
	return types
}

// IsInstanceIn InstanceIn 判断变量是否是某些类型
func IsInstanceIn(v interface{}, types ...reflect.Type) bool {
	for _, e := range types {
		if IsSameStruct(e, v) {
			return true
		}
	}
	return false
}

// EachStructField 遍历结构体的字段
func EachStructField(s interface{}, handler func(reflect.StructField, reflect.Value)) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	for i := 0; i < t.NumField(); i++ {
		handler(t.Field(i), v.Field(i))
	}
}

// GetTypeKey 获取类型唯一字符串
func GetTypeKey(p reflect.Type) string {
	if p.Kind() == reflect.Ptr {
		p = p.Elem()
	}

	return p.PkgPath() + "." + strings.ToTitle(p.Name())
}

// NotNil 尽量不要 nil
func NotNil(args ...interface{}) interface{} {
	for _, arg := range args {
		switch argValue := arg.(type) {
		case contracts.InstanceProvider:
			arg = argValue()
		case func() interface{}:
			arg = argValue()
		}
		if arg != nil {
			return arg
		}
	}
	return nil
}
