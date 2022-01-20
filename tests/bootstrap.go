package tests

import (
	"fmt"
	"github.com/goal-web/cache"
	"github.com/goal-web/contracts"
	"github.com/goal-web/encryption"
	"github.com/goal-web/events"
	"github.com/goal-web/filesystem"
	"github.com/goal-web/redis"
	"github.com/goal-web/supports/logs"
	"github.com/qbhy/goal/application"
	"github.com/qbhy/goal/auth"
	"github.com/qbhy/goal/config"
	"github.com/qbhy/goal/database"
	"github.com/qbhy/goal/examples/helloworld/app/exceptions"
	config2 "github.com/qbhy/goal/examples/helloworld/config"
	"github.com/qbhy/goal/hashing"
	"github.com/qbhy/goal/session"
	"github.com/qbhy/goal/signal"
	"io/ioutil"
	"os"
)

func getApp(path string) contracts.Application {
	env := "testing"
	app := application.Singleton()
	app.Instance("path", path)

	// 设置异常处理器
	app.Singleton("exceptions.handler", func() contracts.ExceptionHandler {
		return exceptions.NewHandler()
	})

	app.RegisterServices(
		&config.ServiceProvider{
			Env:             env,
			Paths:           []string{path},
			Sep:             "=",
			ConfigProviders: config2.Configs(),
		},
		hashing.ServiceProvider{},
		encryption.ServiceProvider{},
		filesystem.ServiceProvider{},
		events.ServiceProvider{},
		redis.ServiceProvider{},
		cache.ServiceProvider{},
		&signal.ServiceProvider{},
		&session.ServiceProvider{},
		auth.ServiceProvider{},
		&database.ServiceProvider{},
		//&http.ServiceProvider{RouteCollectors: []interface{}{
		//	// 路由收集器
		//	routes.V1Routes,
		//}},
	)

	pidPath := path + "/goal.pid"
	// 写入 pid 文件
	_ = ioutil.WriteFile(pidPath, []byte(fmt.Sprintf("%d", os.Getpid())), os.ModePerm)

	go func() {
		if errors := app.Start(); len(errors) > 0 {
			logs.WithField("errors", errors).Fatal("goal 异常关闭!")
		} else {
			_ = os.Remove(pidPath)
			logs.WithInterface(nil).Info("goal 已关闭")
		}
	}()
	return app
}
