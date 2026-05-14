package limiter

import (
	"github.com/redis/go-redis/v9"
	"testing"
)

func TestNewTokenBucketRedis(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	tb := NewTokenBucketRedis(client, "test:bucket", 2, 1)

	if !tb.Allow() {
		t.Error("1й запрос должен пройти")
	}
	if !tb.Allow() {
		t.Error("2й запрос должен пройти")
	}
	if tb.Allow() {
		t.Error("3й запрос должен быть заблокирован")
	}

}
