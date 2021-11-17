package contracts

type Application interface {
	Container

	RegisterServices(provider ...ServiceProvider)

	Start() map[string]error

	Stop()
}

type ServiceProvider interface {
	Register(application Application)
	Start() error
	Stop()
}
