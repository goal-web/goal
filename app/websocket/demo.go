package websocket

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/logs"
)

type DemoController struct {
}

func (d DemoController) OnConnect(request contracts.HttpRequest, fd uint64) error {
	// todo: 绑定用户和 fd
	return nil
}

func (d DemoController) OnMessage(frame contracts.WebSocketFrame) {
	logs.Default().Info("received websocket message:" + frame.RawString())
}

func (d DemoController) OnClose(fd uint64) {
	// todo: 解绑 fd
}
