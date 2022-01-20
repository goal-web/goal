package scheduling

import (
	"fmt"
	"github.com/golang-module/carbon/v2"
	"github.com/goal-web/contracts"
	"strconv"
	"strings"
	"time"
)

func NewEvent(mutex *Mutex, callback interface{}, timezone string) *Event {
	return &Event{
		callback:           callback,
		mutex:              mutex,
		filters:            make([]filter, 0),
		rejects:            make([]filter, 0),
		beforeCallbacks:    make([]func(), 0),
		afterCallbacks:     make([]func(), 0),
		withoutOverlapping: false,
		onOneServer:        false,
		timezone:           timezone,
		expression:         "0 * * * * * *",
		mutexName:          "",
		expiresAt:          0,
	}
}

type filter func() bool

type Event struct {
	callback interface{}

	mutex           *Mutex
	filters         []filter
	rejects         []filter
	beforeCallbacks []func()
	afterCallbacks  []func()

	withoutOverlapping bool
	onOneServer        bool

	timezone   string
	expression string
	mutexName  string
	expiresAt  time.Duration
}

func (this *Event) Years(years ...string) contracts.ScheduleEvent {
	if len(years) > 0 {
		return this.SpliceIntoPosition(6, strings.Join(years, ","))
	}
	return this
}

func (this *Event) Expression() string {
	return this.expression
}

func (this *Event) EveryThirtySeconds() contracts.ScheduleEvent {
	return this.SpliceIntoPosition(0, "0,30")
}

func (this *Event) EveryFifteenSeconds() contracts.ScheduleEvent {
	return this.SpliceIntoPosition(0, "*/15")
}

func (this *Event) EveryTenSeconds() contracts.ScheduleEvent {
	return this.SpliceIntoPosition(0, "*/10")
}

func (this *Event) EveryFiveSeconds() contracts.ScheduleEvent {
	return this.SpliceIntoPosition(0, "*/5")
}

func (this *Event) EveryFourSeconds() contracts.ScheduleEvent {
	return this.SpliceIntoPosition(0, "*/4")
}

func (this *Event) EveryThreeSeconds() contracts.ScheduleEvent {
	return this.SpliceIntoPosition(0, "*/3")
}

func (this *Event) EveryTwoSeconds() contracts.ScheduleEvent {
	return this.SpliceIntoPosition(0, "*/2")
}

func (this *Event) EverySecond() contracts.ScheduleEvent {
	return this.SpliceIntoPosition(0, "*")
}

func (this *Event) WithoutOverlapping(expiresAt int) contracts.ScheduleEvent {
	this.expiresAt = time.Duration(expiresAt) * time.Second
	this.withoutOverlapping = true
	return this.Skip(func() bool {
		return this.mutex.Exists(this)
	})
}

func (this *Event) Run(application contracts.Application) {
	if !this.FiltersPass() {
		return
	}
	defer this.removeMutex()
	if this.withoutOverlapping && !this.mutex.Create(this) {
		return
	}
	application.Call(this.callback)
	return
}

func (this *Event) removeMutex() {
	if this.withoutOverlapping {
		this.mutex.Forget(this)
	}
}
func (this *Event) OnOneServer() contracts.ScheduleEvent {
	this.onOneServer = true
	return this
}

func (this *Event) Timezone(timezone string) contracts.ScheduleEvent {
	this.timezone = timezone
	return this
}

func (this *Event) Days(day string, days ...string) contracts.ScheduleEvent {
	days = append([]string{day}, days...)
	return this.SpliceIntoPosition(5, strings.Join(days, ","))
}

func (this *Event) YearlyOn(month time.Month, dayOfMonth int, timeStr string) contracts.ScheduleEvent {
	this.DailyAt(timeStr)

	return this.SpliceIntoPosition(3, strconv.Itoa(dayOfMonth)).
		SpliceIntoPosition(4, strconv.Itoa(int(month)))
}

func (this *Event) Yearly() contracts.ScheduleEvent {
	return this.SpliceIntoPosition(1, "0").
		SpliceIntoPosition(2, "0").
		SpliceIntoPosition(3, "1").
		SpliceIntoPosition(4, "1")
}

