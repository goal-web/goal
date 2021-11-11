package tests

import (
	"fmt"
	"github.com/qbhy/goal/application"
	"github.com/qbhy/goal/config"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/hashing"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestHashing(t *testing.T) {
	// 初始化容器
	app := application.Singleton()

	path, _ := os.Getwd()

	// 设置运行目录
	app.Instance("path", path)

	// 注册各种服务
	app.RegisterServices(
		// 配置服务
		config.ServiceProvider{
			Paths: []string{path},
			Env:   "testing",
		},
		// 哈希服务
		hashing.ServiceProvider{},
	)

	hashFactory := app.Get("hash").(contracts.HasherFactory)
	value := "goal hashing"

	bcryptHashedValue := hashFactory.Make(value, nil)
	fmt.Println("bcryptHashedValue:", bcryptHashedValue)
	assert.True(t, hashFactory.Check(value, bcryptHashedValue, nil))
	assert.True(t, len(bcryptHashedValue) > 10)

	md5HashedValue := hashFactory.Driver("md5").Make(value, nil)
	fmt.Println("md5HashedValue:", md5HashedValue)
	assert.True(t, hashFactory.Driver("md5").Check(value, md5HashedValue, nil))
}
