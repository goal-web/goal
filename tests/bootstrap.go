package tests

import (
	"github.com/goal-web/application"
	"github.com/goal-web/application/exceptions"
	"github.com/goal-web/application/signal"
	"github.com/goal-web/auth"
	"github.com/goal-web/bloomfilter"
	"github.com/goal-web/cache"
	"github.com/goal-web/config"
	"github.com/goal-web/console"
	"github.com/goal-web/contracts"
	"github.com/goal-web/database"
	"github.com/goal-web/email"
	"github.com/goal-web/encryption"
	"github.com/goal-web/events"
	"github.com/goal-web/filesystem"
	console2 "github.com/goal-web/goal/app/console"
	config2 "github.com/goal-web/goal/config"
	"github.com/goal-web/hashing"
	"github.com/goal-web/redis"
	"github.com/goal-web/session"
	"github.com/goal-web/supports/logs"
)

func initApp(path ...string) contracts.Application {
	runPath := "/Users/qbhy/project/go/goal-web/goal"
	if len(path) > 0 {
		runPath = path[0]
	}
	env := "testing"
	app := application.Singleton()
	app.Instance("path", runPath)

	// 设置异常处理器
	app.Singleton("exceptions.handler", func() contracts.ExceptionHandler {
		return exceptions.DefaultExceptionHandler{}
	})

	app.RegisterServices(
		&config.ServiceProvider{
			Env:             env,
			Paths:           []string{runPath},
			Sep:             "=",
			ConfigProviders: config2.Configs(),
		},
		&console.ServiceProvider{
			ConsoleProvider: func(application contracts.Application) contracts.Console {
				return console2.NewKernel(application)
			},
		},
		hashing.ServiceProvider{},
		encryption.ServiceProvider{},
		filesystem.ServiceProvider{},
		events.ServiceProvider{},
		redis.ServiceProvider{},
		&bloomfilter.ServiceProvider{},
		cache.ServiceProvider{},
		&signal.ServiceProvider{},
		&session.ServiceProvider{},
		auth.ServiceProvider{},
		&email.ServiceProvider{},
		&database.ServiceProvider{},
		//&http.ServiceProvider{RouteCollectors: []interface{}{
		//	// 路由收集器
		//	routes.V1Routes,
		//}},
	)

	go func() {
		if errors := app.Start(); len(errors) > 0 {
			logs.WithField("errors", errors).Fatal("goal 启动异常!")
		}
	}()
	return app
}
