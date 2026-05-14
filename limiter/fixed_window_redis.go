package limiter

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type FixedWindowRedis struct {
	client     *redis.Client
	key        string
	limit      int
	windowSize time.Duration
}

func NewFixedWindowRedis(client *redis.Client, key string, limit int, windowSize time.Duration) *FixedWindowRedis {
	return &FixedWindowRedis{
		client:     client,
		key:        key,
		limit:      limit,
		windowSize: windowSize,
	}
}

var fixedWindowScript = redis.NewScript(`
local count = redis.call('INCR', KEYS[1])
if count == 1 then
    redis.call('PEXPIRE', KEYS[1], ARGV[2])
end
if count <= tonumber(ARGV[1]) then
    return 1
end
return 0
`)

func (fw *FixedWindowRedis) Allow() bool {
	ctx := context.Background()

	result, err := fixedWindowScript.Run(
		ctx,
		fw.client,
		[]string{fw.key},
		fw.limit,
		fw.windowSize.Milliseconds(),
	).Int()

	if err != nil {
		return false
	}
	return result == 1

}
