package cache

import (
	"fmt"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/logs"
	"github.com/qbhy/goal/utils"
)

type CacheFactory struct {
	config           contracts.Config
	exceptionHandler contracts.ExceptionHandler
	stores           map[string]contracts.CacheStore
	drivers          map[string]contracts.CacheStoreProvider
}

func (this *CacheFactory) getDefaultDriver() string {
	return utils.StringOr(this.config.GetString("cache.default"), "default")
}

func (this *CacheFactory) Store(names ...string) contracts.CacheStore {
	var name string
	if len(names) > 0 {
		name = names[0]
	} else {
		name = this.getDefaultDriver()
	}

	if cacheStore, existsStore := this.stores[name]; existsStore {
		return cacheStore
	}

	this.stores[name] = this.get(name)

	return this.stores[name]
}

func (this CacheFactory) getConfig(name string) contracts.Fields {
	return this.config.GetFields(
		utils.IfString(name == "default", "cache", fmt.Sprintf("cache.stores.%s", name)),
	)
}

func (this *CacheFactory) Extend(drive string, cacheStoreProvider contracts.CacheStoreProvider) {
	this.drivers[drive] = cacheStoreProvider
}

func (this *CacheFactory) get(name string) contracts.CacheStore {
	config := this.getConfig(name)
	drive := utils.GetStringField(config, "driver", "redis")
	driveProvider, existsProvider := this.drivers[drive]
	if !existsProvider {
		logs.WithFields(nil).Fatal(fmt.Sprintf("不支持的缓存驱动：%s", drive))
	}
	return driveProvider(config)
}
