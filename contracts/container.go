package contracts

type InstanceProvider func() interface{}

type Container interface {
	Provide(interface{})
	ProvideSingleton(interface{})
	Bind(string, InstanceProvider)
	Instance(string, interface{})
	Singleton(string, InstanceProvider)
	HasBound(string) bool
	Alias(string, string)
	Flush()
	Get(string) interface{}
	Call(interface{}, ...interface{}) []interface{}
	DI(object interface{}, args ...interface{})
}

type Component interface {
	ShouldInject()
}
