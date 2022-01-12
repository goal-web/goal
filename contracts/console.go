package contracts

import "time"

type Console interface {
	Call(command string, arguments CommandArguments) interface{}
	Run(input ConsoleInput) interface{}
	Schedule(schedule Schedule)
	GetSchedule() Schedule
}

type Command interface {
	Handle() interface{}
	InjectArguments(arguments CommandArguments) error
	GetSignature() string
	GetName() string
	GetDescription() string
	GetHelp() string
}

type ConsoleInput interface {
	GetCommand() string
	GetArguments() CommandArguments
}

type CommandArguments interface {
	FieldsProvider
	Getter
	OptionalGetter
	GetArg(index int) string
	GetArgs() []string
	SetOption(key string, value interface{})
	Exists(key string) bool
	StringArrayOption(key string, defaultValue []string) []string
	IntArrayOption(key string, defaultValue []int) []int
	Int64ArrayOption(key string, defaultValue []int64) []int64
	FloatArrayOption(key string, defaultValue []float32) []float32
	Float64ArrayOption(key string, defaultValue []float64) []float64
}

type Schedule interface {
	UseStore(store string)
	Call(callback interface{}, args ...interface{}) CallbackEvent
	Command(command Command, args ...string) CommandEvent
	Exec(command string, args ...string) CommandEvent

	GetEvents() []ScheduleEvent
}

type ScheduleEvent interface {
	Run(application Application)
	WithoutOverlapping(expiresAt int) ScheduleEvent
	OnOneServer() ScheduleEvent
	MutexName() string
	SetMutexName(mutexName string) ScheduleEvent

	Skip(func() bool) ScheduleEvent
	When(func() bool) ScheduleEvent

	SpliceIntoPosition(position int, value string) ScheduleEvent

	// ManagesFrequencies
	Expression() string
	Cron(expression string) ScheduleEvent
	Timezone(timezone string) ScheduleEvent
	Days(day string, days ...string) ScheduleEvent
	Years(years ...string) ScheduleEvent
	Yearly() ScheduleEvent
	YearlyOn(month time.Month, dayOfMonth int, time string) ScheduleEvent
	Quarterly() ScheduleEvent
	LastDayOfMonth(time string) ScheduleEvent
	TwiceMonthly(first, second int, time string) ScheduleEvent
	Monthly() ScheduleEvent
	MonthlyOn(dayOfMonth int, time string) ScheduleEvent
	WeeklyOn(dayOfWeek time.Weekday, time string) ScheduleEvent
	Weekly() ScheduleEvent
	Sundays() ScheduleEvent
	Saturdays() ScheduleEvent
	Fridays() ScheduleEvent
	Thursdays() ScheduleEvent
	Wednesdays() ScheduleEvent
	Tuesdays() ScheduleEvent
	Mondays() ScheduleEvent
	Weekends() ScheduleEvent
	Weekdays() ScheduleEvent
	TwiceDailyAt(first, second, offset int) ScheduleEvent
	TwiceDaily(first, second int) ScheduleEvent
	DailyAt(time string) ScheduleEvent
	Daily() ScheduleEvent
	EverySixHours() ScheduleEvent
	EveryFourHours() ScheduleEvent
	EveryThreeHours() ScheduleEvent
	EveryTwoHours() ScheduleEvent
	HourlyAt(offset ...int) ScheduleEvent
	Hourly() ScheduleEvent
	EveryThirtyMinutes() ScheduleEvent
	EveryFifteenMinutes() ScheduleEvent
	EveryTenMinutes() ScheduleEvent
	EveryFiveMinutes() ScheduleEvent
	EveryFourMinutes() ScheduleEvent
	EveryThreeMinutes() ScheduleEvent
	EveryTwoMinutes() ScheduleEvent
	EveryMinute() ScheduleEvent
	UnlessBetween(startTime, endTime string) ScheduleEvent
	Between(startTime, endTime string) ScheduleEvent

	EveryThirtySeconds() ScheduleEvent
	EveryFifteenSeconds() ScheduleEvent
	EveryTenSeconds() ScheduleEvent
	EveryFiveSeconds() ScheduleEvent
	EveryFourSeconds() ScheduleEvent
	EveryThreeSeconds() ScheduleEvent
	EveryTwoSeconds() ScheduleEvent
	EverySecond() ScheduleEvent
}

type CallbackEvent interface {
	ScheduleEvent
	Description(description string) CallbackEvent
}
type CommandEvent interface {
	ScheduleEvent
}
