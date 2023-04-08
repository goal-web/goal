package policies

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/goal/app/models"
)

var Article contracts.Policy = map[string]contracts.GateChecker{
	"create": func(authorizable contracts.Authorizable, data ...any) bool {
		user, isUser := authorizable.(models.User)
		return isUser && user.Role == "blogger"
	},
	"update": func(authorizable contracts.Authorizable, data ...any) bool {
		user, isUser := authorizable.(models.User)

		if len(data) > 0 && isUser {
			article, isArticle := data[0].(models.Article)

			return isArticle && article.UserId == user.Id
		}

		return false
	},
}
