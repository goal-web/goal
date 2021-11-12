package redis

import (
	"github.com/qbhy/goal/contracts"
)

type ServiceProvider struct {
}

func (this ServiceProvider) OnStop() {

}

func (this ServiceProvider) OnStart() error {
	return nil
}

func (this ServiceProvider) Register(app contracts.Application) {

	app.Singleton("redis.factory", func(config contracts.Config, handler contracts.ExceptionHandler) contracts.RedisFactory {
		return &Factory{
			config:           config,
			exceptionHandler: handler,
			connections:      make(map[string]contracts.RedisConnection),
		}
	})

	app.Singleton("redis.connection", func(factory contracts.RedisFactory) contracts.RedisConnection {
		return factory.Connection()
	})

	app.Singleton("redis", func(redis contracts.RedisConnection) *Connection {
		return redis.(*Connection)
	})
}
