package tests

import (
	"github.com/qbhy/goal/container"
	"github.com/goal-web/contracts"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestRedisFactory(t *testing.T) {
	// 初始化容器

	path, _ := os.Getwd()
	app := getApp(path)

	app.Call(container.NewMagicalFunc(func(factory contracts.RedisFactory) {
		defaultConnection := factory.Connection()

		_, err := defaultConnection.Set("a", "default", time.Minute*5)
		assert.True(t, err == nil)
		aValue, err := defaultConnection.Get("a")
		assert.True(t, err == nil)
		assert.True(t, aValue == "default")
		num, err := defaultConnection.Exists("a")
		assert.True(t, num == 1)
		assert.True(t, err == nil, err)

		cacheConnection := factory.Connection("cache")
		_, err = cacheConnection.Set("a", "cache", time.Minute*5)
		assert.True(t, err == nil)
		aValue, err = cacheConnection.Get("a")
		assert.True(t, err == nil)
		assert.True(t, aValue == "cache")
	}))
}
