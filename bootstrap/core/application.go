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
    "github.com/goal-web/log"
    "github.com/goal-web/migration"
    "github.com/goal-web/queue"
    "github.com/goal-web/ratelimiter"
    "github.com/goal-web/redis"
    "github.com/goal-web/routing"
    "github.com/goal-web/serialization"
    logs "github.com/goal-web/supports/logs"
    "github.com/goal-web/supports/utils"
    "github.com/dromara/carbon/v2"
    "os"
)

type App struct {
	QueueWorker      bool
	SchedulingWorker bool
}

func Application(c ...App) contracts.Application {
	env := config.NewToml(config.File("env.toml"))
	app := application.Singleton(env.GetBool("app.debug"))
	pwd, _ := os.Getwd()
	app.Instance("pwd", pwd)

	conf := utils.DefaultValue(c, App{})
	// 设置异常处理器
	app.Singleton("exceptions.handler", func() contracts.ExceptionHandler {
		return exceptions.NewHandler()
	})

	app.RegisterServices(
		config.NewService(env, config2.GetConfigProviders()),
		log.NewService(),
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

	// 把 supports/logs 的包级调用路由到容器里配置好的 LogManager：
	// log 服务注册之后立即接线，此后的 logs.Default()/With* 等包级调用
	// 都经容器解析到 zap 适配实现（stack 默认通道 = console + file 双通道）。
	// 必须放在 providers.NewApp() 注册之前——它的 Register 会在引导期触发
	// migrate.Auto 并经 logs.* 输出，接线晚了这些记录会落到兜底 logger。
	// 若从上方服务列表移除 log.NewService，此闭包首次调用即 panic
	// （container.Get 对未注册键返回 nil，接口断言失败）——属引导期
	// fail-fast，是有意设计：装配残缺在启动时暴露，而不是静默丢日志。
	// Note: removing log.NewService from the service list above makes this
	// closure panic on its first invocation (container.Get returns nil for an
	// unregistered key and the interface assertion fails) — bootstrap-time
	// fail-fast, intentional: broken wiring surfaces at startup instead of
	// silently dropping logs.
	logs.SetDefault(func() contracts.Logger {
		return app.Get("log").(contracts.LogManager).Channel()
	})

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
