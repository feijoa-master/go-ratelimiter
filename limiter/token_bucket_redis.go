package limiter

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type TokenBucketRedis struct {
	client     *redis.Client
	key        string
	capacity   float64
	refillRate float64
}

func NewTokenBucketRedis(client *redis.Client, key string, capacity float64, refillRate float64) *TokenBucketRedis {
	return &TokenBucketRedis{
		client:     client,
		key:        key,
		capacity:   capacity,
		refillRate: refillRate,
	}
}

var tokenBucketScript = redis.NewScript(`
local tokens = tonumber(redis.call('GET', KEYS[1]))
local last_refill = tonumber(redis.call('GET', KEYS[1] .. ':ts'))
local now = tonumber(ARGV[1])
local capacity = tonumber(ARGV[2])
local rate = tonumber(ARGV[3])

if tokens == nil then
    tokens = capacity
    last_refill = now
end

local elapsed = (now - last_refill) / 1e9
tokens = math.min(capacity, tokens + elapsed * rate)
last_refill = now

if tokens >= 1 then
    tokens = tokens - 1
    redis.call('SET', KEYS[1], tokens, 'EX', 3600)
    redis.call('SET', KEYS[1] .. ':ts', last_refill, 'EX', 3600)
    return 1
end

redis.call('SET', KEYS[1], tokens, 'EX', 3600)
redis.call('SET', KEYS[1] .. ':ts', last_refill, 'EX', 3600)
return 0
`)

func (tb *TokenBucketRedis) Allow() bool {
	ctx := context.Background()
	now := time.Now().UnixNano()

	result, err := tokenBucketScript.Run(
		ctx,
		tb.client,
		[]string{tb.key},
		now,
		tb.capacity,
		tb.refillRate,
	).Int()
	if err != nil {
		return false
	}
	return result == 1
}
