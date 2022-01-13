package drivers

import (
	"github.com/qbhy/goal/contracts"
	"time"
)

func NewRedisCache(redis contracts.RedisConnection, prefix string) contracts.CacheStore {
	return &RedisStore{
		redis:  redis,
		prefix: prefix,
	}
}

type RedisStore struct {
	redis  contracts.RedisConnection
	prefix string
}

func (this *RedisStore) Get(key string) interface{} {
	result, _ := this.redis.Get(this.getKey(key))
	return result
}

func (this *RedisStore) Many(keys []string) []interface{} {
	results, _ := this.redis.MGet(this.getKeys(keys)...)
	return results
}

func (this *RedisStore) Put(key string, value interface{}, seconds time.Duration) error {
	_, err := this.redis.Set(this.getKey(key), value, seconds)
	return err
}

func (this *RedisStore) Add(key string, value interface{}, ttls ...time.Duration) bool {
	var ttl time.Duration
	if len(ttls) > 0 {
		ttl = ttls[0]
	} else {
		ttl = time.Second * 60 * 60 // default 1 hour
	}
	result, _ := this.redis.SetNX(this.getKey(key), value, ttl)

	return result
}

func (this *RedisStore) Pull(key string, defaultValue ...interface{}) interface{} {
	key = this.getKey(key)
	result, err := this.redis.GetDel(key)

	if err != nil {
		result, err = this.redis.Get(key)
		if result != "" {
			_, _ = this.redis.Del(key)
		}
	}

	if (result == "" || err != nil) && len(defaultValue) > 0 {
		return defaultValue[0]
	}

	return result
}

func (this *RedisStore) PutMany(values map[string]interface{}, seconds time.Duration) error {
	data := make(map[string]interface{})
	for key, value := range values {
		data[this.getKey(key)] = value
	}
	_, err := this.redis.MSet(data)

	for key, _ := range data {
		_, _ = this.redis.Expire(key, seconds)
	}

	return err
}

func (this *RedisStore) Increment(key string, value ...int64) (int64, error) {
	key = this.getKey(key)
	if len(value) > 0 {
		return this.redis.IncrBy(key, value[0])
	}
	return this.redis.Incr(key)
}

func (this *RedisStore) Decrement(key string, value ...int64) (int64, error) {
	key = this.getKey(key)
	if len(value) > 0 {
		return this.redis.DecrBy(key, value[0])
	}
	return this.redis.Decr(key)
}

func (this *RedisStore) Forever(key string, value interface{}) error {
	_, err := this.redis.Set(this.getKey(key), value, -1)
	return err
}

func (this *RedisStore) Forget(key string) error {
	_, err := this.redis.Del(this.getKey(key))
	return err
}

func (this *RedisStore) Flush() error {
	_, err := this.redis.FlushDB()
	return err
}

func (this *RedisStore) GetPrefix() string {
	return this.prefix
}

func (this *RedisStore) getKey(key string) string {
	return this.prefix + key
}

func (this *RedisStore) getKeys(keys []string) []string {
	for index, key := range keys {
		keys[index] = this.getKey(key)
	}
	return keys
}

func (this *RedisStore) Remember(key string, ttl time.Duration, provider contracts.InstanceProvider) interface{} {
	result := this.Get(key)
	if result == nil || result == "" {
		_ = this.Put(key, provider(), ttl)
	}
	return result
}

func (this *RedisStore) RememberForever(key string, provider contracts.InstanceProvider) interface{} {
	return this.Remember(key, -1, provider)
}
