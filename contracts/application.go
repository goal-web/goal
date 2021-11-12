package contracts

type Application interface {
	Container

	RegisterServices(provider ...ServiceProvider)

	Start() map[string]error

	OnStop()

	// Listen 添加事件监听
	Listen(key string, handler interface{})
	// Trigger 触发事件
	Trigger(arguments ...interface{})
}

type ServiceProvider interface {
	Register(application Application)
	OnStart() error
	OnStop()
}
