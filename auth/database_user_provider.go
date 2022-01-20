package auth

import (
	"github.com/goal-web/contracts"
)

type DatabaseUserProvider struct {
	db    contracts.DBConnection
	table string
}

func (this *DatabaseUserProvider) RetrieveById(identifier string) contracts.Authorizable {
	panic("implement me")
}

func (this *DatabaseUserProvider) RetrieveByCredentials(credentials contracts.Fields) contracts.Authorizable {
	panic("implement me")
}

func (this *DatabaseUserProvider) ValidateCredentials(user contracts.Authorizable, credentials contracts.Fields) bool {
	panic("implement me")
}
