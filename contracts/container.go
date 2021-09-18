package contracts

type InstanceProvider func() interface{}

type Container interface {
	Provide(interface{})
	ProvideSingleton(interface{})
	Bind(string, InstanceProvider)
	Instance(string, interface{})
	Singleton(string, InstanceProvider)
	Bound(string) bool
	Alias(string, string)
	Flush()
	Get(string) interface{}
	Call(interface{}, ...interface{}) []interface{}
}

type Component interface {
	Construct() interface{}
}
