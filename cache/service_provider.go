package cache

import (
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/utils"
)

type ServiceProvider struct {
}

func (this ServiceProvider) Stop() {

}

func (this ServiceProvider) Start() error {
	return nil
}

func (this ServiceProvider) Register(container contracts.Application) {
	container.Singleton("cache", func(
		config contracts.Config,
		redis contracts.RedisFactory,
		handler contracts.ExceptionHandler) contracts.CacheFactory {
		factory := &Factory{
			config:           config.Get("cache").(Config),
			exceptionHandler: handler,
			stores:           make(map[string]contracts.CacheStore),
			drivers:          make(map[string]contracts.CacheStoreProvider),
		}

		factory.Extend("redis", func(cacheConfig contracts.Fields) contracts.CacheStore {
			return &RedisStore{
				connection: redis.Connection(utils.GetStringField(cacheConfig, "connection")),
				prefix:     utils.GetStringField(cacheConfig, "prefix"),
			}
		})

		return factory
	})
	container.Singleton("cache.store", func(factory contracts.CacheFactory) contracts.CacheStore {
		return factory.Store()
	})
}
