package providers

import (
	"github.com/goal-web/auth/gate"
	"github.com/goal-web/contracts"
	"github.com/goal-web/goal/app/models"
	"github.com/goal-web/goal/app/policies"
)

func Gate() contracts.ServiceProvider {
	return &gate.ServiceProvider{Policies: map[contracts.Class]contracts.Policy{
		models.ArticleClass: policies.Article,
	}}
}
