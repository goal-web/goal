package contracts

type Application interface {
	Container

	RegisterServices(provider ...ServiceProvider)

	Start() map[string]error

	OnStop()
}

type ServiceProvider interface {
	Register(application Application)
	OnStart() error
	OnStop()
}
