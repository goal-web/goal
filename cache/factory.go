package cache

import (
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
	"github.com/qbhy/goal/logs"
)

type Factory struct {
	config           Config
	exceptionHandler contracts.ExceptionHandler
	stores           map[string]contracts.CacheStore
	drivers          map[string]contracts.CacheStoreProvider
}

func (this *Factory) getName(names ...string) string {
	if len(names) > 0 {
		return names[0]
	}
	return this.config.Default

}

func (this Factory) getConfig(name string) contracts.Fields {
	return this.config.Stores[name]
}

func (this *Factory) Store(names ...string) contracts.CacheStore {
	name := this.getName(names...)
	if cacheStore, existsStore := this.stores[name]; existsStore {
		return cacheStore
	}

	this.stores[name] = this.make(name)

	return this.stores[name]
}

func (this *Factory) Extend(driver string, cacheStoreProvider contracts.CacheStoreProvider) {
	this.drivers[driver] = cacheStoreProvider
}

func (this *Factory) make(name string) contracts.CacheStore {
	config := this.getConfig(name)
	driver := utils.GetStringField(config, "driver")
	driveProvider, existsProvider := this.drivers[driver]
	if !existsProvider {
		logs.WithFields(config).Fatal(fmt.Sprintf("不支持的缓存驱动：%s", driver))
	}
	return driveProvider(config)
}
