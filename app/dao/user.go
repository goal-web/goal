package dao

import "github.com/goal-web/goal/app/models"

func FindUser(id any) *models.User {
	user, ok := models.UserQuery().Find(id).(models.User)
	if ok {
		return &user
	}
	return nil
}
