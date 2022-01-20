package cache

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
	"github.com/qbhy/goal/cache/drivers"
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
			return drivers.NewRedisCache(
				redis.Connection(utils.GetStringField(cacheConfig, "connection")),
				utils.GetStringField(cacheConfig, "prefix"),
			)
		})

		return factory
	})
	container.Singleton("cache.store", func(factory contracts.CacheFactory) contracts.CacheStore {
		return factory.Store()
	})
}