func (this *Event) Quarterly() contracts.ScheduleEvent {
	return this.SpliceIntoPosition(1, "0").
		SpliceIntoPosition(2, "0").
		SpliceIntoPosition(3, "1").
		SpliceIntoPosition(4, "1-12/3")
}

func (this *Event) LastDayOfMonth(timeStr string) contracts.ScheduleEvent {
	this.DailyAt(timeStr)

	return this.When(func() bool {
		return carbon.Now(this.timezone).Day() == carbon.Now(this.timezone).EndOfMonth().Day()
	})
}

func (this *Event) TwiceMonthly(first, second int, timeStr string) contracts.ScheduleEvent {
	this.DailyAt(timeStr)
	return this.SpliceIntoPosition(3, fmt.Sprintf("%d,%d", first, second))
}

func (this *Event) MonthlyOn(dayOfMonth int, timeStr string) contracts.ScheduleEvent {
	this.DailyAt(timeStr)
	return this.SpliceIntoPosition(3, strconv.Itoa(dayOfMonth))
}

func (this *Event) Monthly() contracts.ScheduleEvent {
	return this.SpliceIntoPosition(1, "0").
		SpliceIntoPosition(2, "0").
		SpliceIntoPosition(3, "1")
}

func (this *Event) WeeklyOn(dayOfWeek time.Weekday, timeStr string) contracts.ScheduleEvent {
	return this.DailyAt(timeStr).Days(strconv.Itoa(int(dayOfWeek)))
}

func (this *Event) Weekly() contracts.ScheduleEvent {
	return this.SpliceIntoPosition(1, "0").
		SpliceIntoPosition(2, "0").
		SpliceIntoPosition(5, "0")
}

func (this *Event) Sundays() contracts.ScheduleEvent {
	return this.Days(fmt.Sprintf("%d", time.Saturday))
}

func (this *Event) Saturdays() contracts.ScheduleEvent {
	return this.Days(fmt.Sprintf("%d", time.Saturday))
}

func (this *Event) Fridays() contracts.ScheduleEvent {
	return this.Days(fmt.Sprintf("%d", time.Friday))
}

func (this *Event) Thursdays() contracts.ScheduleEvent {
	return this.Days(fmt.Sprintf("%d", time.Thursday))
}

func (this *Event) Wednesdays() contracts.ScheduleEvent {
	return this.Days(fmt.Sprintf("%d", time.Wednesday))
}

func (this *Event) Tuesdays() contracts.ScheduleEvent {
	return this.Days(fmt.Sprintf("%d", time.Tuesday))
}

func (this *Event) Mondays() contracts.ScheduleEvent {
	return this.Days(fmt.Sprintf("%d", time.Monday))
}

func (this *Event) Weekends() contracts.ScheduleEvent {
	return this.Days(fmt.Sprintf("%d,%d", time.Saturday, time.Sunday))
}

func (this *Event) Weekdays() contracts.ScheduleEvent {
	return this.Days(fmt.Sprintf("%d-%d", time.Monday, time.Friday))
}

func (this *Event) TwiceDailyAt(first, second, offset int) contracts.ScheduleEvent {
	return this.SpliceIntoPosition(1, strconv.Itoa(offset)).
		SpliceIntoPosition(2, fmt.Sprintf("%d,%d", first, second))
}

func (this *Event) TwiceDaily(first, second int) contracts.ScheduleEvent {
	return this.TwiceDailyAt(first, second, 0)
}

func (this *Event) DailyAt(timeStr string) contracts.ScheduleEvent {
	segments := strings.Split(timeStr, ":")
	this.SpliceIntoPosition(2, segments[0])

	if len(segments) == 2 {
		return this.SpliceIntoPosition(1, segments[1])
	} else {
		return this.SpliceIntoPosition(1, "0")
	}
}

func (this *Event) Daily() contracts.ScheduleEvent {
	return this.SpliceIntoPosition(1, "0").
		SpliceIntoPosition(2, "0")
}

func (this *Event) EverySixHours() contracts.ScheduleEvent {
	return this.SpliceIntoPosition(1, "0").
		SpliceIntoPosition(2, "*/6")
}

