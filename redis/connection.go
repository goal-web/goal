package redis

import (
	"context"
	goredis "github.com/go-redis/redis/v8"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/exceptions"
	"time"
)

type Connection struct {
	exceptionHandler contracts.ExceptionHandler
	client           *goredis.Client
}

func (this *Connection) Subscribe(channels []string, closure contracts.RedisSubscribeFunc) {
	go func() {
		pubSub := this.client.Subscribe(context.Background(), channels...)
		defer func(pubSub *goredis.PubSub) {
			err := pubSub.Close()
			if err != nil {
				// 处理异常
				this.exceptionHandler.Handle(SubscribeException{
					exceptions.ResolveException(err),
				})
			}
		}(pubSub)

		pubSubChannel := pubSub.Channel()

		for msg := range pubSubChannel {
			closure(msg.Payload, msg.Channel)
		}
	}()
}

func (this *Connection) PSubscribe(channels []string, closure contracts.RedisSubscribeFunc) {
	go func() {
		pubSub := this.client.PSubscribe(context.Background(), channels...)
		defer func(pubSub *goredis.PubSub) {
			err := pubSub.Close()
			if err != nil {
				// 处理异常
				this.exceptionHandler.Handle(SubscribeException{
					exceptions.ResolveException(err),
				})
			}
		}(pubSub)

		pubSubChannel := pubSub.Channel()

		for msg := range pubSubChannel {
			closure(msg.Payload, msg.Channel)
		}
	}()
}

func (this *Connection) Command(method string, args ...interface{}) (interface{}, error) {
	return this.client.Do(context.Background(), append([]interface{}{method}, args...)...).Result()
}

func (this *Connection) PubSubChannels(pattern string) ([]string, error) {
	return this.client.PubSubChannels(context.Background(), pattern).Result()
}

func (this *Connection) PubSubNumSub(channels ...string) (map[string]int64, error) {
	return this.client.PubSubNumSub(context.Background(), channels...).Result()
}

func (this *Connection) PubSubNumPat() (int64, error) {
	return this.client.PubSubNumPat(context.Background()).Result()
}

func (this *Connection) Publish(channel string, message interface{}) (int64, error) {
	return this.client.Publish(context.Background(), channel, message).Result()
}

func (this *Connection) Client() *goredis.Client {
	return this.client
}

// getter start
func (this *Connection) Get(key string) (string, error) {
	return this.client.Get(context.Background(), key).Result()
}

func (this *Connection) GetBit(key string, offset int64) (int64, error) {
	return this.client.GetBit(context.Background(), key, offset).Result()
}

func (this *Connection) GetDel(key string) (string, error) {
	return this.client.GetDel(context.Background(), key).Result()
}

func (this *Connection) GetEx(key string, expiration time.Duration) (string, error) {
	return this.client.GetEx(context.Background(), key, expiration).Result()
}

func (this *Connection) GetRange(key string, start, end int64) (string, error) {
	return this.client.GetRange(context.Background(), key, start, end).Result()
}

func (this *Connection) GetSet(key string, value interface{}) (string, error) {
	return this.client.GetSet(context.Background(), key, value).Result()
}

func (this *Connection) ClientGetName() (string, error) {
	return this.client.ClientGetName(context.Background()).Result()
}

func (this *Connection) StrLen(key string) (int64, error) {
	return this.client.StrLen(context.Background(), key).Result()
}

// getter end
// keys start

func (this *Connection) Keys(pattern string) ([]string, error) {
	return this.client.Keys(context.Background(), pattern).Result()
}

func (this *Connection) Del(keys ...string) (int64, error) {
	return this.client.Del(context.Background(), keys...).Result()
}

func (this *Connection) Dump(key string) (string, error) {
	return this.client.Dump(context.Background(), key).Result()
}

func (this *Connection) Exists(keys ...string) (int64, error) {
	return this.client.Exists(context.Background(), keys...).Result()
}

func (this *Connection) Expire(key string, expiration time.Duration) (bool, error) {
	return this.client.Expire(context.Background(), key, expiration).Result()
}

