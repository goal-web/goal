package core

import (
	"github.com/goal-web/application"
	"github.com/goal-web/bloomfilter"
	"github.com/goal-web/cache"
	"github.com/goal-web/config"
	"github.com/goal-web/console"
	"github.com/goal-web/console/scheduling"
	"github.com/goal-web/contracts"
	"github.com/goal-web/database"
	"github.com/goal-web/email"
	"github.com/goal-web/encryption"
	"github.com/goal-web/events"
	"github.com/goal-web/filesystem"
	"github.com/goal-web/goal/app/exceptions"
	"github.com/goal-web/goal/app/providers"
	config2 "github.com/goal-web/goal/config"
	"github.com/goal-web/hashing"
	"github.com/goal-web/http/sse"
	"github.com/goal-web/http/websocket"
	"github.com/goal-web/migration"
	"github.com/goal-web/queue"
	"github.com/goal-web/ratelimiter"
	"github.com/goal-web/redis"
	"github.com/goal-web/routing"
	"github.com/goal-web/serialization"
	"github.com/goal-web/supports/utils"
	"github.com/golang-module/carbon/v2"
)

type App struct {
	QueueWorker      bool
	SchedulingWorker bool
}

func Application(c ...App) contracts.Application {
	env := config.NewToml(config.File("env.toml"))
	app := application.Singleton(env.GetBool("app.debug"))

	conf := utils.DefaultValue(c, App{})
	// 设置异常处理器
	app.Singleton("exceptions.handler", func() contracts.ExceptionHandler {
		return exceptions.NewHandler()
	})

	app.RegisterServices(
		config.NewService(env, config2.GetConfigProviders()),
		hashing.NewService(),
		encryption.NewService(),
		filesystem.NewService(),
		serialization.NewService(),
		events.NewService(),
		providers.NewEvents(),
		redis.NewService(),
		cache.NewService(),
		bloomfilter.NewService(),
		ratelimiter.NewService(),
		database.NewService(),
		queue.NewService(conf.QueueWorker),
		email.NewService(),
		console.NewService(),
		scheduling.NewService(conf.SchedulingWorker),
		migration.NewService(),
		routing.NewService(),
		sse.NewService(),
		websocket.NewService(),
	)

	app.RegisterServices(
		providers.NewApp(),
	)

	app.Call(func(config contracts.Config, dispatcher contracts.EventDispatcher, console3 contracts.Console) {
		appConfig := config.Get("app").(application.Config)
		carbon.SetLocale(appConfig.Locale)
		carbon.SetTimezone(appConfig.Timezone)
	})

	return app
}
