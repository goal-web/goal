package sse

import (
	"errors"
	"github.com/goal-web/contracts"
)

type DemoController struct {
}

func (d DemoController) OnConnect(request contracts.HttpRequest, fd uint64) error {
	// 伪代码
	if request.GetString("token") != "goal" {
		return errors.New("401")
	}

	// todo: 绑定用户和 fd

	return nil
}

func (d DemoController) OnClose(fd uint64) {
	// todo: 实现解绑
}