func (this *Connection) ExpireAt(key string, tm time.Time) (bool, error) {
	return this.client.ExpireAt(context.Background(), key, tm).Result()
}

func (this *Connection) PExpire(key string, expiration time.Duration) (bool, error) {
	return this.client.PExpire(context.Background(), key, expiration).Result()
}

func (this *Connection) PExpireAt(key string, tm time.Time) (bool, error) {
	return this.client.PExpireAt(context.Background(), key, tm).Result()
}

func (this *Connection) Migrate(host, port, key string, db int, timeout time.Duration) (string, error) {
	return this.client.Migrate(context.Background(), host, port, key, db, timeout).Result()
}

func (this *Connection) Move(key string, db int) (bool, error) {
	return this.client.Move(context.Background(), key, db).Result()
}

func (this *Connection) Persist(key string) (bool, error) {
	return this.client.Persist(context.Background(), key).Result()
}

func (this *Connection) PTTL(key string) (time.Duration, error) {
	return this.client.PTTL(context.Background(), key).Result()
}

func (this *Connection) TTL(key string) (time.Duration, error) {
	return this.client.TTL(context.Background(), key).Result()
}

func (this *Connection) RandomKey() (string, error) {
	return this.client.RandomKey(context.Background()).Result()
}

func (this *Connection) Rename(key, newKey string) (string, error) {
	return this.client.Rename(context.Background(), key, newKey).Result()
}

func (this *Connection) RenameNX(key, newKey string) (bool, error) {
	return this.client.RenameNX(context.Background(), key, newKey).Result()
}

func (this *Connection) Type(key string) (string, error) {
	return this.client.Type(context.Background(), key).Result()
}

func (this *Connection) Wait(numSlaves int, timeout time.Duration) (int64, error) {
	return this.client.Wait(context.Background(), numSlaves, timeout).Result()
}

func (this *Connection) Scan(cursor uint64, match string, count int64) ([]string, uint64, error) {
	return this.client.Scan(context.Background(), cursor, match, count).Result()
}

// keys end

// setter start
func (this *Connection) Set(key string, value interface{}, expiration time.Duration) (string, error) {
	return this.client.Set(context.Background(), key, value, expiration).Result()
}

func (this *Connection) SetNX(key string, value interface{}, expiration time.Duration) (bool, error) {
	return this.client.SetNX(context.Background(), key, value, expiration).Result()
}

func (this *Connection) SetEX(key string, value interface{}, expiration time.Duration) (string, error) {
	return this.client.SetEX(context.Background(), key, value, expiration).Result()
}

func (this *Connection) SetBit(key string, offset int64, value int) (int64, error) {
	return this.client.SetBit(context.Background(), key, offset, value).Result()
}

func (this *Connection) SetRange(key string, offset int64, value string) (int64, error) {
	return this.client.SetRange(context.Background(), key, offset, value).Result()
}

func (this *Connection) Incr(key string) (int64, error) {
	return this.client.Incr(context.Background(), key).Result()
}

func (this *Connection) Decr(key string) (int64, error) {
	return this.client.Decr(context.Background(), key).Result()
}

func (this *Connection) IncrBy(key string, value int64) (int64, error) {
	return this.client.IncrBy(context.Background(), key, value).Result()
}

func (this *Connection) DecrBy(key string, value int64) (int64, error) {
	return this.client.DecrBy(context.Background(), key, value).Result()
}

func (this *Connection) IncrByFloat(key string, value float64) (float64, error) {
	return this.client.IncrByFloat(context.Background(), key, value).Result()
}

// setter end

// hash start
func (this *Connection) HGet(key, field string) (string, error) {
	return this.client.HGet(context.Background(), key, field).Result()
}

func (this *Connection) HGetAll(key string) (map[string]string, error) {
	return this.client.HGetAll(context.Background(), key).Result()
}

func (this *Connection) HMGet(key string, fields ...string) ([]interface{}, error) {
	return this.client.HMGet(context.Background(), key, fields...).Result()
}

func (this *Connection) HKeys(key string) ([]string, error) {
	return this.client.HKeys(context.Background(), key).Result()
}

