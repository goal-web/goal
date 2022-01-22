package console

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/logs"
	"github.com/goal-web/supports/utils"
	"github.com/golang-module/carbon/v2"
	"github.com/gorhill/cronexpr"
	"github.com/qbhy/goal/application"
	"github.com/qbhy/goal/console/inputs"
	"github.com/qbhy/goal/exceptions"
	"reflect"
	"time"
)

type ConsoleProvider func(application contracts.Application) contracts.Console

type ServiceProvider struct {
	ConsoleProvider ConsoleProvider

	stopChan         chan bool
	serverIdChan     chan bool
	app              contracts.Application
	execRecords      map[int]time.Time
	exceptionHandler contracts.ExceptionHandler
}

func (this *ServiceProvider) Register(application contracts.Application) {
	this.app = application
	this.exceptionHandler = application.Get("exceptions.handler").(contracts.ExceptionHandler)

	application.Singleton("console", func() contracts.Console {
		console := this.ConsoleProvider(application)
		console.Schedule(console.GetSchedule())
		return console
	})
	application.Singleton("scheduling", func(console contracts.Console) contracts.Schedule {
		return console.GetSchedule()
	})
	application.Singleton("console.input", func() contracts.ConsoleInput {
		return inputs.NewOSArgsInput()
	})
}

func (this *ServiceProvider) runScheduleEvents(events []contracts.ScheduleEvent) {
	if len(events) > 0 {
		// 并发执行所有事件
		now := time.Now()
		for index, event := range events {
			lastExecTime := this.execRecords[index]
			nextTime := carbon.Time2Carbon(cronexpr.MustParse(event.Expression()).Next(lastExecTime))
			nowCarbon := carbon.Time2Carbon(now)
			if nextTime.DiffInSeconds(nowCarbon) == 0 {
				this.execRecords[index] = now
				go (func(event contracts.ScheduleEvent) {
					defer func() {
						if err := recover(); err != nil {
							this.exceptionHandler.Handle(ScheduleEventException{
								Exception: exceptions.WithRecover(err, contracts.Fields{
									"expression": event.Expression(),
									"mutex_name": event.MutexName(),
									"one_server": event.OnOneServer(),
									"event":      utils.GetTypeKey(reflect.TypeOf(event)),
								}),
							})
						}
					}()
					event.Run(this.app)
				})(event)
			} else if nextTime.Lt(nowCarbon) {
				this.execRecords[index] = now
			}
		}
	}
}

func (this *ServiceProvider) Start() error {
	this.execRecords = make(map[int]time.Time)
	go this.maintainServerId()
	this.app.Call(func(schedule contracts.Schedule) {
		if len(schedule.GetEvents()) > 0 {
			this.stopChan = utils.SetInterval(1, func() {
				this.runScheduleEvents(schedule.GetEvents())
			}, func() {
				logs.Default().Info("the goal scheduling is closed")
			})
		}
	})
	return nil
}

func (this *ServiceProvider) Stop() {
	if this.stopChan != nil {
		this.stopChan <- true
	}
	if this.serverIdChan != nil {
		this.serverIdChan <- true
	}
}

// maintainServerId 维护服务实例ID
func (this *ServiceProvider) maintainServerId() {
	this.app.Call(func(redis contracts.RedisConnection, config contracts.Config, handler contracts.ExceptionHandler) {
		appConfig := config.Get("app").(application.Config)
		this.serverIdChan = utils.SetInterval(1, func() {
			// 维持当前服务心跳
			_, _ = redis.Set("goal.server."+appConfig.ServerId, time.Now().String(), time.Second*2)
		}, func() {
			_, _ = redis.Del("goal.server." + appConfig.ServerId)
		})
	})
}
