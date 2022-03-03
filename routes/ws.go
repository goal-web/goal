package routes

import (
	"fmt"
	"github.com/goal-web/contracts"
	websocket2 "github.com/goal-web/goal/app/websocket"
	"github.com/goal-web/websocket"
)

func WebSocket(router contracts.Router) {
	router.Static("/", "/")

	router.Get("/ws-demo", websocket.New(websocket2.DemoController{}))

	router.Get("/ws", websocket.Default(func(frame contracts.WebSocketFrame) {

		fmt.Println("收到消息", frame.RawString(), frame.Connection().Fd())
		_ = frame.Send("来自服务器的回复1")
		_ = frame.SendBytes([]byte("来自服务器的回复2"))

	}))

}
