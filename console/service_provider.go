package console

import (
	"github.com/golang-module/carbon/v2"
	"github.com/gorhill/cronexpr"
	"github.com/qbhy/goal/console/inputs"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/logs"
	"github.com/qbhy/parallel"
	"time"
)

type ConsoleProvider func(application contracts.Application) contracts.Console

type ServiceProvider struct {
	ConsoleProvider ConsoleProvider

	stopChan    chan bool
	app         contracts.Application
	execRecords map[int]time.Time
}

func (this *ServiceProvider) Register(application contracts.Application) {
	this.stopChan = make(chan bool, 1)
	this.app = application
	this.execRecords = make(map[int]time.Time)

	application.Singleton("console", func() contracts.Console {
		console := this.ConsoleProvider(application)
		console.Schedule(console.GetSchedule())
		return console
	})
	application.Singleton("scheduling", func(console contracts.Console) contracts.Schedule {
		return console.GetSchedule()
	})
	application.Singleton("console.inputs", func() contracts.ConsoleInput {
		return inputs.NewOSArgsInput()
	})
}

func (this *ServiceProvider) runScheduleEvents(events []contracts.ScheduleEvent) {
	if len(events) > 0 {
		// 并发执行所有事件
		parallelInstance := parallel.NewParallel(len(events))
		now := time.Now()
		for index, event := range events {
			lastExecTime := this.execRecords[index]

			nextTime := carbon.Time2Carbon(cronexpr.MustParse(event.Expression()).Next(lastExecTime))
			nowCarbon := carbon.Time2Carbon(now)
			if nextTime.DiffInSeconds(nowCarbon) == 0 {
				this.execRecords[index] = now
				parallelInstance.Add(func() interface{} {
					event.Run(this.app)
					return nil
				})
			} else if nextTime.Lt(nowCarbon) {
				this.execRecords[index] = now
			}
		}
		parallelInstance.Wait()
	}
}

func (this *ServiceProvider) Start() error {
	this.app.Call(func(schedule contracts.Schedule) (err error) {
		ticker := time.NewTicker(time.Second)
		for {
			select {
			case <-ticker.C:
				this.runScheduleEvents(schedule.GetEvents())
			case <-this.stopChan:
				logs.Default().Debug("scheduling closed")
				return
			}
		}
	})
	return nil
}

func (this *ServiceProvider) Stop() {
	this.stopChan <- true
}
