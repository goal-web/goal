package config

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/websocket"
	websocket2 "github.com/gorilla/websocket"
)

func init() {
	configs["websocket"] = func(env contracts.Env) any {
		return websocket.Config{
			Upgrader: websocket2.Upgrader{
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
