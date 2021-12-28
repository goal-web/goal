package contracts

import "reflect"

type InstanceProvider func() interface{}

type Container interface {
	Bind(string, interface{})
	Instance(string, interface{})
	Singleton(string, interface{})
	HasBound(string) bool
	Alias(string, string)
	Flush()
	Get(key string, args ...interface{}) interface{}
	Call(fn interface{}, args ...interface{}) []interface{}
	DI(object interface{}, args ...interface{})
}

// Component 可注入的类
type Component interface {
	ShouldInject()
}

// Injectable 可注入类型
type Injectable interface {
	Call(container Container)
}

// MagicalFunc 可以通过容器调用的魔术方法
type MagicalFunc interface {
	NumOut() int
	NumIn() int
	Call(in []reflect.Value) []reflect.Value
	Arguments() []reflect.Type
	Returns() []reflect.Type
}
