package contracts

import (
	"time"
)

type RedisFactory interface {
	Connection(name ...string) RedisConnection
}

type GeoPos struct {
	Longitude, Latitude float64
}

type BitCount struct {
	Start, End int64
}

type GeoLocation struct {
	Name                      string
	Longitude, Latitude, Dist float64
	GeoHash                   int64
}

type GeoRadiusQuery struct {
	Radius float64
	// Can be m, km, ft, or mi. Default is km.
	Unit        string
	WithCoord   bool
	WithDist    bool
	WithGeoHash bool
	Count       int
	// Can be ASC or DESC. Default is no sort order.
	Sort      string
	Store     string
	StoreDist string
}

type ZStore struct {
	Keys    []string
	Weights []float64
	// Can be SUM, MIN or MAX.
	Aggregate string
}

type Z struct {
	Score  float64
	Member interface{}
}

type ZRangeBy struct {
	Min, Max      string
	Offset, Count int64
}

type RedisSubscribeFunc func(message, channel string)

type RedisConnection interface {
	Subscribe(channels []string, closure RedisSubscribeFunc)
	PSubscribe(channels []string, closure RedisSubscribeFunc)
	Command(method string, args ...interface{}) (interface{}, error)

	PubSubChannels(pattern string) ([]string, error)

	PubSubNumSub(channels ...string) (map[string]int64, error)

	PubSubNumPat() (int64, error)

	Publish(channel string, message interface{}) (int64, error)

	// getter start
	Get(key string) (string, error)

	MGet(keys ...string) ([]interface{}, error)

	GetBit(key string, offset int64) (int64, error)

	BitOpAnd(destKey string, keys ...string) (int64, error)

	BitOpNot(destKey string, key string) (int64, error)

	BitOpOr(destKey string, keys ...string) (int64, error)

	BitOpXor(destKey string, keys ...string) (int64, error)

	GetDel(key string) (string, error)

	GetEx(key string, expiration time.Duration) (string, error)

	GetRange(key string, start, end int64) (string, error)

	GetSet(key string, value interface{}) (string, error)

	ClientGetName() (string, error)

	StrLen(key string) (int64, error)

	// getter end
	// keys start

	Keys(pattern string) ([]string, error)

	Del(keys ...string) (int64, error)

	FlushAll() (string, error)

	FlushDB() (string, error)

	Dump(key string) (string, error)

	Exists(keys ...string) (int64, error)

	Expire(key string, expiration time.Duration) (bool, error)

	ExpireAt(key string, tm time.Time) (bool, error)

	PExpire(key string, expiration time.Duration) (bool, error)

	PExpireAt(key string, tm time.Time) (bool, error)

	Migrate(host, port, key string, db int, timeout time.Duration) (string, error)

	Move(key string, db int) (bool, error)

	Persist(key string) (bool, error)

	PTTL(key string) (time.Duration, error)

	TTL(key string) (time.Duration, error)

	RandomKey() (string, error)

	Rename(key, newKey string) (string, error)

	RenameNX(key, newKey string) (bool, error)

	Type(key string) (string, error)

	Wait(numSlaves int, timeout time.Duration) (int64, error)

	Scan(cursor uint64, match string, count int64) ([]string, uint64, error)

	BitCount(key string, count *BitCount) (int64, error)

	// keys end

	// setter start
	Set(key string, value interface{}, expiration time.Duration) (string, error)

	Append(key, value string) (int64, error)

	MSet(values ...interface{}) (string, error)

	MSetNX(values ...interface{}) (bool, error)

	SetNX(key string, value interface{}, expiration time.Duration) (bool, error)

	SetEX(key string, value interface{}, expiration time.Duration) (string, error)

	SetBit(key string, offset int64, value int) (int64, error)

	BitPos(key string, bit int64, pos ...int64) (int64, error)

	SetRange(key string, offset int64, value string) (int64, error)

	Incr(key string) (int64, error)

	Decr(key string) (int64, error)

	IncrBy(key string, value int64) (int64, error)

	DecrBy(key string, value int64) (int64, error)

	IncrByFloat(key string, value float64) (float64, error)

	// setter end

	// hash start
	HGet(key, field string) (string, error)

	HGetAll(key string) (map[string]string, error)

	HMGet(key string, fields ...string) ([]interface{}, error)

	HKeys(key string) ([]string, error)

	HLen(key string) (int64, error)

	HRandField(key string, count int, withValues bool) ([]string, error)

	HScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error)

	HValues(key string) ([]string, error)

	HSet(key string, values ...interface{}) (int64, error)

	HSetNX(key, field string, value interface{}) (bool, error)

	HMSet(key string, values ...interface{}) (bool, error)

	HDel(key string, fields ...string) (int64, error)

	HExists(key string, field string) (bool, error)

	HIncrBy(key string, field string, value int64) (int64, error)

	HIncrByFloat(key string, field string, value float64) (float64, error)

	// hash end

	// set start
	SAdd(key string, members ...interface{}) (int64, error)

	SCard(key string) (int64, error)

	SDiff(keys ...string) ([]string, error)

	SDiffStore(destination string, keys ...string) (int64, error)

	SInter(keys ...string) ([]string, error)

	SInterStore(destination string, keys ...string) (int64, error)

	SIsMember(key string, member interface{}) (bool, error)

	SMembers(key string) ([]string, error)

	SRem(key string, members ...interface{}) (int64, error)

	SPopN(key string, count int64) ([]string, error)

	SPop(key string) (string, error)

	SRandMemberN(key string, count int64) ([]string, error)

	SMove(source, destination string, member interface{}) (bool, error)

	SRandMember(key string) (string, error)

	SUnion(keys ...string) ([]string, error)

	SUnionStore(destination string, keys ...string) (int64, error)

	// set end

	// geo start

	GeoAdd(key string, geoLocation ...*GeoLocation) (int64, error)

	GeoHash(key string, members ...string) ([]string, error)

	GeoPos(key string, members ...string) ([]*GeoPos, error)

	GeoDist(key string, member1, member2, unit string) (float64, error)

	GeoRadius(key string, longitude, latitude float64, query *GeoRadiusQuery) ([]GeoLocation, error)

	GeoRadiusStore(key string, longitude, latitude float64, query *GeoRadiusQuery) (int64, error)

	GeoRadiusByMember(key, member string, query *GeoRadiusQuery) ([]GeoLocation, error)

	GeoRadiusByMemberStore(key, member string, query *GeoRadiusQuery) (int64, error)

	// geo end

	// lists start

	BLPop(timeout time.Duration, keys ...string) ([]string, error)

	BRPop(timeout time.Duration, keys ...string) ([]string, error)

	BRPopLPush(source, destination string, timeout time.Duration) (string, error)

	LIndex(key string, index int64) (string, error)

	LInsert(key, op string, pivot, value interface{}) (int64, error)

	LLen(key string) (int64, error)

	LPop(key string) (string, error)

	LPush(key string, values ...interface{}) (int64, error)

	LPushX(key string, values ...interface{}) (int64, error)

	LRange(key string, start, stop int64) ([]string, error)

	LRem(key string, count int64, value interface{}) (int64, error)

	LSet(key string, index int64, value interface{}) (string, error)

	LTrim(key string, start, stop int64) (string, error)

	RPop(key string) (string, error)

	RPopCount(key string, count int) ([]string, error)

	RPopLPush(source, destination string) (string, error)

	RPush(key string, values ...interface{}) (int64, error)

	RPushX(key string, values ...interface{}) (int64, error)

	// lists end

	// scripting start
	Eval(script string, keys []string, args ...interface{}) (interface{}, error)

	EvalSha(sha1 string, keys []string, args ...interface{}) (interface{}, error)

	ScriptExists(hashes ...string) ([]bool, error)

	ScriptFlush() (string, error)

	ScriptKill() (string, error)

	ScriptLoad(script string) (string, error)

	// scripting end

	// zset start

	ZAdd(key string, members ...*Z) (int64, error)

	ZCard(key string) (int64, error)

	ZCount(key, min, max string) (int64, error)

	ZIncrBy(key string, increment float64, member string) (float64, error)

	ZInterStore(destination string, store *ZStore) (int64, error)

	ZLexCount(key, min, max string) (int64, error)

	ZPopMax(key string, count ...int64) ([]Z, error)

	ZPopMin(key string, count ...int64) ([]Z, error)

	ZRange(key string, start, stop int64) ([]string, error)

	ZRangeByLex(key string, opt *ZRangeBy) ([]string, error)

	ZRevRangeByLex(key string, opt *ZRangeBy) ([]string, error)

	ZRangeByScore(key string, opt *ZRangeBy) ([]string, error)

	ZRank(key, member string) (int64, error)

	ZRem(key string, members ...interface{}) (int64, error)

	ZRemRangeByLex(key, min, max string) (int64, error)

	ZRemRangeByRank(key string, start, stop int64) (int64, error)

	ZRemRangeByScore(key, min, max string) (int64, error)

	ZRevRange(key string, start, stop int64) ([]string, error)

	ZRevRangeByScore(key string, opt *ZRangeBy) ([]string, error)

	ZRevRank(key, member string) (int64, error)

	ZScore(key, member string) (float64, error)

	ZUnionStore(key string, store *ZStore) (int64, error)

	ZScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error)
}
