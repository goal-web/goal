package tests

import (
	"github.com/qbhy/goal/config"
	"github.com/qbhy/goal/container"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/exceptions"
	"github.com/qbhy/goal/redis"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestRedisFactory(t *testing.T) {
	// 初始化容器
	app := container.Singleton()

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
	)

	app.Call(func(factory contracts.RedisFactory) {
		defaultConnection := factory.Connection()

		_, err := defaultConnection.Set("a", "default", time.Minute * 5)
		assert.True(t, err == nil)
		aValue ,err := defaultConnection.Get("a")
		assert.True(t, err == nil)
		assert.True(t, aValue == "default")


		cacheConnection := factory.Connection("cache")
		_, err = cacheConnection.Set("a", "cache", time.Minute * 5)
		assert.True(t, err == nil)
		aValue ,err = cacheConnection.Get("a")
		assert.True(t, err == nil)
		assert.True(t, aValue == "cache")
	})
}
