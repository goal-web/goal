package dao

import "github.com/goal-web/goal/app/models"

func FindUser(id any) *models.User {
	return models.Users().Find(id)
}
