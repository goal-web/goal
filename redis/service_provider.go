package redis

import (
	"github.com/qbhy/goal/contracts"
)

type ServiceProvider struct {
}

func (this ServiceProvider) Register(app contracts.Container) {

	app.ProvideSingleton(func(config contracts.Config, handler contracts.ExceptionHandler) contracts.RedisFactory {
		return &Factory{
			config:           config,
			exceptionHandler: handler,
			connections:      make(map[string]contracts.RedisConnection),
		}
	})

	app.ProvideSingleton(func(factory contracts.RedisFactory) contracts.RedisConnection {
		return factory.Connection()
	})

	app.ProvideSingleton(func(redis contracts.RedisConnection) *Connection {
		return redis.(*Connection)
	})
}
