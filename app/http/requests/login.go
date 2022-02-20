package requests

import "github.com/goal-web/contracts"

type LoginRequest struct {
	contracts.HttpRequest `di` // 加入 di 标记表示需要注入
}

func (l LoginRequest) Rules() contracts.Fields {
	return contracts.Fields{
		"username": "required",
		"password": "required",
	}
}
