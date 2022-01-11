package scheduling

import (
	"fmt"
	"github.com/qbhy/goal/utils"
	"strings"
	"time"
)

type Event struct {
	mutex           Mutex
	filters         []func() bool
	rejects         []func() bool
	beforeCallbacks []func()
	afterCallbacks  []func()

	withoutOverlapping bool
	onOneServer        bool

	command    string
	expression string
	mutexName  string
	expiresAt  time.Duration
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
func (this *Event) When(filter func() bool) *Event {
	this.filters = append(this.filters, filter)
	return this
}
func (this *Event) Skip(reject func() bool) *Event {
	this.rejects = append(this.rejects, reject)
	return this
}

func (this *Event) Cron(expression string) *Event {
	this.expression = expression
	return this
}

func (this *Event) Between(startTime, endTime string) *Event {
	//this.When(func() bool {
	//	todo: 判断时间
	//})
	return this
}

func (this *Event) UnlessBetween(startTime, endTime string) *Event {
	//this.Skip(func() bool {
	//	todo: 判断时间
	//})
	return this
}

func (this *Event) MutexName() string {
	if this.mutexName != "" {
		return this.mutexName
	}
	return fmt.Sprintf("goal/schedule-%s", utils.Md5(this.expression+this.command))
}

func (this *Event) spliceIntoPosition(position int, value string) *Event {
	segments := strings.Split(" ", this.expression)
	segments[position-1] = value
	return this.Cron(strings.Join(segments, " "))
}

func (this *Event) WithoutOverlapping(expiresAt int) *Event {
	this.expiresAt = time.Duration(expiresAt)
	this.withoutOverlapping = true
	return this.Skip(func() bool {
		return this.mutex.Exists(this)
	})
}