func (this *Connection) HLen(key string) (int64, error) {
	return this.client.HLen(context.Background(), key).Result()
}

func (this *Connection) HRandField(key string, count int, withValues bool) ([]string, error) {
	return this.client.HRandField(context.Background(), key, count, withValues).Result()
}

func (this *Connection) HScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	return this.client.HScan(context.Background(), key, cursor, match, count).Result()
}

func (this *Connection) HValues(key string) ([]string, error) {
	return this.client.HVals(context.Background(), key).Result()
}

func (this *Connection) HSet(key string, values ...interface{}) (int64, error) {
	return this.client.HSet(context.Background(), key, values...).Result()
}

func (this *Connection) HSetNX(key, field string, value interface{}) (bool, error) {
	return this.client.HSetNX(context.Background(), key, field, value).Result()
}

func (this *Connection) HMSet(key string, values ...interface{}) (bool, error) {
	return this.client.HMSet(context.Background(), key, values...).Result()
}

func (this *Connection) HDel(key string, fields ...string) (int64, error) {
	return this.client.HDel(context.Background(), key, fields...).Result()
}

func (this *Connection) HExists(key string, field string) (bool, error) {
	return this.client.HExists(context.Background(), key, field).Result()
}

func (this *Connection) HIncrBy(key string, field string, value int64) (int64, error) {
	return this.client.HIncrBy(context.Background(), key, field, value).Result()
}

func (this *Connection) HIncrByFloat(key string, field string, value float64) (float64, error) {
	return this.client.HIncrByFloat(context.Background(), key, field, value).Result()
}

// hash end

// set start
func (this *Connection) SAdd(key string, members ...interface{}) (int64, error) {
	return this.client.SAdd(context.Background(), key, members...).Result()
}

func (this *Connection) SCard(key string) (int64, error) {
	return this.client.SCard(context.Background(), key).Result()
}

func (this *Connection) SDiff(keys ...string) ([]string, error) {
	return this.client.SDiff(context.Background(), keys...).Result()
}

func (this *Connection) SDiffStore(destination string, keys ...string) (int64, error) {
	return this.client.SDiffStore(context.Background(), destination, keys...).Result()
}

func (this *Connection) SInter(keys ...string) ([]string, error) {
	return this.client.SInter(context.Background(), keys...).Result()
}

func (this *Connection) SInterStore(destination string, keys ...string) (int64, error) {
	return this.client.SInterStore(context.Background(), destination, keys...).Result()
}

func (this *Connection) SIsMember(key string, member interface{}) (bool, error) {
	return this.client.SIsMember(context.Background(), key, member).Result()
}

func (this *Connection) SMembers(key string) ([]string, error) {
	return this.client.SMembers(context.Background(), key).Result()
}

func (this *Connection) SRem(key string, members ...interface{}) (int64, error) {
	return this.client.SRem(context.Background(), key, members...).Result()
}

func (this *Connection) SPopN(key string, count int64) ([]string, error) {
	return this.client.SPopN(context.Background(), key, count).Result()
}

func (this *Connection) SPop(key string) (string, error) {
	return this.client.SPop(context.Background(), key).Result()
}

func (this *Connection) SRandMemberN(key string, count int64) ([]string, error) {
	return this.client.SRandMemberN(context.Background(), key, count).Result()
}

func (this *Connection) SMove(source, destination string, member interface{}) (bool, error) {
	return this.client.SMove(context.Background(), source, destination, member).Result()
}

func (this *Connection) SRandMember(key string) (string, error) {
	return this.client.SRandMember(context.Background(), key).Result()
}

func (this *Connection) SUnion(keys ...string) ([]string, error) {
	return this.client.SUnion(context.Background(), keys...).Result()
}

func (this *Connection) SUnionStore(destination string, keys ...string) (int64, error) {
	return this.client.SUnionStore(context.Background(), destination, keys...).Result()
}

// set end

// geo start

func (this *Connection) GeoAdd(key string, geoLocation ...*goredis.GeoLocation) (int64, error) {
	return this.client.GeoAdd(context.Background(), key, geoLocation...).Result()
}

