package controllers

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/goal/app/controllers/user"
)

// Register 注册路由函数
func Register(router contracts.HttpRouter) {
	user.AuthServiceRouter(router)

	// 在这里添加您的路由注册逻辑
}
