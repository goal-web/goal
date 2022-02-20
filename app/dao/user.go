package dao

import "github.com/goal-web/goal/app/models"

func FindUser(id interface{}) *models.User {
	user, ok := models.UserQuery().Find(id).(models.User)
	if ok {
		return &user
	}
	return nil
}