func (this *Connection) GeoHash(key string, members ...string) ([]string, error) {
	return this.client.GeoHash(context.Background(), key, members...).Result()
}

func (this *Connection) GeoPos(key string, members ...string) ([]*goredis.GeoPos, error) {
	return this.client.GeoPos(context.Background(), key, members...).Result()
}

func (this *Connection) GeoDist(key string, member1, member2, unit string) (float64, error) {
	return this.client.GeoDist(context.Background(), key, member1, member2, unit).Result()
}

func (this *Connection) GeoRadius(key string, longitude, latitude float64, query *goredis.GeoRadiusQuery) ([]goredis.GeoLocation, error) {
	return this.client.GeoRadius(context.Background(), key, longitude, latitude, query).Result()
}

func (this *Connection) GeoRadiusStore(key string, longitude, latitude float64, query *goredis.GeoRadiusQuery) (int64, error) {
	return this.client.GeoRadiusStore(context.Background(), key, longitude, latitude, query).Result()
}

func (this *Connection) GeoRadiusByMember(key, member string, query *goredis.GeoRadiusQuery) ([]goredis.GeoLocation, error) {
	return this.client.GeoRadiusByMember(context.Background(), key, member, query).Result()
}

func (this *Connection) GeoRadiusByMemberStore(key, member string, query *goredis.GeoRadiusQuery) (int64, error) {
	return this.client.GeoRadiusByMemberStore(context.Background(), key, member, query).Result()
}

// geo end

// lists start

func (this *Connection) BLPop(timeout time.Duration, keys ...string) ([]string, error) {
	return this.client.BLPop(context.Background(), timeout, keys...).Result()
}

func (this *Connection) BRPop(timeout time.Duration, keys ...string) ([]string, error) {
	return this.client.BRPop(context.Background(), timeout, keys...).Result()
}

func (this *Connection) BRPopLPush(source, destination string, timeout time.Duration) (string, error) {
	return this.client.BRPopLPush(context.Background(), source, destination, timeout).Result()
}

func (this *Connection) LIndex(key string, index int64) (string, error) {
	return this.client.LIndex(context.Background(), key, index).Result()
}

func (this *Connection) LInsert(key, op string, pivot, value interface{}) (int64, error) {
	return this.client.LInsert(context.Background(), key, op, pivot, value).Result()
}

func (this *Connection) LLen(key string) (int64, error) {
	return this.client.LLen(context.Background(), key).Result()
}

func (this *Connection) LPop(key string) (string, error) {
	return this.client.LPop(context.Background(), key).Result()
}

func (this *Connection) LPush(key string, values ...interface{}) (int64, error) {
	return this.client.LPush(context.Background(), key, values...).Result()
}

func (this *Connection) LPushX(key string, values ...interface{}) (int64, error) {
	return this.client.LPushX(context.Background(), key, values...).Result()
}

func (this *Connection) LRange(key string, start, stop int64) ([]string, error) {
	return this.client.LRange(context.Background(), key, start, stop).Result()
}

func (this *Connection) LRem(key string, count int64, value interface{}) (int64, error) {
	return this.client.LRem(context.Background(), key, count, value).Result()
}

func (this *Connection) LSet(key string, index int64, value interface{}) (string, error) {
	return this.client.LSet(context.Background(), key, index, value).Result()
}

func (this *Connection) LTrim(key string, start, stop int64) (string, error) {
	return this.client.LTrim(context.Background(), key, start, stop).Result()
}

func (this *Connection) RPop(key string) (string, error) {
	return this.client.RPop(context.Background(), key).Result()
}

func (this *Connection) RPopCount(key string, count int) ([]string, error) {
	return this.client.RPopCount(context.Background(), key, count).Result()
}

func (this *Connection) RPopLPush(source, destination string) (string, error) {
	return this.client.RPopLPush(context.Background(), source, destination).Result()
}

func (this *Connection) RPush(key string, values ...interface{}) (int64, error) {
	return this.client.RPush(context.Background(), key, values...).Result()
}

func (this *Connection) RPushX(key string, values ...interface{}) (int64, error) {
	return this.client.RPushX(context.Background(), key, values...).Result()
}

// lists end
