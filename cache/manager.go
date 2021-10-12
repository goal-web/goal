package cache

import (
	"fmt"
	"github.com/qbhy/goal/contracts"
)

type CacheManager struct {
	config           contracts.Config
	events           contracts.EventDispatcher
	exceptionHandler contracts.ExceptionHandler
	stores           map[string]contracts.CacheStore
}

func (this *CacheManager) getDefaultDriver() string {
	return this.config.GetString("cache.default")
}

func (this *CacheManager) Store(names ...string) contracts.CacheStore {
	var name string
	if len(names) > 0 {
		name = names[0]
	} else {
		name = this.getDefaultDriver()
	}

	if repository, existsStore := this.stores[name]; existsStore {
		return repository
	}

	this.stores[name] = this.get(name)

	return this.stores[name]
}

func (this CacheManager) getConfig(name string) contracts.Fields {
	return this.config.GetFields(fmt.Sprintf("cache.stores.%s", name))
}

func (this *CacheManager) get(name string) contracts.CacheStore {
	panic("还没实现")
}
