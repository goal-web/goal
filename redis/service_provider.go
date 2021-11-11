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

	app.ProvideSingleton(func(config contracts.Config, handler contracts.ExceptionHandler) contracts.RedisFactory {
		return &Factory{
			config:           config,
			exceptionHandler: handler,
			connections:      make(map[string]contracts.RedisConnection),
		}
	}, "redis")

	app.ProvideSingleton(func(factory contracts.RedisFactory) contracts.RedisConnection {
		return factory.Connection()
	}, "redis.connection")

	app.ProvideSingleton(func(redis contracts.RedisConnection) *Connection {
		return redis.(*Connection)
	})
}