func (this *Event) EveryFourHours() contracts.ScheduleEvent {
	return this.SpliceIntoPosition(1, "0").
		SpliceIntoPosition(2, "*/4")
}

func (this *Event) EveryThreeHours() contracts.ScheduleEvent {
	return this.SpliceIntoPosition(1, "0").
		SpliceIntoPosition(2, "*/3")
}

func (this *Event) EveryTwoHours() contracts.ScheduleEvent {
	return this.SpliceIntoPosition(1, "0").
		SpliceIntoPosition(2, "*/2")
}

func (this *Event) HourlyAt(offset ...int) contracts.ScheduleEvent {
	offsetStrings := make([]string, 0)
	for _, offsetInt := range offset {
		offsetStrings = append(offsetStrings, strconv.Itoa(offsetInt))
	}
	return this.SpliceIntoPosition(1, strings.Join(offsetStrings, ","))
}

func (this *Event) Hourly() contracts.ScheduleEvent {
	return this.SpliceIntoPosition(1, "0")
}

func (this *Event) EveryThirtyMinutes() contracts.ScheduleEvent {
	return this.SpliceIntoPosition(1, "0,30")
}

func (this *Event) EveryFifteenMinutes() contracts.ScheduleEvent {
	return this.SpliceIntoPosition(1, "*/15")
}

func (this *Event) EveryTenMinutes() contracts.ScheduleEvent {
	return this.SpliceIntoPosition(1, "*/10")
}

func (this *Event) EveryFiveMinutes() contracts.ScheduleEvent {
	return this.SpliceIntoPosition(1, "*/5")
}

func (this *Event) EveryFourMinutes() contracts.ScheduleEvent {
	return this.SpliceIntoPosition(1, "*/4")
}

func (this *Event) EveryThreeMinutes() contracts.ScheduleEvent {
	return this.SpliceIntoPosition(1, "*/3")
}

func (this *Event) EveryTwoMinutes() contracts.ScheduleEvent {
	return this.SpliceIntoPosition(1, "*/2")
}

func (this *Event) EveryMinute() contracts.ScheduleEvent {
	return this.SpliceIntoPosition(1, "*")
}

func (this *Event) FiltersPass() bool {
	for _, filter := range this.filters {
		if !filter() {
			return false
		}
	}
	for _, reject := range this.rejects {
		if reject() {
			return false
		}
	}
	return true
}
func (this *Event) When(filter func() bool) contracts.ScheduleEvent {
	this.filters = append(this.filters, filter)
	return this
}
func (this *Event) Skip(reject func() bool) contracts.ScheduleEvent {
	this.rejects = append(this.rejects, reject)
	return this
}

func (this *Event) Cron(expression string) contracts.ScheduleEvent {
	this.expression = expression
	return this
}

func (this *Event) Between(startTime, endTimeStr string) contracts.ScheduleEvent {
	return this.When(this.inTimeInterval(startTime, endTimeStr))
}

func (this *Event) UnlessBetween(startTime, endTimeStr string) contracts.ScheduleEvent {
	return this.Skip(this.inTimeInterval(startTime, endTimeStr))
}

func (this *Event) inTimeInterval(startTime, endTimeStr string) func() bool {
	var (
		startAt = carbon.Now().ParseByFormat(startTime, "H:i", this.timezone)
		endAt   = carbon.Now().ParseByFormat(endTimeStr, "H:i", this.timezone)
	)

	if endAt.Lt(startAt) {
		if startAt.Gt(carbon.Now(this.timezone).SetYear(0000).SetMonth(1).SetDay(1)) {
			startAt.SubDay()
		} else {
			endAt.AddDay()
		}
	}

	return func() bool {
		now := carbon.Now(this.timezone).SetYear(0000).SetMonth(1).SetDay(1)
		return now.Between(startAt, endAt)
	}
}

func (this *Event) MutexName() string {
	return this.mutexName
}

func (this *Event) SetMutexName(mutexName string) contracts.ScheduleEvent {
	this.mutexName = mutexName
	return this
}

func (this *Event) SpliceIntoPosition(position int, value string) contracts.ScheduleEvent {
	segments := strings.Split(this.expression, " ")
	segments[position] = value
	return this.Cron(strings.Join(segments, " "))
}
