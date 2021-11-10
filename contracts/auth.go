package contracts

type GuardProvider func(config Fields) Guard
type UserProviderProvider func(config Fields) UserProvider

type Auth interface {
	Guard

	ExtendUserProvider(key string, provider UserProviderProvider)
	ExtendGuard(key string, guard GuardProvider)

	Guard(key string) Guard
	UserProvider(key string) UserProvider
}

type Authorizable interface {
	Id() string
}

type Guard interface {
	SetUser(authorizable Authorizable)
	User() Authorizable
	Id() string
	Check() bool
	Guest() bool
	Validate(credentials Fields) bool
}

type UserProvider interface {
	RetrieveById(identifier string) Authorizable

	RetrieveByCredentials(credentials Fields) Authorizable

	ValidateCredentials(user Authorizable, credentials Fields) bool
}
