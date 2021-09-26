package contracts

type InstanceProvider func() interface{}

type Container interface {
	Provide(interface{}, ...string)
	ProvideSingleton(interface{}, ...string)
	Bind(string, InstanceProvider)
	Instance(string, interface{})
	Singleton(string, InstanceProvider)
	HasBound(string) bool
	Alias(string, string)
	Flush()
	Get(string) interface{}
	Call(interface{}, ...interface{}) []interface{}
	DI(object interface{}, args ...interface{})
	RegisterServices(provider ...ServiceProvider)
}

type Component interface {
	ShouldInject()
}

type ServiceProvider interface {
	Register(container Container)
}
