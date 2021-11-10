package contracts

type Auth interface {
	Guard

	ExtendUserProvider(key string, provider UserProvider)
	ExtendGuard(key string, guard Guard)

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
