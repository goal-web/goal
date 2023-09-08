package config

import (
	websocket2 "github.com/fasthttp/websocket"
	"github.com/goal-web/contracts"
	"github.com/goal-web/http/websocket"
)

func init() {
	configs["websocket"] = func(env contracts.Env) any {
		return websocket.Config{
			Upgrader: websocket2.FastHTTPUpgrader{
				HandshakeTimeout: 5,
				//ReadBufferSize:    0,
				//WriteBufferSize:   0,
				//WriteBufferPool:   nil,
				//Subprotocols:      nil,
				//Error:             nil,
				//CheckOrigin:       nil,
				//EnableCompression: false,
			},
		}
	}
}
