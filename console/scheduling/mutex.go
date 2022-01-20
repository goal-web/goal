package scheduling

import (
	"github.com/goal-web/contracts"
)

type Mutex struct {
	redis contracts.RedisFactory
	store string
}

func (this *Mutex) Create(event *Event) bool {
	_, err := this.redis.Connection(this.store).Set(event.MutexName(), "1", event.expiresAt)
	return err == nil
}
func (this *Mutex) Exists(event *Event) bool {
	num, err := this.redis.Connection(this.store).Exists(event.MutexName())
	return err == nil && num > 0
}

func (this *Mutex) Forget(event *Event) {
	_, _ = this.redis.Connection(this.store).Del(event.MutexName())
}
