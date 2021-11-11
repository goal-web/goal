package tests

import (
	"github.com/qbhy/goal/application"
	"github.com/qbhy/goal/cache"
	"github.com/qbhy/goal/config"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/exceptions"
	"github.com/qbhy/goal/redis"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCacheFactory(t *testing.T) {
	// 初始化容器
	app := application.Singleton()

	path, _ := os.Getwd()

	// 设置运行目录
	app.Instance("path", path)

	// 注册异常处理器
	app.ProvideSingleton(func() contracts.ExceptionHandler {
		return exceptions.NewDefaultHandler(nil)
	})

	// 注册各种服务
	app.RegisterServices(
		// 配置服务
		config.ServiceProvider{
			Paths: []string{path},
			Env:   os.Getenv("APP_ENV"),
		},
		// redis 服务
		redis.ServiceProvider{},
		// cache 服务
		cache.ServiceProvider{},
	)

	cacheFactory := app.Get("cache").(contracts.CacheFactory)

	err := cacheFactory.Store().Forever("a", "testing")
	assert.Nil(t, err, err)
	assert.True(t, cacheFactory.Store().Get("a") == "testing")
}
