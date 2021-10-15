package cache

import (
	"github.com/qbhy/goal/contracts"
)

type ServiceProvider struct {
}

func (this *ServiceProvider) Register(container contracts.Container) {
	container.ProvideSingleton(func(config contracts.Config, dispatcher contracts.EventDispatcher, handler contracts.ExceptionHandler) contracts.CacheFactory {
		return &CacheManager{
			config:           config,
			events:           dispatcher,
			exceptionHandler: handler,
			stores:           make(map[string]contracts.CacheStore),
		}
	})
	container.ProvideSingleton(func(factory contracts.CacheFactory) contracts.CacheStore {
		return factory.Store()
	})
}
