package controllers

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/goal/app/http/requests"
	"github.com/goal-web/goal/app/models"
	"github.com/goal-web/validation"
)

func LoginExample(guard contracts.Guard, request requests.LoginRequest) any {

	// 验证不通过将抛异常，如希望自己处理错误，请使用 Form 方法
	validation.VerifyForm(request)

	var age = request.IntOptional("age", -1)
	//  这是伪代码
	var users = models.UserQuery().
		Where("name", request.GetString("username")).
		Where("age", request.GetInt("age")).
		When(age != -1, func(q contracts.QueryBuilder[models.User]) contracts.Query[models.User] {
			return q.Where("age", ">", age)
		}).
		Get() // any

	var user, err = models.UserQuery().Where("name", request.GetString("username")).FirstE() // any

	if err != nil {
		return contracts.Fields{"error": err.Error()}
	}

	return contracts.Fields{
		"token": guard.Login(user), // jwt 返回 token，session 返回 true
		"users": users.ToArray(),
	}
}

func GetCurrentUser(guard contracts.Guard) any {
	return contracts.Fields{
		"user": guard.User(),
	}
}
