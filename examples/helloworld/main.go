package main

import (
	"fmt"
	"github.com/qbhy/goal/application"
	"github.com/qbhy/goal/auth"
	"github.com/qbhy/goal/cache"
	"github.com/qbhy/goal/config"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/database"
	"github.com/qbhy/goal/encryption"
	"github.com/qbhy/goal/events"
	_ "github.com/qbhy/goal/examples/helloworld/config"
	config2 "github.com/qbhy/goal/examples/helloworld/config"
	"github.com/qbhy/goal/examples/helloworld/exceptions"
	"github.com/qbhy/goal/examples/helloworld/routes"
	"github.com/qbhy/goal/filesystemt"
	"github.com/qbhy/goal/hashing"
	"github.com/qbhy/goal/http"
	"github.com/qbhy/goal/logs"
	"github.com/qbhy/goal/redis"
	"github.com/qbhy/goal/session"
	"github.com/qbhy/goal/signal"
	"io/ioutil"
	"os"
)

func main() {
	app := application.Singleton()
	path, _ := os.Getwd()
	app.Instance("path", path)

	// 设置异常处理器
	app.Singleton("exceptions.handler", func() contracts.ExceptionHandler {
		return exceptions.NewHandler()
	})

	app.RegisterServices(
		&config.ServiceProvider{
			Env:             os.Getenv("env"),
			Paths:           []string{path},
			Sep:             "=",
			ConfigProviders: config2.Configs(),
		},
		hashing.ServiceProvider{},
		encryption.ServiceProvider{},
		filesystemt.ServiceProvider{},
		events.ServiceProvider{},
		redis.ServiceProvider{},
		cache.ServiceProvider{},
		&signal.ServiceProvider{},
		&session.ServiceProvider{},
		auth.ServiceProvider{},
		&database.ServiceProvider{},
		&http.ServiceProvider{RouteCollectors: []interface{}{
			// 路由收集器
			routes.V1Routes,
		}},
	)

	pidPath := path + "/goal.pid"
	// 写入 pid 文件
	_ = ioutil.WriteFile(pidPath, []byte(fmt.Sprintf("%d", os.Getpid())), os.ModePerm)

	if errors := app.Start(); len(errors) > 0 {
		logs.WithField("errors", errors).Error("goal 异常关闭!")
	} else {
		_ = os.Remove(pidPath)
		logs.WithInterface(nil).Info("goal 已关闭")
	}
}
